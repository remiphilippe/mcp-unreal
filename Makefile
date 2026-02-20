BINARY = mcp-unreal
VERSION ?= 0.1.0
LDFLAGS = -ldflags "-X main.Version=$(VERSION)"

# macOS defaults for UE paths
UE_EDITOR_PATH ?= /Users/Shared/Epic Games/UE_5.7/Engine/Binaries/Mac/UnrealEditor-Cmd
UE_BUILD_SCRIPT ?= /Users/Shared/Epic Games/UE_5.7/Engine/Build/BatchFiles/Mac/Build.sh

.PHONY: build install test test-race test-cover test-cover-html test-integration vet lint fmt fmt-check check clean build-index plugin-build plugin-install cpp-fmt cpp-fmt-check cpp-tidy cpp-check test-cpp-build test-cpp

build:
	go build $(LDFLAGS) -o $(BINARY) ./cmd/mcp-unreal

install:
	go install $(LDFLAGS) ./cmd/mcp-unreal

test:
	go test -race ./...

test-race: test

test-cover:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -1

test-cover-html:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-integration:
	go test -race -tags integration ./...

vet:
	go vet ./...

lint:
	golangci-lint run

fmt:
	gofmt -w .
	goimports -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "gofmt needed on:"; gofmt -l .; exit 1)

check: fmt-check cpp-fmt-check vet lint test-cover

clean:
	rm -f $(BINARY)
	rm -rf docs/index.bleve
	rm -f coverage.out coverage.html

# Phase 3: build the documentation search index.
build-index: build
	./$(BINARY) --build-index \
		--docs-index ./docs/index.bleve

# Build the UE project (compiles the MCPUnreal plugin).
# Requires UE_EDITOR_PATH and MCP_UNREAL_PROJECT env vars.
plugin-build:
	@test -n "$(UE_EDITOR_PATH)" || (echo "Error: set UE_EDITOR_PATH to the path to UnrealEditor-Cmd"; exit 1)
	@test -n "$(MCP_UNREAL_PROJECT)" || (echo "Error: set MCP_UNREAL_PROJECT to the path to your .uproject file"; exit 1)
	"$(UE_EDITOR_PATH)" "$(MCP_UNREAL_PROJECT)" \
		-build -TargetType=Editor \
		-platform=Mac -configuration=Development

# Copy plugin source into a UE project's Plugins/ directory.
# Requires MCP_UNREAL_PROJECT env var.
plugin-install:
	@test -n "$(MCP_UNREAL_PROJECT)" || (echo "Error: set MCP_UNREAL_PROJECT to the path to your .uproject file"; exit 1)
	@PROJ_DIR="$$(dirname "$(MCP_UNREAL_PROJECT)")" && \
	mkdir -p "$$PROJ_DIR/Plugins/MCPUnreal" && \
	cp -r plugin/* "$$PROJ_DIR/Plugins/MCPUnreal/" && \
	echo "Plugin installed to $$PROJ_DIR/Plugins/MCPUnreal/"

# ---------------------------------------------------------------------------
# C++ plugin tests (requires UE 5.7 installed)
# ---------------------------------------------------------------------------

# Build the test project (compiles MCPUnreal plugin + tests module).
test-cpp-build:
	@mkdir -p test-project/Plugins
	@test -L test-project/Plugins/MCPUnreal || ln -s ../../plugin test-project/Plugins/MCPUnreal
	"$(UE_BUILD_SCRIPT)" MCPTestProjectEditor Mac Development \
		-project="$(shell pwd)/test-project/MCPTestProject.uproject"

# Build and run all MCPUnreal.* automation tests headlessly.
TEST_LOG := $(shell pwd)/test-project/Saved/Logs/test-run.log
test-cpp: test-cpp-build
	@rm -f "$(TEST_LOG)"
	"$(UE_EDITOR_PATH)" "$(shell pwd)/test-project/MCPTestProject.uproject" \
		-ExecCmds="Automation RunTests MCPUnreal;Quit" \
		-nullrhi -nopause -unattended -nosplash -nosound \
		-ABSLOG="$(TEST_LOG)" \
		-testexit="Automation Test Queue Empty" \
		> /dev/null 2>&1 || true
	@sleep 1
	@if [ ! -s "$(TEST_LOG)" ]; then echo "ERROR: No log output captured." >&2; exit 1; fi
	@echo "--- Test Results ---"
	@grep -E 'LogAutomationController|LogAutomationCommandLine' "$(TEST_LOG)" || echo "(no automation output found)"
	@if grep -q 'Test Completed. Result={Fail' "$(TEST_LOG)"; then echo "FAIL: Some tests failed" >&2; exit 1; fi
	@echo "PASS: All tests passed"

# ---------------------------------------------------------------------------
# C++ static analysis (plugin)
# ---------------------------------------------------------------------------

CPP_SOURCES = $(shell find plugin/Source -name '*.cpp' -o -name '*.h')

# Format all C++ plugin files in-place.
cpp-fmt:
	clang-format -i $(CPP_SOURCES)

# Dry-run format check (CI-friendly, exits non-zero on violations).
cpp-fmt-check:
	clang-format --dry-run -Werror $(CPP_SOURCES)

# Run clang-tidy (local-only, requires compile_commands.json from UBT).
cpp-tidy:
	@test -f compile_commands.json || (echo "Error: compile_commands.json not found — generate it with UBT first"; exit 1)
	clang-tidy $(CPP_SOURCES) -p .

# Run cppcheck with UE-specific suppressions (local-only).
cpp-check:
	cppcheck --enable=warning,style,performance \
		--suppressions-list=plugin/.cppcheck-suppressions.txt \
		--error-exitcode=1 \
		$(CPP_SOURCES)
