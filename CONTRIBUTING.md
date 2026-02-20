# Contributing to mcp-unreal

Thank you for your interest in contributing! This document covers the development workflow, coding standards, and how to add new features.

## Development Setup

### Prerequisites

- **Go 1.25+** — [install](https://go.dev/dl/)
- **golangci-lint** — `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
- **clang-format** — for C++ formatting (`brew install clang-format` on macOS, included on Ubuntu)
- **Unreal Engine 5.7** — for integration testing (not required for Go-only work)

### Build and Test

```bash
# Build the binary
make build

# Run all tests with race detector
make test

# Go lint
make lint

# C++ format check
make cpp-fmt-check

# Run all checks (Go format, C++ format, vet, lint, tests)
make check

# Clean build artifacts
make clean

# C++ plugin tests (requires UE 5.7 installed, not run in CI)
make test-cpp
```

### Project Structure

See `IMPLEMENTATION.md` for the full architecture. Key directories:

- `cmd/mcp-unreal/` — Entry point, CLI flags, tool registration
- `internal/config/` — Environment config, project detection
- `internal/status/` — Status tool, health checks
- `internal/headless/` — Build, test, cook (subprocess-based)
- `internal/editor/` — Editor communication via HTTP (RC API + plugin)
- `internal/docs/` — Documentation index and lookup tools
- `plugin/` — UE 5.7 C++ editor plugin

## Pull Request Process

1. **Fork** the repository and create a branch: `feat/tool-name`, `fix/issue-description`, `docs/topic`
2. **Write tests** for new functionality
3. **Run checks**: `make check` (or individually: `make test && make lint && make cpp-fmt-check`)
4. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/):
   - `feat: add run_tests tool with structured JSON output`
   - `fix: handle empty UE log file gracefully`
   - `docs: add AActor class reference to doc index`
   - `chore: update MCP SDK to v0.5.0`
   - `refactor: extract log parser into standalone function`
   - `test: add table-driven tests for Blueprint modify operations`
5. **Open a PR** against `main` with a clear description of what and why
6. One logical change per commit — squash fixups before merge

## Code Standards

### Go

- **Format**: `gofmt` / `goimports` on all files
- **Lint**: `golangci-lint` must pass (see `.golangci.yml`)
- **Errors**: Wrap with context — `fmt.Errorf("doing X: %w", err)`
- **Logging**: `log/slog` to stderr only. Never `fmt.Println` (stdout is the MCP transport)
- **Testing**: Table-driven tests in `_test.go` files. Every `internal/` package needs tests

### C++ (UE Plugin)

- **Style**: Google C++ Style Guide, enforced by `clang-format` (config: `plugin/.clang-format`)
- **Format**: `make cpp-fmt` to format in-place, `make cpp-fmt-check` to verify (CI-enforced)
- **Lint (local-only)**: `make cpp-tidy` (requires `compile_commands.json`) and `make cpp-check`
- Prefix classes: `U` (UObject), `A` (Actor), `F` (struct), `E` (enum), `I` (interface)
- Validate all HTTP input JSON before acting on it
- No raw `new`/`delete` — use UE memory management

## How to Add a New Tool

1. **Choose the right package**: `internal/headless/` for subprocess tools, `internal/editor/` for HTTP tools, `internal/docs/` for doc tools
2. **Define input/output types** following the naming convention:
   ```go
   type MyToolInput struct {
       Param string `json:"param" jsonschema:"description of this parameter"`
   }
   type MyToolOutput struct {
       Result string `json:"result" jsonschema:"description of the result"`
   }
   ```
3. **Implement the handler** with signature:
   ```go
   func (h *Handler) MyTool(ctx context.Context, req *mcp.CallToolRequest, input MyToolInput) (*mcp.CallToolResult, MyToolOutput, error)
   ```
4. **Register** in `cmd/mcp-unreal/main.go`:
   ```go
   mcp.AddTool(server, &mcp.Tool{
       Name:        "my_tool",
       Description: "Clear, specific description of what this tool does.",
   }, handler.MyTool)
   ```
5. **Write tests** — table-driven, cover success and error cases
6. **Update IMPLEMENTATION.md** if adding a new tool to the inventory

## How to Add Documentation Entries

1. Create a markdown file in the appropriate `docs/` subdirectory
2. Use structured headings for class references (name, parent, properties, functions)
3. Rebuild the index: `make build-index`
4. Test with `lookup_docs` and `lookup_class` tools

## Code Review

- All PRs require at least one review
- CI must pass (tests, lint, build)
- Keep PRs focused — one feature or fix per PR
- Respond to review feedback promptly

## Questions?

Open a [discussion](https://github.com/remiphilippe/mcp-unreal/discussions) or file an issue.
