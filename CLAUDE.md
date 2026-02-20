# CLAUDE.md — mcp-unreal

> Model Context Protocol server for Unreal Engine 5.7. Single Go binary.
> Repository: https://github.com/remiphilippe/mcp-unreal
> License: Apache-2.0

---

## Project Overview

This is an MCP (Model Context Protocol) server written in Go that gives AI coding agents (Claude Code, Cursor, etc.) complete autonomous control over an Unreal Engine 5.7 project — builds, tests, editor manipulation, Blueprint editing, procedural mesh generation, and documentation lookup.

**Read `IMPLEMENTATION.md` before writing any code.** It contains the full architecture, tool inventory, communication paths, and design decisions.

### Key Facts

- **Language**: Go (primary), C++ (UE editor plugin only)
- **MCP SDK**: `github.com/modelcontextprotocol/go-sdk/mcp` (official, co-maintained with Google)
- **Transport**: stdio (JSON-RPC 2.0) — Claude Code launches the binary as a subprocess
- **Doc search**: `github.com/blevesearch/bleve/v2` (compiled into the binary, no external service)
- **Target UE version**: 5.7 (macOS primary, Windows and Linux secondary)
- **Go version**: 1.25+

---

## Repository Structure

```
mcp-unreal/
├── cmd/
│   └── mcp-unreal/
│       └── main.go                     # Entry point, CLI flags, tool registration
├── internal/
│   ├── config/
│   │   └── config.go                   # Environment config, project detection, path resolution
│   ├── headless/
│   │   ├── build.go                    # build_project, cook_project, generate_project_files
│   │   ├── build_test.go
│   │   ├── test.go                     # run_tests, run_visual_tests, list_tests
│   │   ├── test_test.go
│   │   └── log.go                      # get_test_log
│   ├── editor/
│   │   ├── client.go                   # HTTP client for RC API + plugin
│   │   ├── client_test.go
│   │   ├── actors.go                   # Actor CRUD tools
│   │   ├── actors_test.go
│   │   ├── properties.go              # set_property, get_property, call_function
│   │   ├── blueprints.go             # blueprint_query, blueprint_modify
│   │   ├── anim_blueprints.go        # anim_blueprint_query, anim_blueprint_modify
│   │   ├── assets.go                  # search_assets, get_asset_info
│   │   ├── materials.go              # material_ops
│   │   ├── characters.go             # character_config
│   │   ├── input.go                   # input_ops
│   │   ├── mesh.go                    # procedural_mesh, realtime_mesh
│   │   ├── levels.go                  # level_ops
│   │   └── utilities.go              # console cmd, output log, viewport, scripts
│   ├── docs/
│   │   ├── index.go                   # Bleve index open/create/query
│   │   ├── index_test.go
│   │   ├── lookup.go                  # lookup_docs, lookup_class tool handlers
│   │   ├── ingest.go                  # Markdown → bleve doc entries
│   │   └── class_parser.go           # Parse UE class reference markdown
│   └── status/
│       └── status.go                  # status tool, connectivity checks
├── docs/                               # Documentation source files for the index
│   ├── ue5.7/
│   ├── realtimemesh/
│   └── README.md                      # How to add/update doc entries
├── plugin/                             # UE 5.7 C++ editor plugin
│   ├── MCPUnreal.uplugin
│   ├── Source/
│   │   └── MCPUnreal/
│   │       ├── MCPUnreal.Build.cs
│   │       └── ...
│   └── README.md                      # Plugin build/install instructions
├── .github/
│   ├── workflows/
│   │   ├── ci.yml                     # Go lint + test + build
│   │   └── release.yml                # GoReleaser on tag push
│   └── ISSUE_TEMPLATE/
│       ├── bug_report.md
│       └── feature_request.md
├── IMPLEMENTATION.md                   # Full architecture document
├── CLAUDE.md                          # This file
├── CONTRIBUTING.md
├── LICENSE                            # Apache-2.0
├── README.md                          # User-facing: install, configure, use
├── SECURITY.md
├── go.mod
├── go.sum
├── Makefile
└── .goreleaser.yml
```

---

## Code Standards

### Go

- **Format**: `gofmt` / `goimports` on all files. No exceptions.
- **Lint**: `golangci-lint` with the config below. CI must pass.
- **Naming**: Follow Effective Go. Exported types use `PascalCase`. Unexported use `camelCase`. Acronyms keep case (`HTTP`, `URL`, `API`, not `Http`).
- **Errors**: Always wrap with context using `fmt.Errorf("doing X: %w", err)`. Never discard errors silently. Use sentinel errors in `internal/` packages only when callers need to match on them.
- **Logging**: Use `log/slog` (structured logging). Never `fmt.Println` for operational output. The MCP transport uses stdout — any stray prints corrupt the JSON-RPC stream.
- **Context**: All tool handlers receive `context.Context`. Propagate it to subprocess calls and HTTP requests. Respect cancellation.
- **Testing**: Every package under `internal/` must have `_test.go` files. Use table-driven tests. Use `testify/assert` only if the team agrees — standard library `testing` is fine.
- **Dependencies**: Minimize. Currently allowed:
  - `github.com/modelcontextprotocol/go-sdk` — MCP SDK
  - `github.com/blevesearch/bleve/v2` — Full-text search
  - Standard library for everything else (especially `net/http`, `os/exec`, `encoding/json`)
  - Do **not** add frameworks, DI containers, or ORM-like abstractions.

### golangci-lint config (`.golangci.yml`)

```yaml
linters:
  enable:
    - errcheck
    - govet
    - staticcheck
    - unused
    - gosimple
    - ineffassign
    - misspell
    - gofmt
    - goimports
    - revive
    - gosec
    - bodyclose

linters-settings:
  revive:
    rules:
      - name: exported
        arguments:
          - checkPrivateReceivers
      - name: blank-imports
      - name: context-as-argument
      - name: error-return
      - name: error-strings
      - name: increment-decrement

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
```

### C++ (UE Plugin)

- **Style**: Google C++ Style Guide, enforced by `clang-format`. Config in `plugin/.clang-format`.
- **Format**: `make cpp-fmt` to format in-place, `make cpp-fmt-check` for dry-run (CI + pre-commit hook).
- **Overrides from Google baseline**: 100-column limit (UE types are verbose), no include sorting (UE PCH order), namespace indentation on.
- **Lint (local-only)**: `clang-tidy` via `make cpp-tidy` (requires `compile_commands.json` from UBT). Config in `plugin/.clang-tidy`.
- **Static analysis (local-only)**: `cppcheck` via `make cpp-check`. Suppressions in `plugin/.cppcheck-suppressions.txt`.
- **CI**: Only `clang-format` runs in CI (`cpp-format` job). `clang-tidy` and `cppcheck` require UE headers unavailable on ubuntu-latest.
- **Naming**: Prefix classes `U` (UObject), `A` (Actor), `F` (structs), `E` (enums), `I` (interfaces), `T` (templates).
- Use `GENERATED_BODY()` in all UCLASS/USTRUCT.
- All HTTP route handlers must validate input JSON before acting on it.
- No raw `new`/`delete`. Use UE memory management (`NewObject`, `CreateDefaultSubobject`, smart pointers).
- Plugin module name: `MCPUnreal`. Log category: `LogMCPUnreal`.

---

## Security

### Critical Rules

1. **stdout is sacred.** The MCP protocol uses stdout for JSON-RPC messages. Any writes to stdout outside of the MCP SDK will corrupt the transport and break the connection. Use `slog` with an stderr handler for all logging.

2. **Never expose network ports from the Go binary.** The MCP server communicates exclusively via stdio. The Go binary does not open any listening sockets. It only makes outbound HTTP requests to localhost (RC API on 30010, plugin on 8090).

3. **Localhost only.** The editor plugin HTTP server must bind to `127.0.0.1`, never `0.0.0.0`. The Remote Control API defaults to localhost — do not change it. Document this clearly.

4. **Input validation.** Every tool handler must validate its input before passing to UE:
   - File paths: Reject path traversal (`..`, absolute paths outside project root). Use `filepath.Clean` and check prefix.
   - Object paths: Must match UE path format (`/Game/...`, `/Engine/...`). Reject shell metacharacters.
   - Console commands: Log all commands executed. The `run_console_command` tool description must warn that it can execute arbitrary UE commands.
   - Script execution: `execute_script` is the most dangerous tool. The plugin must log the full script before execution. Claude Code's built-in permission system provides user approval.

5. **No credentials in code or config.** This project does not require API keys, tokens, or auth. If future features need credentials (e.g., remote editor), use environment variables, never hardcoded values. Add a `.env.example` file, never a `.env` file.

6. **Dependency supply chain.** Run `go mod verify` in CI. Pin dependency versions in `go.mod` (no floating). Review new dependencies before adding — prefer standard library.

7. **Subprocess execution.** All `exec.Command` calls must:
   - Use explicit argument arrays, never shell expansion (`exec.Command("sh", "-c", userInput)` is forbidden)
   - Set a timeout via `context.WithTimeout`
   - Capture both stdout and stderr
   - Log the command being executed (at debug level)

8. **Sensitive data in logs.** Never log file contents, user code, or script bodies at levels above `Debug`. Default log level is `Info`.

### SECURITY.md

The repo must include a `SECURITY.md` with:
- Responsible disclosure instructions (GitHub Security Advisories)
- Scope: what is considered a vulnerability (RCE via MCP tools, path traversal, etc.)
- Out of scope: UE editor bugs, issues requiring local access (the whole system is local-only)

---

## Documentation Requirements

This project is open source. Documentation is not optional.

### README.md

Must include:
1. **One-line description** + badges (Go version, license, CI status)
2. **What this does** — 2-3 sentences explaining the MCP server concept
3. **Quick start** — `go install`, configure Claude Code, verify with `status`
4. **Prerequisites** — Go 1.25+, UE 5.7, Remote Control API plugin enabled
5. **Installation** — building from source, installing the UE plugin, registering with Claude Code
6. **Configuration** — environment variables table, `.mcp.json` example
7. **Available tools** — grouped table with name, description, mode (headless/editor)
8. **Documentation index** — how to build, how to add custom docs
9. **Architecture diagram** — ASCII art from IMPLEMENTATION.md
10. **Contributing** — link to CONTRIBUTING.md
11. **License** — Apache-2.0

### Code Documentation

- Every exported function and type must have a Go doc comment.
- Tool handlers: doc comment must describe what the tool does, what parameters it takes, and what it returns. This becomes the MCP tool description that Claude sees.
- Package-level doc comments in every package.
- Complex logic (log parsing, Blueprint JSON translation, mesh data serialization) must have inline comments explaining the UE-specific behavior.

### Tool Descriptions

MCP tool descriptions are critical — they're what Claude reads to decide which tool to use. Follow these rules:

```go
// GOOD: Specific, actionable, mentions prerequisites
mcp.AddTool(server, &mcp.Tool{
    Name:        "run_tests",
    Description: "Run Pelorus headless automation tests using UE 5.7 (-nullrhi). " +
        "Returns structured JSON with per-test pass/fail results and failure event details. " +
        "Does not require the editor to be running. " +
        "Call build_project first if you've edited C++ files.",
}, testRunner.RunTests)

// BAD: Vague, no context
mcp.AddTool(server, &mcp.Tool{
    Name:        "run_tests",
    Description: "Runs tests",
}, testRunner.RunTests)
```

- Tool names: `snake_case`, verb-first when possible (`run_tests`, `spawn_actor`, `lookup_docs`)
- Descriptions: 1-3 sentences. State what it does, when to use it, and any prerequisites.
- Parameter descriptions: Use `jsonschema` struct tags. Include defaults and valid values.

### CONTRIBUTING.md

Must include:
1. How to set up the dev environment
2. How to run tests (`go test ./...`)
3. How to run linting (`golangci-lint run`)
4. PR process (fork, branch, test, PR against `main`)
5. Commit message format (Conventional Commits: `feat:`, `fix:`, `docs:`, `chore:`)
6. Code review expectations
7. How to add a new tool (step-by-step)
8. How to add documentation entries to the index
9. CLA note if applicable

### Changelog

Use GitHub Releases with auto-generated release notes. Tag format: `v0.1.0`, `v0.2.0`, etc. Follow semver.

---

## Implementation Order

Follow this sequence. Each phase should be a PR (or small set of PRs) that passes CI.

### Phase 1: Project Skeleton
**Goal**: Binary compiles, registers with Claude Code, `status` tool works.

1. `go mod init github.com/remiphilippe/mcp-unreal`
2. Add MCP SDK dependency
3. `cmd/mcp-unreal/main.go` — CLI flags, server creation, stdio transport
4. `internal/config/config.go` — env var parsing, project root detection
5. `internal/status/status.go` — `status` tool (check UE path, editor connectivity, project file)
6. `Makefile` with `build`, `test`, `lint` targets
7. `.github/workflows/ci.yml` — Go test + lint on push/PR
8. `README.md` skeleton, `LICENSE`, `CONTRIBUTING.md`, `SECURITY.md`
9. `.golangci.yml`

**Verify**: `go build && claude mcp add mcp-unreal -- ./mcp-unreal` → `status` returns project info.

### Phase 2: Headless — Build & Test
**Goal**: Claude Code can build the project and run tests autonomously.

1. `internal/headless/build.go` — `build_project`, `generate_project_files`
2. `internal/headless/test.go` — `run_tests`, `list_tests` (port bash script log parsing to Go)
3. `internal/headless/log.go` — `get_test_log`
4. Tests for log parsing (`test_test.go` with sample UE log fixtures)
5. Register tools in `main.go`

**Verify**: Edit C++ → `build_project` → `run_tests` → see structured results.

### Phase 3: Documentation System
**Goal**: Claude Code can look up UE APIs before writing code.

1. `internal/docs/index.go` — Bleve index creation and querying
2. `internal/docs/ingest.go` — Parse markdown files into `DocEntry` structs, index them
3. `internal/docs/lookup.go` — `lookup_docs`, `lookup_class` tool handlers
4. `internal/docs/class_parser.go` — Parse structured class reference markdown
5. `docs/ue5.7/` — Curate initial set of 20-30 key class references
6. `docs/realtimemesh/` — RMC API docs
7. CLI flag `--build-index` to run indexing
8. Auto-index project `CLAUDE.md` at startup

**Verify**: `lookup_docs(query: "how to spawn an actor")` returns relevant AActor/UWorld docs.

### Phase 4: Editor Communication — Core
**Goal**: Claude Code can read/write actor properties and call functions in the running editor.

1. `internal/editor/client.go` — HTTP client with RC API and plugin endpoints
2. `internal/editor/properties.go` — `set_property`, `get_property`, `call_function` (RC API only, no plugin needed)
3. `internal/editor/actors.go` — `get_level_actors`, `spawn_actor`, `delete_actors`, `move_actor`
4. `internal/editor/utilities.go` — `run_console_command`
5. Tests with mock HTTP server

**Verify**: With editor running + RC API enabled, `get_level_actors` returns actors.

### Phase 5: UE Plugin
**Goal**: The C++ plugin exposes editor internals that RC API cannot.

1. `plugin/MCPUnreal.uplugin` — plugin descriptor
2. `plugin/Source/MCPUnreal/MCPUnreal.Build.cs` — module definition
3. `plugin/Source/MCPUnreal/MCPUnrealModule.cpp` — HTTP server startup on editor load
4. Actor routes (list, spawn, delete — beyond what RC API provides)
5. Blueprint routes (query graphs, modify nodes, connect pins)
6. Output log + viewport capture routes
7. `plugin/README.md` — build and install instructions

**Verify**: Plugin loads in UE 5.7, `http://localhost:8090/api/status` returns JSON.

### Phase 6: Blueprint & Animation Blueprint Tools
**Goal**: Full Blueprint editing from Claude Code.

1. `internal/editor/blueprints.go` — `blueprint_query`, `blueprint_modify`
2. `internal/editor/anim_blueprints.go` — `anim_blueprint_query`, `anim_blueprint_modify`
3. Corresponding C++ routes in plugin

**Verify**: Claude Code can create a Blueprint, add variables, add nodes, connect pins.

### Phase 7: Assets, Materials, Characters, Input, Levels
**Goal**: Remaining editor tools.

1. `internal/editor/assets.go`
2. `internal/editor/materials.go`
3. `internal/editor/characters.go`
4. `internal/editor/input.go`
5. `internal/editor/levels.go`
6. Corresponding C++ routes

### Phase 8: Mesh Tools
**Goal**: Procedural and RealtimeMesh support.

1. C++ plugin routes for ProceduralMeshComponent and RealtimeMeshComponent
2. `internal/editor/mesh.go` — `procedural_mesh`, `realtime_mesh`
3. Conditional compilation in plugin for RMC (optional dependency)

### Phase 9: Polish & Release
**Goal**: Production-ready open source release.

1. `run_visual_tests`, `cook_project`, `live_compile` tools
2. MCP async task support for long-running operations
3. `.goreleaser.yml` for cross-platform binary releases
4. Full README with architecture diagram, tool table, examples
5. Comprehensive CONTRIBUTING.md with "How to add a new tool" guide
6. GitHub Actions release workflow
7. Documentation index build automation

---

## Conventions

### Git

- Branch naming: `feat/tool-name`, `fix/issue-description`, `docs/topic`
- Commit messages: Conventional Commits
  - `feat: add run_tests tool with structured JSON output`
  - `fix: handle empty UE log file gracefully`
  - `docs: add AActor class reference to doc index`
  - `chore: update MCP SDK to v0.5.0`
  - `refactor: extract log parser into standalone function`
  - `test: add table-driven tests for Blueprint modify operations`
- One logical change per commit. Squash fixups before merge.
- `main` branch is always releasable. All work goes through PRs.

### MCP Tool Input/Output Types

Follow this pattern for all tools:

```go
// internal/editor/actors.go

// SpawnActorInput defines the parameters for the spawn_actor tool.
type SpawnActorInput struct {
    ClassName string     `json:"class_name" jsonschema:"required,UE class name (e.g. StaticMeshActor, PointLight, CameraActor)"`
    Name      string     `json:"name,omitempty" jsonschema:"Optional display name for the actor"`
    Location  [3]float64 `json:"location,omitempty" jsonschema:"[X,Y,Z] world position in centimeters"`
    Rotation  [3]float64 `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] in degrees"`
    Scale     [3]float64 `json:"scale,omitempty" jsonschema:"[X,Y,Z] scale factor. Default [1,1,1]"`
}

// SpawnActorOutput is returned by the spawn_actor tool.
type SpawnActorOutput struct {
    ActorPath string `json:"actor_path"`
    ActorName string `json:"actor_name"`
    Class     string `json:"class"`
}
```

- Input types: named `<ToolName>Input`
- Output types: named `<ToolName>Output`
- JSON tags: `snake_case`, match the MCP parameter names
- `jsonschema` tags: provide descriptions, defaults, valid values
- `omitempty` on optional fields
- Never use `interface{}` / `any` in output types — always define concrete structs

### Error Handling in Tools

Tools should return user-friendly errors, not stack traces:

```go
func (h *Handler) RunTests(ctx context.Context, req *mcp.CallToolRequest, input RunTestsInput) (*mcp.CallToolResult, RunTestsOutput, error) {
    if _, err := os.Stat(h.config.UEEditorPath); err != nil {
        return nil, RunTestsOutput{}, fmt.Errorf(
            "UnrealEditor-Cmd not found at %s — set UE_EDITOR_PATH env var or install UE 5.7",
            h.config.UEEditorPath,
        )
    }
    // ...
}
```

### HTTP Client Calls

All calls to the UE editor must handle the editor being offline:

```go
resp, err := h.client.PluginCall("/api/actors/list", input)
if err != nil {
    return nil, GetLevelActorsOutput{}, fmt.Errorf(
        "editor unreachable at %s — ensure UE is running with the MCPUnreal plugin loaded: %w",
        h.client.PluginURL(), err,
    )
}
```

---

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `UE_EDITOR_PATH` | Platform-dependent (see config.go) | Path to `UnrealEditor-Cmd` binary |
| `MCP_UNREAL_PROJECT` | Auto-detected from cwd | Path to `.uproject` file or project root |
| `RC_API_PORT` | `30010` | UE Remote Control API HTTP port |
| `PLUGIN_PORT` | `8090` | MCPUnreal editor plugin HTTP port |
| `MCP_UNREAL_LOG_LEVEL` | `info` | Log level: `debug`, `info`, `warn`, `error` |
| `MCP_UNREAL_DOCS_INDEX` | `./docs/index.bleve` | Path to bleve documentation index |

Platform defaults for `UE_EDITOR_PATH`:
- macOS: `/Users/Shared/Epic Games/UE_5.7/Engine/Binaries/Mac/UnrealEditor-Cmd`
- Windows: `C:\Program Files\Epic Games\UE_5.7\Engine\Binaries\Win64\UnrealEditor-Cmd.exe`
- Linux: `/opt/UnrealEngine/Engine/Binaries/Linux/UnrealEditor-Cmd`

---

## Testing

### Go Tests

```bash
# Run all tests
go test ./...

# Run with race detector
go test -race ./...

# Run specific package
go test ./internal/headless/...

# Verbose with coverage
go test -v -cover ./...
```

### Test Fixtures

Place sample UE log files in `internal/headless/testdata/`:
- `testdata/passing_tests.log` — All tests pass
- `testdata/failing_tests.log` — Some tests fail with events
- `testdata/empty.log` — Empty log (UE failed to start)
- `testdata/no_tests_found.log` — Zero tests matched filter

### Integration Tests

Tag integration tests that need UE running with `//go:build integration`:

```go
//go:build integration

package editor_test

func TestGetLevelActors_LiveEditor(t *testing.T) {
    // Only runs with: go test -tags integration ./internal/editor/...
}
```

### CI Pipeline

```yaml
# .github/workflows/ci.yml
name: CI
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.25'
      - run: go mod verify
      - run: go vet ./...
      - uses: golangci/golangci-lint-action@v6
      - run: go test -race -coverprofile=coverage.out ./...
      - uses: codecov/codecov-action@v4
        with:
          file: coverage.out

  cpp-format:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Check C++ formatting
        run: find plugin/Source -name '*.cpp' -o -name '*.h' | xargs clang-format --dry-run -Werror

  build:
    runs-on: ubuntu-latest
    needs: test
    strategy:
      matrix:
        goos: [darwin, linux, windows]
        goarch: [amd64, arm64]
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.25'
      - run: GOOS=${{ matrix.goos }} GOARCH=${{ matrix.goarch }} go build -o mcp-unreal-${{ matrix.goos }}-${{ matrix.goarch }} ./cmd/mcp-unreal
```

---

## Release Process

1. Update version in `main.go` (`const Version = "0.x.0"`)
2. Update CHANGELOG section in README or GitHub Release notes
3. Tag: `git tag v0.x.0 && git push --tags`
4. GoReleaser builds binaries for darwin/linux/windows × amd64/arm64
5. GitHub Release auto-created with binaries attached
6. Users install via `go install github.com/remiphilippe/mcp-unreal/cmd/mcp-unreal@latest`

---

## Open Source Checklist

Before the first public release, ensure:

- [ ] `LICENSE` file (Apache-2.0) present at repo root
- [ ] `README.md` with install, configure, use instructions
- [ ] `CONTRIBUTING.md` with dev setup and PR process
- [ ] `SECURITY.md` with responsible disclosure instructions
- [ ] `CODE_OF_CONDUCT.md` (Contributor Covenant)
- [ ] `.github/ISSUE_TEMPLATE/` with bug report and feature request templates
- [ ] `.github/PULL_REQUEST_TEMPLATE.md`
- [ ] CI passes on all PRs
- [ ] No hardcoded paths, credentials, or internal references
- [ ] All dependencies are permissively licensed (check with `go-licenses`)
- [ ] Binary builds for macOS (amd64, arm64), Linux (amd64, arm64), Windows (amd64)
- [ ] GoDoc comments on all exported symbols
- [ ] At least one example in README showing end-to-end usage
