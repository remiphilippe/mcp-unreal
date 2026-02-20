// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- pcg_ops ---

// PCGOpsInput defines parameters for the pcg_ops tool.
type PCGOpsInput struct {
	Operation      string          `json:"operation" jsonschema:"required,One of: execute, cleanup, get_graph_info, set_parameter, add_node, connect_nodes, remove_node"`
	ActorPath      string          `json:"actor_path,omitempty" jsonschema:"Actor with UPCGComponent — required for execute, cleanup, set_parameter"`
	GraphPath      string          `json:"graph_path,omitempty" jsonschema:"PCG graph asset path — required for get_graph_info, add_node, connect_nodes, remove_node"`
	ParameterName  string          `json:"parameter_name,omitempty" jsonschema:"User parameter name for set_parameter"`
	ParameterValue any             `json:"parameter_value,omitempty" jsonschema:"User parameter value for set_parameter"`
	NodeType       string          `json:"node_type,omitempty" jsonschema:"PCG settings class name for add_node (e.g. PCGSurfaceSamplerSettings)"`
	NodeLabel      string          `json:"node_label,omitempty" jsonschema:"Optional display label for add_node"`
	NodeID         string          `json:"node_id,omitempty" jsonschema:"Source node FName for connect_nodes or remove_node"`
	TargetNodeID   string          `json:"target_node_id,omitempty" jsonschema:"Target node FName for connect_nodes"`
	SourcePinLabel string          `json:"source_pin_label,omitempty" jsonschema:"Source pin label for connect_nodes (defaults to first output pin)"`
	TargetPinLabel string          `json:"target_pin_label,omitempty" jsonschema:"Target pin label for connect_nodes (defaults to first input pin)"`
	Params         json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
	World          string          `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// PCGOpsOutput is returned by the pcg_ops tool.
type PCGOpsOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterPCG adds the pcg_ops tool to the MCP server.
func (h *Handler) RegisterPCG(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "pcg_ops",
		Description: "Manage Procedural Content Generation (PCG) graphs and components in the UE editor. Operations:\n" +
			"- execute: Run PCG generation on an actor's PCGComponent (requires actor_path)\n" +
			"- cleanup: Remove generated PCG output from an actor (requires actor_path)\n" +
			"- get_graph_info: List nodes and edges in a PCG graph (requires graph_path)\n" +
			"- set_parameter: Set a user parameter on a PCGComponent (requires actor_path, parameter_name, parameter_value)\n" +
			"- add_node: Add a node to a PCG graph by settings class name (requires graph_path, node_type)\n" +
			"- connect_nodes: Connect two nodes in a PCG graph (requires graph_path, node_id, target_node_id)\n" +
			"- remove_node: Remove a node from a PCG graph (requires graph_path, node_id)\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090). PCG plugin must be enabled in the project.",
	}, h.PCGOps)
}

// PCGOps implements the pcg_ops tool.
func (h *Handler) PCGOps(ctx context.Context, req *mcp.CallToolRequest, input PCGOpsInput) (*mcp.CallToolResult, PCGOpsOutput, error) {
	if input.Operation == "" {
		return nil, PCGOpsOutput{}, fmt.Errorf("operation is required (execute, cleanup, get_graph_info, set_parameter, add_node, connect_nodes, remove_node)")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.ActorPath != "" {
		body["actor_path"] = input.ActorPath
	}
	if input.GraphPath != "" {
		body["graph_path"] = input.GraphPath
	}
	if input.ParameterName != "" {
		body["parameter_name"] = input.ParameterName
	}
	if input.ParameterValue != nil {
		body["parameter_value"] = input.ParameterValue
	}
	if input.NodeType != "" {
		body["node_type"] = input.NodeType
	}
	if input.NodeLabel != "" {
		body["node_label"] = input.NodeLabel
	}
	if input.NodeID != "" {
		body["node_id"] = input.NodeID
	}
	if input.TargetNodeID != "" {
		body["target_node_id"] = input.TargetNodeID
	}
	if input.SourcePinLabel != "" {
		body["source_pin_label"] = input.SourcePinLabel
	}
	if input.TargetPinLabel != "" {
		body["target_pin_label"] = input.TargetPinLabel
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

	resp, err := h.Client.PluginCall(ctx, "/api/pcg/ops", body)
	if err != nil {
		return nil, PCGOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, PCGOpsOutput{}, fmt.Errorf("parsing pcg_ops response: %w", err)
	}

	return nil, PCGOpsOutput{Result: result}, nil
}
