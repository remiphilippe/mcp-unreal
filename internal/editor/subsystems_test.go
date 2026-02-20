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

func newSubsystemTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestSubsystemQuery_World(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/subsystems/query" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["type"] != "world" {
			t.Errorf("expected type world, got %v", body["type"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(SubsystemQueryOutput{
			Count: 2,
			Subsystems: []SubsystemInfo{
				{Class: "PelorusBaseMapSubsystem", Type: "world", Initialized: true},
				{Class: "UWorldPartitionSubsystem", Type: "world", Initialized: true},
			},
		})
	}))
	defer ts.Close()

	h := newSubsystemTestHandler(ts)
	_, out, err := h.SubsystemQuery(context.Background(), &mcp.CallToolRequest{}, SubsystemQueryInput{
		Type: "world",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 2 {
		t.Errorf("expected 2, got %d", out.Count)
	}
	if out.Subsystems[0].Class != "PelorusBaseMapSubsystem" {
		t.Errorf("unexpected class: %s", out.Subsystems[0].Class)
	}
	if !out.Subsystems[0].Initialized {
		t.Error("expected initialized")
	}
}

func TestSubsystemQuery_All(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["type"] != "all" {
			t.Errorf("expected all, got %v", body["type"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(SubsystemQueryOutput{
			Count: 5,
			Subsystems: []SubsystemInfo{
				{Class: "PelorusBaseMapSubsystem", Type: "world", Initialized: true},
				{Class: "PelorusUISubsystem", Type: "game_instance", Initialized: true},
				{Class: "UAssetEditorSubsystem", Type: "editor", Initialized: true},
				{Class: "UEngineSubsystem1", Type: "engine", Initialized: true},
				{Class: "ULocalPlayerSubsystem1", Type: "local_player", Initialized: false},
			},
		})
	}))
	defer ts.Close()

	h := newSubsystemTestHandler(ts)
	_, out, err := h.SubsystemQuery(context.Background(), &mcp.CallToolRequest{}, SubsystemQueryInput{
		Type: "all",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 5 {
		t.Errorf("expected 5, got %d", out.Count)
	}
}

func TestSubsystemQuery_MissingType(t *testing.T) {
	h := newSubsystemTestHandler(nil)
	_, _, err := h.SubsystemQuery(context.Background(), &mcp.CallToolRequest{}, SubsystemQueryInput{})
	if err == nil {
		t.Fatal("expected error for missing type")
	}
	if !strings.Contains(err.Error(), "type is required") {
		t.Errorf("expected 'type is required', got: %v", err)
	}
}

func TestSubsystemQuery_PluginOffline(t *testing.T) {
	h := newSubsystemTestHandler(nil)
	_, _, err := h.SubsystemQuery(context.Background(), &mcp.CallToolRequest{}, SubsystemQueryInput{
		Type: "world",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
