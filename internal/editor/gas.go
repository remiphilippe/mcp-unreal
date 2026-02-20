package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- gas_ops ---

// GASOpsInput defines parameters for the gas_ops tool.
type GASOpsInput struct {
	Operation      string          `json:"operation" jsonschema:"required,One of: grant_ability, revoke_ability, list_abilities, apply_effect, get_attributes, set_attribute"`
	ActorPath      string          `json:"actor_path,omitempty" jsonschema:"Actor with AbilitySystemComponent — required for all operations"`
	AbilityClass   string          `json:"ability_class,omitempty" jsonschema:"Gameplay ability class path for grant_ability/revoke_ability"`
	AbilityTag     string          `json:"ability_tag,omitempty" jsonschema:"Gameplay tag string for revoke_ability (alternative to ability_class)"`
	EffectClass    string          `json:"effect_class,omitempty" jsonschema:"Gameplay effect class path for apply_effect"`
	EffectLevel    *float64        `json:"effect_level,omitempty" jsonschema:"Effect level for apply_effect (default 1). Pointer type: zero is a valid value"`
	AttributeSet   string          `json:"attribute_set,omitempty" jsonschema:"Attribute set class name filter for get_attributes"`
	AttributeName  string          `json:"attribute_name,omitempty" jsonschema:"Attribute name for set_attribute (e.g. Health, MaxHealth)"`
	AttributeValue *float64        `json:"attribute_value,omitempty" jsonschema:"Attribute value for set_attribute. Pointer type: zero is a valid value"`
	Params         json.RawMessage `json:"params,omitempty" jsonschema:"Additional operation-specific parameters"`
}

// GASOpsOutput is returned by the gas_ops tool.
type GASOpsOutput struct {
	Result any `json:"result" jsonschema:"Operation results (structure depends on operation)"`
}

// RegisterGAS adds the gas_ops tool to the MCP server.
func (h *Handler) RegisterGAS(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "gas_ops",
		Description: "Manage the Gameplay Ability System (GAS) on actors in the UE editor. Operations:\n" +
			"- grant_ability: Grant a gameplay ability to an actor (requires actor_path, ability_class)\n" +
			"- revoke_ability: Remove a granted ability (requires actor_path, + ability_class or ability_tag)\n" +
			"- list_abilities: List all granted abilities and their status (requires actor_path)\n" +
			"- apply_effect: Apply a gameplay effect (requires actor_path, effect_class)\n" +
			"- get_attributes: List all attribute sets and current values (requires actor_path)\n" +
			"- set_attribute: Set an attribute's base value (requires actor_path, attribute_name, attribute_value)\n" +
			"The actor must have an AbilitySystemComponent (via IAbilitySystemInterface or as a component).\n" +
			"Requires the editor running with MCPUnreal plugin (port 8090). GameplayAbilities plugin must be enabled.",
	}, h.GASOps)
}

// GASOps implements the gas_ops tool.
func (h *Handler) GASOps(ctx context.Context, req *mcp.CallToolRequest, input GASOpsInput) (*mcp.CallToolResult, GASOpsOutput, error) {
	if input.Operation == "" {
		return nil, GASOpsOutput{}, fmt.Errorf("operation is required (grant_ability, revoke_ability, list_abilities, apply_effect, get_attributes, set_attribute)")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.ActorPath != "" {
		body["actor_path"] = input.ActorPath
	}
	if input.AbilityClass != "" {
		body["ability_class"] = input.AbilityClass
	}
	if input.AbilityTag != "" {
		body["ability_tag"] = input.AbilityTag
	}
	if input.EffectClass != "" {
		body["effect_class"] = input.EffectClass
	}
	if input.EffectLevel != nil {
		body["effect_level"] = *input.EffectLevel
	}
	if input.AttributeSet != "" {
		body["attribute_set"] = input.AttributeSet
	}
	if input.AttributeName != "" {
		body["attribute_name"] = input.AttributeName
	}
	if input.AttributeValue != nil {
		body["attribute_value"] = *input.AttributeValue
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

	resp, err := h.Client.PluginCall(ctx, "/api/gas/ops", body)
	if err != nil {
		return nil, GASOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var result any
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, GASOpsOutput{}, fmt.Errorf("parsing gas_ops response: %w", err)
	}

	return nil, GASOpsOutput{Result: result}, nil
}
