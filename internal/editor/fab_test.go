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

func newFabTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestFabOps_ListCache(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/fab/ops" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_cache" {
			t.Errorf("expected list_cache, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(FabOpsOutput{
			Success:    true,
			AssetCount: 2,
			Assets: []FabCachedAsset{
				{AssetID: "abc123", FilePath: "/tmp/fab/abc123.zip"},
				{AssetID: "def456", FilePath: "/tmp/fab/def456.zip"},
			},
		})
	}))
	defer ts.Close()

	h := newFabTestHandler(ts)
	_, out, err := h.FabOps(context.Background(), &mcp.CallToolRequest{}, FabOpsInput{
		Operation: "list_cache",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.AssetCount != 2 {
		t.Errorf("expected 2 assets, got %d", out.AssetCount)
	}
	if len(out.Assets) != 2 {
		t.Fatalf("expected 2 assets in list, got %d", len(out.Assets))
	}
	if out.Assets[0].AssetID != "abc123" {
		t.Errorf("expected abc123, got %s", out.Assets[0].AssetID)
	}
}

func TestFabOps_CacheInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "cache_info" {
			t.Errorf("expected cache_info, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(FabOpsOutput{
			Success:        true,
			CacheLocation:  "/tmp/FabLibrary",
			CacheSize:      "2.3 GB",
			CacheSizeBytes: 2469606195,
			AssetCount:     15,
		})
	}))
	defer ts.Close()

	h := newFabTestHandler(ts)
	_, out, err := h.FabOps(context.Background(), &mcp.CallToolRequest{}, FabOpsInput{
		Operation: "cache_info",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.CacheLocation != "/tmp/FabLibrary" {
		t.Errorf("expected /tmp/FabLibrary, got %s", out.CacheLocation)
	}
	if out.CacheSizeBytes != 2469606195 {
		t.Errorf("expected cache size bytes, got %d", out.CacheSizeBytes)
	}
}

func TestFabOps_Import(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "import" {
			t.Errorf("expected import, got %v", body["operation"])
		}
		if body["asset_id"] != "abc123" {
			t.Errorf("expected asset_id abc123, got %v", body["asset_id"])
		}
		if body["destination"] != "/Game/Assets/Castle" {
			t.Errorf("expected destination /Game/Assets/Castle, got %v", body["destination"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(FabOpsOutput{
			Success:      true,
			ImportedPath: "/Game/Assets/Castle",
			Message:      "Imported 12 assets",
		})
	}))
	defer ts.Close()

	h := newFabTestHandler(ts)
	_, out, err := h.FabOps(context.Background(), &mcp.CallToolRequest{}, FabOpsInput{
		Operation:   "import",
		AssetID:     "abc123",
		Destination: "/Game/Assets/Castle",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.ImportedPath != "/Game/Assets/Castle" {
		t.Errorf("expected /Game/Assets/Castle, got %s", out.ImportedPath)
	}
}

func TestFabOps_ClearCache(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "clear_cache" {
			t.Errorf("expected clear_cache, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(FabOpsOutput{
			Success: true,
			Message: "Cache cleared",
		})
	}))
	defer ts.Close()

	h := newFabTestHandler(ts)
	_, out, err := h.FabOps(context.Background(), &mcp.CallToolRequest{}, FabOpsInput{
		Operation: "clear_cache",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestFabOps_MissingOperation(t *testing.T) {
	h := newFabTestHandler(nil)
	_, _, err := h.FabOps(context.Background(), &mcp.CallToolRequest{}, FabOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required' in error, got: %v", err)
	}
}

func TestFabOps_PluginOffline(t *testing.T) {
	h := newFabTestHandler(nil)
	_, _, err := h.FabOps(context.Background(), &mcp.CallToolRequest{}, FabOpsInput{
		Operation: "list_cache",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
	if !strings.Contains(err.Error(), "editor unreachable") {
		t.Errorf("expected 'editor unreachable' in error, got: %v", err)
	}
}
