// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package docs

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"
)

func testIngestLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

// --- isMarkdown ---

func TestIsMarkdown(t *testing.T) {
	tests := []struct {
		path string
		want bool
	}{
		{"file.md", true},
		{"file.markdown", true},
		{"dir/sub/readme.md", true},
		{"dir/sub/readme.markdown", true},
		// Extensions are lowercased, so uppercase variants work.
		{"file.MD", true},
		{"file.MARKDOWN", true},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			if got := isMarkdown(tt.path); got != tt.want {
				t.Errorf("isMarkdown(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestIsNotMarkdown(t *testing.T) {
	paths := []string{"file.txt", "file.go", "file.cpp", "file.json", "noext"}
	for _, p := range paths {
		t.Run(p, func(t *testing.T) {
			if isMarkdown(p) {
				t.Errorf("isMarkdown(%q) = true, want false", p)
			}
		})
	}
}

func TestExtractTitle(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    string
	}{
		{"simple heading", "# AActor\nSome content", "AActor"},
		{"heading with spaces", "# A Player Controller\n", "A Player Controller"},
		{"no heading", "Just content here", ""},
		{"heading after blank line", "\n# Delayed Heading\n", "Delayed Heading"},
		{"second-level heading only", "## Not a title\nContent", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractTitle(tt.content)
			if got != tt.want {
				t.Errorf("extractTitle() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestInferCategory(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		content string
		want    string
	}{
		{"actor from path", "/docs/actor/AActor.md", "", "actor"},
		{"blueprint from content", "/docs/general.md", "This is about Blueprint visual scripting", "blueprint"},
		{"realtimemesh from path", "/docs/realtimemesh/URealtimeMesh.md", "", "realtimemesh"},
		{"general fallback", "/docs/other.md", "Nothing special here at all", "general"},
		{"input from content", "/docs/system.md", "Enhanced Input setup and action mapping", "input"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := inferCategory(tt.path, tt.content)
			if got != tt.want {
				t.Errorf("inferCategory() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExtractClassNames(t *testing.T) {
	content := "The AActor class extends UObject. Use FVector for positions. UPROPERTY is a macro."
	classes := extractClassNames(content)

	found := make(map[string]bool)
	for _, c := range classes {
		found[c] = true
	}

	if !found["AActor"] {
		t.Error("expected AActor in class names")
	}
	if !found["UObject"] {
		t.Error("expected UObject in class names")
	}
	if !found["FVector"] {
		t.Error("expected FVector in class names")
	}
	// UPROPERTY should be filtered as false positive.
	if found["UPROPERTY"] {
		t.Error("UPROPERTY should be filtered out")
	}
}

func TestIsLikelyClassName(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"AActor", true},
		{"UObject", true},
		{"FVector", true},
		{"ECollisionChannel", true},
		{"UPROPERTY", false},
		{"UCLASS", false},
		{"UINT32", false},
		{"ENGINE", false},
		{"AB", false},  // too short
		{"ABC", false}, // too short
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isLikelyClassName(tt.name)
			if got != tt.want {
				t.Errorf("isLikelyClassName(%q) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestExtractClassNames_Deduplication(t *testing.T) {
	content := "AActor is great. AActor is used everywhere. AActor again."
	classes := extractClassNames(content)
	if len(classes) != 1 {
		t.Errorf("expected 1 deduplicated class, got %d: %v", len(classes), classes)
	}
}

func TestExtractClassNames_NoneFound(t *testing.T) {
	content := "This content has no class names at all, just plain words."
	classes := extractClassNames(content)
	if len(classes) != 0 {
		t.Errorf("expected 0 classes, got %d: %v", len(classes), classes)
	}
}

func TestInferCategory_ContentSampleLimit(t *testing.T) {
	// Keyword appears well past the 500-char sample limit.
	longPrefix := make([]byte, 600)
	for i := range longPrefix {
		longPrefix[i] = 'x'
	}
	content := string(longPrefix) + " actor spawning"
	got := inferCategory("docs/misc/file.md", content)
	if got != "general" {
		t.Errorf("inferCategory with late keyword = %q, want %q", got, "general")
	}
}

func TestExtractTitle_MultipleHeadings(t *testing.T) {
	content := "Some preamble\n# First Title\n# Second Title\nContent"
	got := extractTitle(content)
	if got != "First Title" {
		t.Errorf("extractTitle = %q, want %q", got, "First Title")
	}
}

func TestParseClassDoc(t *testing.T) {
	content := `# AActor

**Parent**: UObject
**Module**: Engine

AActor is the base class for all actors.

## Key Properties

- ` + "`RootComponent`" + ` — The root scene component
- ` + "`bReplicates`" + ` — Whether this actor replicates

## Key Functions

- ` + "`BeginPlay()`" + ` — Called when the game starts
- ` + "`Tick(float DeltaTime)`" + ` — Called every frame
`

	info := ParseClassDoc("AActor", content)

	if info.Name != "AActor" {
		t.Errorf("Name = %q, want AActor", info.Name)
	}
	if info.Parent != "UObject" {
		t.Errorf("Parent = %q, want UObject", info.Parent)
	}
	if info.Module != "Engine" {
		t.Errorf("Module = %q, want Engine", info.Module)
	}
	if info.Description == "" {
		t.Error("expected non-empty description")
	}
	if len(info.KeyProps) != 2 {
		t.Errorf("expected 2 properties, got %d", len(info.KeyProps))
	}
	if len(info.KeyFuncs) != 2 {
		t.Errorf("expected 2 functions, got %d", len(info.KeyFuncs))
	}
}

// --- parseMarkdownDoc ---

func TestParseMarkdownDoc(t *testing.T) {
	content := "# AActor Reference\n\nAActor is the base class for all actors in UE.\nUse FVector for positions.\n"
	entry := parseMarkdownDoc("/docs/ue5.7/actors/aactor.md", content, "ue5.7")

	if entry.Title != "AActor Reference" {
		t.Errorf("Title = %q, want %q", entry.Title, "AActor Reference")
	}
	if entry.Source != "ue5.7" {
		t.Errorf("Source = %q, want %q", entry.Source, "ue5.7")
	}
	if entry.Category != "actor" {
		t.Errorf("Category = %q, want %q", entry.Category, "actor")
	}
	if entry.Content != content {
		t.Error("Content mismatch")
	}

	// Check classes extracted.
	foundAActor := false
	foundFVector := false
	for _, c := range entry.Classes {
		if c == "AActor" {
			foundAActor = true
		}
		if c == "FVector" {
			foundFVector = true
		}
	}
	if !foundAActor {
		t.Error("expected AActor in classes")
	}
	if !foundFVector {
		t.Error("expected FVector in classes")
	}

	// ID stability: same path should produce same ID.
	entry2 := parseMarkdownDoc("/docs/ue5.7/actors/aactor.md", content, "ue5.7")
	if entry.ID != entry2.ID {
		t.Errorf("ID not stable: %q vs %q", entry.ID, entry2.ID)
	}

	// Different path produces different ID.
	entry3 := parseMarkdownDoc("/docs/ue5.7/actors/other.md", content, "ue5.7")
	if entry.ID == entry3.ID {
		t.Error("different paths produced same ID")
	}
}

func TestParseMarkdownDoc_NoTitle(t *testing.T) {
	content := "No heading here, just content."
	entry := parseMarkdownDoc("/docs/myfile.md", content, "test")

	// Should fall back to filename without extension.
	if entry.Title != "myfile" {
		t.Errorf("Title fallback = %q, want %q", entry.Title, "myfile")
	}
}

// --- IngestFile ---

func TestIngestFile_Success(t *testing.T) {
	idx := createTestIndex(t)

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "test_doc.md")
	if err := os.WriteFile(mdPath, []byte("# Test Document\n\nAActor spawning guide."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	if err := IngestFile(idx, mdPath, "test"); err != nil {
		t.Fatalf("IngestFile: %v", err)
	}

	count, _ := idx.DocCount()
	if count != 1 {
		t.Errorf("DocCount = %d, want 1", count)
	}
}

func TestIngestFile_NonExistent(t *testing.T) {
	idx := createTestIndex(t)

	err := IngestFile(idx, "/nonexistent/path/doc.md", "test")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestIngestFile_EmptyContent(t *testing.T) {
	idx := createTestIndex(t)

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "empty.md")
	if err := os.WriteFile(mdPath, []byte(""), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	if err := IngestFile(idx, mdPath, "test"); err != nil {
		t.Fatalf("IngestFile should not error on empty file: %v", err)
	}

	count, _ := idx.DocCount()
	if count != 0 {
		t.Errorf("DocCount = %d, want 0 for empty file", count)
	}
}

func TestIngestFile_WhitespaceOnly(t *testing.T) {
	idx := createTestIndex(t)

	dir := t.TempDir()
	mdPath := filepath.Join(dir, "whitespace.md")
	if err := os.WriteFile(mdPath, []byte("   \n  \n  "), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	if err := IngestFile(idx, mdPath, "test"); err != nil {
		t.Fatalf("IngestFile should not error on whitespace-only file: %v", err)
	}

	count, _ := idx.DocCount()
	if count != 0 {
		t.Errorf("DocCount = %d, want 0 for whitespace-only file", count)
	}
}

// --- IngestDirectory ---

func TestIngestDirectory_Success(t *testing.T) {
	idx := createTestIndex(t)
	logger := testIngestLogger()

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "doc1.md"), []byte("# Doc One\n\nFirst document about actors."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "doc2.md"), []byte("# Doc Two\n\nSecond document about blueprints."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "notes.txt"), []byte("This is a text file, not markdown."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	count, err := IngestDirectory(idx, dir, "test", logger)
	if err != nil {
		t.Fatalf("IngestDirectory: %v", err)
	}
	if count != 2 {
		t.Errorf("IngestDirectory returned count=%d, want 2", count)
	}

	docCount, _ := idx.DocCount()
	if docCount != 2 {
		t.Errorf("DocCount = %d, want 2", docCount)
	}
}

func TestIngestDirectory_SkipsNonMarkdown(t *testing.T) {
	idx := createTestIndex(t)
	logger := testIngestLogger()

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("Not markdown"), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "code.go"), []byte("package main"), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "data.json"), []byte(`{"key":"value"}`), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	count, err := IngestDirectory(idx, dir, "test", logger)
	if err != nil {
		t.Fatalf("IngestDirectory: %v", err)
	}
	if count != 0 {
		t.Errorf("IngestDirectory indexed %d non-markdown files, want 0", count)
	}
}

func TestIngestDirectory_SkipsEmptyMarkdown(t *testing.T) {
	idx := createTestIndex(t)
	logger := testIngestLogger()

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "empty.md"), []byte(""), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "content.md"), []byte("# Has Content\nSome real content here."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	count, err := IngestDirectory(idx, dir, "test", logger)
	if err != nil {
		t.Fatalf("IngestDirectory: %v", err)
	}
	if count != 1 {
		t.Errorf("IngestDirectory returned count=%d, want 1 (skip empty)", count)
	}
}

func TestIngestDirectory_Recursive(t *testing.T) {
	idx := createTestIndex(t)
	logger := testIngestLogger()

	dir := t.TempDir()
	subdir := filepath.Join(dir, "subdir")
	if err := os.Mkdir(subdir, 0750); err != nil {
		t.Fatalf("Mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "top.md"), []byte("# Top Level\nContent."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	if err := os.WriteFile(filepath.Join(subdir, "nested.md"), []byte("# Nested Doc\nMore content."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	count, err := IngestDirectory(idx, dir, "test", logger)
	if err != nil {
		t.Fatalf("IngestDirectory: %v", err)
	}
	if count != 2 {
		t.Errorf("IngestDirectory returned count=%d, want 2 (recursive)", count)
	}
}

func TestIngestDirectory_NonExistent(t *testing.T) {
	idx := createTestIndex(t)
	logger := testIngestLogger()

	_, err := IngestDirectory(idx, "/nonexistent/directory", "test", logger)
	if err == nil {
		t.Error("expected error for nonexistent directory")
	}
}

func TestIngestDirectory_SkipsMetaFiles(t *testing.T) {
	idx := createTestIndex(t)
	logger := testIngestLogger()

	dir := t.TempDir()
	// README.md should be skipped.
	if err := os.WriteFile(filepath.Join(dir, "README.md"), []byte("# How to add docs\nMeta info."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	// CHANGELOG.md should be skipped.
	if err := os.WriteFile(filepath.Join(dir, "CHANGELOG.md"), []byte("# Changelog\nChanges."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
	// Actual doc should be indexed.
	if err := os.WriteFile(filepath.Join(dir, "AActor.md"), []byte("# AActor\nActor docs."), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	count, err := IngestDirectory(idx, dir, "test", logger)
	if err != nil {
		t.Fatalf("IngestDirectory: %v", err)
	}
	if count != 1 {
		t.Errorf("IngestDirectory returned count=%d, want 1 (skip meta files)", count)
	}
}
