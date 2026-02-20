package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- data_asset_ops ---

// DataTableRow represents a single row in a DataTable.
type DataTableRow struct {
	RowName string         `json:"row_name" jsonschema:"row identifier"`
	Data    map[string]any `json:"data" jsonschema:"row column values"`
}

// DataTableInfo describes a DataTable asset.
type DataTableInfo struct {
	Asset     string `json:"asset" jsonschema:"asset path"`
	Name      string `json:"name" jsonschema:"table name"`
	RowStruct string `json:"row_struct" jsonschema:"row struct type (e.g. FItemData)"`
	RowCount  int    `json:"row_count" jsonschema:"number of rows"`
}

// DataAssetOpsInput defines parameters for the data_asset_ops tool.
type DataAssetOpsInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: list_tables, get_table, add_row, update_row, delete_row, create_table, import_csv"`
	// For most operations: target asset.
	Asset string `json:"asset,omitempty" jsonschema:"DataTable asset path (e.g. /Game/Data/DT_Items). For get_table, add_row, update_row, delete_row."`
	// For list_tables.
	Path string `json:"path,omitempty" jsonschema:"Content path to list DataTables from (e.g. /Game/Data/). For list_tables."`
	// For add_row, update_row, delete_row.
	RowName string         `json:"row_name,omitempty" jsonschema:"Row name/key. For add_row, update_row, delete_row."`
	Data    map[string]any `json:"data,omitempty" jsonschema:"Row data as JSON key-value pairs. For add_row, update_row."`
	// For create_table.
	Destination string `json:"destination,omitempty" jsonschema:"Destination asset path for new table. For create_table."`
	RowStruct   string `json:"row_struct,omitempty" jsonschema:"Row struct type (e.g. FItemData). For create_table."`
	// For import_csv.
	SourcePath string `json:"source_path,omitempty" jsonschema:"CSV file path to import. For import_csv."`
}

// DataAssetOpsOutput is returned by the data_asset_ops tool.
type DataAssetOpsOutput struct {
	Success  bool            `json:"success" jsonschema:"whether the operation succeeded"`
	Asset    string          `json:"asset,omitempty" jsonschema:"affected DataTable asset path"`
	Tables   []DataTableInfo `json:"tables,omitempty" jsonschema:"list of DataTables (for list_tables)"`
	Rows     []DataTableRow  `json:"rows,omitempty" jsonschema:"table rows (for get_table)"`
	RowCount int             `json:"row_count,omitempty" jsonschema:"number of rows after operation"`
	Message  string          `json:"message,omitempty" jsonschema:"status message"`
}

// RegisterDataAssets adds the data_asset_ops tool to the MCP server.
func (h *Handler) RegisterDataAssets(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "data_asset_ops",
		Description: "Manage UE DataTables: list tables in a folder, read all rows, add/update/delete rows, " +
			"create new tables with a specified row struct, and import from CSV. " +
			"Operations: list_tables, get_table, add_row, update_row, delete_row, create_table, import_csv. " +
			"Row data is serialized as JSON key-value pairs. " +
			"Requires the editor running with the MCPUnreal plugin loaded (port 8090).",
	}, h.DataAssetOps)
}

// DataAssetOps implements the data_asset_ops tool.
func (h *Handler) DataAssetOps(ctx context.Context, req *mcp.CallToolRequest, input DataAssetOpsInput) (*mcp.CallToolResult, DataAssetOpsOutput, error) {
	if input.Operation == "" {
		return nil, DataAssetOpsOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.Asset != "" {
		body["asset"] = input.Asset
	}
	if input.Path != "" {
		body["path"] = input.Path
	}
	if input.RowName != "" {
		body["row_name"] = input.RowName
	}
	if input.Data != nil {
		body["data"] = input.Data
	}
	if input.Destination != "" {
		body["destination"] = input.Destination
	}
	if input.RowStruct != "" {
		body["row_struct"] = input.RowStruct
	}
	if input.SourcePath != "" {
		body["source_path"] = input.SourcePath
	}

	resp, err := h.Client.PluginCall(ctx, "/api/data/ops", body)
	if err != nil {
		return nil, DataAssetOpsOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out DataAssetOpsOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, DataAssetOpsOutput{}, fmt.Errorf("parsing data asset response: %w", err)
	}

	return nil, out, nil
}
