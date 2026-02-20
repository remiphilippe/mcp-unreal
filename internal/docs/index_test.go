package docs

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTestIndex(t *testing.T) *Index {
	t.Helper()
	dir := t.TempDir()
	idx, err := CreateIndex(filepath.Join(dir, "test.bleve"))
	if err != nil {
		t.Fatalf("CreateIndex: %v", err)
	}
	t.Cleanup(func() { _ = idx.Close() })
	return idx
}

func TestCreateAndOpenIndex(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.bleve")

	// Create.
	idx, err := CreateIndex(path)
	if err != nil {
		t.Fatalf("CreateIndex: %v", err)
	}
	if err := idx.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	// Reopen.
	idx2, err := OpenIndex(path)
	if err != nil {
		t.Fatalf("OpenIndex: %v", err)
	}
	defer func() { _ = idx2.Close() }()

	count, err := idx2.DocCount()
	if err != nil {
		t.Fatalf("DocCount: %v", err)
	}
	if count != 0 {
		t.Errorf("expected 0 docs, got %d", count)
	}
}

func TestIndexAndSearch(t *testing.T) {
	idx := createTestIndex(t)

	// Index some test docs.
	entries := []DocEntry{
		{
			ID:       "actor-1",
			Title:    "AActor Class Reference",
			Category: "actor",
			Source:   "ue5.7",
			Content:  "AActor is the base class for all actors in UE. Actors can be spawned into levels.",
			Classes:  []string{"AActor"},
		},
		{
			ID:       "blueprint-1",
			Title:    "Blueprint Visual Scripting",
			Category: "blueprint",
			Source:   "ue5.7",
			Content:  "Blueprints are visual scripting graphs that allow you to create gameplay logic without C++.",
			Classes:  []string{"UBlueprintGeneratedClass"},
		},
		{
			ID:       "mesh-1",
			Title:    "URealtimeMeshSimple",
			Category: "realtimemesh",
			Source:   "realtimemesh",
			Content:  "URealtimeMeshSimple provides a simplified API for creating runtime meshes with LOD support.",
			Classes:  []string{"URealtimeMeshSimple", "URealtimeMeshComponent"},
		},
	}

	if err := idx.IndexBatch(entries); err != nil {
		t.Fatalf("IndexBatch: %v", err)
	}

	count, _ := idx.DocCount()
	if count != 3 {
		t.Fatalf("expected 3 docs, got %d", count)
	}

	// Test lookup_docs: general query.
	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "actor spawn",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Error("expected results for 'actor spawn'")
	}
	if out.Results[0].Title != "AActor Class Reference" {
		t.Errorf("expected AActor result first, got %q", out.Results[0].Title)
	}

	// Test lookup_docs: category filter.
	_, out, err = idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query:    "runtime mesh",
		Category: "realtimemesh",
	})
	if err != nil {
		t.Fatalf("LookupDocs with category: %v", err)
	}
	if out.Total == 0 {
		t.Error("expected results for 'runtime mesh' with realtimemesh category")
	}

	// Test lookup_class: by class name.
	_, classOut, err := idx.LookupClass(context.Background(), nil, LookupClassInput{
		ClassName: "AActor",
	})
	if err != nil {
		t.Fatalf("LookupClass: %v", err)
	}
	if !classOut.Found {
		t.Error("expected AActor to be found")
	}
	if classOut.Class.Name != "AActor" {
		t.Errorf("class name = %q, want AActor", classOut.Class.Name)
	}

	// Test lookup_class: not found.
	_, classOut, err = idx.LookupClass(context.Background(), nil, LookupClassInput{
		ClassName: "UNonExistentClass",
	})
	if err != nil {
		t.Fatalf("LookupClass not found: %v", err)
	}
	if classOut.Found {
		t.Error("expected UNonExistentClass to not be found")
	}
}

func TestLookupDocs_TokenBudget(t *testing.T) {
	idx := createTestIndex(t)

	// Index a doc with large content.
	longContent := ""
	for i := 0; i < 200; i++ {
		longContent += "This is a line of documentation about actors and spawning. "
	}

	if err := idx.IndexDoc(DocEntry{
		ID:       "big-doc",
		Title:    "Big Actor Doc",
		Category: "actor",
		Source:   "ue5.7",
		Content:  longContent,
		Classes:  []string{"AActor"},
	}); err != nil {
		t.Fatalf("IndexDoc: %v", err)
	}

	// Query with tiny token budget.
	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query:     "actor",
		MaxTokens: 50,
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}

	// Result should be truncated.
	if out.Total > 0 && len(out.Results[0].Snippet) > 250 {
		t.Errorf("snippet too long for 50 token budget: %d chars", len(out.Results[0].Snippet))
	}
}

func TestLookupDocs_EmptyQuery(t *testing.T) {
	idx := createTestIndex(t)
	_, _, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "",
	})
	if err == nil {
		t.Error("expected error for empty query")
	}
}

func TestOpenOrCreate_NewIndex(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "new.bleve")
	idx, err := OpenOrCreate(path)
	if err != nil {
		t.Fatalf("OpenOrCreate: %v", err)
	}
	defer func() { _ = idx.Close() }()
	count, _ := idx.DocCount()
	if count != 0 {
		t.Errorf("expected 0 docs in new index, got %d", count)
	}
}

func TestOpenOrCreate_ExistingIndex(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "existing.bleve")

	// Create and populate.
	idx, err := CreateIndex(path)
	if err != nil {
		t.Fatalf("CreateIndex: %v", err)
	}
	if err := idx.IndexDoc(DocEntry{ID: "test-1", Title: "Test Doc", Content: "test content", Source: "test"}); err != nil {
		t.Fatalf("IndexDoc: %v", err)
	}
	if err := idx.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	// Reopen via OpenOrCreate.
	idx2, err := OpenOrCreate(path)
	if err != nil {
		t.Fatalf("OpenOrCreate: %v", err)
	}
	defer func() { _ = idx2.Close() }()
	count, _ := idx2.DocCount()
	if count != 1 {
		t.Errorf("expected 1 doc in existing index, got %d", count)
	}
}

func TestLookupClass_EmptyClassName(t *testing.T) {
	idx := createTestIndex(t)
	_, _, err := idx.LookupClass(context.Background(), nil, LookupClassInput{
		ClassName: "",
	})
	if err == nil {
		t.Error("expected error for empty class name")
	}
}

// ---------------------------------------------------------------------------
// Content verification tests — uses realistic doc entries to verify
// that bleve returns sensible results for common MCP tool queries.
// ---------------------------------------------------------------------------

func createPopulatedIndex(t *testing.T) *Index {
	t.Helper()
	idx := createTestIndex(t)

	entries := []DocEntry{
		{
			ID:       "aactor",
			Title:    "AActor",
			Category: "actor",
			Source:   "ue5.7",
			Content: "AActor is the base class for all actors that can be placed or spawned in a level. " +
				"Actors support 3D transformations (location, rotation, scale), component attachment hierarchies, " +
				"replication for networking, and lifecycle events. " +
				"Key functions: BeginPlay(), Tick(), SetActorLocation(), SetActorRotation(), SetActorScale3D(), Destroy().",
			Classes: []string{"AActor"},
		},
		{
			ID:       "umaterial",
			Title:    "UMaterial",
			Category: "material",
			Source:   "ue5.7",
			Content: "Material asset defining the visual appearance of surfaces in Unreal Engine. " +
				"Controls how light interacts with a surface through shading models, blend modes, and material domains. " +
				"Key properties: ShadingModel, BlendMode, MaterialDomain, OpacityMaskClipValue, bTwoSided.",
			Classes: []string{"UMaterial", "UMaterialInterface"},
		},
		{
			ID:       "upcg-component",
			Title:    "UPCGComponent",
			Category: "gameplay",
			Source:   "ue5.7",
			Content: "Procedural Content Generation component that runs PCG graphs on actors. " +
				"Attach this component to any actor to execute a UPCGGraph in the context of that actor's transform. " +
				"Returns array of AActor* instances spawned by generation. " +
				"Key properties: Graph, GenerationTrigger, Seed, InputType.",
			Classes: []string{"UPCGComponent", "UPCGGraph", "AActor"},
		},
		{
			ID:       "ulevel",
			Title:    "ULevel",
			Category: "actor",
			Source:   "ue5.7",
			Content: "Container for all actors in a map. ULevel holds the persistent level and streaming sublevels. " +
				"Key functions: GetWorldSettings(), GetLevelScriptActor(), OwningWorld.",
			Classes: []string{"ULevel", "AWorldSettings"},
		},
		{
			ID:       "uworld",
			Title:    "UWorld",
			Category: "actor",
			Source:   "ue5.7",
			Content: "Top-level object for a game world. Holds the persistent level and streaming levels. " +
				"SpawnActor creates new actors. GetWorld() is available on most UObject subclasses. " +
				"Key functions: SpawnActor(), GetTimerManager(), GetAuthGameMode().",
			Classes: []string{"UWorld"},
		},
		{
			ID:       "blueprint-guide",
			Title:    "Blueprint Visual Scripting",
			Category: "blueprint",
			Source:   "ue5.7",
			Content: "Blueprints are visual scripting graphs that allow you to create gameplay logic without C++. " +
				"Event Graph handles gameplay events. Construction Script runs at spawn time. " +
				"Blueprint variables can be exposed to the editor.",
			Classes: []string{"UBlueprintGeneratedClass", "UBlueprint"},
		},
		{
			ID:       "niagara-system",
			Title:    "UNiagaraSystem",
			Category: "rendering",
			Source:   "ue5.7",
			Content: "Niagara particle system asset containing one or more emitters. " +
				"Controls emitter lifecycle, spawn rates, and system-level parameters. " +
				"Use UNiagaraComponent to add to actors.",
			Classes: []string{"UNiagaraSystem", "UNiagaraComponent", "UNiagaraEmitter"},
		},
		{
			ID:       "anim-instance",
			Title:    "UAnimInstance",
			Category: "animation",
			Source:   "ue5.7",
			Content: "Animation instance running on a skeletal mesh component. " +
				"Manages blend spaces, montages, and state machine transitions. " +
				"NativeUpdateAnimation() is called every frame for animation logic.",
			Classes: []string{"UAnimInstance", "UAnimMontage"},
		},
		{
			ID:       "gas-ability",
			Title:    "UGameplayAbility",
			Category: "gameplay",
			Source:   "ue5.7",
			Content: "Base class for gameplay abilities in the Gameplay Ability System (GAS). " +
				"Abilities are activated via UAbilitySystemComponent. " +
				"Override ActivateAbility() and EndAbility() for custom logic. " +
				"Gameplay effects modify attributes through UGameplayEffect.",
			Classes: []string{"UGameplayAbility", "UAbilitySystemComponent", "UGameplayEffect"},
		},
		{
			ID:       "enhanced-input",
			Title:    "Enhanced Input System",
			Category: "input",
			Source:   "ue5.7",
			Content: "Enhanced Input provides data-driven input handling with triggers and modifiers. " +
				"Input Actions define abstract actions. Input Mapping Contexts bind actions to physical inputs. " +
				"Supports dead zones, hold triggers, and chorded actions.",
			Classes: []string{"UInputAction", "UInputMappingContext"},
		},
	}

	if err := idx.IndexBatch(entries); err != nil {
		t.Fatalf("IndexBatch: %v", err)
	}
	return idx
}

func TestBleveContent_SpawnActorQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "how to spawn an actor in the level",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for 'spawn actor' query")
	}
	// Top result should mention AActor or UWorld (both reference spawning)
	topSnippet := out.Results[0].Snippet
	if !containsAny(topSnippet, "AActor", "spawn", "SpawnActor") {
		t.Errorf("top result for 'spawn actor' should reference actor spawning, got: %.100s", topSnippet)
	}
}

func TestBleveContent_MaterialQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "material shading model blend mode",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for material query")
	}
	found := false
	for _, r := range out.Results {
		if r.Title == "UMaterial" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected UMaterial in results for 'material shading model blend mode'")
	}
}

func TestBleveContent_PCGQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "procedural content generation graph",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for PCG query")
	}
	topSnippet := out.Results[0].Snippet
	if !containsAny(topSnippet, "PCG", "Procedural Content Generation", "UPCGGraph") {
		t.Errorf("top result should reference PCG, got: %.100s", topSnippet)
	}
}

func TestBleveContent_BlueprintQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "blueprint visual scripting event graph",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for blueprint query")
	}
	found := false
	for _, r := range out.Results {
		if r.Category == "blueprint" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected a blueprint category result")
	}
}

func TestBleveContent_NiagaraQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "niagara particle emitter system",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for niagara query")
	}
	topSnippet := out.Results[0].Snippet
	if !containsAny(topSnippet, "Niagara", "particle", "emitter") {
		t.Errorf("top result should reference Niagara, got: %.100s", topSnippet)
	}
}

func TestBleveContent_GASQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "gameplay ability system activate",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for GAS query")
	}
	topSnippet := out.Results[0].Snippet
	if !containsAny(topSnippet, "ability", "Gameplay Ability", "ActivateAbility") {
		t.Errorf("top result should reference GAS, got: %.100s", topSnippet)
	}
}

func TestBleveContent_InputQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "enhanced input action mapping",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for input query")
	}
	found := false
	for _, r := range out.Results {
		if r.Category == "input" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected an input category result")
	}
}

func TestBleveContent_AnimationQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "animation montage blend state machine",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for animation query")
	}
	topSnippet := out.Results[0].Snippet
	if !containsAny(topSnippet, "animation", "montage", "state machine", "AnimInstance") {
		t.Errorf("top result should reference animation, got: %.100s", topSnippet)
	}
}

func TestBleveContent_LevelQuery(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "level streaming sublevel persistent",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for level query")
	}
	topSnippet := out.Results[0].Snippet
	if !containsAny(topSnippet, "level", "Level", "streaming", "persistent") {
		t.Errorf("top result should reference levels, got: %.100s", topSnippet)
	}
}

func TestBleveContent_ClassLookup_AActor(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupClass(context.Background(), nil, LookupClassInput{
		ClassName: "AActor",
	})
	if err != nil {
		t.Fatalf("LookupClass: %v", err)
	}
	if !out.Found {
		t.Fatal("expected AActor to be found")
	}
	if out.Class.Name != "AActor" {
		t.Errorf("expected name AActor, got %s", out.Class.Name)
	}
	if out.Class.Description == "" {
		t.Error("expected non-empty description for AActor")
	}
	if !containsAny(out.Class.Description, "actor", "base class", "spawned") {
		t.Errorf("AActor description should mention actors: %s", out.Class.Description)
	}
	// Regression: AActor lookup must return the AActor doc, not UPCGComponent
	// or another doc that merely references AActor in its classes list.
	if containsAny(out.Class.Description, "Procedural Content Generation", "PCG") {
		t.Errorf("AActor lookup returned wrong doc (UPCGComponent?): %s", out.Class.Description)
	}
}

func TestBleveContent_ClassLookup_PrefersExactTitle(t *testing.T) {
	// Regression test: when multiple docs reference "AActor" in their classes
	// field, lookup_class must prefer the doc whose title matches the class name.
	idx := createPopulatedIndex(t)

	classes := []struct {
		name       string
		wantInDesc string
		dontWant   string
	}{
		{"AActor", "base class", "PCG"},
		{"UMaterial", "material", "PCG"},
		{"UPCGComponent", "PCG", "material shading"},
		{"UNiagaraSystem", "particle", "material"},
		{"UGameplayAbility", "ability", "material"},
		{"UAnimInstance", "animation", "material"},
	}

	for _, tc := range classes {
		t.Run(tc.name, func(t *testing.T) {
			_, out, err := idx.LookupClass(context.Background(), nil, LookupClassInput{
				ClassName: tc.name,
			})
			if err != nil {
				t.Fatalf("LookupClass(%s): %v", tc.name, err)
			}
			if !out.Found {
				t.Fatalf("%s not found", tc.name)
			}
			if !containsAny(out.Class.Description, tc.wantInDesc) {
				t.Errorf("%s description should mention %q, got: %.100s", tc.name, tc.wantInDesc, out.Class.Description)
			}
			if tc.dontWant != "" && containsAny(out.Class.Description, tc.dontWant) {
				t.Errorf("%s description should NOT mention %q — wrong doc returned: %.100s", tc.name, tc.dontWant, out.Class.Description)
			}
		})
	}
}

func TestBleveContent_ClassLookup_UNiagaraSystem(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupClass(context.Background(), nil, LookupClassInput{
		ClassName: "UNiagaraSystem",
	})
	if err != nil {
		t.Fatalf("LookupClass: %v", err)
	}
	if !out.Found {
		t.Fatal("expected UNiagaraSystem to be found")
	}
	if out.Class.Name != "UNiagaraSystem" {
		t.Errorf("expected name UNiagaraSystem, got %s", out.Class.Name)
	}
}

func TestBleveContent_ClassLookup_UGameplayAbility(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupClass(context.Background(), nil, LookupClassInput{
		ClassName: "UGameplayAbility",
	})
	if err != nil {
		t.Fatalf("LookupClass: %v", err)
	}
	if !out.Found {
		t.Fatal("expected UGameplayAbility to be found")
	}
}

func TestBleveContent_CategoryFilter(t *testing.T) {
	idx := createPopulatedIndex(t)

	// Query for "component" but filter to only gameplay category
	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query:    "component graph",
		Category: "gameplay",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	// All results should be in the gameplay category
	for _, r := range out.Results {
		if r.Category != "gameplay" {
			t.Errorf("category filter should restrict to gameplay, got %s for %s", r.Category, r.Title)
		}
	}
}

func TestBleveContent_MultipleResults(t *testing.T) {
	idx := createPopulatedIndex(t)

	// Generic query that should match multiple docs
	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "component actor",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total < 2 {
		t.Errorf("expected multiple results for 'component actor', got %d", out.Total)
	}
}

func TestBleveContent_ResultsHaveSource(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "actor",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	for _, r := range out.Results {
		if r.Source == "" {
			t.Errorf("result %q has empty source", r.Title)
		}
	}
}

func TestBleveContent_SnippetHasContent(t *testing.T) {
	idx := createPopulatedIndex(t)

	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "material",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Fatal("expected results for 'material'")
	}
	if len(out.Results[0].Snippet) < 20 {
		t.Errorf("snippet too short: %q", out.Results[0].Snippet)
	}
}

func TestBleveContent_IngestDirectory(t *testing.T) {
	idx := createTestIndex(t)

	// Create a temp directory with test markdown files.
	dir := t.TempDir()
	writeTestDoc(t, dir, "TestActor.md", "# TestActor\n\n**Parent**: AActor\n**Module**: Engine\n\nA test actor for unit tests.\n\n## Key Functions\n\n- `DoSomething()` — Does something\n")
	writeTestDoc(t, dir, "TestComponent.md", "# TestComponent\n\n**Parent**: UActorComponent\n**Module**: Engine\n\nA test component.\n")
	writeTestDoc(t, dir, "README.md", "# README\n\nThis should be skipped.\n")

	count, err := IngestDirectory(idx, dir, "test", testSlogLogger())
	if err != nil {
		t.Fatalf("IngestDirectory: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 indexed (README skipped), got %d", count)
	}

	docCount, _ := idx.DocCount()
	if docCount != 2 {
		t.Errorf("expected 2 docs in index, got %d", docCount)
	}

	// Verify content is searchable
	_, out, err := idx.LookupDocs(context.Background(), nil, LookupDocsInput{
		Query: "test actor unit tests",
	})
	if err != nil {
		t.Fatalf("LookupDocs: %v", err)
	}
	if out.Total == 0 {
		t.Error("expected ingested docs to be searchable")
	}
}

// --- helpers ---

func containsAny(s string, substrs ...string) bool {
	lower := strings.ToLower(s)
	for _, sub := range substrs {
		if strings.Contains(lower, strings.ToLower(sub)) {
			return true
		}
	}
	return false
}

func writeTestDoc(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
		t.Fatalf("writing test doc %s: %v", name, err)
	}
}

func testSlogLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
