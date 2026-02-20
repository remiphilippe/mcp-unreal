package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- get_actor_components ---

// GetActorComponentsInput defines parameters for the get_actor_components tool.
type GetActorComponentsInput struct {
	ActorPath         string `json:"actor_path,omitempty" jsonschema:"Full object path of the actor"`
	ActorName         string `json:"actor_name,omitempty" jsonschema:"Display name of the actor"`
	IncludeTransforms bool   `json:"include_transforms,omitempty" jsonschema:"Include relative transform for each component. Default false."`
	World             string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// ComponentTransform describes a component's relative transform.
type ComponentTransform struct {
	Location [3]float64 `json:"location" jsonschema:"[X,Y,Z] relative position"`
	Rotation [3]float64 `json:"rotation" jsonschema:"[Pitch,Yaw,Roll] relative rotation in degrees"`
	Scale    [3]float64 `json:"scale" jsonschema:"[X,Y,Z] relative scale"`
}

// ComponentInfo describes a single scene component and its children.
// Children is typed as []any to avoid a recursive type cycle that the
// MCP SDK schema generator cannot handle.  At runtime the elements are
// ComponentInfo-shaped maps produced by json.Unmarshal.
type ComponentInfo struct {
	Name          string              `json:"name" jsonschema:"component name"`
	Class         string              `json:"class" jsonschema:"UE class name"`
	Visible       bool                `json:"visible" jsonschema:"whether the component is visible"`
	InstanceCount *int                `json:"instance_count,omitempty" jsonschema:"number of instances (ISM/HISM components only)"`
	Mesh          string              `json:"mesh,omitempty" jsonschema:"mesh asset path (mesh components only)"`
	Transform     *ComponentTransform `json:"transform,omitempty" jsonschema:"relative transform (if include_transforms is true)"`
	Children      []any               `json:"children,omitempty" jsonschema:"child scene components (recursive ComponentInfo objects)"`
}

// NonSceneComponentInfo describes a non-scene component (no transform/hierarchy).
type NonSceneComponentInfo struct {
	Name     string `json:"name" jsonschema:"component name"`
	Class    string `json:"class" jsonschema:"UE class name"`
	IsActive bool   `json:"is_active" jsonschema:"whether the component is active"`
}

// GetActorComponentsOutput is returned by the get_actor_components tool.
type GetActorComponentsOutput struct {
	Actor              string                  `json:"actor" jsonschema:"actor display name"`
	Class              string                  `json:"class" jsonschema:"actor UE class name"`
	Path               string                  `json:"path" jsonschema:"actor full object path"`
	Components         []ComponentInfo         `json:"components,omitempty" jsonschema:"scene component hierarchy"`
	NonSceneComponents []NonSceneComponentInfo `json:"non_scene_components,omitempty" jsonschema:"non-scene components (no transform)"`
	TotalComponents    int                     `json:"total_components" jsonschema:"total number of components on the actor"`
}

// RegisterComponents adds the component introspection tool to the MCP server.
func (h *Handler) RegisterComponents(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "get_actor_components",
		Description: "Get the full component hierarchy for a specific actor, including component types, " +
			"visibility, mesh references, ISM instance counts, and optionally transforms. " +
			"Requires the editor to be running with the MCPUnreal plugin loaded (port 8090). " +
			"Find actors first with get_level_actors, then pass the actor_path or actor_name here.",
	}, h.GetActorComponents)
}

// GetActorComponents implements the get_actor_components tool.
func (h *Handler) GetActorComponents(ctx context.Context, req *mcp.CallToolRequest, input GetActorComponentsInput) (*mcp.CallToolResult, GetActorComponentsOutput, error) {
	if input.ActorPath == "" && input.ActorName == "" {
		return nil, GetActorComponentsOutput{}, fmt.Errorf("either actor_path or actor_name is required")
	}

	body := map[string]any{
		"include_transforms": input.IncludeTransforms,
	}
	if input.ActorPath != "" {
		body["actor_path"] = input.ActorPath
	}
	if input.ActorName != "" {
		body["actor_name"] = input.ActorName
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/actors/components", body)
	if err != nil {
		return nil, GetActorComponentsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out GetActorComponentsOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, GetActorComponentsOutput{}, fmt.Errorf("parsing component response: %w", err)
	}

	return nil, out, nil
}
