package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- search_assets ---

// SearchAssetsInput defines parameters for the search_assets tool.
type SearchAssetsInput struct {
	ClassFilter   string `json:"class_filter,omitempty" jsonschema:"Filter by asset class (e.g. Blueprint, StaticMesh, Material, Texture2D)"`
	PathFilter    string `json:"path_filter,omitempty" jsonschema:"Filter by path prefix (e.g. /Game/Blueprints)"`
	NameFilter    string `json:"name_filter,omitempty" jsonschema:"Filter by asset name substring"`
	RecursivePath bool   `json:"recursive_path,omitempty" jsonschema:"Search subdirectories (default true)"`
}

// AssetEntry describes a single asset in the registry.
type AssetEntry struct {
	Name    string `json:"name" jsonschema:"asset name"`
	Path    string `json:"path" jsonschema:"full object path"`
	Class   string `json:"class" jsonschema:"asset class"`
	Package string `json:"package" jsonschema:"package name"`
}

// SearchAssetsOutput is returned by the search_assets tool.
type SearchAssetsOutput struct {
	Assets []AssetEntry `json:"assets" jsonschema:"matching assets"`
	Total  int          `json:"total" jsonschema:"number of matching assets"`
}

// --- get_asset_info ---

// GetAssetInfoInput defines parameters for the get_asset_info tool.
type GetAssetInfoInput struct {
	AssetPath string `json:"asset_path" jsonschema:"required,Full asset path (e.g. /Game/Blueprints/BP_Player.BP_Player)"`
}

// GetAssetInfoOutput is returned by the get_asset_info tool.
type GetAssetInfoOutput struct {
	Name         string            `json:"name" jsonschema:"asset name"`
	Path         string            `json:"path" jsonschema:"full object path"`
	Class        string            `json:"class" jsonschema:"asset class path"`
	Package      string            `json:"package" jsonschema:"package name"`
	PackageFlags float64           `json:"package_flags" jsonschema:"package flags bitmask"`
	Tags         map[string]string `json:"tags,omitempty" jsonschema:"asset registry tags"`
	Dependencies []string          `json:"dependencies,omitempty" jsonschema:"package dependencies"`
	Referencers  []string          `json:"referencers,omitempty" jsonschema:"packages that reference this asset"`
}

// RegisterAssets adds asset query tools to the MCP server.
func (h *Handler) RegisterAssets(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "search_assets",
		Description: "Search for assets in the UE Asset Registry by class, path, or name. " +
			"Returns asset names, paths, and classes. " +
			"Requires the editor running with the MCPUnreal plugin (port 8090).",
	}, h.SearchAssets)

	mcp.AddTool(server, &mcp.Tool{
		Name: "get_asset_info",
		Description: "Get detailed information about a specific asset including metadata, " +
			"AssetRegistry tags, dependencies, and referencers. " +
			"Requires the editor running with the MCPUnreal plugin (port 8090).",
	}, h.GetAssetInfo)
}

// SearchAssets implements the search_assets tool.
func (h *Handler) SearchAssets(ctx context.Context, req *mcp.CallToolRequest, input SearchAssetsInput) (*mcp.CallToolResult, SearchAssetsOutput, error) {
	body := map[string]any{}
	if input.ClassFilter != "" {
		body["class_filter"] = input.ClassFilter
	}
	if input.PathFilter != "" {
		body["path_filter"] = input.PathFilter
	}
	if input.NameFilter != "" {
		body["name_filter"] = input.NameFilter
	}
	body["recursive_path"] = input.RecursivePath

	resp, err := h.Client.PluginCall(ctx, "/api/assets/search", body)
	if err != nil {
		return nil, SearchAssetsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	// The plugin returns a raw JSON array of asset objects.
	var assets []AssetEntry
	if err := json.Unmarshal(resp, &assets); err != nil {
		return nil, SearchAssetsOutput{}, fmt.Errorf("parsing search response: %w", err)
	}

	return nil, SearchAssetsOutput{
		Assets: assets,
		Total:  len(assets),
	}, nil
}

// GetAssetInfo implements the get_asset_info tool.
func (h *Handler) GetAssetInfo(ctx context.Context, req *mcp.CallToolRequest, input GetAssetInfoInput) (*mcp.CallToolResult, GetAssetInfoOutput, error) {
	if input.AssetPath == "" {
		return nil, GetAssetInfoOutput{}, fmt.Errorf("asset_path is required")
	}

	// Fetch base info.
	infoResp, err := h.Client.PluginCall(ctx, "/api/assets/info", map[string]string{
		"asset_path": input.AssetPath,
	})
	if err != nil {
		return nil, GetAssetInfoOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out GetAssetInfoOutput
	if err := json.Unmarshal(infoResp, &out); err != nil {
		return nil, GetAssetInfoOutput{}, fmt.Errorf("parsing asset info response: %w", err)
	}

	// Fetch dependencies.
	depsResp, err := h.Client.PluginCall(ctx, "/api/assets/dependencies", map[string]string{
		"asset_path": input.AssetPath,
	})
	if err == nil {
		var deps struct {
			Dependencies []string `json:"dependencies"`
		}
		if json.Unmarshal(depsResp, &deps) == nil {
			out.Dependencies = deps.Dependencies
		}
	}

	// Fetch referencers.
	refsResp, err := h.Client.PluginCall(ctx, "/api/assets/referencers", map[string]string{
		"asset_path": input.AssetPath,
	})
	if err == nil {
		var refs struct {
			Referencers []string `json:"referencers"`
		}
		if json.Unmarshal(refsResp, &refs) == nil {
			out.Referencers = refs.Referencers
		}
	}

	return nil, out, nil
}
