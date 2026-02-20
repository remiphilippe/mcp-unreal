---
name: go-mcp-tool
description: How to add a new MCP tool to mcp-unreal
---
## Adding a New Tool

1. Create input/output types in the appropriate `internal/` package
2. Implement the handler method on the package's Handler struct
3. Register in `cmd/mcp-unreal/main.go` with descriptive tool name + description
4. Add tests with table-driven patterns
5. Update README.md tool table
6. Run `go test ./... && golangci-lint run`

## Template
```go
type MyToolInput struct {
    Param string `json:"param" jsonschema:"required,Description here"`
}
type MyToolOutput struct {
    Result string `json:"result"`
}
func (h *Handler) MyTool(ctx context.Context, req *mcp.CallToolRequest, input MyToolInput) (*mcp.CallToolResult, MyToolOutput, error) {
    // validate, execute, return
}
```
