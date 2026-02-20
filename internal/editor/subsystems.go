package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- subsystem_query ---

// SubsystemInfo describes a UE subsystem.
type SubsystemInfo struct {
	Class       string `json:"class" jsonschema:"C++ class name"`
	Type        string `json:"type" jsonschema:"subsystem type: world, game_instance, engine, editor, local_player"`
	Initialized bool   `json:"initialized" jsonschema:"whether the subsystem is currently initialized"`
}

// SubsystemQueryInput defines parameters for the subsystem_query tool.
type SubsystemQueryInput struct {
	Type  string `json:"type" jsonschema:"required,Subsystem type to query: world, game_instance, engine, editor, local_player, or all"`
	World string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// SubsystemQueryOutput is returned by the subsystem_query tool.
type SubsystemQueryOutput struct {
	Subsystems []SubsystemInfo `json:"subsystems" jsonschema:"list of active subsystems"`
	Count      int             `json:"count" jsonschema:"number of subsystems found"`
}

// RegisterSubsystems adds the subsystem_query tool to the MCP server.
func (h *Handler) RegisterSubsystems(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "subsystem_query",
		Description: "List active UE subsystems by type. " +
			"Types: world (UWorldSubsystem), game_instance (UGameInstanceSubsystem), " +
			"engine (UEngineSubsystem), editor (UEditorSubsystem), " +
			"local_player (ULocalPlayerSubsystem), or 'all' for everything. " +
			"Returns class names and initialization status. " +
			"Useful for understanding modular architecture and debugging initialization order. " +
			"Requires the editor running with the MCPUnreal plugin loaded (port 8090).",
	}, h.SubsystemQuery)
}

// SubsystemQuery implements the subsystem_query tool.
func (h *Handler) SubsystemQuery(ctx context.Context, req *mcp.CallToolRequest, input SubsystemQueryInput) (*mcp.CallToolResult, SubsystemQueryOutput, error) {
	if input.Type == "" {
		return nil, SubsystemQueryOutput{}, fmt.Errorf("type is required (world, game_instance, engine, editor, local_player, or all)")
	}

	body := map[string]any{
		"type": input.Type,
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/subsystems/query", body)
	if err != nil {
		return nil, SubsystemQueryOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out SubsystemQueryOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, SubsystemQueryOutput{}, fmt.Errorf("parsing subsystem query response: %w", err)
	}

	return nil, out, nil
}
