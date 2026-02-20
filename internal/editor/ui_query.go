package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- ui_query ---

// WidgetInfo describes a Slate/UMG widget.
// Children is []any to avoid a recursive type cycle that the MCP SDK
// schema generator cannot handle.  Elements are WidgetInfo-shaped maps
// after JSON round-tripping.
type WidgetInfo struct {
	Type     string        `json:"type" jsonschema:"widget class name (e.g. SCheckBox, STextBlock)"`
	Name     string        `json:"name,omitempty" jsonschema:"widget debug name or ID"`
	Visible  bool          `json:"visible" jsonschema:"whether the widget is visible"`
	Enabled  bool          `json:"enabled" jsonschema:"whether the widget is enabled"`
	Bounds   *WidgetBounds `json:"bounds,omitempty" jsonschema:"screen-space bounds"`
	Children []any         `json:"children,omitempty" jsonschema:"child widgets (recursive WidgetInfo objects)"`
}

// WidgetBounds describes widget screen-space bounds.
type WidgetBounds struct {
	X      float64 `json:"x" jsonschema:"top-left X"`
	Y      float64 `json:"y" jsonschema:"top-left Y"`
	Width  float64 `json:"width" jsonschema:"width in pixels"`
	Height float64 `json:"height" jsonschema:"height in pixels"`
}

// UIQueryInput defines parameters for the ui_query tool.
type UIQueryInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: tree, find, get_widget, umg_list"`
	// For find: widget class filter.
	Class string `json:"class,omitempty" jsonschema:"Widget class name to search for (e.g. SCheckBox). For find."`
	// For get_widget: widget path.
	Path string `json:"path,omitempty" jsonschema:"Widget path in hierarchy. For get_widget."`
	// For tree: depth limit.
	MaxDepth int    `json:"max_depth,omitempty" jsonschema:"Maximum tree depth to return (default unlimited). For tree."`
	World    string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// UIQueryOutput is returned by the ui_query tool.
type UIQueryOutput struct {
	Widgets []WidgetInfo `json:"widgets,omitempty" jsonschema:"widget tree or search results"`
	Widget  *WidgetInfo  `json:"widget,omitempty" jsonschema:"single widget (for get_widget)"`
	Count   int          `json:"count,omitempty" jsonschema:"number of widgets found"`
}

// RegisterUIQuery adds the ui_query tool to the MCP server.
func (h *Handler) RegisterUIQuery(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "ui_query",
		Description: "Introspect Slate and UMG widget hierarchy. " +
			"Operations: tree (full widget tree with visibility/bounds), " +
			"find (search for widgets by class name), " +
			"get_widget (get details of a specific widget by path), " +
			"umg_list (list UMG UserWidget instances). " +
			"Supports max_depth to limit tree traversal. " +
			"Requires the editor running with the MCPUnreal plugin loaded (port 8090).",
	}, h.UIQuery)
}

// UIQuery implements the ui_query tool.
func (h *Handler) UIQuery(ctx context.Context, req *mcp.CallToolRequest, input UIQueryInput) (*mcp.CallToolResult, UIQueryOutput, error) {
	if input.Operation == "" {
		return nil, UIQueryOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.Class != "" {
		body["class"] = input.Class
	}
	if input.Path != "" {
		body["path"] = input.Path
	}
	if input.MaxDepth > 0 {
		body["max_depth"] = input.MaxDepth
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/ui/query", body)
	if err != nil {
		return nil, UIQueryOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out UIQueryOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, UIQueryOutput{}, fmt.Errorf("parsing UI query response: %w", err)
	}

	return nil, out, nil
}
