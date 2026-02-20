// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- level_ops ---

// LevelOpsInput defines parameters for the level_ops tool.
type LevelOpsInput struct {
	Operation   string          `json:"operation" jsonschema:"required,One of: get_current, list_levels, load_level, save_level, new_level, add_sublevel, remove_sublevel, set_streaming_method"`
	LevelPath   string          `json:"level_path,omitempty" jsonschema:"Level asset path (e.g. /Game/Maps/MainLevel)"`
	LevelName   string          `json:"level_name,omitempty" jsonschema:"Level name for new_level"`
	PackagePath string          `json:"package_path,omitempty" jsonschema:"Package path for new_level (e.g. /Game/Maps)"`
	Template    string          `json:"template,omitempty" jsonschema:"Template for new_level (Default, TimeOfDay, VR-Basic)"`
	Streaming   string          `json:"streaming,omitempty" jsonschema:"For set_streaming_method: AlwaysLoaded, Blueprint, or Distance"`
	Params      json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
	World       string          `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// LevelOpsOutput is returned by the level_ops tool.
type LevelOpsOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterLevels adds the level_ops tool to the MCP server.
func (h *Handler) RegisterLevels(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "level_ops",
		Description: "Manage UE levels (maps) in the editor. Operations:\n" +
			"- get_current: Get the current persistent level info\n" +
			"- list_levels: List all level assets in the project\n" +
			"- load_level: Open a level in the editor (requires level_path)\n" +
			"- save_level: Save the current level\n" +
			"- new_level: Create a new level (requires level_name, package_path; optional template)\n" +
			"- add_sublevel: Add a streaming sublevel (requires level_path)\n" +
			"- remove_sublevel: Remove a streaming sublevel (requires level_path)\n" +
			"- set_streaming_method: Set sublevel streaming (requires level_path, streaming: AlwaysLoaded/Blueprint/Distance)\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.LevelOps)
}

// LevelOps implements the level_ops tool.
func (h *Handler) LevelOps(ctx context.Context, req *mcp.CallToolRequest, input LevelOpsInput) (*mcp.CallToolResult, LevelOpsOutput, error) {
	if input.Operation == "" {
		return nil, LevelOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.LevelPath != "" {
		body["level_path"] = input.LevelPath
	}
	if input.LevelName != "" {
		body["level_name"] = input.LevelName
	}
	if input.PackagePath != "" {
		body["package_path"] = input.PackagePath
	}
	if input.Template != "" {
		body["template"] = input.Template
	}
	if input.Streaming != "" {
		body["streaming"] = input.Streaming
	}

	if input.World != "" {
		body["world"] = input.World
	}

	// Merge extra params.
	if len(input.Params) > 0 {
		var extra map[string]any
		if err := json.Unmarshal(input.Params, &extra); err == nil {
			for k, v := range extra {
				body[k] = v
			}
		}
	}

	resp, err := h.Client.PluginCall(ctx, "/api/levels/ops", body)
	if err != nil {
		return nil, LevelOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, LevelOpsOutput{}, fmt.Errorf("parsing level_ops response: %w", err)
	}

	return nil, LevelOpsOutput{Result: result}, nil
}
