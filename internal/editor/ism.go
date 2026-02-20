package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- ism_ops ---

// ISMTransform describes an instance transform for ISM operations.
type ISMTransform struct {
	Location [3]float64 `json:"location" jsonschema:"[X,Y,Z] world position in centimeters"`
	Rotation [3]float64 `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] in degrees"`
	Scale    [3]float64 `json:"scale,omitempty" jsonschema:"[X,Y,Z] scale factor"`
}

// ISMOpsInput defines parameters for the ism_ops tool.
type ISMOpsInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: create, add_instances, clear_instances, get_instance_count, update_instance, remove_instance, set_material"`
	// For create: actor to add the ISM component to.
	ActorPath string `json:"actor_path,omitempty" jsonschema:"Actor object path (for create operation)"`
	ActorName string `json:"actor_name,omitempty" jsonschema:"Actor display name (for create operation)"`
	// For create: mesh and material to use.
	Mesh     string `json:"mesh,omitempty" jsonschema:"Static mesh asset path (e.g. /Engine/BasicShapes/Cube). For create."`
	Material string `json:"material,omitempty" jsonschema:"Material asset path. For create or set_material."`
	// For operations on existing ISM: component name.
	ComponentName string `json:"component_name,omitempty" jsonschema:"ISM component name on the actor. Required for add_instances, clear_instances, etc."`
	// For add_instances: batch of transforms.
	Transforms []ISMTransform `json:"transforms,omitempty" jsonschema:"Array of instance transforms (for add_instances)"`
	// For update_instance / remove_instance: instance index.
	InstanceIndex *int          `json:"instance_index,omitempty" jsonschema:"Instance index (for update_instance, remove_instance)"`
	Transform     *ISMTransform `json:"transform,omitempty" jsonschema:"Single transform (for update_instance)"`
	// For create: use HISM instead of ISM.
	UseHISM bool   `json:"use_hism,omitempty" jsonschema:"Use HierarchicalInstancedStaticMeshComponent instead of ISM. Default false."`
	World   string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// ISMOpsOutput is returned by the ism_ops tool.
type ISMOpsOutput struct {
	Success       bool   `json:"success" jsonschema:"whether the operation succeeded"`
	ComponentName string `json:"component_name,omitempty" jsonschema:"ISM component name"`
	InstanceCount int    `json:"instance_count,omitempty" jsonschema:"current instance count after operation"`
	AddedCount    int    `json:"added_count,omitempty" jsonschema:"number of instances added (for add_instances)"`
}

// RegisterISM adds the ism_ops tool to the MCP server.
func (h *Handler) RegisterISM(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "ism_ops",
		Description: "Manage InstancedStaticMesh (ISM) and HierarchicalInstancedStaticMesh (HISM) components. " +
			"Operations: create (add ISM to actor), add_instances (batch add transforms), " +
			"clear_instances, get_instance_count, update_instance, remove_instance, set_material. " +
			"Requires the editor to be running with the MCPUnreal plugin loaded (port 8090). " +
			"Use for efficient rendering of thousands of identical objects (vegetation, markers, debris).",
	}, h.ISMOps)
}

// ISMOps implements the ism_ops tool.
func (h *Handler) ISMOps(ctx context.Context, req *mcp.CallToolRequest, input ISMOpsInput) (*mcp.CallToolResult, ISMOpsOutput, error) {
	if input.Operation == "" {
		return nil, ISMOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.ActorPath != "" {
		body["actor_path"] = input.ActorPath
	}
	if input.ActorName != "" {
		body["actor_name"] = input.ActorName
	}
	if input.Mesh != "" {
		body["mesh"] = input.Mesh
	}
	if input.Material != "" {
		body["material"] = input.Material
	}
	if input.ComponentName != "" {
		body["component_name"] = input.ComponentName
	}
	if len(input.Transforms) > 0 {
		body["transforms"] = input.Transforms
	}
	if input.InstanceIndex != nil {
		body["instance_index"] = *input.InstanceIndex
	}
	if input.Transform != nil {
		body["transform"] = input.Transform
	}
	if input.UseHISM {
		body["use_hism"] = true
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/ism/ops", body)
	if err != nil {
		return nil, ISMOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out ISMOpsOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, ISMOpsOutput{}, fmt.Errorf("parsing ISM response: %w", err)
	}

	return nil, out, nil
}
