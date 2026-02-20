# mcp-unreal

[![CI](https://github.com/remiphilippe/mcp-unreal/actions/workflows/ci.yml/badge.svg)](https://github.com/remiphilippe/mcp-unreal/actions/workflows/ci.yml)
[![Go 1.25+](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go)](https://go.dev)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

MCP (Model Context Protocol) server that gives AI coding agents complete autonomous control over an Unreal Engine 5.7 project. Single Go binary, zero external dependencies.

Build, test, manipulate the editor, edit Blueprints, generate procedural meshes, and look up UE API documentation — all through 49 MCP tools that any MCP-compatible agent can call directly.

## Quick Start

```bash
# Install
go install github.com/remiphilippe/mcp-unreal/cmd/mcp-unreal@latest

# Build the documentation index
mcp-unreal --build-index

# Register with your AI coding tool (see Setup section below)
```

## Prerequisites

- **Go 1.25+** — [install](https://go.dev/dl/)
- **Unreal Engine 5.7** — for build/test tools and editor interaction
- **Remote Control API plugin** — enabled in UE (see [Enabling Remote Control API](#enabling-remote-control-api) below)
- **MCPUnreal editor plugin** — for advanced editor tools (see [plugin/README.md](plugin/README.md))

## Installation

### From Source

```bash
git clone https://github.com/remiphilippe/mcp-unreal.git
cd mcp-unreal
make build
```

### Pre-built Binaries

Download from [GitHub Releases](https://github.com/remiphilippe/mcp-unreal/releases) — available for macOS, Linux, and Windows on both amd64 and arm64.

### Enabling Remote Control API

The Remote Control API is a built-in UE plugin that provides HTTP access to actor properties and function calls on port 30010. mcp-unreal uses it for `get_property`, `set_property`, `call_function`, and `move_actor` tools.

1. Open your project in Unreal Editor
2. Go to **Edit > Plugins**
3. Search for **"Remote Control API"**
4. Check the **Enabled** checkbox
5. Restart the editor when prompted

To verify it's running:

```bash
curl http://localhost:30010/remote/info
```

You should get a JSON response with the API version. If the connection is refused, the plugin is not enabled or the editor isn't running.

### Install the UE Plugin

Copy the `plugin/` directory into your project's `Plugins/` folder:

```bash
cp -r plugin/ /path/to/YourProject/Plugins/MCPUnreal/
```

Rebuild the project. The plugin starts an HTTP server on port 8090 when the editor loads. See [plugin/README.md](plugin/README.md) for details.

## Setup by Client

mcp-unreal works with any MCP-compatible AI coding tool. Below are setup instructions for each.

> **Note**: Replace `/path/to/mcp-unreal` with the actual path to your built binary (e.g. the output of `go install` or `make build`). Set `MCP_UNREAL_PROJECT` to your `.uproject` file path or project root directory.

### Claude Code

**CLI registration:**

```bash
claude mcp add mcp-unreal -- /path/to/mcp-unreal
```

**Config file** (`.mcp.json` at project root, or `~/.claude.json` for global):

```json
{
  "mcpServers": {
    "mcp-unreal": {
      "type": "stdio",
      "command": "/path/to/mcp-unreal",
      "env": {
        "MCP_UNREAL_PROJECT": "/path/to/MyProject/MyProject.uproject"
      }
    }
  }
}
```

### Cursor

**Config file** (`.cursor/mcp.json` at project root, or `~/.cursor/mcp.json` for global):

```json
{
  "mcpServers": {
    "mcp-unreal": {
      "command": "/path/to/mcp-unreal",
      "env": {
        "MCP_UNREAL_PROJECT": "/path/to/MyProject/MyProject.uproject"
      }
    }
  }
}
```

After saving, enable via **Cursor Settings > MCP Servers**.

### OpenAI Codex CLI

**CLI registration:**

```bash
codex mcp add mcp-unreal -- /path/to/mcp-unreal
```

**Config file** (`.codex/config.toml` at project root, or `~/.codex/config.toml` for global):

```toml
[mcp_servers.mcp-unreal]
command = "/path/to/mcp-unreal"
args = []
tool_timeout_sec = 120

[mcp_servers.mcp-unreal.env]
MCP_UNREAL_PROJECT = "/path/to/MyProject/MyProject.uproject"
```

> Increase `tool_timeout_sec` for `build_project` and `run_tests` which can take 60+ seconds.

### VS Code with GitHub Copilot

Requires VS Code 1.99+ with GitHub Copilot Agent Mode enabled.

**Config file** (`.vscode/mcp.json` at project root):

```json
{
  "servers": {
    "mcp-unreal": {
      "type": "stdio",
      "command": "/path/to/mcp-unreal",
      "env": {
        "MCP_UNREAL_PROJECT": "${workspaceFolder}"
      }
    }
  }
}
```

### Windsurf

**Config file** (`~/.codeium/windsurf/mcp_config.json`):

```json
{
  "mcpServers": {
    "mcp-unreal": {
      "command": "/path/to/mcp-unreal",
      "env": {
        "MCP_UNREAL_PROJECT": "/path/to/MyProject/MyProject.uproject"
      }
    }
  }
}
```

### Cline (VS Code Extension)

Open the Cline panel > MCP Servers icon > Configure tab > **Advanced MCP Settings**, then add:

```json
{
  "mcpServers": {
    "mcp-unreal": {
      "command": "/path/to/mcp-unreal",
      "env": {
        "MCP_UNREAL_PROJECT": "/path/to/MyProject/MyProject.uproject"
      }
    }
  }
}
```

### Zed

**Config file** (`~/.config/zed/settings.json`):

```json
{
  "context_servers": {
    "mcp-unreal": {
      "command": "/path/to/mcp-unreal",
      "env": {
        "MCP_UNREAL_PROJECT": "/path/to/MyProject/MyProject.uproject"
      }
    }
  }
}
```

### Quick Reference

| Client | Config File | Top-level Key |
|--------|-------------|---------------|
| Claude Code | `.mcp.json` or `~/.claude.json` | `mcpServers` |
| Cursor | `.cursor/mcp.json` or `~/.cursor/mcp.json` | `mcpServers` |
| Codex CLI | `.codex/config.toml` or `~/.codex/config.toml` | `[mcp_servers.<name>]` |
| VS Code Copilot | `.vscode/mcp.json` | `servers` |
| Windsurf | `~/.codeium/windsurf/mcp_config.json` | `mcpServers` |
| Cline | Extension-managed settings | `mcpServers` |
| Zed | `~/.config/zed/settings.json` | `context_servers` |

## Recommended System Prompt

For best results, add the following to your project's system prompt (e.g. `CLAUDE.md`, `.cursorrules`, `.codex/instructions.md`, or equivalent). This tells the AI agent how to use the MCP tools effectively.

```markdown
## Unreal Engine MCP Tools

This project uses mcp-unreal for AI-driven Unreal Engine development. The following MCP tools are available:

### Workflow Guidelines

1. **Always check status first**: Call `status` to verify the editor and plugin connections before attempting editor operations.
2. **Look up docs before writing UE code**: Use `lookup_class` to get class references (inheritance, properties, functions) and `lookup_docs` for API patterns before writing C++ or Blueprint logic.
3. **Build-test cycle**: After editing C++ files, call `build_project` to compile, then `run_tests` to verify. Read build errors carefully — use `lookup_docs` to understand UE APIs.
4. **Actor workflow**: Use `get_level_actors` to discover existing actors, `spawn_actor` to create new ones, `move_actor` to position them, and `set_property`/`get_property` to read/write properties.
5. **Blueprint workflow**: Use `blueprint_query` with operation `list` to find Blueprints, then `list_variables`/`list_functions`/`list_nodes` to inspect them. Use `blueprint_modify` to make changes, and always `compile` after modifications.
6. **Material workflow**: Use `material_ops` with `create` to make new materials, `get_parameters` to inspect existing ones, and `set_parameter` or `set_texture` to modify them.
7. **Save your work**: The editor does not auto-save. After making changes, use `level_ops` with `save_level` or remind the user to save.
8. **PIE testing**: Use `pie_control` with operation `start` to begin a Play In Editor session, `status` to verify it started, and `stop` to end it. Use `capture_viewport` with `include_ui=true` during PIE to see the game view with HUD/menus. Use `player_control` to move the player pawn (`teleport`, `set_rotation`) or reposition the editor viewport camera (`get_camera`, `set_camera`).

### Tool Modes

- **Headless tools** (`build_project`, `run_tests`, `cook_project`, etc.) do NOT require the editor to be running. They invoke UnrealEditor-Cmd directly.
- **Editor tools** (actors, blueprints, materials, etc.) require the Unreal Editor to be open with the MCPUnreal plugin loaded.
- **Documentation tools** (`lookup_docs`, `lookup_class`) are always available — use them liberally.

### Object Paths

Many tools require UE object paths. The format is:
- Actor paths: `/Game/Maps/MapName.MapName:PersistentLevel.ActorName`
- Asset paths: `/Game/Folder/AssetName`
- Engine assets: `/Engine/BasicShapes/Cube`

Use `get_level_actors` to discover actor paths, and `search_assets` to find asset paths.
```

> **Tip**: For Claude Code, save this as `CLAUDE.md` in your project root. For Cursor, save as `.cursorrules`. For Codex, save as `.codex/instructions.md`.

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `UE_EDITOR_PATH` | Platform-dependent | Path to `UnrealEditor-Cmd` binary |
| `MCP_UNREAL_PROJECT` | Auto-detected from cwd | Path to `.uproject` file or project root |
| `RC_API_PORT` | `30010` | UE Remote Control API HTTP port |
| `PLUGIN_PORT` | `8090` | MCPUnreal editor plugin HTTP port |
| `MCP_UNREAL_LOG_LEVEL` | `info` | Log level: `debug`, `info`, `warn`, `error` |
| `MCP_UNREAL_DOCS_INDEX` | `./docs/index.bleve` | Path to bleve documentation index |

Platform defaults for `UE_EDITOR_PATH`:
- **macOS**: `/Users/Shared/Epic Games/UE_5.7/Engine/Binaries/Mac/UnrealEditor-Cmd`
- **Windows**: `C:\Program Files\Epic Games\UE_5.7\Engine\Binaries\Win64\UnrealEditor-Cmd.exe`
- **Linux**: `/opt/UnrealEngine/Engine/Binaries/Linux/UnrealEditor-Cmd`

## Architecture

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
│  AI Coding   │   JSON-RPC   │              ││     │  ┌─────────────────┐  │
│    Agent     │◄────────────►│  mcp-unreal  │├────►│  │ MCPUnreal       │  │
│ (Claude Code │              │ (Go binary)  ││     │  │ Plugin (port    │  │
│  Cursor, etc)│              │              ││     │  │ 8090)           │  │
└──────────────┘              │ 49 tools     │┘     │  │ • Actors        │  │
                              │ doc index    │      │  │ • Blueprints    │  │
                              │              │      │  │ • Materials     │  │
                              │ ┌──────────┐ │      │  │ • PCG / GAS    │  │
                              │ │ Headless │ │      │  │ • Niagara       │  │
                              │ │ exec.Cmd │─┼──────│──│ • Mesh ops      │  │
                              │ └──────────┘ │      │  │ • Levels        │  │
                              │              │      │  └─────────────────┘  │
                              │ ┌──────────┐ │      │                       │
                              │ │ Bleve    │ │      └───────────────────────┘
                              │ │ Doc Index│ │
                              │ └──────────┘ │      ┌───────────────────────┐
                              │              │──────│ docs/                 │
                              └──────────────┘      │ ├── ue5.7/            │
                                                    │ ├── realtimemesh/     │
                                                    │ └── index.bleve       │
                                                    └───────────────────────┘
```

Three communication paths:
1. **Headless** — `exec.Command` invokes `UnrealEditor-Cmd` for builds, tests, and cooking. No editor needed.
2. **Remote Control API** — HTTP PUT to `localhost:30010` for property access and function calls. Built into UE.
3. **MCPUnreal Plugin** — HTTP POST to `localhost:8090` for Blueprint editing, asset queries, mesh ops, and other deep editor internals.

See [IMPLEMENTATION.md](IMPLEMENTATION.md) for the full architecture document.

## Available Tools (48)

### Build & Compile (Headless)

| Tool | Description |
|------|-------------|
| `build_project` | Build the UE project via UnrealEditor-Cmd / UBT. Returns structured JSON with error count and error details. |
| `cook_project` | Cook (package) content for a target platform using RunUAT. Supports iterative cooks. |
| `generate_project_files` | Regenerate IDE project files (.xcworkspace / .sln) after adding or removing C++ modules. |

### Project & Config (Headless)

| Tool | Description |
|------|-------------|
| `project_ops` | Read and modify the .uproject file: get project info, list/enable/disable plugins, add modules, set target platforms. |
| `config_ops` | Read and write UE project .ini config files (DefaultEngine.ini, DefaultGame.ini, etc.): get, set, delete keys, list sections. |

### Test Automation (Headless)

| Tool | Description |
|------|-------------|
| `run_tests` | Run headless automation tests (-nullrhi) with per-test pass/fail results and failure details. |
| `run_visual_tests` | Run automation tests with GPU rendering enabled (no -nullrhi) for visual/rendering tests. |
| `list_tests` | List available automation test names matching a filter pattern. |
| `get_test_log` | Read raw UE log files with line limits, offsets, and keyword filtering. |

### Actors & Properties (Editor)

| Tool | Description |
|------|-------------|
| `get_level_actors` | List all actors in the current level with class, path, location, rotation, scale. |
| `spawn_actor` | Spawn an actor by class name at a given location/rotation/scale. |
| `delete_actors` | Delete one or more actors by object path or display name. |
| `move_actor` | Set location, rotation, and/or scale of an actor via RC API function calls (`K2_SetActorLocation`, `K2_SetActorRotation`, `SetActorScale3D`). |
| `get_property` | Read a property value from an actor via the Remote Control API. |
| `set_property` | Set a property value on an actor via the Remote Control API. |
| `call_function` | Call a function on an object via the Remote Control API. |

### Blueprint Editing (Editor)

| Tool | Description |
|------|-------------|
| `blueprint_query` | Query Blueprint structure: list all BPs, inspect a BP's variables/functions/graphs, get graph nodes and connections. |
| `blueprint_modify` | Modify Blueprints: create, add/remove variables, add/remove functions, add/delete nodes, connect/disconnect pins, set pin values, compile. |

### Animation Blueprint Editing (Editor)

| Tool | Description |
|------|-------------|
| `anim_blueprint_query` | Query AnimBP state machines and states: list state machines, inspect states/transitions/blend spaces. |
| `anim_blueprint_modify` | Modify AnimBPs: add/remove states, add/remove transitions, set blend spaces, set state animations, set transition rules, add notify states. |

### Assets (Editor)

| Tool | Description |
|------|-------------|
| `search_assets` | Search assets by class, path, name, or tag. Returns paginated results. |
| `get_asset_info` | Get detailed asset info including dependencies and referencers. |

### Materials (Editor)

| Tool | Description |
|------|-------------|
| `material_ops` | Material operations: create, create_instance, set_parameter, set_texture, get_parameters, list_parameters. |

### Characters & Input (Editor)

| Tool | Description |
|------|-------------|
| `character_config` | Character configuration: get_config, set_movement, set_capsule, set_mesh, set_camera, get_movement_modes. |
| `input_ops` | Enhanced Input operations: list_actions, list_contexts, add_action, remove_action, add_context, bind_action, unbind_action, get_bindings. |

### Procedural Content Generation (Editor)

| Tool | Description |
|------|-------------|
| `pcg_ops` | PCG graph operations: list_components, get_graph, execute, set_parameter, add_node, connect_nodes, get_results. |

### Gameplay Ability System (Editor)

| Tool | Description |
|------|-------------|
| `gas_ops` | GAS operations: list_abilities, grant_ability, activate, get_attributes, set_attribute, apply_effect, remove_effect, list_active_effects. |

### Niagara VFX (Editor)

| Tool | Description |
|------|-------------|
| `niagara_ops` | Niagara operations: spawn_system, set_parameter, get_parameter, activate, deactivate, list_emitters, set_emitter_enabled, get_system_info. |

### Mesh Generation (Editor)

| Tool | Description |
|------|-------------|
| `procedural_mesh` | Create and modify ProceduralMeshComponent geometry: create_section, update_section, clear, set_material. |
| `realtime_mesh` | Create and modify RealtimeMeshComponent geometry (requires RMC plugin): create_lod, create_section_group, create_section, update_mesh_data, set_material_slot, setup_collision. |

### Levels (Editor)

| Tool | Description |
|------|-------------|
| `level_ops` | Level management: get_current, list_levels, load_level, save_level, new_level, add_sublevel, remove_sublevel, set_streaming_method. |

### Editor Utilities (Editor)

| Tool | Description |
|------|-------------|
| `run_console_command` | Execute an Unreal console command in the editor. |
| `get_output_log` | Read the editor output log with optional category and verbosity filtering. |
| `capture_viewport` | Capture a screenshot of the active viewport. Returns MCP ImageContent (viewable by the LLM) or saves to file. Set `include_ui=true` to capture with Slate/UMG overlays (HUD, menus) — requires PIE. |
| `execute_script` | Execute a Python script in the editor's Python environment. |
| `live_compile` | Trigger Live Coding (hot reload) compilation without restarting the editor. |
| `pie_control` | Control Play In Editor (PIE) sessions: start, stop, or check status. Supports map override and Simulate In Editor mode. Start/stop are async — use status to verify. |
| `player_control` | Control player pawn and editor viewport camera. Operations: `get_info` (player state), `teleport` (move pawn), `set_rotation` (set view direction), `set_view_target` (change camera target) — require PIE. `get_camera`/`set_camera` move the editor viewport camera without PIE. |

### Components & Instancing (Editor)

| Tool | Description |
|------|-------------|
| `get_actor_components` | Get the full component hierarchy for an actor, including types, visibility, mesh references, and transforms. |
| `ism_ops` | Manage ISM/HISM components: create, add/remove/update instances, set materials, get instance counts. |

### Textures (Editor)

| Tool | Description |
|------|-------------|
| `texture_ops` | Import images as texture assets, query texture info, assign textures to material parameters, list textures. |

### DataTables (Editor)

| Tool | Description |
|------|-------------|
| `data_asset_ops` | Manage DataTables: list, read rows, add/update/delete rows, create tables, import from CSV. |

### Fab Marketplace (Editor)

| Tool | Description |
|------|-------------|
| `fab_ops` | Manage Fab marketplace asset cache: list downloaded assets, import into project, clear cache. |

### Subsystems & UI (Editor)

| Tool | Description |
|------|-------------|
| `subsystem_query` | List active UE subsystems by type (world, game_instance, engine, editor, local_player). |
| `ui_query` | Introspect Slate and UMG widget hierarchy: tree, find by class, get details, list bindings. |

### Network Debug (Editor)

| Tool | Description |
|------|-------------|
| `network_debug` | Introspect active HTTP requests, WebSocket connections, and recent request/response logs. |

### Documentation (Always Available)

| Tool | Description |
|------|-------------|
| `status` | Check server health, UE installation path, project info, and editor connectivity. |
| `lookup_docs` | Search UE 5.7 API docs, RealtimeMesh docs, and project docs by natural language query. |
| `lookup_class` | Get structured class reference (inheritance, properties, functions) for a specific UE class. |

## Documentation Index

Build the local documentation search index:

```bash
mcp-unreal --build-index
```

This indexes markdown files from `docs/ue5.7/` and `docs/realtimemesh/`, plus your project's `CLAUDE.md`. The index is stored at `./docs/index.bleve` (configurable via `MCP_UNREAL_DOCS_INDEX`).

To add custom documentation, place markdown files in the `docs/` directory and rebuild the index. See [docs/README.md](docs/README.md).

## Example Usage

Once registered, your AI agent can autonomously:

```
You: "Build the project, fix any errors, and run the tests"

Agent calls:
  1. status()                    → editor + plugin online
  2. build_project()             → finds 2 compile errors
  3. lookup_class("AActor")      → looks up the correct API
  4. (edits the C++ source)
  5. build_project()             → clean build
  6. run_tests(filter: "Game")   → 14/14 tests pass
```

```
You: "Add a Niagara fire effect to the torch actor"

Agent calls:
  1. get_level_actors(name_filter: "Torch")  → finds actor path
  2. lookup_class("UNiagaraComponent")       → checks API
  3. niagara_ops(operation: "spawn_system", system_path: "/Game/FX/NS_Fire", ...)
  4. niagara_ops(operation: "set_parameter", parameter_name: "SpawnRate", value: 50)
```

```
You: "Create a PCG scatter for placing rocks on the landscape"

Agent calls:
  1. lookup_docs("PCG scatter landscape")    → reads PCG docs
  2. pcg_ops(operation: "list_components")   → finds existing PCG actors
  3. spawn_actor(class: "Actor", name: "RockScatter")
  4. pcg_ops(operation: "add_node", node_type: "SurfaceSampler")
  5. pcg_ops(operation: "execute")           → generates rocks
```

## Development

```bash
# Build
make build

# Run all tests with race detector
make test

# Go lint
make lint

# C++ format check (requires clang-format)
make cpp-fmt-check

# Run all checks (Go format, C++ format, vet, lint, tests)
make check
```

### C++ Plugin Tests

The MCPUnreal plugin includes automation tests that run inside the UE editor headlessly. These require UE 5.7 installed locally and are **not** part of CI — they are opt-in for local development.

```bash
# Build the test project (first run ~2-5 min, incremental after)
make test-cpp-build

# Build and run all MCPUnreal.* automation tests
make test-cpp
```

This uses a minimal UE project in `test-project/` that symlinks the plugin source and runs all `MCPUnreal.*` tests via `UnrealEditor-Cmd -nullrhi`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup, coding standards, and how to add new tools.

## License

Apache-2.0 — see [LICENSE](LICENSE).
