// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

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

func newNetworkDebugTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestNetworkDebug_ListActive(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/network/debug" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_active" {
			t.Errorf("expected list_active, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(NetworkDebugOutput{
			Count: 2,
			ActiveRequests: []HTTPRequestInfo{
				{URL: "https://api.weather.gov/alerts", Method: "GET"},
			},
			WebSockets: []WebSocketInfo{
				{URL: "wss://stream.aisstream.io/v0/stream", State: "connected"},
			},
		})
	}))
	defer ts.Close()

	h := newNetworkDebugTestHandler(ts)
	_, out, err := h.NetworkDebug(context.Background(), &mcp.CallToolRequest{}, NetworkDebugInput{
		Operation: "list_active",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 2 {
		t.Errorf("expected 2, got %d", out.Count)
	}
	if len(out.ActiveRequests) != 1 {
		t.Errorf("expected 1 active request, got %d", len(out.ActiveRequests))
	}
	if len(out.WebSockets) != 1 {
		t.Errorf("expected 1 websocket, got %d", len(out.WebSockets))
	}
}

func TestNetworkDebug_RecentRequests(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["last_n"] != float64(10) {
			t.Errorf("expected last_n 10, got %v", body["last_n"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(NetworkDebugOutput{
			Count: 3,
			RecentRequests: []HTTPRequestInfo{
				{URL: "https://tiles.noaa.gov/tile/1/0/0.png", Method: "GET", StatusCode: 200, DurationMs: 150},
				{URL: "https://api.weather.gov/stations", Method: "GET", StatusCode: 200, DurationMs: 230},
				{URL: "https://api.weather.gov/alerts", Method: "GET", StatusCode: 503, Error: "Service Unavailable"},
			},
		})
	}))
	defer ts.Close()

	h := newNetworkDebugTestHandler(ts)
	_, out, err := h.NetworkDebug(context.Background(), &mcp.CallToolRequest{}, NetworkDebugInput{
		Operation: "recent_requests",
		LastN:     10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 3 {
		t.Errorf("expected 3, got %d", out.Count)
	}
	if out.RecentRequests[2].StatusCode != 503 {
		t.Errorf("expected 503, got %d", out.RecentRequests[2].StatusCode)
	}
}

func TestNetworkDebug_WebSocketStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "websocket_status" {
			t.Errorf("expected websocket_status, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"websockets": []any{},
			"count":      0,
			"note":       "WebSocket tracking requires per-connection instrumentation.",
		})
	}))
	defer ts.Close()

	h := newNetworkDebugTestHandler(ts)
	_, out, err := h.NetworkDebug(context.Background(), &mcp.CallToolRequest{}, NetworkDebugInput{
		Operation: "websocket_status",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 0 {
		t.Errorf("expected 0, got %d", out.Count)
	}
}

func TestNetworkDebug_Summary(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "summary" {
			t.Errorf("expected summary, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"total_tracked":42,"active_count":3,"auto_tracking_enabled":true}`))
	}))
	defer ts.Close()

	h := newNetworkDebugTestHandler(ts)
	_, _, err := h.NetworkDebug(context.Background(), &mcp.CallToolRequest{}, NetworkDebugInput{
		Operation: "summary",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNetworkDebug_MissingOperation(t *testing.T) {
	h := newNetworkDebugTestHandler(nil)
	_, _, err := h.NetworkDebug(context.Background(), &mcp.CallToolRequest{}, NetworkDebugInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required', got: %v", err)
	}
}

func TestNetworkDebug_PluginOffline(t *testing.T) {
	h := newNetworkDebugTestHandler(nil)
	_, _, err := h.NetworkDebug(context.Background(), &mcp.CallToolRequest{}, NetworkDebugInput{
		Operation: "list_active",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
