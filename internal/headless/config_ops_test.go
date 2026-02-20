// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package headless

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/remiphilippe/mcp-unreal/internal/config"
)

// createTestConfig sets up a temp project with a Config/ directory and returns
// a Handler pointing at it.
func createTestConfig(t *testing.T, fileName, content string) *Handler {
	t.Helper()
	dir := t.TempDir()
	configDir := filepath.Join(dir, "Config")
	if err := os.MkdirAll(configDir, 0o750); err != nil { //nolint:gosec // test helper
		t.Fatal(err)
	}
	if content != "" {
		if err := os.WriteFile(filepath.Join(configDir, fileName+".ini"), []byte(content), 0o600); err != nil { //nolint:gosec // test helper
			t.Fatal(err)
		}
	}
	return &Handler{
		Config: &config.Config{ProjectRoot: dir},
	}
}

const sampleINI = `[/Script/Engine.RendererSettings]
r.DefaultFeature.AutoExposure=False
r.DefaultFeature.MotionBlur=False

[Pelorus.BaseMap]
DemGridSize=2048
UseEnc=true
EncMetersToUnreal=100.0
+EncDatasetPath=US3CA52M.zip
+EncDatasetPath=US5OAKFG.zip

[Pelorus.UI]
CoordinateFormat=DD
DistanceUnit=NM
`

func TestConfigOps_Get(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "DemGridSize",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Value != "2048" {
		t.Errorf("expected 2048, got %q", out.Value)
	}
}

func TestConfigOps_Get_RendererSettings(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "/Script/Engine.RendererSettings",
		Key:       "r.DefaultFeature.AutoExposure",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Value != "False" {
		t.Errorf("expected False, got %q", out.Value)
	}
}

func TestConfigOps_Get_MissingKey(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "NonExistentKey",
	})
	if err == nil {
		t.Fatal("expected error for missing key")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("expected 'not found' in error, got: %v", err)
	}
}

func TestConfigOps_Get_MissingSection(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "NonExistent.Section",
		Key:       "SomeKey",
	})
	if err == nil {
		t.Fatal("expected error for missing section")
	}
}

func TestConfigOps_Set_ExistingKey(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "set",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "DemGridSize",
		Value:     "4096",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}

	// Verify the value was persisted.
	_, out2, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "DemGridSize",
	})
	if err != nil {
		t.Fatalf("unexpected error on re-read: %v", err)
	}
	if out2.Value != "4096" {
		t.Errorf("expected 4096 after set, got %q", out2.Value)
	}
}

func TestConfigOps_Set_NewKey(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "set",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "NewSetting",
		Value:     "hello",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "NewSetting",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Value != "hello" {
		t.Errorf("expected hello, got %q", out.Value)
	}
}

func TestConfigOps_Set_NewSection(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "set",
		File:      "DefaultEngine",
		Section:   "MyPlugin.Settings",
		Key:       "Enabled",
		Value:     "true",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "MyPlugin.Settings",
		Key:       "Enabled",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Value != "true" {
		t.Errorf("expected true, got %q", out.Value)
	}
}

func TestConfigOps_Delete(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "delete",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "UseEnc",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify it's gone.
	_, _, err = h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "UseEnc",
	})
	if err == nil {
		t.Fatal("expected error after delete, key should be gone")
	}
}

func TestConfigOps_Delete_NonExistent(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "delete",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "GhostKey",
	})
	if err == nil {
		t.Fatal("expected error deleting non-existent key")
	}
}

func TestConfigOps_List(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "list",
		File:      "DefaultEngine",
		Section:   "Pelorus.UI",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if len(out.Values) != 2 {
		t.Errorf("expected 2 keys in Pelorus.UI, got %d: %v", len(out.Values), out.Values)
	}
	if out.Values["CoordinateFormat"] != "DD" {
		t.Errorf("expected DD, got %q", out.Values["CoordinateFormat"])
	}
}

func TestConfigOps_ListSections(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "list_sections",
		File:      "DefaultEngine",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Sections) != 3 {
		t.Errorf("expected 3 sections, got %d: %v", len(out.Sections), out.Sections)
	}
}

func TestConfigOps_ArrayValues(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	// UE uses +Key=Value for array append. The parser keeps the last value
	// for the same key, but the +prefix is preserved.
	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultEngine",
		Section:   "Pelorus.BaseMap",
		Key:       "+EncDatasetPath",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Last +EncDatasetPath wins in the map.
	if out.Value != "US5OAKFG.zip" {
		t.Errorf("expected US5OAKFG.zip (last array value), got %q", out.Value)
	}
}

func TestConfigOps_MissingFile(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultNonExistent",
		Section:   "Foo",
		Key:       "Bar",
	})
	if err == nil {
		t.Fatal("expected error for missing file")
	}
	if !strings.Contains(err.Error(), "not found") {
		t.Errorf("expected 'not found' in error, got: %v", err)
	}
}

func TestConfigOps_PathTraversal(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "../../etc/passwd",
		Section:   "Foo",
		Key:       "Bar",
	})
	if err == nil {
		t.Fatal("expected error for path traversal")
	}
	if !strings.Contains(err.Error(), "invalid") {
		t.Errorf("expected 'invalid' in error, got: %v", err)
	}
}

func TestConfigOps_NoProjectRoot(t *testing.T) {
	h := &Handler{Config: &config.Config{ProjectRoot: ""}}
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "list_sections",
		File:      "DefaultEngine",
	})
	if err == nil {
		t.Fatal("expected error when no project root")
	}
	if !strings.Contains(err.Error(), "no UE project root") {
		t.Errorf("expected project root error, got: %v", err)
	}
}

func TestConfigOps_PreservesComments(t *testing.T) {
	ini := `; This is a comment
[Section1]
; Another comment
Key1=Value1
Key2=Value2

[Section2]
Key3=Value3
`
	h := createTestConfig(t, "DefaultEngine", ini)
	ctx := context.Background()

	// Set Key1 to a new value.
	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "set",
		File:      "DefaultEngine",
		Section:   "Section1",
		Key:       "Key1",
		Value:     "NewValue",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Read the file back and verify comments are preserved.
	iniPath := filepath.Join(h.Config.ProjectRoot, "Config", "DefaultEngine.ini")
	data, err := os.ReadFile(iniPath) //nolint:gosec // test reads known temp path
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, "; This is a comment") {
		t.Error("top-level comment was lost")
	}
	if !strings.Contains(content, "; Another comment") {
		t.Error("section comment was lost")
	}
	if !strings.Contains(content, "Key1=NewValue") {
		t.Error("Key1 was not updated")
	}
	if !strings.Contains(content, "Key2=Value2") {
		t.Error("Key2 was lost during edit")
	}
}

func TestConfigOps_Set_CreatesFile(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	// Set a value in a file that doesn't exist yet.
	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "set",
		File:      "DefaultGame",
		Section:   "MyGame.Settings",
		Key:       "Difficulty",
		Value:     "Hard",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the file was created and is readable.
	_, out, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "get",
		File:      "DefaultGame",
		Section:   "MyGame.Settings",
		Key:       "Difficulty",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Value != "Hard" {
		t.Errorf("expected Hard, got %q", out.Value)
	}
}

func TestConfigOps_InvalidOperation(t *testing.T) {
	h := createTestConfig(t, "DefaultEngine", sampleINI)
	ctx := context.Background()

	_, _, err := h.ConfigOps(ctx, nil, ConfigOpsInput{
		Operation: "explode",
		File:      "DefaultEngine",
	})
	if err == nil {
		t.Fatal("expected error for invalid operation")
	}
}
