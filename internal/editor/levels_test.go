package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestLevelOps_GetCurrent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/levels/ops" {
			t.Errorf("expected /api/levels/ops, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "get_current" {
			t.Errorf("expected get_current, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"level_name":"MainLevel","level_path":"/Game/Maps/MainLevel"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.LevelOps(context.Background(), &mcp.CallToolRequest{}, LevelOpsInput{
		Operation: "get_current",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestLevelOps_NewLevel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "new_level" {
			t.Errorf("expected new_level, got %v", body["operation"])
		}
		if body["level_name"] != "TestLevel" {
			t.Errorf("expected level_name TestLevel, got %v", body["level_name"])
		}
		if body["package_path"] != "/Game/Maps" {
			t.Errorf("expected package_path /Game/Maps, got %v", body["package_path"])
		}
		if body["template"] != "TimeOfDay" {
			t.Errorf("expected template TimeOfDay, got %v", body["template"])
		}
		if body["streaming"] != "Blueprint" {
			t.Errorf("expected streaming Blueprint, got %v", body["streaming"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.LevelOps(context.Background(), &mcp.CallToolRequest{}, LevelOpsInput{
		Operation:   "new_level",
		LevelName:   "TestLevel",
		PackagePath: "/Game/Maps",
		Template:    "TimeOfDay",
		Streaming:   "Blueprint",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLevelOps_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["custom_flag"] != true {
			t.Errorf("expected custom_flag=true from Params merge, got %v", body["custom_flag"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.LevelOps(context.Background(), &mcp.CallToolRequest{}, LevelOpsInput{
		Operation: "get_current",
		Params:    json.RawMessage(`{"custom_flag":true}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLevelOps_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.LevelOps(context.Background(), &mcp.CallToolRequest{}, LevelOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestLevelOps_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.LevelOps(context.Background(), &mcp.CallToolRequest{}, LevelOpsInput{
		Operation: "get_current",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
