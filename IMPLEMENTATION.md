# mcp-unreal — Complete Autonomous UE 5.7 Development Server

## Single Go Binary. Zero External Dependencies. Full Autonomy.

This is the definitive architecture for a Go MCP server that gives Claude Code **complete autonomous control** over a UE 5.7 project — builds, tests, editor manipulation, Blueprint editing, procedural mesh generation, and documentation lookup — all from one `go build`.

---

## 1. What This Replaces

| Before | After |
|--------|-------|
| UnrealClaude Node.js MCP bridge | Go MCP server (stdio) |
| Bash test scripts | `run_tests` / `run_visual_tests` tools |
| Manual build → test → fix cycle | Autonomous edit → build → test → fix loop |
| Tabbing to UE docs | Built-in `lookup_docs` tool with local index |
| Separate Context7 MCP for docs | Embedded doc search in same binary |
| No procedural mesh support | RealtimeMesh + ProceduralMesh tools |

The only thing that still runs inside UE is a **C++ editor plugin** (the HTTP server that exposes Blueprint graphs, output log, viewport capture, etc.). This is unavoidable — you need code inside the UE process to access editor internals. But the MCP server itself is a **single Go binary** with no Node.js, no Python, no npm.

---

## 2. Architecture

```
                                                    ┌───────────────────────┐
                                                    │   UE 5.7 Editor       │
                                                    │                       │
                                                    │  ┌─────────────────┐  │
                                                    │  │ Remote Control  │  │
                                              ┌────►│  │ API (port 30010)│  │
                                              │     │  │ (built-in)      │  │
                                              │     │  └─────────────────┘  │
┌──────────────┐    stdio     ┌──────────────┐│     │                       │
│              │   JSON-RPC   │              ││     │  ┌─────────────────┐  │
│  Claude Code │◄────────────►│  mcp-unreal  │├────►│  │ MCPUnreal       │  │
│              │              │ (Go binary)  ││     │  │ Plugin (port    │  │
└──────────────┘              │              ││     │  │ 8090)           │  │
                              │ 30+ tools    │┘     │  │ • BP graphs     │  │
                              │ 6 resources  │      │  │ • Anim BP       │  │
                              │ doc index    │      │  │ • Output log    │  │
                              │              │      │  │ • Viewport cap  │  │
                              │ ┌──────────┐ │      │  │ • Script exec   │  │
                              │ │ Headless │ │      │  │ • Asset queries │  │
                              │ │ exec.Cmd │─┼──────│──│ • RMC bridge    │  │
                              │ └──────────┘ │      │  └─────────────────┘  │
                              │              │      │                       │
                              │ ┌──────────┐ │      └───────────────────────┘
                              │ │ Bleve    │ │
                              │ │ Doc Index│ │      ┌───────────────────────┐
                              │ └──────────┘ │      │ docs/                 │
                              │              │──────│ ├── ue5.7/            │
                              └──────────────┘      │ ├── realtimemesh/     │
                                                    │ ├── project/          │
                                                    │ └── index.bleve       │
                                                    └───────────────────────┘
```

### Communication Paths

| Path | Protocol | What For |
|------|----------|----------|
| Claude Code ↔ mcp-unreal | stdio / JSON-RPC 2.0 | MCP tool calls & responses |
| mcp-unreal → UnrealEditor-Cmd | `exec.Command` (subprocess) | Build, test, cook (headless, no editor needed) |
| mcp-unreal → Remote Control API | HTTP PUT `localhost:30010` | Actor properties, function calls, asset search |
| mcp-unreal → MCPUnreal Editor Plugin | HTTP `localhost:8090` | Blueprint graphs, anim BP, output log, viewport, scripts, RMC |
| mcp-unreal → docs/ | Filesystem + bleve | Documentation lookup |

---

## 3. Complete Tool Inventory (34 tools)

### 3.1 Build & Compile (4 tools)

| Tool | Description | Mode |
|------|-------------|------|
| `build_project` | Full compile via Build.sh / UBT. Params: target, config, platform, clean | Headless |
| `cook_project` | Cook content for target platform via RunUAT | Headless |
| `live_compile` | Trigger Live Coding hot-reload in running editor | Editor |
| `generate_project_files` | Regenerate .xcworkspace / VS solution after adding modules | Headless |

### 3.2 Test Automation (4 tools)

| Tool | Description | Mode |
|------|-------------|------|
| `run_tests` | Headless automation tests (-nullrhi). Structured JSON results with per-test pass/fail and failure events | Headless |
| `run_visual_tests` | GPU-rendered visual tests. Returns screenshot paths | Headless+Display |
| `list_tests` | List available test names matching a filter | Headless |
| `get_test_log` | Read raw UE log from a specific test run, with line limit | Filesystem |

### 3.3 Actor & Level (7 tools)

| Tool | Description | Backend |
|------|-------------|---------|
| `get_level_actors` | List all actors, filterable by class/name/tag | Plugin |
| `spawn_actor` | Spawn actor by class with transform | Plugin |
| `delete_actors` | Delete actors by name or path (batch) | Plugin |
| `move_actor` | Set location/rotation/scale | RC API |
| `set_property` | Set any UPROPERTY on any UObject | RC API |
| `get_property` | Read any UPROPERTY from any UObject | RC API |
| `call_function` | Call any UFUNCTION on any UObject | RC API |

### 3.4 Blueprint Editing (2 tools, operation-based)

| Tool | Description |
|------|-------------|
| `blueprint_query` | Operations: `list`, `inspect`, `get_graph`, `get_node_types`. Read-only introspection |
| `blueprint_modify` | Operations: `create`, `add_variable`, `remove_variable`, `add_function`, `remove_function`, `add_node`, `add_nodes_batch`, `delete_node`, `connect_pins`, `disconnect_pins`, `set_pin_value`, `compile`. Auto-compiles after mutation |

### 3.5 Animation Blueprint (2 tools)

| Tool | Description |
|------|-------------|
| `anim_blueprint_query` | Operations: `list_state_machines`, `inspect_state_machine`, `list_states`, `list_transitions` |
| `anim_blueprint_modify` | Operations: `create_state_machine`, `delete_state_machine`, `rename_state_machine`, `set_entry_state`, `create_state`, `delete_state`, `rename_state`, `create_transition`, `delete_transition`, `add_anim_node`, `delete_anim_node`. Auto-compiles |

### 3.6 Asset Management (2 tools)

| Tool | Description | Backend |
|------|-------------|---------|
| `search_assets` | Search asset registry by name, class, path. Paginated | RC API |
| `get_asset_info` | Metadata, dependencies, referencers for a specific asset | Plugin |

### 3.7 Materials (1 tool)

| Tool | Description |
|------|-------------|
| `material_ops` | Operations: `create`, `create_instance`, `set_scalar_param`, `set_vector_param`, `set_texture_param`, `get_params`, `list_instances` |

### 3.8 Character & Enhanced Input (2 tools)

| Tool | Description |
|------|-------------|
| `character_config` | Operations: `get`, `set`. Read/write character movement settings (max walk speed, jump height, gravity, etc.) |
| `input_ops` | Operations: `create_action`, `create_mapping_context`, `bind_action`, `list_actions`, `list_mappings` |

### 3.9 Procedural Mesh / RealtimeMesh (2 tools)

| Tool | Description |
|------|-------------|
| `procedural_mesh` | Operations: `create_section`, `update_section`, `clear`, `set_material`. Works with UProceduralMeshComponent — the built-in UE component. Feed vertices, triangles, normals, UVs as JSON arrays |
| `realtime_mesh` | Operations: `create_lod`, `create_section_group`, `create_section`, `update_mesh_data`, `set_material_slot`, `setup_collision`. Works with URealtimeMeshComponent (requires RMC plugin). Supports LODs, section groups, material slots, collision |

These tools send mesh data (vertex positions, triangle indices, normals, tangents, UVs, vertex colors) as JSON arrays to the editor plugin, which constructs the mesh data structures and applies them to the component. Claude Code generates the geometry procedurally and feeds it through.

**Example: Claude Code creating a procedural terrain**
```
1. spawn_actor(class: "Actor", name: "Terrain")
2. procedural_mesh(operation: "create_section", actor: "Terrain", section: 0, data: {
     vertices: [[0,0,0], [100,0,0], [0,100,0], ...],
     triangles: [0,1,2, ...],
     normals: [[0,0,1], ...],
     uvs: [[0,0], [1,0], [0,1], ...]
   })
3. material_ops(operation: "create_instance", parent: "/Engine/BasicShapes/BasicShapeMaterial", name: "TerrainMat")
4. procedural_mesh(operation: "set_material", actor: "Terrain", section: 0, material: "/Game/TerrainMat")
```

**RealtimeMesh specifics**: The RMC has a richer API than PMC. Key differences the tool exposes:
- **LODs**: `create_lod(mesh, lod_index, screen_size)` — each LOD can have different geometry
- **Section Groups**: Organizational layer between LOD and Section
- **Collision**: `setup_collision(mesh, collision_mesh_data)` — separate from render mesh
- **Material Slots**: Named material slots, not just section indices

The C++ plugin side wraps `URealtimeMeshSimple` which is the simplified API that handles mesh data via `FRealtimeMeshStreamBuilder`.

### 3.10 Level Management (1 tool)

| Tool | Description |
|------|-------------|
| `level_ops` | Operations: `open`, `create`, `save`, `save_as`, `list`, `get_current` |

### 3.11 Editor Utilities (7 tools)

| Tool | Description | Backend |
|------|-------------|---------|
| `run_console_command` | Any UE console command (stat fps, obj list, etc.) | RC API |
| `get_output_log` | Recent output log entries, filterable by category/verbosity | Plugin |
| `capture_viewport` | Screenshot active viewport. Returns base64 or file path. `include_ui=true` captures with Slate/UMG overlays (requires PIE) | Plugin |
| `execute_script` | Run Python script in editor context | Plugin |
| `pie_control` | Start/stop/status for Play In Editor sessions. Supports map override and Simulate In Editor mode | Plugin |
| `player_control` | Control player pawn (get_info, teleport, set_rotation, set_view_target) and editor viewport camera (get_camera, set_camera). Player ops require PIE; camera ops work without PIE | Plugin |
| `status` | Server health: editor online, project path, UE version, plugin version, enabled features | Both |

### 3.12 Documentation Lookup (2 tools)

| Tool | Description |
|------|-------------|
| `lookup_docs` | Search UE 5.7 API docs, RealtimeMesh docs, and project-specific docs. Returns relevant snippets with source references. Params: `query`, `category` (optional: actor, blueprint, material, animation, input, realtimemesh, gameplay, rendering, networking), `max_tokens` (default 3000) |
| `lookup_class` | Get full class reference for a specific UE class: inheritance chain, key properties, key functions, usage notes. Param: `class_name` (e.g. "AActor", "URealtimeMeshSimple", "UCharacterMovementComponent") |

---

## 4. Documentation System

This is what makes autonomous development viable. Without docs, Claude hallucinates UE APIs. With a local search index, it can look up exact function signatures and usage patterns before writing code.

### 4.1 Architecture

Use [bleve](https://github.com/blevesearch/bleve) — a Go-native full-text search library. No external services, no API keys, builds into the binary.

```go
import "github.com/blevesearch/bleve/v2"

type DocEntry struct {
    ID       string   `json:"id"`
    Title    string   `json:"title"`
    Category string   `json:"category"`   // "actor", "blueprint", "realtimemesh", etc.
    Source   string   `json:"source"`      // "ue5.7", "realtimemesh", "project"
    Content  string   `json:"content"`     // The actual doc text
    Classes  []string `json:"classes"`     // Related class names for cross-referencing
    URL      string   `json:"url"`         // Original doc URL for attribution
}
```

### 4.2 Doc Sources to Index

| Source | What to Scrape | How |
|--------|---------------|-----|
| **UE 5.7 API** | Key class references (AActor, UActorComponent, UCharacterMovementComponent, UMaterialInterface, UBlueprintGeneratedClass, etc.) | Scrape from dev.epicgames.com/documentation or extract from Engine/Documentation XML |
| **UE 5.7 Guides** | Automation testing, Remote Control API, Blueprint Visual Scripting, Enhanced Input, Gameplay Framework, Animation | Markdown conversion of key doc pages |
| **RealtimeMesh** | URealtimeMesh, URealtimeMeshSimple, URealtimeMeshComponent, FRealtimeMeshStreamBuilder, LOD system, section groups, collision | GitHub wiki + header comments from TriAxis-Games/RealtimeMeshComponent |
| **ProceduralMeshComponent** | UProceduralMeshComponent API: CreateMeshSection, UpdateMeshSection, ClearMeshSection, SetMaterial | UE 5.7 docs + source header comments |
| **User Project** | Your project's CLAUDE.md, module structure, custom class references, test naming conventions | Auto-indexed from project root at startup |

### 4.3 Index Build Process

```bash
# One-time: build the doc index (run during setup, not at runtime)
mcp-unreal --build-index \
    --ue-docs ./docs/ue5.7/ \
    --rmc-docs ./docs/realtimemesh/ \
    --project-docs ./docs/project/ \
    --output ./docs/index.bleve
```

The index ships alongside the binary (or gets built on first run). At runtime, `lookup_docs` queries are sub-millisecond.

### 4.4 Implementation

```go
package docs

import (
    "context"
    "fmt"
    "strings"

    "github.com/blevesearch/bleve/v2"
    "github.com/blevesearch/bleve/v2/search/highlight/highlighter/ansi"
    "github.com/modelcontextprotocol/go-sdk/mcp"
)

type DocIndex struct {
    index bleve.Index
}

func OpenIndex(path string) (*DocIndex, error) {
    idx, err := bleve.Open(path)
    if err != nil {
        return nil, fmt.Errorf("failed to open doc index at %s: %w", path, err)
    }
    return &DocIndex{index: idx}, nil
}

type LookupDocsInput struct {
    Query     string `json:"query" jsonschema:"Natural language query about UE APIs, classes, or patterns"`
    Category  string `json:"category,omitempty" jsonschema:"Optional filter: actor|blueprint|material|animation|input|realtimemesh|gameplay|rendering|networking"`
    MaxTokens int    `json:"max_tokens,omitempty" jsonschema:"Max tokens to return,default=3000"`
}

type DocResult struct {
    Title    string  `json:"title"`
    Source   string  `json:"source"`
    Category string  `json:"category"`
    Snippet  string  `json:"snippet"`
    URL      string  `json:"url,omitempty"`
    Score    float64 `json:"score"`
}

type LookupDocsOutput struct {
    Results []DocResult `json:"results"`
    Total   int         `json:"total"`
}

func (d *DocIndex) LookupDocs(ctx context.Context, req *mcp.CallToolRequest, input LookupDocsInput) (*mcp.CallToolResult, LookupDocsOutput, error) {
    maxTokens := input.MaxTokens
    if maxTokens == 0 {
        maxTokens = 3000
    }

    queryStr := input.Query
    if input.Category != "" {
        queryStr = fmt.Sprintf("+category:%s %s", input.Category, queryStr)
    }

    searchReq := bleve.NewSearchRequest(bleve.NewQueryStringQuery(queryStr))
    searchReq.Size = 10
    searchReq.Fields = []string{"title", "source", "category", "content", "url"}
    searchReq.Highlight = bleve.NewHighlightWithStyle(ansi.Name)

    result, err := d.index.Search(searchReq)
    if err != nil {
        return nil, LookupDocsOutput{}, fmt.Errorf("search failed: %w", err)
    }

    var results []DocResult
    totalTokens := 0
    for _, hit := range result.Hits {
        content, _ := hit.Fields["content"].(string)
        // Rough token estimate: 1 token ≈ 4 chars
        tokens := len(content) / 4
        if totalTokens+tokens > maxTokens {
            // Truncate to fit budget
            remaining := (maxTokens - totalTokens) * 4
            if remaining > 0 && len(content) > remaining {
                content = content[:remaining] + "..."
            } else {
                break
            }
        }
        totalTokens += tokens

        results = append(results, DocResult{
            Title:    strField(hit.Fields, "title"),
            Source:   strField(hit.Fields, "source"),
            Category: strField(hit.Fields, "category"),
            Snippet:  content,
            URL:      strField(hit.Fields, "url"),
            Score:    hit.Score,
        })
    }

    return nil, LookupDocsOutput{Results: results, Total: len(results)}, nil
}

type LookupClassInput struct {
    ClassName string `json:"class_name" jsonschema:"UE class name e.g. AActor, URealtimeMeshSimple, UCharacterMovementComponent"`
}

type ClassInfo struct {
    Name        string   `json:"name"`
    Parent      string   `json:"parent"`
    Module      string   `json:"module"`
    Description string   `json:"description"`
    KeyProps    []string `json:"key_properties"`
    KeyFuncs    []string `json:"key_functions"`
    Source      string   `json:"source"`
    URL         string   `json:"url,omitempty"`
}

type LookupClassOutput struct {
    Found bool      `json:"found"`
    Class ClassInfo `json:"class,omitempty"`
}

func (d *DocIndex) LookupClass(ctx context.Context, req *mcp.CallToolRequest, input LookupClassInput) (*mcp.CallToolResult, LookupClassOutput, error) {
    query := fmt.Sprintf("+classes:%s", input.ClassName)
    searchReq := bleve.NewSearchRequest(bleve.NewQueryStringQuery(query))
    searchReq.Size = 1
    searchReq.Fields = []string{"title", "source", "content", "url", "classes"}

    result, err := d.index.Search(searchReq)
    if err != nil || result.Total == 0 {
        return nil, LookupClassOutput{Found: false}, nil
    }

    hit := result.Hits[0]
    content, _ := hit.Fields["content"].(string)

    // Parse structured class info from content
    info := parseClassDoc(input.ClassName, content)
    info.Source = strField(hit.Fields, "source")
    info.URL = strField(hit.Fields, "url")

    return nil, LookupClassOutput{Found: true, Class: info}, nil
}

func strField(fields map[string]interface{}, key string) string {
    if v, ok := fields[key].(string); ok {
        return v
    }
    return ""
}

func parseClassDoc(name, content string) ClassInfo {
    // Parse the doc content into structured fields
    // This is a simplified version — real implementation would parse
    // the indexed markdown/xml format
    return ClassInfo{
        Name:        name,
        Description: content,
    }
}
```

### 4.5 Why Not Context7?

Context7 is great for web frameworks but has no UE 5.7 coverage. You'd need to add UE docs to their platform and depend on a remote API. By embedding bleve, you get:
- **Offline**: Works in air-gapped environments, no API keys
- **UE-specific**: Index only what matters — no noise from React/Next.js docs
- **Project-aware**: Index your own project's docs and conventions
- **Fast**: Sub-millisecond queries, no network round-trip
- **Single binary**: The index is a directory on disk, loaded at startup

---

## 5. The C++ Editor Plugin (MCPUnreal Editor Plugin)

The Go binary handles MCP protocol, headless operations, and doc search. But for live editor interaction beyond what the Remote Control API provides, you need a C++ plugin running inside UE. This is a thin HTTP server — the Go binary is the brain.

### 5.1 Endpoints the Plugin Must Expose

```
POST /api/status                    → { project, version, capabilities }
POST /api/actors/list               → [{ name, class, path, location, rotation, scale }]
POST /api/actors/spawn              → { actor_path, actor_name }
POST /api/actors/delete             → { deleted_count }

POST /api/blueprints/list           → [{ name, path, parent_class }]
POST /api/blueprints/inspect        → { variables, functions, graphs }
POST /api/blueprints/get_graph      → { nodes, connections }
POST /api/blueprints/modify         → { success, compiled }

POST /api/anim_blueprints/query     → { state_machines, states, transitions }
POST /api/anim_blueprints/modify    → { success, compiled }

POST /api/assets/info               → { metadata, dependencies, referencers }
POST /api/assets/dependencies       → { deps[] }
POST /api/assets/referencers        → { refs[] }

POST /api/materials/modify          → { success }

POST /api/character/config          → { movement_settings }
POST /api/input/manage              → { success }

POST /api/mesh/procedural           → { success, component_path }
POST /api/mesh/realtime             → { success, component_path }

POST /api/editor/output_log         → { entries[] }
POST /api/editor/capture_viewport   → { image_base64 | file_path }
POST /api/editor/execute_script     → { output, success }
POST /api/editor/pie_control        → { success, pie_active, pie_map }
POST /api/editor/player_control     → { success, location, rotation, ... }

POST /api/levels/manage             → { success, level_path }
```

### 5.2 Plugin Architecture

```cpp
// MCPUnrealModule.h
UCLASS()
class UMCPUnrealSubsystem : public UEditorSubsystem
{
    GENERATED_BODY()
public:
    virtual void Initialize(FSubsystemCollectionBase& Collection) override;
    virtual void Deinitialize() override;
private:
    TSharedPtr<FHttpServerModule> HttpServer;
    void RegisterRoutes();

    // Route handlers
    FHttpRequestHandler HandleActorsList();
    FHttpRequestHandler HandleActorsSpawn();
    FHttpRequestHandler HandleBlueprintModify();
    FHttpRequestHandler HandleRealtimeMesh();
    // ... etc
};
```

The plugin uses UE's built-in `FHttpServerModule` (available since UE 4.25) to run on port 8090. No external HTTP library needed.

### 5.3 RealtimeMesh Plugin Integration

The editor plugin conditionally compiles RealtimeMesh support:

```cpp
// In Build.cs
if (Target.bBuildEditor)
{
    PrivateDependencyModuleNames.Add("RealtimeMeshComponent");
}
```

The `/api/mesh/realtime` endpoint handler:

```cpp
void UMCPUnrealSubsystem::HandleRealtimeMesh(const FHttpServerRequest& Request, ...)
{
    // Parse JSON body
    FString Operation = JsonBody->GetStringField("operation");

    if (Operation == "create_section")
    {
        AActor* Actor = FindActorByName(JsonBody->GetStringField("actor"));
        auto* RMC = Actor->FindComponentByClass<URealtimeMeshComponent>();
        auto Mesh = RMC->GetRealtimeMeshAs<URealtimeMeshSimple>();

        FRealtimeMeshStreamSet StreamSet;
        // ... populate from JSON vertices/triangles/normals/uvs
        
        FRealtimeMeshSectionGroupKey GroupKey = FRealtimeMeshSectionGroupKey::Create(0, 0);
        FRealtimeMeshSectionKey SectionKey = FRealtimeMeshSectionKey::Create(GroupKey, 0);
        Mesh->CreateSectionGroup(GroupKey);
        Mesh->CreateSection(SectionKey, FRealtimeMeshSectionConfig(), StreamSet);
    }
    // ... handle update_mesh_data, create_lod, setup_collision, etc.
}
```

---

## 6. Project Structure

```
mcp-unreal/
├── cmd/
│   └── mcp-unreal/
│       └── main.go                 # Entry: server setup, all tool registration
├── internal/
│   ├── config/
│   │   └── config.go               # Env vars, project detection, paths
│   ├── headless/
│   │   ├── build.go                # build_project, cook_project, generate_project_files
│   │   ├── test.go                 # run_tests, run_visual_tests, list_tests
│   │   └── log.go                  # get_test_log
│   ├── editor/
│   │   ├── client.go               # HTTP client: RC API + plugin endpoints
│   │   ├── actors.go               # get_level_actors, spawn_actor, delete_actors, move_actor
│   │   ├── properties.go           # set_property, get_property, call_function
│   │   ├── blueprints.go           # blueprint_query, blueprint_modify
│   │   ├── anim_blueprints.go      # anim_blueprint_query, anim_blueprint_modify
│   │   ├── assets.go               # search_assets, get_asset_info
│   │   ├── materials.go            # material_ops
│   │   ├── characters.go           # character_config
│   │   ├── input.go                # input_ops
│   │   ├── mesh.go                 # procedural_mesh, realtime_mesh
│   │   ├── levels.go               # level_ops
│   │   └── utilities.go            # run_console_command, get_output_log, capture_viewport, execute_script
│   ├── docs/
│   │   ├── index.go                # Bleve index open/create/search
│   │   ├── lookup.go               # lookup_docs, lookup_class tool handlers
│   │   ├── ingest.go               # Doc ingestion: markdown → bleve entries
│   │   └── class_parser.go         # Parse UE class references into structured data
│   └── status/
│       └── status.go               # status tool, health checks
├── docs/                            # Doc source files (committed to repo)
│   ├── ue5.7/                       # Curated UE 5.7 API markdown
│   │   ├── actors.md
│   │   ├── blueprints.md
│   │   ├── animation.md
│   │   ├── materials.md
│   │   ├── input.md
│   │   ├── gameplay_framework.md
│   │   ├── rendering.md
│   │   ├── networking.md
│   │   ├── automation_testing.md
│   │   ├── remote_control_api.md
│   │   └── classes/                 # Per-class reference files
│   │       ├── AActor.md
│   │       ├── UActorComponent.md
│   │       ├── UCharacterMovementComponent.md
│   │       ├── UBlueprintGeneratedClass.md
│   │       ├── UMaterialInterface.md
│   │       ├── UProceduralMeshComponent.md
│   │       └── ... (50-100 key classes)
│   ├── realtimemesh/
│   │   ├── overview.md
│   │   ├── URealtimeMesh.md
│   │   ├── URealtimeMeshSimple.md
│   │   ├── URealtimeMeshComponent.md
│   │   ├── lods_and_sections.md
│   │   └── collision.md
│   └── project/                     # Auto-populated from project
│       └── (generated at index time)
├── plugin/                          # UE C++ editor plugin source
│   ├── MCPUnreal.uplugin
│   ├── Source/
│   │   └── MCPUnreal/
│   │       ├── MCPUnreal.Build.cs
│   │       ├── MCPUnrealModule.h / .cpp
│   │       ├── HttpServer.h / .cpp
│   │       ├── ActorRoutes.h / .cpp
│   │       ├── BlueprintRoutes.h / .cpp
│   │       ├── AnimBlueprintRoutes.h / .cpp
│   │       ├── AssetRoutes.h / .cpp
│   │       ├── MaterialRoutes.h / .cpp
│   │       ├── MeshRoutes.h / .cpp      # ProceduralMesh + RealtimeMesh
│   │       ├── EditorRoutes.h / .cpp
│   │       └── LevelRoutes.h / .cpp
│   └── Config/
│       └── DefaultMCPUnreal.ini
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

---

## 7. Build & Distribution

```makefile
# Makefile

BINARY=mcp-unreal
VERSION=0.1.0

.PHONY: build build-index install clean

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY) ./cmd/mcp-unreal

build-index: build
	./$(BINARY) --build-index \
		--ue-docs ./docs/ue5.7 \
		--rmc-docs ./docs/realtimemesh \
		--project-root ../ \
		--output ./docs/index.bleve

# Install: build binary + index, copy plugin to project
install: build build-index
	cp $(BINARY) ../tools/$(BINARY)
	cp -r docs/index.bleve ../tools/docs-index/
	@echo "Plugin must be manually copied: cp -r plugin/ ../Plugins/MCPUnreal/"
	@echo ""
	@echo "Register with Claude Code:"
	@echo "  claude mcp add mcp-unreal -- ../tools/mcp-unreal"

clean:
	rm -f $(BINARY)
	rm -rf docs/index.bleve
```

### Team Distribution via `.mcp.json`

```json
{
  "mcpServers": {
    "mcp-unreal": {
      "command": "./tools/mcp-unreal",
      "args": ["--docs-index", "./tools/docs-index"],
      "env": {
        "UE_EDITOR_PATH": "/Users/Shared/Epic Games/UE_5.7/Engine/Binaries/Mac/UnrealEditor-Cmd",
        "MCP_UNREAL_PROJECT": ".",
        "RC_API_PORT": "30010",
        "PLUGIN_PORT": "8090"
      }
    }
  }
}
```

---

## 8. CLAUDE.md (Project Instructions for Claude Code)

```markdown
# MyProject — UE 5.7 Project

## MCP Server: mcp-unreal
Single Go binary with 34 tools for complete autonomous UE development.
Check `status` first to see what's available.

## Autonomous Workflow

### After editing C++ source:
1. `build_project(target: "MyProjectEditor", config: "Development")`
2. `run_tests(filter: "MyProject")` — check structured results
3. If failures: read `events` field, fix source, goto 1
4. If all pass: proceed with next task

### Before using any UE API:
1. `lookup_docs(query: "how to ...", category: "...")` — get correct API signatures
2. `lookup_class(class_name: "UWhatever")` — get class reference
3. Then write code using verified APIs

### Scene/Level work:
1. `status` — confirm editor is running
2. `get_level_actors` — see what exists
3. `spawn_actor` / `move_actor` / `set_property` — build scene
4. `capture_viewport` — verify visual result

### Blueprint work:
1. `blueprint_query(operation: "list")` — find target BP
2. `blueprint_query(operation: "inspect", path: "...")` — see variables/functions
3. `blueprint_modify(operation: "add_variable", ...)` — make changes
4. Auto-compiles after each mutation

### Procedural mesh:
1. `spawn_actor` with empty actor
2. `procedural_mesh` or `realtime_mesh` — feed vertex/triangle data
3. `material_ops` — create and assign materials
4. For RealtimeMesh: use LODs for performance (`realtime_mesh(operation: "create_lod")`)

### PIE testing:
1. `pie_control(operation: "start")` — begin Play In Editor
2. `pie_control(operation: "status")` — verify PIE is active
3. `player_control(operation: "get_info")` — get player pawn location, rotation, camera
4. `player_control(operation: "teleport", location: [X,Y,Z])` — move player pawn
5. `capture_viewport(include_ui: true)` — see game view with HUD
6. `pie_control(operation: "stop")` — end session

### Debugging:
- `get_output_log(category: "LogMyProject")` — project-specific logs
- `run_console_command("stat fps")` — performance
- `run_console_command("obj list class=AMyProjectActor")` — object introspection
- `capture_viewport` — visual state

## Test Naming
- `MyProject.*` — All tests
- `MyProject.Visual.*` — Visual/screenshot tests (need GPU)
- `MyProject.Unit.*` — Unit tests (headless OK)
- `MyProject.Integration.*` — Integration tests

## Key Project Classes
- `AMyProjectGameMode` — Main game mode
- `AMyProjectPlayerController` — Player controller
- `UMyProjectAbilityComponent` — GAS integration
(Claude: use `lookup_class` to get full API for any of these)

## Module Dependencies
- `RealtimeMeshComponent` — Runtime mesh generation
- `EnhancedInput` — Input system
(Claude: use `lookup_docs` before using unfamiliar module APIs)
```

---

## 9. Implementation Roadmap

### Phase 1 — Skeleton + Headless (3 days)
- [ ] Go module with official MCP SDK + bleve
- [ ] Config system (env vars, project detection)
- [ ] `status` tool
- [ ] `build_project` + `generate_project_files`
- [ ] `run_tests` + `list_tests` + `get_test_log` (port bash logic to Go)
- [ ] stdio transport, register with Claude Code
- [ ] Verify autonomous build → test → fix loop works

### Phase 2 — Documentation System (2 days)
- [ ] Curate UE 5.7 docs: top 50-100 classes as markdown
- [ ] Curate RealtimeMesh docs from GitHub headers + wiki
- [ ] Bleve indexing pipeline (`--build-index` CLI mode)
- [ ] `lookup_docs` + `lookup_class` tools
- [ ] Auto-index project CLAUDE.md and source headers at startup

### Phase 3 — Editor Plugin (4-5 days)
- [ ] C++ plugin: `FHttpServerModule` on port 8090
- [ ] Actor routes: list, spawn, delete
- [ ] Blueprint routes: query + modify (most complex)
- [ ] Anim Blueprint routes
- [ ] Asset info routes
- [ ] Output log + viewport capture + script execution routes

### Phase 4 — Go Editor Tools (3 days)
- [ ] HTTP client for RC API (set_property, get_property, call_function)
- [ ] HTTP client for plugin endpoints
- [ ] All actor tools
- [ ] All blueprint tools
- [ ] All anim blueprint tools
- [ ] Asset, material, character, input tools
- [ ] Console command, output log, capture viewport, execute script
- [ ] Level ops

### Phase 5 — Mesh + Polish (3 days)
- [ ] ProceduralMesh routes in C++ plugin + Go tool
- [ ] RealtimeMesh routes in C++ plugin + Go tool
- [ ] `run_visual_tests` + `cook_project`
- [ ] MCP async tasks for long-running operations (build, cook, test)
- [ ] `.mcp.json` + `CLAUDE.md` finalization
- [ ] Error recovery, graceful degradation, token budget management
- [ ] `live_compile` tool (trigger Live Coding)

### Total: ~15 days for a production-quality server

---

## 10. Key Design Principles

### Single binary, zero dependencies
`go build` produces one executable. Bleve compiles in. No npm, no Python, no Java. The doc index is a directory that ships alongside.

### Graceful degradation
Every tool checks what's available before running. Headless tools always work. Editor tools return `"editor_offline"` error with clear instructions. Doc tools always work (local index).

### Token budget awareness
Claude Code warns at 10K tokens per tool response. Every tool that could return large data has `max_results` or `max_tokens` params. Default to concise. The `lookup_docs` tool caps at 3000 tokens by default — enough for 2-3 class references without flooding context.

### Operation-based mega-tools vs. many small tools
Blueprint editing uses 2 tools (query + modify) with an `operation` field rather than 12 separate tools. This keeps the tool list manageable (34 total) while still being discoverable — the tool description lists all operations.

### Security
Everything runs on localhost. RC API: 30010. Plugin: 8090. Both loopback only. `execute_script` is the only dangerous tool — the plugin should log all scripts and Claude Code's permission system handles approval.

---

## 11. References

| Resource | URL |
|----------|-----|
| MCP Spec v2025-11-25 | https://modelcontextprotocol.io/specification/2025-11-25 |
| Official Go MCP SDK | https://github.com/modelcontextprotocol/go-sdk |
| Go SDK API Docs | https://pkg.go.dev/github.com/modelcontextprotocol/go-sdk/mcp |
| Bleve (Go search engine) | https://github.com/blevesearch/bleve |
| Claude Code MCP Docs | https://code.claude.com/docs/en/mcp |
| UnrealClaude (reference impl) | https://github.com/Natfii/UnrealClaude |
| UnrealClaude MCP Bridge | https://lobehub.com/mcp/natfii-ue5-mcp-bridge |
| UE 5.7 Remote Control API | https://dev.epicgames.com/documentation/en-us/unreal-engine/remote-control-api-http-reference-for-unreal-engine |
| UE 5.7 Remote Control Quick Start | https://dev.epicgames.com/documentation/en-us/unreal-engine/remote-control-quick-start-for-unreal-engine |
| UE 5.7 WebSocket API | https://dev.epicgames.com/documentation/en-us/unreal-engine/remote-control-api-websocket-reference-for-unreal-engine |
| UE 5.7 Automation Tests | https://dev.epicgames.com/documentation/en-us/unreal-engine/run-automation-tests-in-unreal-engine |
| UE 5.7 Test Configuration | https://dev.epicgames.com/documentation/en-us/unreal-engine/configure-automation-tests-in-unreal-engine |
| UE 5.7 Gauntlet | https://dev.epicgames.com/documentation/en-us/unreal-engine/running-gauntlet-tests-in-unreal-engine |
| UE 5.7 ProceduralMeshComponent | https://dev.epicgames.com/documentation/en-us/unreal-engine/API/Plugins/ProceduralMeshComponent/UProceduralMeshComponent |
| RealtimeMesh GitHub | https://github.com/TriAxis-Games/RealtimeMeshComponent |
| RealtimeMesh Fab Page | https://www.unrealengine.com/marketplace/en-US/product/runtime-mesh-component |
| Context7 (architecture reference) | https://github.com/upstash/context7 |
| chongdashu/unreal-mcp | https://github.com/chongdashu/unreal-mcp |
| ChiR24/Unreal_mcp | https://github.com/ChiR24/Unreal_mcp |
| mirno-ehf/ue5-mcp | https://github.com/mirno-ehf/ue5-mcp |
| MCP Year in Review | https://www.pento.ai/blog/a-year-of-mcp-2025-review |
