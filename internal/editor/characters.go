package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- character_config ---

// CharacterConfigInput defines parameters for the character_config tool.
type CharacterConfigInput struct {
	Operation     string          `json:"operation" jsonschema:"required,One of: get_config, set_movement, set_capsule, set_mesh, set_camera, get_movement_modes"`
	BlueprintPath string          `json:"blueprint_path" jsonschema:"required,Character Blueprint path (e.g. /Game/Characters/BP_PlayerCharacter)"`
	Params        json.RawMessage `json:"params,omitempty" jsonschema:"Operation-specific parameters"`
	// Common movement fields.
	MaxWalkSpeed        *float64 `json:"max_walk_speed,omitempty" jsonschema:"Maximum walking speed (cm/s)"`
	MaxAcceleration     *float64 `json:"max_acceleration,omitempty" jsonschema:"Maximum acceleration (cm/s^2)"`
	JumpZVelocity       *float64 `json:"jump_z_velocity,omitempty" jsonschema:"Jump velocity (cm/s)"`
	GravityScale        *float64 `json:"gravity_scale,omitempty" jsonschema:"Gravity multiplier"`
	AirControl          *float64 `json:"air_control,omitempty" jsonschema:"Air control factor (0-1)"`
	BrakingDeceleration *float64 `json:"braking_deceleration,omitempty" jsonschema:"Braking deceleration (cm/s^2)"`
	// Capsule.
	CapsuleRadius     *float64 `json:"capsule_radius,omitempty" jsonschema:"Capsule collision radius"`
	CapsuleHalfHeight *float64 `json:"capsule_half_height,omitempty" jsonschema:"Capsule collision half-height"`
	// Mesh.
	SkeletalMeshPath  string `json:"skeletal_mesh_path,omitempty" jsonschema:"Skeletal mesh asset path to set"`
	AnimBlueprintPath string `json:"anim_blueprint_path,omitempty" jsonschema:"Animation Blueprint to assign"`
}

// CharacterConfigOutput is returned by the character_config tool.
type CharacterConfigOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterCharacters adds the character_config tool to the MCP server.
func (h *Handler) RegisterCharacters(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "character_config",
		Description: "Configure character movement, capsule, mesh, and camera settings in the UE editor. Operations:\n" +
			"- get_config: Get current character configuration\n" +
			"- set_movement: Set CharacterMovementComponent properties (max_walk_speed, jump_z_velocity, etc.)\n" +
			"- set_capsule: Set capsule collision size (capsule_radius, capsule_half_height)\n" +
			"- set_mesh: Set skeletal mesh and animation Blueprint\n" +
			"- set_camera: Set camera boom/follow settings\n" +
			"- get_movement_modes: List available movement modes\n" +
			"Requires blueprint_path. Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.CharacterConfig)
}

// CharacterConfig implements the character_config tool.
func (h *Handler) CharacterConfig(ctx context.Context, req *mcp.CallToolRequest, input CharacterConfigInput) (*mcp.CallToolResult, CharacterConfigOutput, error) {
	if input.Operation == "" {
		return nil, CharacterConfigOutput{}, fmt.Errorf("operation is required")
	}
	if input.BlueprintPath == "" {
		return nil, CharacterConfigOutput{}, fmt.Errorf("blueprint_path is required")
	}

	body := map[string]any{
		"operation":      input.Operation,
		"blueprint_path": input.BlueprintPath,
	}
	if input.MaxWalkSpeed != nil {
		body["max_walk_speed"] = *input.MaxWalkSpeed
	}
	if input.MaxAcceleration != nil {
		body["max_acceleration"] = *input.MaxAcceleration
	}
	if input.JumpZVelocity != nil {
		body["jump_z_velocity"] = *input.JumpZVelocity
	}
	if input.GravityScale != nil {
		body["gravity_scale"] = *input.GravityScale
	}
	if input.AirControl != nil {
		body["air_control"] = *input.AirControl
	}
	if input.BrakingDeceleration != nil {
		body["braking_deceleration"] = *input.BrakingDeceleration
	}
	if input.CapsuleRadius != nil {
		body["capsule_radius"] = *input.CapsuleRadius
	}
	if input.CapsuleHalfHeight != nil {
		body["capsule_half_height"] = *input.CapsuleHalfHeight
	}
	if input.SkeletalMeshPath != "" {
		body["skeletal_mesh_path"] = input.SkeletalMeshPath
	}
	if input.AnimBlueprintPath != "" {
		body["anim_blueprint_path"] = input.AnimBlueprintPath
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

	resp, err := h.Client.PluginCall(ctx, "/api/characters/config", body)
	if err != nil {
		return nil, CharacterConfigOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, CharacterConfigOutput{}, fmt.Errorf("parsing character_config response: %w", err)
	}

	return nil, CharacterConfigOutput{Result: result}, nil
}
