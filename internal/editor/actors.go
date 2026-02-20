// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- get_level_actors ---

// GetLevelActorsInput defines parameters for the get_level_actors tool.
type GetLevelActorsInput struct {
	ClassFilter string `json:"class_filter,omitempty" jsonschema:"Filter by UE class name (e.g. StaticMeshActor, PointLight)"`
	NameFilter  string `json:"name_filter,omitempty" jsonschema:"Filter by actor display name substring"`
	TagFilter   string `json:"tag_filter,omitempty" jsonschema:"Filter by actor tag"`
	World       string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// ActorInfo describes a single actor in the level.
type ActorInfo struct {
	Name     string     `json:"name" jsonschema:"actor display name"`
	Class    string     `json:"class" jsonschema:"UE class name"`
	Path     string     `json:"path" jsonschema:"full object path"`
	Location [3]float64 `json:"location" jsonschema:"[X,Y,Z] world position"`
	Rotation [3]float64 `json:"rotation" jsonschema:"[Pitch,Yaw,Roll] in degrees"`
	Scale    [3]float64 `json:"scale" jsonschema:"[X,Y,Z] scale factor"`
}

// GetLevelActorsOutput is returned by the get_level_actors tool.
type GetLevelActorsOutput struct {
	Actors []ActorInfo `json:"actors" jsonschema:"list of actors in the level"`
	Total  int         `json:"total" jsonschema:"total number of actors returned"`
}

// --- spawn_actor ---

// SpawnActorInput defines parameters for the spawn_actor tool.
type SpawnActorInput struct {
	ClassName string     `json:"class_name" jsonschema:"required,UE class name (e.g. StaticMeshActor, PointLight, CameraActor)"`
	Name      string     `json:"name,omitempty" jsonschema:"Optional display name for the actor"`
	Location  [3]float64 `json:"location,omitempty" jsonschema:"[X,Y,Z] world position in centimeters"`
	Rotation  [3]float64 `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] in degrees"`
	Scale     [3]float64 `json:"scale,omitempty" jsonschema:"[X,Y,Z] scale factor — default [1,1,1]"`
	World     string     `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// SpawnActorOutput is returned by the spawn_actor tool.
type SpawnActorOutput struct {
	ActorPath string `json:"actor_path" jsonschema:"full object path of the spawned actor"`
	ActorName string `json:"actor_name" jsonschema:"display name of the spawned actor"`
	Class     string `json:"class" jsonschema:"UE class of the spawned actor"`
}

// --- delete_actors ---

// DeleteActorsInput defines parameters for the delete_actors tool.
type DeleteActorsInput struct {
	ActorPaths []string `json:"actor_paths,omitempty" jsonschema:"Object paths of actors to delete"`
	ActorNames []string `json:"actor_names,omitempty" jsonschema:"Display names of actors to delete"`
	World      string   `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// DeleteActorsOutput is returned by the delete_actors tool.
type DeleteActorsOutput struct {
	DeletedCount int      `json:"deleted_count" jsonschema:"number of actors deleted"`
	Deleted      []string `json:"deleted,omitempty" jsonschema:"names of deleted actors"`
}

// --- move_actor ---

// MoveActorInput defines parameters for the move_actor tool.
type MoveActorInput struct {
	ObjectPath string      `json:"object_path" jsonschema:"required,Full object path of the actor to move"`
	Location   *[3]float64 `json:"location,omitempty" jsonschema:"[X,Y,Z] world position in centimeters"`
	Rotation   *[3]float64 `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] in degrees"`
	Scale      *[3]float64 `json:"scale,omitempty" jsonschema:"[X,Y,Z] scale factor"`
}

// MoveActorOutput is returned by the move_actor tool.
type MoveActorOutput struct {
	Success    bool       `json:"success" jsonschema:"whether the transform was applied"`
	ObjectPath string     `json:"object_path" jsonschema:"the actor that was moved"`
	Location   [3]float64 `json:"location,omitempty" jsonschema:"new location if set"`
	Rotation   [3]float64 `json:"rotation,omitempty" jsonschema:"new rotation if set"`
	Scale      [3]float64 `json:"scale,omitempty" jsonschema:"new scale if set"`
}

// RegisterActors adds the actor CRUD tools to the MCP server.
func (h *Handler) RegisterActors(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "get_level_actors",
		Description: "List all actors in the current level, filterable by class, name, or tag. " +
			"Requires the editor to be running with the MCPUnreal plugin loaded (port 8090). " +
			"Returns actor names, classes, object paths, and transforms. " +
			"Supports world parameter to target PIE game world or editor world (default: auto).",
	}, h.GetLevelActors)

	mcp.AddTool(server, &mcp.Tool{
		Name: "spawn_actor",
		Description: "Spawn a new actor in the current level by UE class name with an optional transform. " +
			"Requires the editor to be running with the MCPUnreal plugin loaded (port 8090). " +
			"Scale defaults to [1,1,1] if not provided.",
	}, h.SpawnActor)

	mcp.AddTool(server, &mcp.Tool{
		Name: "delete_actors",
		Description: "Delete one or more actors from the current level by object path or display name. " +
			"Requires the editor to be running with the MCPUnreal plugin loaded (port 8090). " +
			"Supports batch deletion by providing multiple paths or names.",
	}, h.DeleteActors)

	mcp.AddTool(server, &mcp.Tool{
		Name: "move_actor",
		Description: "Set the location, rotation, and/or scale of an actor via the Remote Control API. " +
			"Requires the editor to be running with the Remote Control API enabled (port 30010). " +
			"Only the provided transform components are modified — omit fields to leave them unchanged.",
	}, h.MoveActor)
}

// GetLevelActors implements the get_level_actors tool.
func (h *Handler) GetLevelActors(ctx context.Context, req *mcp.CallToolRequest, input GetLevelActorsInput) (*mcp.CallToolResult, GetLevelActorsOutput, error) {
	body := map[string]any{}
	if input.ClassFilter != "" {
		body["class_filter"] = input.ClassFilter
	}
	if input.NameFilter != "" {
		body["name_filter"] = input.NameFilter
	}
	if input.TagFilter != "" {
		body["tag_filter"] = input.TagFilter
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/actors/list", body)
	if err != nil {
		return nil, GetLevelActorsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var actors []ActorInfo
	if err := json.Unmarshal(resp, &actors); err != nil {
		return nil, GetLevelActorsOutput{}, fmt.Errorf("parsing actor list response: %w", err)
	}

	return nil, GetLevelActorsOutput{Actors: actors, Total: len(actors)}, nil
}

// SpawnActor implements the spawn_actor tool.
func (h *Handler) SpawnActor(ctx context.Context, req *mcp.CallToolRequest, input SpawnActorInput) (*mcp.CallToolResult, SpawnActorOutput, error) {
	if input.ClassName == "" {
		return nil, SpawnActorOutput{}, fmt.Errorf("class_name is required")
	}

	// Default scale to [1,1,1] when not provided (zero-value means unset).
	scale := input.Scale
	if scale == [3]float64{} {
		scale = [3]float64{1, 1, 1}
	}

	body := map[string]any{
		"class_name": input.ClassName,
		"location":   input.Location,
		"rotation":   input.Rotation,
		"scale":      scale,
	}
	if input.Name != "" {
		body["name"] = input.Name
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/actors/spawn", body)
	if err != nil {
		return nil, SpawnActorOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out SpawnActorOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, SpawnActorOutput{}, fmt.Errorf("parsing spawn response: %w", err)
	}

	return nil, out, nil
}

// DeleteActors implements the delete_actors tool.
func (h *Handler) DeleteActors(ctx context.Context, req *mcp.CallToolRequest, input DeleteActorsInput) (*mcp.CallToolResult, DeleteActorsOutput, error) {
	if len(input.ActorPaths) == 0 && len(input.ActorNames) == 0 {
		return nil, DeleteActorsOutput{}, fmt.Errorf("at least one of actor_paths or actor_names is required")
	}

	body := map[string]any{}
	if len(input.ActorPaths) > 0 {
		body["actor_paths"] = input.ActorPaths
	}
	if len(input.ActorNames) > 0 {
		body["actor_names"] = input.ActorNames
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/actors/delete", body)
	if err != nil {
		return nil, DeleteActorsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out DeleteActorsOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, DeleteActorsOutput{}, fmt.Errorf("parsing delete response: %w", err)
	}

	return nil, out, nil
}

// MoveActor implements the move_actor tool. It uses the RC API function call
// endpoint to set actor transforms via K2_SetActorLocation, K2_SetActorRotation,
// and SetActorScale3D, only modifying the components that are provided.
func (h *Handler) MoveActor(ctx context.Context, req *mcp.CallToolRequest, input MoveActorInput) (*mcp.CallToolResult, MoveActorOutput, error) {
	if input.ObjectPath == "" {
		return nil, MoveActorOutput{}, fmt.Errorf("object_path is required")
	}
	if input.Location == nil && input.Rotation == nil && input.Scale == nil {
		return nil, MoveActorOutput{}, fmt.Errorf("at least one of location, rotation, or scale must be provided")
	}

	out := MoveActorOutput{
		Success:    true,
		ObjectPath: input.ObjectPath,
	}

	if input.Location != nil {
		body := map[string]any{
			"objectPath":   input.ObjectPath,
			"functionName": "K2_SetActorLocation",
			"parameters": map[string]any{
				"NewLocation": map[string]float64{
					"X": input.Location[0], "Y": input.Location[1], "Z": input.Location[2],
				},
				"bSweep":    false,
				"bTeleport": true,
			},
		}
		if _, err := h.Client.RCAPICall(ctx, "/remote/object/call", body); err != nil {
			return nil, MoveActorOutput{}, fmt.Errorf("setting location on %s: %w", input.ObjectPath, err)
		}
		out.Location = *input.Location
	}

	if input.Rotation != nil {
		body := map[string]any{
			"objectPath":   input.ObjectPath,
			"functionName": "K2_SetActorRotation",
			"parameters": map[string]any{
				"NewRotation": map[string]float64{
					"Pitch": input.Rotation[0], "Yaw": input.Rotation[1], "Roll": input.Rotation[2],
				},
				"bTeleportPhysics": true,
			},
		}
		if _, err := h.Client.RCAPICall(ctx, "/remote/object/call", body); err != nil {
			return nil, MoveActorOutput{}, fmt.Errorf("setting rotation on %s: %w", input.ObjectPath, err)
		}
		out.Rotation = *input.Rotation
	}

	if input.Scale != nil {
		body := map[string]any{
			"objectPath":   input.ObjectPath,
			"functionName": "SetActorScale3D",
			"parameters": map[string]any{
				"NewScale3D": map[string]float64{
					"X": input.Scale[0], "Y": input.Scale[1], "Z": input.Scale[2],
				},
			},
		}
		if _, err := h.Client.RCAPICall(ctx, "/remote/object/call", body); err != nil {
			return nil, MoveActorOutput{}, fmt.Errorf("setting scale on %s: %w", input.ObjectPath, err)
		}
		out.Scale = *input.Scale
	}

	return nil, out, nil
}
