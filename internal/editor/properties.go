// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- set_property ---

// SetPropertyInput defines parameters for the set_property tool.
type SetPropertyInput struct {
	ObjectPath    string          `json:"object_path" jsonschema:"required,Full UObject path (e.g. /Game/Maps/MyMap.MyMap:PersistentLevel.MyActor)"`
	PropertyName  string          `json:"property_name" jsonschema:"required,UPROPERTY name (e.g. RelativeLocation, bHidden, StaticMesh)"`
	PropertyValue json.RawMessage `json:"property_value" jsonschema:"required,Value to set — type depends on property (number, string, object, array)"`
}

// SetPropertyOutput is returned by the set_property tool.
type SetPropertyOutput struct {
	Success bool   `json:"success" jsonschema:"whether the property was set successfully"`
	Message string `json:"message,omitempty" jsonschema:"description of what was set"`
}

// --- get_property ---

// GetPropertyInput defines parameters for the get_property tool.
type GetPropertyInput struct {
	ObjectPath   string `json:"object_path" jsonschema:"required,Full UObject path (e.g. /Game/Maps/MyMap.MyMap:PersistentLevel.MyActor)"`
	PropertyName string `json:"property_name" jsonschema:"required,UPROPERTY name to read (e.g. RelativeLocation, bHidden)"`
}

// GetPropertyOutput is returned by the get_property tool.
type GetPropertyOutput struct {
	ObjectPath    string          `json:"object_path" jsonschema:"the queried object path"`
	PropertyName  string          `json:"property_name" jsonschema:"the queried property name"`
	PropertyValue json.RawMessage `json:"property_value" jsonschema:"the property value as JSON"`
}

// --- call_function ---

// CallFunctionInput defines parameters for the call_function tool.
type CallFunctionInput struct {
	ObjectPath   string         `json:"object_path" jsonschema:"required,Full UObject path to call the function on"`
	FunctionName string         `json:"function_name" jsonschema:"required,UFUNCTION name to call"`
	Parameters   map[string]any `json:"parameters,omitempty" jsonschema:"Function parameters as key-value pairs"`
}

// CallFunctionOutput is returned by the call_function tool.
type CallFunctionOutput struct {
	ReturnValue json.RawMessage `json:"return_value" jsonschema:"function return value as JSON"`
}

// RegisterProperties adds the set_property, get_property, and call_function
// tools to the MCP server. These use the UE Remote Control API (port 30010)
// and do not require the MCPUnreal plugin.
func (h *Handler) RegisterProperties(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "set_property",
		Description: "Set any UPROPERTY on any UObject in the running editor via the Remote Control API. " +
			"Requires the editor to be running with the Remote Control API plugin enabled (port 30010). " +
			"Example: set_property(object_path: '/Game/Maps/Main.Main:PersistentLevel.PointLight_0', " +
			"property_name: 'Intensity', property_value: 5000)",
	}, h.SetProperty)

	mcp.AddTool(server, &mcp.Tool{
		Name: "get_property",
		Description: "Read any UPROPERTY from any UObject in the running editor via the Remote Control API. " +
			"Requires the editor to be running with the Remote Control API plugin enabled (port 30010). " +
			"Returns the property value as JSON.",
	}, h.GetProperty)

	mcp.AddTool(server, &mcp.Tool{
		Name: "call_function",
		Description: "Call any UFUNCTION on any UObject in the running editor via the Remote Control API. " +
			"Requires the editor to be running with the Remote Control API plugin enabled (port 30010). " +
			"Pass parameters as key-value pairs matching the UFUNCTION signature.",
	}, h.CallFunction)
}

// SetProperty implements the set_property tool.
func (h *Handler) SetProperty(ctx context.Context, req *mcp.CallToolRequest, input SetPropertyInput) (*mcp.CallToolResult, SetPropertyOutput, error) {
	if input.ObjectPath == "" {
		return nil, SetPropertyOutput{}, fmt.Errorf("object_path is required")
	}
	if input.PropertyName == "" {
		return nil, SetPropertyOutput{}, fmt.Errorf("property_name is required")
	}

	body := map[string]any{
		"objectPath":    input.ObjectPath,
		"propertyName":  input.PropertyName,
		"propertyValue": input.PropertyValue,
	}

	_, err := h.Client.RCAPICall(ctx, "/remote/object/property", body)
	if err != nil {
		return nil, SetPropertyOutput{}, fmt.Errorf("setting property %s on %s: %w", input.PropertyName, input.ObjectPath, err)
	}

	return nil, SetPropertyOutput{
		Success: true,
		Message: fmt.Sprintf("set %s on %s", input.PropertyName, input.ObjectPath),
	}, nil
}

// GetProperty implements the get_property tool.
func (h *Handler) GetProperty(ctx context.Context, req *mcp.CallToolRequest, input GetPropertyInput) (*mcp.CallToolResult, GetPropertyOutput, error) {
	if input.ObjectPath == "" {
		return nil, GetPropertyOutput{}, fmt.Errorf("object_path is required")
	}
	if input.PropertyName == "" {
		return nil, GetPropertyOutput{}, fmt.Errorf("property_name is required")
	}

	body := map[string]any{
		"objectPath":   input.ObjectPath,
		"propertyName": input.PropertyName,
		"access":       "READ_ACCESS",
	}

	resp, err := h.Client.RCAPICall(ctx, "/remote/object/property", body)
	if err != nil {
		return nil, GetPropertyOutput{}, fmt.Errorf("reading property %s from %s: %w", input.PropertyName, input.ObjectPath, err)
	}

	return nil, GetPropertyOutput{
		ObjectPath:    input.ObjectPath,
		PropertyName:  input.PropertyName,
		PropertyValue: resp,
	}, nil
}

// CallFunction implements the call_function tool.
func (h *Handler) CallFunction(ctx context.Context, req *mcp.CallToolRequest, input CallFunctionInput) (*mcp.CallToolResult, CallFunctionOutput, error) {
	if input.ObjectPath == "" {
		return nil, CallFunctionOutput{}, fmt.Errorf("object_path is required")
	}
	if input.FunctionName == "" {
		return nil, CallFunctionOutput{}, fmt.Errorf("function_name is required")
	}

	body := map[string]any{
		"objectPath":   input.ObjectPath,
		"functionName": input.FunctionName,
	}
	if input.Parameters != nil {
		body["parameters"] = input.Parameters
	}

	resp, err := h.Client.RCAPICall(ctx, "/remote/object/call", body)
	if err != nil {
		return nil, CallFunctionOutput{}, fmt.Errorf("calling %s on %s: %w", input.FunctionName, input.ObjectPath, err)
	}

	return nil, CallFunctionOutput{ReturnValue: resp}, nil
}
