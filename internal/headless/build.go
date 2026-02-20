// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

// Package headless implements MCP tools that invoke UnrealEditor-Cmd as
// a subprocess for builds, tests, and project file generation.
//
// These tools work without the editor running — they use the headless
// command-line binary. See IMPLEMENTATION.md §3.1 and §3.2.
//
// All exec.Command calls follow CLAUDE.md Security §7:
//   - Explicit argument arrays (no shell expansion)
//   - Timeout via context.WithTimeout
//   - Capture stdout and stderr
//   - Log the command at debug level
package headless

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/remiphilippe/mcp-unreal/internal/config"
)

// Handler holds references needed by headless tools.
type Handler struct {
	Config *config.Config
	Logger *slog.Logger
}

// --- build_project ---

// BuildInput defines parameters for the build_project tool.
type BuildInput struct {
	Target   string `json:"target,omitempty" jsonschema:"Build target name (e.g. MyProjectEditor). Defaults to the project editor target."`
	Config   string `json:"config,omitempty" jsonschema:"Build configuration: Development, DebugGame, Shipping. Default Development."`
	Platform string `json:"platform,omitempty" jsonschema:"Target platform: Mac, Win64, Linux. Defaults to current platform."`
	Clean    bool   `json:"clean,omitempty" jsonschema:"If true, clean before building."`
}

// BuildOutput is returned by the build_project tool.
type BuildOutput struct {
	Success      bool     `json:"success" jsonschema:"whether the build succeeded"`
	ExitCode     int      `json:"exit_code" jsonschema:"process exit code"`
	Duration     string   `json:"duration" jsonschema:"build duration"`
	ErrorCount   int      `json:"error_count" jsonschema:"number of build errors"`
	WarningCount int      `json:"warning_count" jsonschema:"number of build warnings"`
	Errors       []string `json:"errors,omitempty" jsonschema:"build error messages"`
	Warnings     []string `json:"warnings,omitempty" jsonschema:"build warning messages (first 20)"`
	LogTail      string   `json:"log_tail,omitempty" jsonschema:"last 50 lines of build output for context"`
}

// Register adds the build and generate tools to the MCP server.
func (h *Handler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "build_project",
		Description: "Build the Unreal Engine project using UnrealEditor-Cmd / UBT. " +
			"Returns structured JSON with success/failure, error count, and error details. " +
			"Does not require the editor to be running. " +
			"Set UE_EDITOR_PATH if UnrealEditor-Cmd is not at the default location.",
	}, h.BuildProject)

	mcp.AddTool(server, &mcp.Tool{
		Name: "generate_project_files",
		Description: "Regenerate IDE project files (.xcworkspace on Mac, .sln on Windows) " +
			"after adding or removing C++ modules. Does not require the editor to be running.",
	}, h.GenerateProjectFiles)
}

// BuildProject implements the build_project tool.
func (h *Handler) BuildProject(ctx context.Context, req *mcp.CallToolRequest, input BuildInput) (*mcp.CallToolResult, BuildOutput, error) {
	editorPath := h.Config.UEEditorPath
	if _, err := os.Stat(editorPath); err != nil {
		return nil, BuildOutput{}, fmt.Errorf(
			"UnrealEditor-Cmd not found at %s — set UE_EDITOR_PATH env var or install UE 5.7",
			editorPath,
		)
	}

	projectFile := h.Config.UProjectFile
	if projectFile == "" {
		return nil, BuildOutput{}, fmt.Errorf(
			"no .uproject file found — set MCP_UNREAL_PROJECT or run from inside a UE project directory",
		)
	}

	cfg := input.Config
	if cfg == "" {
		cfg = "Development"
	}

	platform := input.Platform
	if platform == "" {
		platform = defaultPlatform()
	}

	target := input.Target
	if target == "" {
		// Derive editor target from project name: MyProject -> MyProjectEditor.
		base := strings.TrimSuffix(filepath.Base(projectFile), ".uproject")
		target = base + "Editor"
	}

	args := []string{
		projectFile,
		"-Target=" + target,
		"-Configuration=" + cfg,
		"-Platform=" + platform,
	}
	if input.Clean {
		args = append(args, "-Clean")
	}
	args = append(args, "-NoP4", "-UTF8Output")

	start := time.Now()
	stdout, stderr, exitCode, err := h.runCommand(ctx, editorPath, args, 30*time.Minute)
	duration := time.Since(start)

	if err != nil && exitCode == -1 {
		return nil, BuildOutput{}, fmt.Errorf("failed to run build command: %w", err)
	}

	combined := stdout + "\n" + stderr
	errors := parseBuildErrors(combined)
	warnings := parseBuildWarnings(combined)

	out := BuildOutput{
		Success:      exitCode == 0,
		ExitCode:     exitCode,
		Duration:     duration.Round(time.Second).String(),
		ErrorCount:   len(errors),
		WarningCount: len(warnings),
		Errors:       errors,
		LogTail:      lastNLines(combined, 50),
	}

	// Cap warnings to avoid flooding context (IMPLEMENTATION.md §10: token budget).
	if len(warnings) > 20 {
		out.Warnings = warnings[:20]
	} else {
		out.Warnings = warnings
	}

	return nil, out, nil
}

// --- cook_project ---

// CookInput defines parameters for the cook_project tool.
type CookInput struct {
	Platform  string `json:"platform,omitempty" jsonschema:"Target platform: Mac, Win64, Linux, IOS, Android. Defaults to current platform."`
	Config    string `json:"config,omitempty" jsonschema:"Build configuration: Development, Shipping. Default Development."`
	Iterative bool   `json:"iterative,omitempty" jsonschema:"If true, only cook changed content (faster). Default false."`
}

// CookOutput is returned by the cook_project tool.
type CookOutput struct {
	Success  bool   `json:"success" jsonschema:"whether cooking succeeded"`
	ExitCode int    `json:"exit_code" jsonschema:"process exit code"`
	Duration string `json:"duration" jsonschema:"cook duration"`
	LogTail  string `json:"log_tail,omitempty" jsonschema:"last 50 lines of cook output"`
}

// RegisterCook adds the cook_project tool to the MCP server.
func (h *Handler) RegisterCook(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "cook_project",
		Description: "Cook (package) content for a target platform using RunUAT. " +
			"This bakes all assets for deployment. Can take several minutes for large projects. " +
			"Does not require the editor to be running. " +
			"Set iterative=true for incremental cooks (only changed content).",
	}, h.CookProject)
}

// CookProject implements the cook_project tool.
func (h *Handler) CookProject(ctx context.Context, req *mcp.CallToolRequest, input CookInput) (*mcp.CallToolResult, CookOutput, error) {
	projectFile := h.Config.UProjectFile
	if projectFile == "" {
		return nil, CookOutput{}, fmt.Errorf(
			"no .uproject file found — set MCP_UNREAL_PROJECT or run from inside a UE project directory",
		)
	}

	// Find RunUAT script.
	runUAT := findRunUATScript(h.Config.UEEditorPath)
	if runUAT == "" {
		return nil, CookOutput{}, fmt.Errorf(
			"RunUAT script not found — ensure UE 5.7 is installed and UE_EDITOR_PATH is correct",
		)
	}

	platform := input.Platform
	if platform == "" {
		platform = defaultPlatform()
	}

	cfg := input.Config
	if cfg == "" {
		cfg = "Development"
	}

	args := []string{
		"BuildCookRun",
		"-project=" + projectFile,
		"-noP4",
		"-platform=" + platform,
		"-clientconfig=" + cfg,
		"-cook",
		"-skipstage",
		"-utf8output",
	}
	if input.Iterative {
		args = append(args, "-iterate")
	}

	start := time.Now()
	stdout, stderr, exitCode, err := h.runCommand(ctx, runUAT, args, 60*time.Minute)
	duration := time.Since(start)

	if err != nil && exitCode == -1 {
		return nil, CookOutput{}, fmt.Errorf("failed to run cook command: %w", err)
	}

	combined := stdout + "\n" + stderr
	return nil, CookOutput{
		Success:  exitCode == 0,
		ExitCode: exitCode,
		Duration: duration.Round(time.Second).String(),
		LogTail:  lastNLines(combined, 50),
	}, nil
}

// findRunUATScript locates RunUAT relative to the UE editor binary.
func findRunUATScript(editorPath string) string {
	engineDir := editorPath
	for i := 0; i < 3; i++ {
		engineDir = filepath.Dir(engineDir)
	}

	candidates := []string{
		filepath.Join(engineDir, "Build", "BatchFiles", "RunUAT.sh"),
		filepath.Join(engineDir, "Build", "BatchFiles", "RunUAT.bat"),
	}

	for _, c := range candidates {
		if fileExists(c) {
			return c
		}
	}
	return ""
}

// --- generate_project_files ---

// GenerateProjectFilesInput defines parameters for generate_project_files.
type GenerateProjectFilesInput struct{}

// GenerateProjectFilesOutput is returned by generate_project_files.
type GenerateProjectFilesOutput struct {
	Success  bool   `json:"success" jsonschema:"whether generation succeeded"`
	ExitCode int    `json:"exit_code" jsonschema:"process exit code"`
	Duration string `json:"duration" jsonschema:"generation duration"`
	Output   string `json:"output,omitempty" jsonschema:"command output (last 30 lines)"`
}

// GenerateProjectFiles implements the generate_project_files tool.
func (h *Handler) GenerateProjectFiles(ctx context.Context, req *mcp.CallToolRequest, input GenerateProjectFilesInput) (*mcp.CallToolResult, GenerateProjectFilesOutput, error) {
	projectFile := h.Config.UProjectFile
	if projectFile == "" {
		return nil, GenerateProjectFilesOutput{}, fmt.Errorf(
			"no .uproject file found — set MCP_UNREAL_PROJECT or run from inside a UE project directory",
		)
	}

	// Find GenerateProjectFiles script relative to UE editor path.
	script := findGenerateProjectFilesScript(h.Config.UEEditorPath)
	if script == "" {
		return nil, GenerateProjectFilesOutput{}, fmt.Errorf(
			"GenerateProjectFiles script not found — ensure UE 5.7 is installed and UE_EDITOR_PATH is correct",
		)
	}

	args := []string{"-project=" + projectFile, "-game", "-engine"}

	start := time.Now()
	stdout, stderr, exitCode, err := h.runCommand(ctx, script, args, 10*time.Minute)
	duration := time.Since(start)

	if err != nil && exitCode == -1 {
		return nil, GenerateProjectFilesOutput{}, fmt.Errorf("failed to run GenerateProjectFiles: %w", err)
	}

	combined := stdout + "\n" + stderr
	return nil, GenerateProjectFilesOutput{
		Success:  exitCode == 0,
		ExitCode: exitCode,
		Duration: duration.Round(time.Second).String(),
		Output:   lastNLines(combined, 30),
	}, nil
}

// --- helpers ---

// runCommand executes a command with timeout, capturing stdout and stderr.
// Returns (stdout, stderr, exitCode, error). exitCode is -1 if the process
// could not be started.
func (h *Handler) runCommand(ctx context.Context, name string, args []string, timeout time.Duration) (string, string, int, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	h.Logger.Debug("executing command", "cmd", name, "args", args)

	cmd := exec.CommandContext(ctx, name, args...)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return "", "", -1, err
		}
	}

	return stdoutBuf.String(), stderrBuf.String(), exitCode, nil
}

var (
	buildErrorRe   = regexp.MustCompile(`(?m)^.+\(\d+\):\s+error\s+.*$`)
	buildWarningRe = regexp.MustCompile(`(?m)^.+\(\d+\):\s+warning\s+.*$`)
)

func parseBuildErrors(output string) []string {
	matches := buildErrorRe.FindAllString(output, -1)
	return dedup(matches)
}

func parseBuildWarnings(output string) []string {
	matches := buildWarningRe.FindAllString(output, -1)
	return dedup(matches)
}

func dedup(items []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" && !seen[trimmed] {
			seen[trimmed] = true
			result = append(result, trimmed)
		}
	}
	return result
}

func lastNLines(s string, n int) string {
	lines := strings.Split(s, "\n")
	if len(lines) <= n {
		return strings.TrimSpace(s)
	}
	return strings.TrimSpace(strings.Join(lines[len(lines)-n:], "\n"))
}

func defaultPlatform() string {
	switch {
	case fileExists("/System/Library/CoreServices/SystemVersion.plist"):
		return "Mac"
	case fileExists(`C:\Windows\System32`):
		return "Win64"
	default:
		return "Linux"
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// findGenerateProjectFilesScript locates the GenerateProjectFiles script
// relative to the UE editor binary path.
func findGenerateProjectFilesScript(editorPath string) string {
	// Walk up from editor binary to Engine root, then look for the script.
	// macOS: .../Engine/Binaries/Mac/UnrealEditor-Cmd -> .../Engine/Build/BatchFiles/Mac/GenerateProjectFiles.sh
	// Windows: .../Engine/Binaries/Win64/UnrealEditor-Cmd.exe -> .../Engine/Build/BatchFiles/GenerateProjectFiles.bat
	// Linux: .../Engine/Binaries/Linux/UnrealEditor-Cmd -> .../Engine/Build/BatchFiles/Linux/GenerateProjectFiles.sh

	engineDir := editorPath
	for i := 0; i < 3; i++ {
		engineDir = filepath.Dir(engineDir)
	}

	candidates := []string{
		filepath.Join(engineDir, "Build", "BatchFiles", "Mac", "GenerateProjectFiles.sh"),
		filepath.Join(engineDir, "Build", "BatchFiles", "GenerateProjectFiles.bat"),
		filepath.Join(engineDir, "Build", "BatchFiles", "Linux", "GenerateProjectFiles.sh"),
	}

	for _, c := range candidates {
		if fileExists(c) {
			return c
		}
	}
	return ""
}
