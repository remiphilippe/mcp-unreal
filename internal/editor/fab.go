package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- fab_ops ---

// FabCachedAsset describes a cached Fab asset.
type FabCachedAsset struct {
	AssetID  string `json:"asset_id" jsonschema:"Fab asset identifier"`
	FilePath string `json:"file_path,omitempty" jsonschema:"local file path of cached asset"`
}

// FabOpsInput defines parameters for the fab_ops tool.
type FabOpsInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: list_cache, cache_info, import, clear_cache"`
	// For import: which cached asset to import.
	AssetID string `json:"asset_id,omitempty" jsonschema:"Fab asset ID to import (for import operation)"`
	// For import: destination content path.
	Destination string `json:"destination,omitempty" jsonschema:"UE content path to import to, e.g. /Game/Assets/ (for import operation)"`
}

// FabOpsOutput is returned by the fab_ops tool.
type FabOpsOutput struct {
	Success        bool             `json:"success" jsonschema:"whether the operation succeeded"`
	CacheLocation  string           `json:"cache_location,omitempty" jsonschema:"local file system path of the Fab cache"`
	CacheSize      string           `json:"cache_size,omitempty" jsonschema:"formatted cache size (e.g. '2.3 GB')"`
	CacheSizeBytes int64            `json:"cache_size_bytes,omitempty" jsonschema:"cache size in bytes"`
	AssetCount     int              `json:"asset_count,omitempty" jsonschema:"number of cached assets"`
	Assets         []FabCachedAsset `json:"assets,omitempty" jsonschema:"list of cached assets"`
	ImportedPath   string           `json:"imported_path,omitempty" jsonschema:"content path where assets were imported"`
	Message        string           `json:"message,omitempty" jsonschema:"status message"`
}

// RegisterFab adds the fab_ops tool to the MCP server.
func (h *Handler) RegisterFab(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "fab_ops",
		Description: "Manage Fab marketplace asset cache and imports. " +
			"Operations: list_cache (list downloaded assets), cache_info (cache location and size), " +
			"import (import a cached asset into the project), clear_cache (purge download cache). " +
			"Only works with already-downloaded/cached assets — does not search or download from Fab. " +
			"Requires the editor running with the MCPUnreal plugin and the Fab plugin enabled.",
	}, h.FabOps)
}

// FabOps implements the fab_ops tool.
func (h *Handler) FabOps(ctx context.Context, req *mcp.CallToolRequest, input FabOpsInput) (*mcp.CallToolResult, FabOpsOutput, error) {
	if input.Operation == "" {
		return nil, FabOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.AssetID != "" {
		body["asset_id"] = input.AssetID
	}
	if input.Destination != "" {
		body["destination"] = input.Destination
	}

	resp, err := h.Client.PluginCall(ctx, "/api/fab/ops", body)
	if err != nil {
		return nil, FabOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out FabOpsOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, FabOpsOutput{}, fmt.Errorf("parsing Fab response: %w", err)
	}

	return nil, out, nil
}
