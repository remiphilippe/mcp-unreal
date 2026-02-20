package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- material_ops ---

// MaterialOpsInput defines parameters for the material_ops tool.
type MaterialOpsInput struct {
	Operation    string          `json:"operation" jsonschema:"required,One of: create, set_parameter, get_parameters, set_texture, create_instance, list_parameters"`
	MaterialPath string          `json:"material_path,omitempty" jsonschema:"Material asset path — required for most operations"`
	ParentPath   string          `json:"parent_path,omitempty" jsonschema:"Parent material path for create_instance"`
	PackagePath  string          `json:"package_path,omitempty" jsonschema:"Package path for create (e.g. /Game/Materials)"`
	MaterialName string          `json:"material_name,omitempty" jsonschema:"Name for new material/instance"`
	Params       json.RawMessage `json:"params,omitempty" jsonschema:"Operation-specific parameters"`
	// Common parameter fields.
	ParameterName  string     `json:"parameter_name,omitempty" jsonschema:"Material parameter name"`
	ParameterValue any        `json:"parameter_value,omitempty" jsonschema:"Parameter value (type depends on parameter)"`
	TexturePath    string     `json:"texture_path,omitempty" jsonschema:"Texture asset path for set_texture"`
	Color          [4]float64 `json:"color,omitempty" jsonschema:"[R,G,B,A] for vector parameters (0-1 range)"`
}

// MaterialOpsOutput is returned by the material_ops tool.
type MaterialOpsOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterMaterials adds the material_ops tool to the MCP server.
func (h *Handler) RegisterMaterials(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "material_ops",
		Description: "Create and modify materials in the UE editor. Operations:\n" +
			"- create: Create a new material (requires package_path, material_name)\n" +
			"- create_instance: Create a material instance (requires parent_path, package_path, material_name)\n" +
			"- set_parameter: Set a scalar/vector parameter (requires material_path, parameter_name, parameter_value)\n" +
			"- set_texture: Set a texture parameter (requires material_path, parameter_name, texture_path)\n" +
			"- get_parameters: Get all parameter values (requires material_path)\n" +
			"- list_parameters: List available parameter names and types (requires material_path)\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.MaterialOps)
}

// MaterialOps implements the material_ops tool.
func (h *Handler) MaterialOps(ctx context.Context, req *mcp.CallToolRequest, input MaterialOpsInput) (*mcp.CallToolResult, MaterialOpsOutput, error) {
	if input.Operation == "" {
		return nil, MaterialOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.MaterialPath != "" {
		body["material_path"] = input.MaterialPath
	}
	if input.ParentPath != "" {
		body["parent_path"] = input.ParentPath
	}
	if input.PackagePath != "" {
		body["package_path"] = input.PackagePath
	}
	if input.MaterialName != "" {
		body["material_name"] = input.MaterialName
	}
	if input.ParameterName != "" {
		body["parameter_name"] = input.ParameterName
	}
	if input.ParameterValue != nil {
		body["parameter_value"] = input.ParameterValue
	}
	if input.TexturePath != "" {
		body["texture_path"] = input.TexturePath
	}
	if input.Color != [4]float64{} {
		body["color"] = input.Color
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

	resp, err := h.Client.PluginCall(ctx, "/api/materials/ops", body)
	if err != nil {
		return nil, MaterialOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, MaterialOpsOutput{}, fmt.Errorf("parsing material_ops response: %w", err)
	}

	return nil, MaterialOpsOutput{Result: result}, nil
}
