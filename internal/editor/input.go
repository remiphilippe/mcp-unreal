// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- input_ops ---

// InputOpsInput defines parameters for the input_ops tool.
type InputOpsInput struct {
	Operation   string          `json:"operation" jsonschema:"required,One of: list_actions, list_contexts, add_action, remove_action, add_context, bind_action, unbind_action, get_bindings"`
	AssetPath   string          `json:"asset_path,omitempty" jsonschema:"Input Action or Mapping Context asset path"`
	ActionName  string          `json:"action_name,omitempty" jsonschema:"Input Action name for create/bind/unbind"`
	ContextName string          `json:"context_name,omitempty" jsonschema:"Input Mapping Context name"`
	ValueType   string          `json:"value_type,omitempty" jsonschema:"For add_action: value type (bool, float, Vector2D, Vector3D)"`
	Key         string          `json:"key,omitempty" jsonschema:"For bind_action: input key name (e.g. W, SpaceBar, LeftMouseButton, Gamepad_LeftThumbstick_X)"`
	Modifiers   []string        `json:"modifiers,omitempty" jsonschema:"For bind_action: modifier names (Negate, Swizzle, DeadZone, etc.)"`
	Triggers    []string        `json:"triggers,omitempty" jsonschema:"For bind_action: trigger names (Pressed, Released, Hold, Tap, etc.)"`
	PackagePath string          `json:"package_path,omitempty" jsonschema:"Package path for create operations"`
	Params      json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
}

// InputOpsOutput is returned by the input_ops tool.
type InputOpsOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterInput adds the input_ops tool to the MCP server.
func (h *Handler) RegisterInput(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "input_ops",
		Description: "Manage UE 5.7 Enhanced Input system — Input Actions, Mapping Contexts, and key bindings. Operations:\n" +
			"- list_actions: List all Input Action assets\n" +
			"- list_contexts: List all Input Mapping Context assets\n" +
			"- add_action: Create a new Input Action (requires action_name, value_type, package_path)\n" +
			"- remove_action: Delete an Input Action (requires asset_path)\n" +
			"- add_context: Create a new Input Mapping Context (requires context_name, package_path)\n" +
			"- bind_action: Bind a key to an action in a context (requires asset_path for context, action_name, key)\n" +
			"- unbind_action: Remove a key binding (requires asset_path for context, action_name)\n" +
			"- get_bindings: Get all bindings in a context (requires asset_path)\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.InputOps)
}

// InputOps implements the input_ops tool.
func (h *Handler) InputOps(ctx context.Context, req *mcp.CallToolRequest, input InputOpsInput) (*mcp.CallToolResult, InputOpsOutput, error) {
	if input.Operation == "" {
		return nil, InputOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.AssetPath != "" {
		body["asset_path"] = input.AssetPath
	}
	if input.ActionName != "" {
		body["action_name"] = input.ActionName
	}
	if input.ContextName != "" {
		body["context_name"] = input.ContextName
	}
	if input.ValueType != "" {
		body["value_type"] = input.ValueType
	}
	if input.Key != "" {
		body["key"] = input.Key
	}
	if len(input.Modifiers) > 0 {
		body["modifiers"] = input.Modifiers
	}
	if len(input.Triggers) > 0 {
		body["triggers"] = input.Triggers
	}
	if input.PackagePath != "" {
		body["package_path"] = input.PackagePath
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

	resp, err := h.Client.PluginCall(ctx, "/api/input/ops", body)
	if err != nil {
		return nil, InputOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, InputOpsOutput{}, fmt.Errorf("parsing input_ops response: %w", err)
	}

	return nil, InputOpsOutput{Result: result}, nil
}
