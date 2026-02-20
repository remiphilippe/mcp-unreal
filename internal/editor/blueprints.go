// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- blueprint_query ---

// BlueprintQueryInput defines parameters for the blueprint_query tool.
type BlueprintQueryInput struct {
	Operation        string `json:"operation" jsonschema:"required,One of: list, inspect, get_graph, get_node_types"`
	Path             string `json:"path,omitempty" jsonschema:"Blueprint asset path (e.g. /Game/Blueprints/BP_Player) — required for inspect and get_graph"`
	GraphName        string `json:"graph_name,omitempty" jsonschema:"Graph name for get_graph (from inspect results)"`
	StateMachineName string `json:"state_machine_name,omitempty" jsonschema:"State machine name for inspect_state_machine"`
}

// BlueprintQueryOutput is returned by the blueprint_query tool.
type BlueprintQueryOutput struct {
	Result any `json:"result" jsonschema:"Query results (structure depends on operation)"`
}

// --- blueprint_modify ---

// BlueprintModifyInput defines parameters for the blueprint_modify tool.
type BlueprintModifyInput struct {
	Operation     string          `json:"operation" jsonschema:"required,One of: create, add_variable, remove_variable, add_function, remove_function, add_node, delete_node, connect_pins, disconnect_pins, set_pin_value, compile"`
	BlueprintPath string          `json:"blueprint_path,omitempty" jsonschema:"Blueprint asset path — required for all operations except create"`
	Params        json.RawMessage `json:"params,omitempty" jsonschema:"Operation-specific parameters (see tool description for details)"`
	// Common fields forwarded directly.
	ClassName     string `json:"class_name,omitempty" jsonschema:"For create: parent class name (default Actor)"`
	PackagePath   string `json:"package_path,omitempty" jsonschema:"For create: package path (e.g. /Game/Blueprints)"`
	BlueprintName string `json:"blueprint_name,omitempty" jsonschema:"For create: new Blueprint name"`
	VariableName  string `json:"variable_name,omitempty" jsonschema:"For add_variable/remove_variable: variable name"`
	VariableType  string `json:"variable_type,omitempty" jsonschema:"For add_variable: type (bool, int, float, string, FVector, etc.)"`
	FunctionName  string `json:"function_name,omitempty" jsonschema:"For add_function/remove_function: function name"`
	NodeClass     string `json:"node_class,omitempty" jsonschema:"For add_node: K2Node class name"`
	NodeID        string `json:"node_id,omitempty" jsonschema:"For delete_node/set_pin_value: node GUID"`
	GraphName     string `json:"graph_name,omitempty" jsonschema:"Target graph name"`
	SourceNodeID  string `json:"source_node_id,omitempty" jsonschema:"For connect_pins: source node GUID"`
	SourcePinName string `json:"source_pin_name,omitempty" jsonschema:"For connect_pins: source pin name"`
	TargetNodeID  string `json:"target_node_id,omitempty" jsonschema:"For connect_pins: target node GUID"`
	TargetPinName string `json:"target_pin_name,omitempty" jsonschema:"For connect_pins: target pin name"`
	PinName       string `json:"pin_name,omitempty" jsonschema:"For set_pin_value/disconnect_pins: pin name"`
	PinValue      string `json:"pin_value,omitempty" jsonschema:"For set_pin_value: value to set"`
}

// BlueprintModifyOutput is returned by the blueprint_modify tool.
type BlueprintModifyOutput struct {
	Success  bool   `json:"success" jsonschema:"whether the operation succeeded"`
	Compiled bool   `json:"compiled" jsonschema:"whether the Blueprint was compiled after the operation"`
	Path     string `json:"path,omitempty" jsonschema:"path to created Blueprint (for create operation)"`
	NodeID   string `json:"node_id,omitempty" jsonschema:"GUID of created node (for add_node operation)"`
	Message  string `json:"message,omitempty" jsonschema:"additional info about the operation"`
}

// RegisterBlueprints adds the blueprint_query and blueprint_modify tools to the MCP server.
func (h *Handler) RegisterBlueprints(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "blueprint_query",
		Description: "Query Blueprint assets in the UE editor. Operations:\n" +
			"- list: List all Blueprint assets in the project\n" +
			"- inspect: Get variables, functions, and event graphs of a Blueprint (requires path)\n" +
			"- get_graph: Serialize a graph's nodes, pins, and connections (requires path + graph_name)\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.BlueprintQuery)

	mcp.AddTool(server, &mcp.Tool{
		Name: "blueprint_modify",
		Description: "Modify Blueprints in the UE editor. Operations:\n" +
			"- create: Create a new Blueprint (requires blueprint_name, package_path; optional class_name)\n" +
			"- add_variable: Add a variable (requires blueprint_path, variable_name, variable_type)\n" +
			"- remove_variable: Remove a variable (requires blueprint_path, variable_name)\n" +
			"- add_function: Add a custom function (requires blueprint_path, function_name)\n" +
			"- remove_function: Remove a custom function (requires blueprint_path, function_name)\n" +
			"- add_node: Add a node to a graph (requires blueprint_path, node_class, graph_name)\n" +
			"- delete_node: Delete a node (requires blueprint_path, node_id, graph_name)\n" +
			"- connect_pins: Connect two pins (requires blueprint_path, source_node_id, source_pin_name, target_node_id, target_pin_name)\n" +
			"- disconnect_pins: Disconnect a pin (requires blueprint_path, node_id, pin_name)\n" +
			"- set_pin_value: Set a pin's default value (requires blueprint_path, node_id, pin_name, pin_value)\n" +
			"- compile: Force-compile the Blueprint (requires blueprint_path)\n" +
			"Auto-compiles after mutations. Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.BlueprintModify)
}

// BlueprintQuery implements the blueprint_query tool.
func (h *Handler) BlueprintQuery(ctx context.Context, req *mcp.CallToolRequest, input BlueprintQueryInput) (*mcp.CallToolResult, BlueprintQueryOutput, error) {
	if input.Operation == "" {
		return nil, BlueprintQueryOutput{}, fmt.Errorf("operation is required (list, inspect, get_graph)")
	}

	var endpoint string
	body := map[string]any{}

	switch input.Operation {
	case "list":
		endpoint = "/api/blueprints/list"
	case "inspect":
		if input.Path == "" {
			return nil, BlueprintQueryOutput{}, fmt.Errorf("path is required for inspect operation")
		}
		endpoint = "/api/blueprints/inspect"
		body["blueprint_path"] = input.Path
	case "get_graph":
		if input.Path == "" {
			return nil, BlueprintQueryOutput{}, fmt.Errorf("path is required for get_graph operation")
		}
		if input.GraphName == "" {
			return nil, BlueprintQueryOutput{}, fmt.Errorf("graph_name is required for get_graph operation")
		}
		endpoint = "/api/blueprints/get_graph"
		body["blueprint_path"] = input.Path
		body["graph_name"] = input.GraphName
	default:
		return nil, BlueprintQueryOutput{}, fmt.Errorf("unknown operation %q — use list, inspect, or get_graph", input.Operation)
	}

	resp, err := h.Client.PluginCall(ctx, endpoint, body)
	if err != nil {
		return nil, BlueprintQueryOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, BlueprintQueryOutput{}, fmt.Errorf("parsing blueprint_query response: %w", err)
	}

	return nil, BlueprintQueryOutput{Result: result}, nil
}

// BlueprintModify implements the blueprint_modify tool.
func (h *Handler) BlueprintModify(ctx context.Context, req *mcp.CallToolRequest, input BlueprintModifyInput) (*mcp.CallToolResult, BlueprintModifyOutput, error) {
	if input.Operation == "" {
		return nil, BlueprintModifyOutput{}, fmt.Errorf("operation is required")
	}

	// Build the request body, forwarding all fields.
	body := map[string]any{
		"operation": input.Operation,
	}
	if input.BlueprintPath != "" {
		body["blueprint_path"] = input.BlueprintPath
	}
	if input.ClassName != "" {
		body["parent_class"] = input.ClassName
	}
	if input.PackagePath != "" {
		body["package_path"] = input.PackagePath
	}
	if input.BlueprintName != "" {
		body["name"] = input.BlueprintName
	}
	if input.VariableName != "" {
		body["variable_name"] = input.VariableName
	}
	if input.VariableType != "" {
		body["variable_type"] = input.VariableType
	}
	if input.FunctionName != "" {
		body["function_name"] = input.FunctionName
	}
	if input.NodeClass != "" {
		body["node_class"] = input.NodeClass
	}
	if input.NodeID != "" {
		body["node_id"] = input.NodeID
	}
	if input.GraphName != "" {
		body["graph_name"] = input.GraphName
	}
	if input.SourceNodeID != "" {
		body["source_node_id"] = input.SourceNodeID
	}
	if input.SourcePinName != "" {
		body["source_pin"] = input.SourcePinName
	}
	if input.TargetNodeID != "" {
		body["target_node_id"] = input.TargetNodeID
	}
	if input.TargetPinName != "" {
		body["target_pin"] = input.TargetPinName
	}
	if input.PinName != "" {
		body["pin_name"] = input.PinName
	}
	if input.PinValue != "" {
		body["value"] = input.PinValue
	}

	// Merge any extra params from the raw JSON field.
	if len(input.Params) > 0 {
		var extra map[string]any
		if err := json.Unmarshal(input.Params, &extra); err == nil {
			for k, v := range extra {
				body[k] = v
			}
		}
	}

	resp, err := h.Client.PluginCall(ctx, "/api/blueprints/modify", body)
	if err != nil {
		return nil, BlueprintModifyOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out BlueprintModifyOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, BlueprintModifyOutput{}, fmt.Errorf("parsing modify response: %w", err)
	}

	return nil, out, nil
}
