package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func newDataAssetTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestDataAssetOps_ListTables(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/data/ops" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_tables" {
			t.Errorf("expected list_tables, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(DataAssetOpsOutput{
			Success: true,
			Tables: []DataTableInfo{
				{Asset: "/Game/Data/DT_Items", Name: "DT_Items", RowStruct: "FItemData", RowCount: 25},
				{Asset: "/Game/Data/DT_Enemies", Name: "DT_Enemies", RowStruct: "FEnemyData", RowCount: 10},
			},
		})
	}))
	defer ts.Close()

	h := newDataAssetTestHandler(ts)
	_, out, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{
		Operation: "list_tables",
		Path:      "/Game/Data/",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if len(out.Tables) != 2 {
		t.Fatalf("expected 2 tables, got %d", len(out.Tables))
	}
	if out.Tables[0].RowStruct != "FItemData" {
		t.Errorf("expected FItemData, got %s", out.Tables[0].RowStruct)
	}
}

func TestDataAssetOps_GetTable(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(DataAssetOpsOutput{
			Success:  true,
			Asset:    "/Game/Data/DT_Items",
			RowCount: 2,
			Rows: []DataTableRow{
				{RowName: "Sword_01", Data: map[string]any{"Damage": float64(50), "Weight": 3.5}},
				{RowName: "Shield_01", Data: map[string]any{"Damage": float64(5), "Weight": 8.0}},
			},
		})
	}))
	defer ts.Close()

	h := newDataAssetTestHandler(ts)
	_, out, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{
		Operation: "get_table",
		Asset:     "/Game/Data/DT_Items",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.RowCount != 2 {
		t.Errorf("expected 2 rows, got %d", out.RowCount)
	}
	if out.Rows[0].RowName != "Sword_01" {
		t.Errorf("expected Sword_01, got %s", out.Rows[0].RowName)
	}
}

func TestDataAssetOps_AddRow(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["row_name"] != "Axe_01" {
			t.Errorf("expected row_name Axe_01, got %v", body["row_name"])
		}
		data, ok := body["data"].(map[string]any)
		if !ok {
			t.Error("expected data object")
		}
		if data["Damage"] != float64(75) {
			t.Errorf("expected Damage 75, got %v", data["Damage"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(DataAssetOpsOutput{
			Success:  true,
			Asset:    "/Game/Data/DT_Items",
			RowCount: 3,
			Message:  "Row added",
		})
	}))
	defer ts.Close()

	h := newDataAssetTestHandler(ts)
	_, out, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{
		Operation: "add_row",
		Asset:     "/Game/Data/DT_Items",
		RowName:   "Axe_01",
		Data:      map[string]any{"Damage": float64(75), "Weight": 5.0},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.RowCount != 3 {
		t.Errorf("expected 3 rows, got %d", out.RowCount)
	}
}

func TestDataAssetOps_DeleteRow(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["row_name"] != "Sword_01" {
			t.Errorf("expected Sword_01, got %v", body["row_name"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(DataAssetOpsOutput{
			Success:  true,
			RowCount: 1,
			Message:  "Row deleted",
		})
	}))
	defer ts.Close()

	h := newDataAssetTestHandler(ts)
	_, out, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{
		Operation: "delete_row",
		Asset:     "/Game/Data/DT_Items",
		RowName:   "Sword_01",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.RowCount != 1 {
		t.Errorf("expected 1 row, got %d", out.RowCount)
	}
}

func TestDataAssetOps_ImportCSV(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["source_path"] != "/tmp/items.csv" {
			t.Errorf("expected source_path, got %v", body["source_path"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(DataAssetOpsOutput{
			Success:  true,
			Asset:    "/Game/Data/DT_Items",
			RowCount: 50,
			Message:  "Imported 50 rows from CSV",
		})
	}))
	defer ts.Close()

	h := newDataAssetTestHandler(ts)
	_, out, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{
		Operation:  "import_csv",
		Asset:      "/Game/Data/DT_Items",
		SourcePath: "/tmp/items.csv",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.RowCount != 50 {
		t.Errorf("expected 50 rows, got %d", out.RowCount)
	}
}

func TestDataAssetOps_MissingOperation(t *testing.T) {
	h := newDataAssetTestHandler(nil)
	_, _, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required', got: %v", err)
	}
}

func TestDataAssetOps_PluginOffline(t *testing.T) {
	h := newDataAssetTestHandler(nil)
	_, _, err := h.DataAssetOps(context.Background(), &mcp.CallToolRequest{}, DataAssetOpsInput{
		Operation: "list_tables",
		Path:      "/Game/Data/",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
