package headless

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/remiphilippe/mcp-unreal/internal/config"
)

// createTestUProject writes a sample .uproject file and returns the path.
func createTestUProject(t *testing.T) (string, *Handler) {
	t.Helper()
	tmpDir := t.TempDir()

	projFile := filepath.Join(tmpDir, "TestProject.uproject")
	content := `{
	"FileVersion": 3,
	"EngineAssociation": "5.7",
	"Category": "",
	"Description": "",
	"Modules": [
		{
			"Name": "TestProject",
			"Type": "Runtime",
			"LoadingPhase": "Default"
		}
	],
	"Plugins": [
		{
			"Name": "RemoteControl",
			"Enabled": true
		},
		{
			"Name": "MCPUnreal",
			"Enabled": true
		}
	]
}`
	if err := os.WriteFile(projFile, []byte(content), 0o600); err != nil { //nolint:gosec // test helper
		t.Fatalf("writing test .uproject: %v", err)
	}

	h := &Handler{
		Config: &config.Config{ProjectRoot: tmpDir},
		Logger: testLogger(),
	}
	return projFile, h
}

func TestProjectOps_GetInfo(t *testing.T) {
	_, h := createTestUProject(t)
	ctx := context.Background()

	_, out, err := h.ProjectOps(ctx, nil, ProjectOpsInput{
		Operation: "get_info",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.ProjectName != "TestProject" {
		t.Errorf("expected TestProject, got %s", out.ProjectName)
	}
	if out.EngineVersion != "5.7" {
		t.Errorf("expected 5.7, got %s", out.EngineVersion)
	}
	if len(out.Modules) != 1 {
		t.Fatalf("expected 1 module, got %d", len(out.Modules))
	}
	if out.Modules[0].Name != "TestProject" {
		t.Errorf("expected TestProject module, got %s", out.Modules[0].Name)
	}
	if len(out.Plugins) != 2 {
		t.Fatalf("expected 2 plugins, got %d", len(out.Plugins))
	}
}

func TestProjectOps_ListPlugins(t *testing.T) {
	_, h := createTestUProject(t)

	_, out, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "list_plugins",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Plugins) != 2 {
		t.Fatalf("expected 2 plugins, got %d", len(out.Plugins))
	}
	if out.Plugins[0].Name != "RemoteControl" {
		t.Errorf("expected RemoteControl, got %s", out.Plugins[0].Name)
	}
}

func TestProjectOps_EnablePlugin(t *testing.T) {
	projFile, h := createTestUProject(t)

	// Enable a new plugin.
	_, out, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "enable_plugin",
		Name:      "RealtimeMeshComponent",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}

	// Verify the file was updated.
	data, _ := os.ReadFile(projFile) //nolint:gosec // test reads known temp path
	if !strings.Contains(string(data), "RealtimeMeshComponent") {
		t.Error("expected RealtimeMeshComponent in .uproject")
	}

	// Verify backup was created.
	bakPath := projFile + ".bak"
	if _, err := os.Stat(bakPath); os.IsNotExist(err) {
		t.Error("expected .uproject.bak backup file")
	}
}

func TestProjectOps_DisablePlugin(t *testing.T) {
	projFile, h := createTestUProject(t)

	_, out, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "disable_plugin",
		Name:      "RemoteControl",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}

	// Verify the plugin is disabled.
	data, _ := os.ReadFile(projFile) //nolint:gosec // test reads known temp path
	var proj struct {
		Plugins []struct {
			Name    string `json:"Name"`
			Enabled bool   `json:"Enabled"`
		} `json:"Plugins"`
	}
	_ = json.Unmarshal(data, &proj)
	for _, p := range proj.Plugins {
		if p.Name == "RemoteControl" && p.Enabled {
			t.Error("expected RemoteControl to be disabled")
		}
	}
}

func TestProjectOps_AddModule(t *testing.T) {
	projFile, h := createTestUProject(t)

	_, out, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "add_module",
		Name:      "TestProjectEditor",
		Type:      "Editor",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}

	// Verify.
	data, _ := os.ReadFile(projFile) //nolint:gosec // test reads known temp path
	if !strings.Contains(string(data), "TestProjectEditor") {
		t.Error("expected TestProjectEditor in .uproject")
	}
}

func TestProjectOps_AddModuleDuplicate(t *testing.T) {
	_, h := createTestUProject(t)

	_, out, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "add_module",
		Name:      "TestProject",
		Type:      "Runtime",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.Message, "already exists") {
		t.Errorf("expected 'already exists' message, got: %s", out.Message)
	}
}

func TestProjectOps_SetTargetPlatforms(t *testing.T) {
	projFile, h := createTestUProject(t)

	_, out, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "set_target_platforms",
		Platforms: []string{"Mac", "Win64"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}

	data, _ := os.ReadFile(projFile) //nolint:gosec // test reads known temp path
	if !strings.Contains(string(data), "Mac") {
		t.Error("expected Mac in TargetPlatforms")
	}
}

func TestProjectOps_MissingOperation(t *testing.T) {
	_, h := createTestUProject(t)
	_, _, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required', got: %v", err)
	}
}

func TestProjectOps_NoProjectRoot(t *testing.T) {
	h := &Handler{
		Config: &config.Config{ProjectRoot: ""},
		Logger: testLogger(),
	}
	_, _, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "get_info",
	})
	if err == nil {
		t.Fatal("expected error for missing project root")
	}
}

func TestProjectOps_EnablePluginMissingName(t *testing.T) {
	_, h := createTestUProject(t)
	_, _, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "enable_plugin",
	})
	if err == nil {
		t.Fatal("expected error for missing name")
	}
}

func TestProjectOps_UnknownOperation(t *testing.T) {
	_, h := createTestUProject(t)
	_, _, err := h.ProjectOps(context.Background(), nil, ProjectOpsInput{
		Operation: "bogus",
	})
	if err == nil {
		t.Fatal("expected error for unknown operation")
	}
}
