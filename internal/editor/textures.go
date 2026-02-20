package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- texture_ops ---

// TextureInfo describes a texture asset.
type TextureInfo struct {
	Asset       string `json:"asset" jsonschema:"asset path"`
	Name        string `json:"name" jsonschema:"texture name"`
	Width       int    `json:"width" jsonschema:"width in pixels"`
	Height      int    `json:"height" jsonschema:"height in pixels"`
	Format      string `json:"format,omitempty" jsonschema:"pixel format (e.g. PF_B8G8R8A8)"`
	MipCount    int    `json:"mip_count,omitempty" jsonschema:"number of mip levels"`
	Compression string `json:"compression,omitempty" jsonschema:"compression setting (e.g. TC_Default)"`
	SizeKB      int    `json:"size_kb,omitempty" jsonschema:"approximate size in KB"`
}

// TextureOpsInput defines parameters for the texture_ops tool.
type TextureOpsInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: import, get_info, set_material_texture, list"`
	// For import: source file path and destination asset path.
	SourcePath  string `json:"source_path,omitempty" jsonschema:"Local file path to import (PNG, TIFF, EXR, JPG). For import."`
	Destination string `json:"destination,omitempty" jsonschema:"UE content path for imported texture (e.g. /Game/Textures/MyTex). For import."`
	Compression string `json:"compression,omitempty" jsonschema:"Compression setting: TC_Default, TC_Normalmap, TC_Masks, TC_HDR, TC_VectorDisplacementmap. For import."`
	// For get_info and set_material_texture.
	Asset string `json:"asset,omitempty" jsonschema:"Texture asset path (e.g. /Game/Textures/MyTex). For get_info."`
	// For set_material_texture.
	MaterialInstance string `json:"material_instance,omitempty" jsonschema:"Material instance asset path. For set_material_texture."`
	ParamName        string `json:"param_name,omitempty" jsonschema:"Texture parameter name in the material. For set_material_texture."`
	Texture          string `json:"texture,omitempty" jsonschema:"Texture asset path to assign. For set_material_texture."`
	// For list.
	Path string `json:"path,omitempty" jsonschema:"Content path to list textures from (e.g. /Game/Textures/). For list."`
}

// TextureOpsOutput is returned by the texture_ops tool.
type TextureOpsOutput struct {
	Success  bool          `json:"success" jsonschema:"whether the operation succeeded"`
	Asset    string        `json:"asset,omitempty" jsonschema:"imported or queried texture asset path"`
	Info     *TextureInfo  `json:"info,omitempty" jsonschema:"texture details (for get_info)"`
	Textures []TextureInfo `json:"textures,omitempty" jsonschema:"list of textures (for list)"`
	Count    int           `json:"count,omitempty" jsonschema:"number of textures found"`
	Message  string        `json:"message,omitempty" jsonschema:"status message"`
}

// RegisterTextures adds the texture_ops tool to the MCP server.
func (h *Handler) RegisterTextures(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "texture_ops",
		Description: "Manage texture assets: import images (PNG, TIFF, EXR, JPG) as texture assets, " +
			"query texture info (dimensions, format, compression), assign textures to material parameters, " +
			"and list texture assets in a folder. " +
			"Operations: import, get_info, set_material_texture, list. " +
			"Requires the editor running with the MCPUnreal plugin loaded (port 8090).",
	}, h.TextureOps)
}

// TextureOps implements the texture_ops tool.
func (h *Handler) TextureOps(ctx context.Context, req *mcp.CallToolRequest, input TextureOpsInput) (*mcp.CallToolResult, TextureOpsOutput, error) {
	if input.Operation == "" {
		return nil, TextureOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.SourcePath != "" {
		body["source_path"] = input.SourcePath
	}
	if input.Destination != "" {
		body["destination"] = input.Destination
	}
	if input.Compression != "" {
		body["compression"] = input.Compression
	}
	if input.Asset != "" {
		body["asset"] = input.Asset
	}
	if input.MaterialInstance != "" {
		body["material_instance"] = input.MaterialInstance
	}
	if input.ParamName != "" {
		body["param_name"] = input.ParamName
	}
	if input.Texture != "" {
		body["texture"] = input.Texture
	}
	if input.Path != "" {
		body["path"] = input.Path
	}

	resp, err := h.Client.PluginCall(ctx, "/api/textures/ops", body)
	if err != nil {
		return nil, TextureOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out TextureOpsOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, TextureOpsOutput{}, fmt.Errorf("parsing texture response: %w", err)
	}

	return nil, out, nil
}
