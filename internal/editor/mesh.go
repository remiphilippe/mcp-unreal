// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- procedural_mesh ---

// ProceduralMeshInput defines parameters for the procedural_mesh tool.
type ProceduralMeshInput struct {
	Operation    string          `json:"operation" jsonschema:"required,One of: create_section, update_section, clear, set_material"`
	ActorPath    string          `json:"actor_path,omitempty" jsonschema:"ProceduralMeshActor object path — required for update/clear/set_material"`
	ActorName    string          `json:"actor_name,omitempty" jsonschema:"Display name for the actor (used with create_section to spawn)"`
	SectionIndex int             `json:"section_index,omitempty" jsonschema:"Mesh section index (default 0)"`
	Vertices     [][3]float64    `json:"vertices,omitempty" jsonschema:"Array of [X,Y,Z] vertex positions"`
	Triangles    []int           `json:"triangles,omitempty" jsonschema:"Array of triangle vertex indices (3 per triangle)"`
	Normals      [][3]float64    `json:"normals,omitempty" jsonschema:"Array of [X,Y,Z] vertex normals"`
	UVs          [][2]float64    `json:"uvs,omitempty" jsonschema:"Array of [U,V] texture coordinates"`
	Colors       [][4]float64    `json:"colors,omitempty" jsonschema:"Array of [R,G,B,A] vertex colors (0-1 range)"`
	MaterialPath string          `json:"material_path,omitempty" jsonschema:"Material asset path for set_material"`
	Location     [3]float64      `json:"location,omitempty" jsonschema:"[X,Y,Z] spawn location for create_section"`
	Params       json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
	World        string          `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// ProceduralMeshOutput is returned by the procedural_mesh tool.
type ProceduralMeshOutput struct {
	Result any `json:"result" jsonschema:"Operation results"`
}

// --- realtime_mesh ---

// RealtimeMeshInput defines parameters for the realtime_mesh tool.
type RealtimeMeshInput struct {
	Operation          string          `json:"operation" jsonschema:"required,One of: create_lod, create_section_group, create_section, update_mesh_data, set_material_slot, setup_collision"`
	ActorPath          string          `json:"actor_path,omitempty" jsonschema:"RealtimeMeshActor object path"`
	ActorName          string          `json:"actor_name,omitempty" jsonschema:"Display name for the actor"`
	LODIndex           int             `json:"lod_index,omitempty" jsonschema:"LOD level index"`
	ScreenSize         float64         `json:"screen_size,omitempty" jsonschema:"LOD screen size threshold (0-1)"`
	SectionGroupKey    string          `json:"section_group_key,omitempty" jsonschema:"Section group identifier"`
	SectionKey         string          `json:"section_key,omitempty" jsonschema:"Section identifier within group"`
	Vertices           [][3]float64    `json:"vertices,omitempty" jsonschema:"Array of [X,Y,Z] vertex positions"`
	Triangles          []int           `json:"triangles,omitempty" jsonschema:"Array of triangle vertex indices"`
	Normals            [][3]float64    `json:"normals,omitempty" jsonschema:"Array of [X,Y,Z] vertex normals"`
	Tangents           [][3]float64    `json:"tangents,omitempty" jsonschema:"Array of [X,Y,Z] vertex tangents"`
	UVs                [][2]float64    `json:"uvs,omitempty" jsonschema:"Array of [U,V] texture coordinates"`
	Colors             [][4]float64    `json:"colors,omitempty" jsonschema:"Array of [R,G,B,A] vertex colors"`
	MaterialSlotName   string          `json:"material_slot_name,omitempty" jsonschema:"Named material slot for set_material_slot"`
	MaterialPath       string          `json:"material_path,omitempty" jsonschema:"Material asset path"`
	CollisionVertices  [][3]float64    `json:"collision_vertices,omitempty" jsonschema:"Simplified collision mesh vertices"`
	CollisionTriangles []int           `json:"collision_triangles,omitempty" jsonschema:"Simplified collision mesh indices"`
	Location           [3]float64      `json:"location,omitempty" jsonschema:"[X,Y,Z] spawn location"`
	Params             json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
	World              string          `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// RealtimeMeshOutput is returned by the realtime_mesh tool.
type RealtimeMeshOutput struct {
	Result any `json:"result" jsonschema:"Operation results"`
}

// RegisterMesh adds the procedural_mesh and realtime_mesh tools to the MCP server.
func (h *Handler) RegisterMesh(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "procedural_mesh",
		Description: "Create and modify ProceduralMeshComponent geometry in the UE editor. Operations:\n" +
			"- create_section: Create a mesh section with vertices, triangles, normals, UVs (spawns actor if needed)\n" +
			"- update_section: Update an existing mesh section's geometry\n" +
			"- clear: Clear all mesh sections from a ProceduralMeshActor\n" +
			"- set_material: Assign a material to a mesh section\n\n" +
			"Vertex data format: vertices=[[X,Y,Z],...], triangles=[0,1,2,...], normals=[[X,Y,Z],...], uvs=[[U,V],...]\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.ProceduralMesh)

	mcp.AddTool(server, &mcp.Tool{
		Name: "realtime_mesh",
		Description: "Create and modify RealtimeMeshComponent geometry in the UE editor (requires RealtimeMesh plugin). Operations:\n" +
			"- create_lod: Create a LOD level with screen size threshold\n" +
			"- create_section_group: Create a section group within a LOD\n" +
			"- create_section: Create a mesh section with vertex stream data\n" +
			"- update_mesh_data: Update section geometry\n" +
			"- set_material_slot: Assign a material to a named slot\n" +
			"- setup_collision: Configure collision mesh from simplified geometry\n\n" +
			"Supports multi-LOD meshes with section groups. RealtimeMesh plugin must be installed.\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.RealtimeMesh)
}

// ProceduralMesh implements the procedural_mesh tool.
func (h *Handler) ProceduralMesh(ctx context.Context, req *mcp.CallToolRequest, input ProceduralMeshInput) (*mcp.CallToolResult, ProceduralMeshOutput, error) {
	if input.Operation == "" {
		return nil, ProceduralMeshOutput{}, fmt.Errorf("operation is required (create_section, update_section, clear, set_material)")
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
	if input.SectionIndex != 0 {
		body["section_index"] = input.SectionIndex
	}
	if len(input.Vertices) > 0 {
		body["vertices"] = input.Vertices
	}
	if len(input.Triangles) > 0 {
		body["triangles"] = input.Triangles
	}
	if len(input.Normals) > 0 {
		body["normals"] = input.Normals
	}
	if len(input.UVs) > 0 {
		body["uvs"] = input.UVs
	}
	if len(input.Colors) > 0 {
		body["colors"] = input.Colors
	}
	if input.MaterialPath != "" {
		body["material_path"] = input.MaterialPath
	}
	if input.Location != [3]float64{} {
		body["location"] = input.Location
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/mesh/procedural", body)
	if err != nil {
		return nil, ProceduralMeshOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, ProceduralMeshOutput{}, fmt.Errorf("parsing procedural_mesh response: %w", err)
	}

	return nil, ProceduralMeshOutput{Result: result}, nil
}

// RealtimeMesh implements the realtime_mesh tool.
func (h *Handler) RealtimeMesh(ctx context.Context, req *mcp.CallToolRequest, input RealtimeMeshInput) (*mcp.CallToolResult, RealtimeMeshOutput, error) {
	if input.Operation == "" {
		return nil, RealtimeMeshOutput{}, fmt.Errorf("operation is required (create_lod, create_section_group, create_section, update_mesh_data, set_material_slot, setup_collision)")
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
	if input.LODIndex != 0 {
		body["lod_index"] = input.LODIndex
	}
	if input.ScreenSize != 0 {
		body["screen_size"] = input.ScreenSize
	}
	if input.SectionGroupKey != "" {
		body["section_group_key"] = input.SectionGroupKey
	}
	if input.SectionKey != "" {
		body["section_key"] = input.SectionKey
	}
	if len(input.Vertices) > 0 {
		body["vertices"] = input.Vertices
	}
	if len(input.Triangles) > 0 {
		body["triangles"] = input.Triangles
	}
	if len(input.Normals) > 0 {
		body["normals"] = input.Normals
	}
	if len(input.Tangents) > 0 {
		body["tangents"] = input.Tangents
	}
	if len(input.UVs) > 0 {
		body["uvs"] = input.UVs
	}
	if len(input.Colors) > 0 {
		body["colors"] = input.Colors
	}
	if input.MaterialSlotName != "" {
		body["material_slot_name"] = input.MaterialSlotName
	}
	if input.MaterialPath != "" {
		body["material_path"] = input.MaterialPath
	}
	if len(input.CollisionVertices) > 0 {
		body["collision_vertices"] = input.CollisionVertices
	}
	if len(input.CollisionTriangles) > 0 {
		body["collision_triangles"] = input.CollisionTriangles
	}
	if input.Location != [3]float64{} {
		body["location"] = input.Location
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/mesh/realtime", body)
	if err != nil {
		return nil, RealtimeMeshOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded and RealtimeMesh plugin installed: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, RealtimeMeshOutput{}, fmt.Errorf("parsing realtime_mesh response: %w", err)
	}

	return nil, RealtimeMeshOutput{Result: result}, nil
}
