// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package config

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDefaultUEEditorPath(t *testing.T) {
	path := defaultUEEditorPath()
	if path == "" {
		t.Fatal("defaultUEEditorPath returned empty string")
	}

	switch runtime.GOOS {
	case "darwin":
		if path != "/Users/Shared/Epic Games/UE_5.7/Engine/Binaries/Mac/UnrealEditor-Cmd" {
			t.Errorf("unexpected macOS default: %s", path)
		}
	case "windows":
		if path != `C:\Program Files\Epic Games\UE_5.7\Engine\Binaries\Win64\UnrealEditor-Cmd.exe` {
			t.Errorf("unexpected Windows default: %s", path)
		}
	case "linux":
		if path != "/opt/UnrealEngine/Engine/Binaries/Linux/UnrealEditor-Cmd" {
			t.Errorf("unexpected Linux default: %s", path)
		}
	}
}

func TestEnvOrDefault(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		setVal   string
		fallback string
		want     string
	}{
		{"uses env when set", "TEST_CONFIG_A", "custom", "default", "custom"},
		{"uses fallback when unset", "TEST_CONFIG_B_UNSET", "", "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVal != "" {
				t.Setenv(tt.key, tt.setVal)
			}
			got := envOrDefault(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("envOrDefault(%q) = %q, want %q", tt.key, got, tt.want)
			}
		})
	}
}

func TestEnvIntOrDefault(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		setVal   string
		fallback int
		want     int
	}{
		{"uses env int", "TEST_INT_A", "8080", 30010, 8080},
		{"uses fallback when unset", "TEST_INT_B_UNSET", "", 30010, 30010},
		{"uses fallback on invalid", "TEST_INT_C", "notanumber", 30010, 30010},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setVal != "" {
				t.Setenv(tt.key, tt.setVal)
			}
			got := envIntOrDefault(tt.key, tt.fallback)
			if got != tt.want {
				t.Errorf("envIntOrDefault(%q) = %d, want %d", tt.key, got, tt.want)
			}
		})
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input string
		want  slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"warning", slog.LevelWarn},
		{"error", slog.LevelError},
		{"DEBUG", slog.LevelDebug},
		{"", slog.LevelInfo},
		{"unknown", slog.LevelInfo},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseLogLevel(tt.input)
			if got != tt.want {
				t.Errorf("parseLogLevel(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFindUProjectFile(t *testing.T) {
	// Create a temp dir with a .uproject file.
	dir := t.TempDir()
	uproject := filepath.Join(dir, "MyProject.uproject")
	if err := os.WriteFile(uproject, []byte(`{"EngineAssociation":"5.7"}`), 0600); err != nil {
		t.Fatal(err)
	}

	got := findUProjectFile(dir)
	if got != uproject {
		t.Errorf("findUProjectFile(%q) = %q, want %q", dir, got, uproject)
	}

	// Non-existent directory returns empty.
	got = findUProjectFile("/nonexistent/path/12345")
	if got != "" {
		t.Errorf("findUProjectFile(nonexistent) = %q, want empty", got)
	}
}

func TestLoadRespectsEnvVars(t *testing.T) {
	t.Setenv("UE_EDITOR_PATH", "/custom/editor")
	t.Setenv("RC_API_PORT", "9999")
	t.Setenv("PLUGIN_PORT", "7777")
	t.Setenv("MCP_UNREAL_LOG_LEVEL", "debug")
	t.Setenv("MCP_UNREAL_DOCS_INDEX", "/custom/index.bleve")

	cfg := Load()

	if cfg.UEEditorPath != "/custom/editor" {
		t.Errorf("UEEditorPath = %q, want /custom/editor", cfg.UEEditorPath)
	}
	if cfg.RCAPIPort != 9999 {
		t.Errorf("RCAPIPort = %d, want 9999", cfg.RCAPIPort)
	}
	if cfg.PluginPort != 7777 {
		t.Errorf("PluginPort = %d, want 7777", cfg.PluginPort)
	}
	if cfg.LogLevel != slog.LevelDebug {
		t.Errorf("LogLevel = %v, want debug", cfg.LogLevel)
	}
	if cfg.DocsIndexPath != "/custom/index.bleve" {
		t.Errorf("DocsIndexPath = %q, want /custom/index.bleve", cfg.DocsIndexPath)
	}
}

func TestConfigURLs(t *testing.T) {
	cfg := &Config{RCAPIPort: 30010, PluginPort: 8090}

	if got := cfg.RCAPIURL(); got != "http://127.0.0.1:30010" {
		t.Errorf("RCAPIURL() = %q", got)
	}
	if got := cfg.PluginURL(); got != "http://127.0.0.1:8090" {
		t.Errorf("PluginURL() = %q", got)
	}
}
