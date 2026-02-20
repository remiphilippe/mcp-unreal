#!/usr/bin/env bash
set -euo pipefail

# Read JSON input from stdin
INPUT=$(cat)
COMMAND=$(echo "$INPUT" | jq -r '.tool_input.command // ""')

# Only gate git commit commands
if [[ "$COMMAND" != *"git commit"* ]]; then
  exit 0
fi

cd "$CLAUDE_PROJECT_DIR"
go test ./... 2>&1 || { echo "Tests failed — fix before committing" >&2; exit 2; }
golangci-lint run 2>&1 || { echo "Lint failed — fix before committing" >&2; exit 2; }
if command -v clang-format >/dev/null 2>&1; then
  make cpp-fmt-check 2>&1 || { echo "C++ format check failed — run 'make cpp-fmt' to fix" >&2; exit 2; }
fi
exit 0
