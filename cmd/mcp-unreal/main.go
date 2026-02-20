// Command mcp-unreal is an MCP (Model Context Protocol) server that gives
// AI coding agents complete autonomous control over a UE 5.7 project.
//
// It communicates via stdio (JSON-RPC 2.0) and provides tools for builds,
// tests, editor manipulation, Blueprint editing, mesh generation, and
// documentation lookup.
//
// See IMPLEMENTATION.md §2 for the architecture and CLAUDE.md for the
// full project specification.
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/remiphilippe/mcp-unreal/internal/config"
	"github.com/remiphilippe/mcp-unreal/internal/docs"
	"github.com/remiphilippe/mcp-unreal/internal/editor"
	"github.com/remiphilippe/mcp-unreal/internal/headless"
	"github.com/remiphilippe/mcp-unreal/internal/status"
)

// Version is set at build time via -ldflags.
var Version = "0.2.0"

func main() {
	// Parse CLI flags.
	buildIndex := flag.Bool("build-index", false, "Build the documentation search index and exit")
	docsIndex := flag.String("docs-index", "", "Path to the bleve documentation index (overrides MCP_UNREAL_DOCS_INDEX)")
	logLevel := flag.String("log-level", "", "Log level: debug, info, warn, error (overrides MCP_UNREAL_LOG_LEVEL)")
	showVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println("mcp-unreal", Version)
		os.Exit(0)
	}

	// Load configuration from environment (IMPLEMENTATION.md §6, config.go).
	cfg := config.Load()

	// CLI flags override env vars.
	if *logLevel != "" {
		if err := os.Setenv("MCP_UNREAL_LOG_LEVEL", *logLevel); err == nil {
			cfg = config.Load() // reload to pick up new level
		}
	}
	if *docsIndex != "" {
		cfg.DocsIndexPath = *docsIndex
	}

	// All logging goes to stderr — stdout is sacred (CLAUDE.md Security §1).
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}))
	slog.SetDefault(logger)

	// Handle --build-index mode (IMPLEMENTATION.md §4.3).
	if *buildIndex {
		if err := buildDocsIndex(cfg, logger); err != nil {
			logger.Error("failed to build index", "error", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	// Create MCP server (IMPLEMENTATION.md §2).
	server := mcp.NewServer(
		&mcp.Implementation{Name: "mcp-unreal", Version: Version},
		&mcp.ServerOptions{
			Instructions: "MCP server for Unreal Engine 5.7. " +
				"Use the 'status' tool first to check connectivity and available features.",
			Logger: logger,
		},
	)

	// Register tools.
	registerTools(server, cfg, logger)

	// Set up graceful shutdown.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	logger.Info("starting mcp-unreal", "version", Version, "project", cfg.ProjectRoot)

	// Run stdio transport (blocks until client disconnects).
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		logger.Error("server exited with error", "error", err)
		os.Exit(1)
	}
}

// registerTools wires up all MCP tool handlers.
// Tools are added in phases — see IMPLEMENTATION.md §9 for the roadmap.
func registerTools(server *mcp.Server, cfg *config.Config, logger *slog.Logger) {
	// Phase 1: Status tool.
	statusHandler := &status.Handler{Config: cfg, Version: Version}
	statusHandler.Register(server)

	// Phase 2: Headless build & test tools.
	headlessHandler := &headless.Handler{Config: cfg, Logger: logger}
	headlessHandler.Register(server)
	headlessHandler.RegisterTests(server)
	headlessHandler.RegisterLog(server)
	headlessHandler.RegisterCook(server)
	headlessHandler.RegisterConfig(server)
	headlessHandler.RegisterProject(server)

	// Phase 3: Documentation lookup tools (IMPLEMENTATION.md §4).
	docIdx, err := docs.OpenOrCreate(cfg.DocsIndexPath)
	if err != nil {
		logger.Warn("documentation index unavailable, lookup tools disabled", "path", cfg.DocsIndexPath, "error", err)
	} else {
		docIdx.Register(server)
		logger.Debug("documentation index loaded", "path", cfg.DocsIndexPath)
	}

	// Phase 4: Editor communication tools (IMPLEMENTATION.md §3.3–§3.11).
	editorClient := editor.NewClient(cfg, logger)
	editorHandler := &editor.Handler{Client: editorClient, Logger: logger}
	editorHandler.RegisterProperties(server)
	editorHandler.RegisterActors(server)
	editorHandler.RegisterUtilities(server)

	// Phase 6: Blueprint & Animation Blueprint tools.
	editorHandler.RegisterBlueprints(server)
	editorHandler.RegisterAnimBlueprints(server)

	// Phase 7: Assets, materials, characters, input, levels, editor utilities.
	editorHandler.RegisterAssets(server)
	editorHandler.RegisterMaterials(server)
	editorHandler.RegisterCharacters(server)
	editorHandler.RegisterInput(server)
	editorHandler.RegisterLevels(server)
	editorHandler.RegisterEditorUtils(server)

	// Phase 8: Mesh tools.
	editorHandler.RegisterMesh(server)

	// Phase 10+: Component introspection (issue #40).
	editorHandler.RegisterComponents(server)

	// Phase 10+: ISM/HISM management (issue #41).
	editorHandler.RegisterISM(server)

	// Phase 10+: Fab marketplace cache/import (issue #42).
	editorHandler.RegisterFab(server)

	// Phase 10+: Texture management (issue #44).
	editorHandler.RegisterTextures(server)

	// Phase 10+: Subsystem introspection (issue #45).
	editorHandler.RegisterSubsystems(server)

	// Phase 10+: DataTable/DataAsset management (issue #46).
	editorHandler.RegisterDataAssets(server)

	// Phase 10+: UI widget introspection (issue #47).
	editorHandler.RegisterUIQuery(server)

	// Phase 10+: Network debug (issue #48).
	editorHandler.RegisterNetworkDebug(server)

	// Phase 10: PCG, GAS, Niagara tools.
	editorHandler.RegisterPCG(server)
	editorHandler.RegisterGAS(server)
	editorHandler.RegisterNiagara(server)

	logger.Debug("registered tools", "count", 49)
}

// buildDocsIndex creates or rebuilds the documentation search index
// from markdown source files (IMPLEMENTATION.md §4.3).
func buildDocsIndex(cfg *config.Config, logger *slog.Logger) error {
	logger.Info("building documentation index", "output", cfg.DocsIndexPath)

	// Remove existing index to rebuild from scratch.
	if err := os.RemoveAll(cfg.DocsIndexPath); err != nil {
		return fmt.Errorf("removing old index: %w", err)
	}

	idx, err := docs.CreateIndex(cfg.DocsIndexPath)
	if err != nil {
		return fmt.Errorf("creating index: %w", err)
	}
	defer func() { _ = idx.Close() }()

	total := 0

	// Find the docs directory relative to the binary or working directory.
	docsRoot := findDocsRoot()
	if docsRoot == "" {
		logger.Warn("no docs/ directory found, index will be empty")
	} else {
		// Ingest UE 5.7 docs.
		ueDir := filepath.Join(docsRoot, "ue5.7")
		if info, err := os.Stat(ueDir); err == nil && info.IsDir() {
			n, err := docs.IngestDirectory(idx, ueDir, "ue5.7", logger)
			if err != nil {
				return fmt.Errorf("ingesting UE docs: %w", err)
			}
			total += n
			logger.Info("indexed UE 5.7 docs", "count", n)
		}

		// Ingest RealtimeMesh docs.
		rmcDir := filepath.Join(docsRoot, "realtimemesh")
		if info, err := os.Stat(rmcDir); err == nil && info.IsDir() {
			n, err := docs.IngestDirectory(idx, rmcDir, "realtimemesh", logger)
			if err != nil {
				return fmt.Errorf("ingesting RealtimeMesh docs: %w", err)
			}
			total += n
			logger.Info("indexed RealtimeMesh docs", "count", n)
		}
	}

	logger.Info("documentation index built", "total_docs", total, "path", cfg.DocsIndexPath)
	return nil
}

// findDocsRoot locates the docs/ directory by checking common locations.
func findDocsRoot() string {
	// Check relative to working directory.
	if info, err := os.Stat("docs"); err == nil && info.IsDir() {
		return "docs"
	}

	// Check relative to the executable.
	if exe, err := os.Executable(); err == nil {
		docsDir := filepath.Join(filepath.Dir(exe), "docs")
		if info, err := os.Stat(docsDir); err == nil && info.IsDir() {
			return docsDir
		}
	}

	return ""
}
