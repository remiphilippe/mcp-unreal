package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- niagara_ops ---

// NiagaraOpsInput defines parameters for the niagara_ops tool.
type NiagaraOpsInput struct {
	Operation      string          `json:"operation" jsonschema:"required,One of: spawn_system, set_parameter, get_system_info, add_emitter, remove_emitter, activate, deactivate"`
	SystemPath     string          `json:"system_path,omitempty" jsonschema:"Niagara system asset path — required for spawn_system, get_system_info, add_emitter, remove_emitter"`
	ActorPath      string          `json:"actor_path,omitempty" jsonschema:"Actor with NiagaraComponent — required for set_parameter, activate, deactivate"`
	Location       [3]float64      `json:"location,omitempty" jsonschema:"[X,Y,Z] spawn location for spawn_system"`
	Rotation       [3]float64      `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] spawn rotation in degrees"`
	Scale          [3]float64      `json:"scale,omitempty" jsonschema:"[X,Y,Z] scale factor. Default [1,1,1]"`
	ParameterName  string          `json:"parameter_name,omitempty" jsonschema:"Parameter name for set_parameter"`
	ParameterValue any             `json:"parameter_value,omitempty" jsonschema:"Parameter value for set_parameter"`
	ParameterType  string          `json:"parameter_type,omitempty" jsonschema:"Parameter type for set_parameter: int, float, bool, vector, color"`
	EmitterPath    string          `json:"emitter_path,omitempty" jsonschema:"Emitter asset path for add_emitter"`
	EmitterName    string          `json:"emitter_name,omitempty" jsonschema:"Emitter name for remove_emitter"`
	AutoActivate   *bool           `json:"auto_activate,omitempty" jsonschema:"Auto-activate on spawn. Pointer type: false is meaningful"`
	ActorName      string          `json:"actor_name,omitempty" jsonschema:"Display name for spawned actor"`
	Params         json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
	World          string          `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// NiagaraOpsOutput is returned by the niagara_ops tool.
type NiagaraOpsOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterNiagara adds the niagara_ops tool to the MCP server.
func (h *Handler) RegisterNiagara(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "niagara_ops",
		Description: "Manage Niagara VFX systems and components in the UE editor. Operations:\n" +
			"- spawn_system: Spawn a Niagara system actor at a location (requires system_path)\n" +
			"- set_parameter: Set a parameter on a NiagaraComponent (requires actor_path, parameter_name, parameter_value, parameter_type)\n" +
			"- get_system_info: List emitters and exposed parameters of a Niagara system (requires system_path)\n" +
			"- add_emitter: Add an emitter to a Niagara system (requires system_path, emitter_path)\n" +
			"- remove_emitter: Remove an emitter by name (requires system_path, emitter_name)\n" +
			"- activate: Activate a NiagaraComponent (requires actor_path)\n" +
			"- deactivate: Deactivate a NiagaraComponent (requires actor_path)\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090). Niagara plugin must be enabled.",
	}, h.NiagaraOps)
}

// NiagaraOps implements the niagara_ops tool.
func (h *Handler) NiagaraOps(ctx context.Context, req *mcp.CallToolRequest, input NiagaraOpsInput) (*mcp.CallToolResult, NiagaraOpsOutput, error) {
	if input.Operation == "" {
		return nil, NiagaraOpsOutput{}, fmt.Errorf("operation is required (spawn_system, set_parameter, get_system_info, add_emitter, remove_emitter, activate, deactivate)")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.SystemPath != "" {
		body["system_path"] = input.SystemPath
	}
	if input.ActorPath != "" {
		body["actor_path"] = input.ActorPath
	}
	if input.Location != [3]float64{} {
		body["location"] = input.Location
	}
	if input.Rotation != [3]float64{} {
		body["rotation"] = input.Rotation
	}
	if input.Scale != [3]float64{} {
		body["scale"] = input.Scale
	}
	if input.ParameterName != "" {
		body["parameter_name"] = input.ParameterName
	}
	if input.ParameterValue != nil {
		body["parameter_value"] = input.ParameterValue
	}
	if input.ParameterType != "" {
		body["parameter_type"] = input.ParameterType
	}
	if input.EmitterPath != "" {
		body["emitter_path"] = input.EmitterPath
	}
	if input.EmitterName != "" {
		body["emitter_name"] = input.EmitterName
	}
	if input.AutoActivate != nil {
		body["auto_activate"] = *input.AutoActivate
	}
	if input.ActorName != "" {
		body["actor_name"] = input.ActorName
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

	resp, err := h.Client.PluginCall(ctx, "/api/niagara/ops", body)
	if err != nil {
		return nil, NiagaraOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, NiagaraOpsOutput{}, fmt.Errorf("parsing niagara_ops response: %w", err)
	}

	return nil, NiagaraOpsOutput{Result: result}, nil
}
