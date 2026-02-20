// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

// Package config handles environment configuration, UE path detection,
// and project root discovery for the mcp-unreal server.
//
// All configuration is read from environment variables with sensible
// platform-dependent defaults. See IMPLEMENTATION.md §6 and CLAUDE.md
// Environment Variables table for the full list.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Config holds all runtime configuration for the MCP server.
type Config struct {
	// UEEditorPath is the path to the UnrealEditor-Cmd binary.
	UEEditorPath string

	// ProjectRoot is the path to the UE project root (directory containing .uproject).
	ProjectRoot string

	// UProjectFile is the full path to the .uproject file, if found.
	UProjectFile string

	// RCAPIPort is the UE Remote Control API HTTP port (default 30010).
	RCAPIPort int

	// PluginPort is the MCPUnreal editor plugin HTTP port (default 8090).
	PluginPort int

	// LogLevel is the slog level for the server.
	LogLevel slog.Level

	// DocsIndexPath is the path to the bleve documentation index.
	DocsIndexPath string
}

// Load reads configuration from environment variables and applies
// platform-dependent defaults. It does not fail on missing values —
// tools check availability at call time (graceful degradation per
// IMPLEMENTATION.md §10).
func Load() *Config {
	cfg := &Config{
		UEEditorPath:  envOrDefault("UE_EDITOR_PATH", defaultUEEditorPath()),
		RCAPIPort:     envIntOrDefault("RC_API_PORT", 30010),
		PluginPort:    envIntOrDefault("PLUGIN_PORT", 8090),
		LogLevel:      parseLogLevel(envOrDefault("MCP_UNREAL_LOG_LEVEL", "info")),
		DocsIndexPath: envOrDefault("MCP_UNREAL_DOCS_INDEX", "./docs/index.bleve"),
	}

	// Project root: explicit env var or auto-detect from cwd.
	// MCP_UNREAL_PROJECT can be a .uproject file path or a project directory.
	projectEnv := os.Getenv("MCP_UNREAL_PROJECT")
	if projectEnv != "" {
		if strings.HasSuffix(projectEnv, ".uproject") {
			// Given a .uproject file path directly.
			if _, err := os.Stat(projectEnv); err == nil {
				cfg.UProjectFile = projectEnv
				cfg.ProjectRoot = filepath.Dir(projectEnv)
			} else {
				cfg.ProjectRoot = filepath.Dir(projectEnv)
			}
		} else {
			cfg.ProjectRoot = projectEnv
			cfg.UProjectFile = findUProjectFile(projectEnv)
		}
	} else {
		root, uproject := detectProjectRoot()
		cfg.ProjectRoot = root
		cfg.UProjectFile = uproject
	}

	return cfg
}

// RCAPIURL returns the base URL for the UE Remote Control API.
func (c *Config) RCAPIURL() string {
	return fmt.Sprintf("http://127.0.0.1:%d", c.RCAPIPort)
}

// PluginURL returns the base URL for the MCPUnreal editor plugin.
func (c *Config) PluginURL() string {
	return fmt.Sprintf("http://127.0.0.1:%d", c.PluginPort)
}

// defaultUEEditorPath returns the platform-dependent default path to
// UnrealEditor-Cmd. See CLAUDE.md Environment Variables table.
func defaultUEEditorPath() string {
	switch runtime.GOOS {
	case "darwin":
		return "/Users/Shared/Epic Games/UE_5.7/Engine/Binaries/Mac/UnrealEditor-Cmd"
	case "windows":
		return `C:\Program Files\Epic Games\UE_5.7\Engine\Binaries\Win64\UnrealEditor-Cmd.exe`
	case "linux":
		return "/opt/UnrealEngine/Engine/Binaries/Linux/UnrealEditor-Cmd"
	default:
		return "UnrealEditor-Cmd"
	}
}

// detectProjectRoot walks up from the current working directory looking
// for a .uproject file. Returns (root, uprojectPath) or ("", "") if not found.
func detectProjectRoot() (string, string) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", ""
	}

	dir := cwd
	for {
		uproject := findUProjectFile(dir)
		if uproject != "" {
			return dir, uproject
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached filesystem root
		}
		dir = parent
	}

	return cwd, ""
}

// findUProjectFile looks for a .uproject file in the given directory.
// Returns the full path to the first one found, or "" if none.
func findUProjectFile(dir string) string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return ""
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".uproject") {
			return filepath.Join(dir, e.Name())
		}
	}
	return ""
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envIntOrDefault(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}

func parseLogLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
