package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- anim_blueprint_query ---

// AnimBlueprintQueryInput defines parameters for the anim_blueprint_query tool.
type AnimBlueprintQueryInput struct {
	Operation        string `json:"operation" jsonschema:"required,One of: list_state_machines, inspect_state_machine"`
	BlueprintPath    string `json:"blueprint_path" jsonschema:"required,Animation Blueprint asset path (e.g. /Game/Animations/ABP_Character)"`
	StateMachineName string `json:"state_machine_name,omitempty" jsonschema:"State machine name for inspect_state_machine"`
}

// AnimBlueprintQueryOutput is returned by the anim_blueprint_query tool.
type AnimBlueprintQueryOutput struct {
	Result any `json:"result" jsonschema:"Query results (structure depends on operation)"`
}

// --- anim_blueprint_modify ---

// AnimBlueprintModifyInput defines parameters for the anim_blueprint_modify tool.
type AnimBlueprintModifyInput struct {
	Operation        string `json:"operation" jsonschema:"required,One of: create_state_machine, delete_state_machine, rename_state_machine, set_entry_state, create_state, delete_state, rename_state, create_transition, delete_transition, add_anim_node, delete_anim_node"`
	BlueprintPath    string `json:"blueprint_path" jsonschema:"required,Animation Blueprint asset path"`
	StateMachineName string `json:"state_machine_name,omitempty" jsonschema:"State machine name (required for most operations)"`
	StateName        string `json:"state_name,omitempty" jsonschema:"State name for create_state/delete_state/rename_state/set_entry_state"`
	OldName          string `json:"old_name,omitempty" jsonschema:"Current name for rename operations"`
	NewName          string `json:"new_name,omitempty" jsonschema:"New name for rename operations"`
	FromState        string `json:"from_state,omitempty" jsonschema:"Source state for create_transition"`
	ToState          string `json:"to_state,omitempty" jsonschema:"Target state for create_transition"`
	TransitionID     string `json:"transition_id,omitempty" jsonschema:"Transition node ID for delete_transition"`
	NodeClass        string `json:"node_class,omitempty" jsonschema:"Anim node class for add_anim_node"`
	NodeID           string `json:"node_id,omitempty" jsonschema:"Node ID for delete_anim_node"`
}

// AnimBlueprintModifyOutput is returned by the anim_blueprint_modify tool.
type AnimBlueprintModifyOutput struct {
	Success  bool `json:"success" jsonschema:"whether the operation succeeded"`
	Compiled bool `json:"compiled" jsonschema:"whether the AnimBP was compiled after the operation"`
}

// RegisterAnimBlueprints adds the anim_blueprint_query and anim_blueprint_modify tools to the MCP server.
func (h *Handler) RegisterAnimBlueprints(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "anim_blueprint_query",
		Description: "Query Animation Blueprint state machines in the UE editor. Operations:\n" +
			"- list_state_machines: List all state machines with state/transition counts\n" +
			"- inspect_state_machine: Get states and transitions for a specific state machine\n" +
			"Requires blueprint_path. Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.AnimBlueprintQuery)

	mcp.AddTool(server, &mcp.Tool{
		Name: "anim_blueprint_modify",
		Description: "Modify Animation Blueprint state machines in the UE editor. Operations:\n" +
			"- create_state_machine: Create a new state machine\n" +
			"- delete_state_machine: Delete a state machine\n" +
			"- rename_state_machine: Rename a state machine (requires old_name, new_name)\n" +
			"- set_entry_state: Set the entry state (requires state_machine_name, state_name)\n" +
			"- create_state: Add a state (requires state_machine_name, state_name)\n" +
			"- delete_state: Remove a state (requires state_machine_name, state_name)\n" +
			"- rename_state: Rename a state (requires state_machine_name, old_name, new_name)\n" +
			"- create_transition: Add a transition (requires state_machine_name, from_state, to_state)\n" +
			"- delete_transition: Remove a transition (requires state_machine_name, transition_id)\n" +
			"- add_anim_node: Add an animation node (requires state_machine_name, node_class)\n" +
			"- delete_anim_node: Remove an animation node (requires state_machine_name, node_id)\n" +
			"Auto-compiles after modifications. Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.AnimBlueprintModify)
}

// AnimBlueprintQuery implements the anim_blueprint_query tool.
func (h *Handler) AnimBlueprintQuery(ctx context.Context, req *mcp.CallToolRequest, input AnimBlueprintQueryInput) (*mcp.CallToolResult, AnimBlueprintQueryOutput, error) {
	if input.Operation == "" {
		return nil, AnimBlueprintQueryOutput{}, fmt.Errorf("operation is required (list_state_machines, inspect_state_machine)")
	}
	if input.BlueprintPath == "" {
		return nil, AnimBlueprintQueryOutput{}, fmt.Errorf("blueprint_path is required")
	}

	body := map[string]any{
		"operation":      input.Operation,
		"blueprint_path": input.BlueprintPath,
	}
	if input.StateMachineName != "" {
		body["state_machine_name"] = input.StateMachineName
	}

	resp, err := h.Client.PluginCall(ctx, "/api/anim_blueprints/query", body)
	if err != nil {
		return nil, AnimBlueprintQueryOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, AnimBlueprintQueryOutput{}, fmt.Errorf("parsing anim_blueprint_query response: %w", err)
	}

	return nil, AnimBlueprintQueryOutput{Result: result}, nil
}

// AnimBlueprintModify implements the anim_blueprint_modify tool.
func (h *Handler) AnimBlueprintModify(ctx context.Context, req *mcp.CallToolRequest, input AnimBlueprintModifyInput) (*mcp.CallToolResult, AnimBlueprintModifyOutput, error) {
	if input.Operation == "" {
		return nil, AnimBlueprintModifyOutput{}, fmt.Errorf("operation is required")
	}
	if input.BlueprintPath == "" {
		return nil, AnimBlueprintModifyOutput{}, fmt.Errorf("blueprint_path is required")
	}

	body := map[string]any{
		"operation":      input.Operation,
		"blueprint_path": input.BlueprintPath,
	}
	if input.StateMachineName != "" {
		body["state_machine_name"] = input.StateMachineName
	}
	if input.StateName != "" {
		body["state_name"] = input.StateName
	}
	if input.OldName != "" {
		body["old_name"] = input.OldName
	}
	if input.NewName != "" {
		body["new_name"] = input.NewName
	}
	if input.FromState != "" {
		body["from_state"] = input.FromState
	}
	if input.ToState != "" {
		body["to_state"] = input.ToState
	}
	if input.TransitionID != "" {
		body["transition_id"] = input.TransitionID
	}
	if input.NodeClass != "" {
		body["node_class"] = input.NodeClass
	}
	if input.NodeID != "" {
		body["node_id"] = input.NodeID
	}

	resp, err := h.Client.PluginCall(ctx, "/api/anim_blueprints/modify", body)
	if err != nil {
		return nil, AnimBlueprintModifyOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out AnimBlueprintModifyOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, AnimBlueprintModifyOutput{}, fmt.Errorf("parsing modify response: %w", err)
	}

	return nil, out, nil
}
