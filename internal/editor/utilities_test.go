// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestRunConsoleCommand_PluginSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/editor/console_command" {
			t.Errorf("expected /api/editor/console_command, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["command"] != "stat fps" {
			t.Errorf("expected command 'stat fps', got %v", body["command"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]string{"output": "FPS: 60.0"})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.RunConsoleCommand(context.Background(), &mcp.CallToolRequest{}, RunConsoleCommandInput{
		Command: "stat fps",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Command != "stat fps" {
		t.Errorf("expected command 'stat fps', got %s", out.Command)
	}
	if out.Output != "FPS: 60.0" {
		t.Errorf("expected output 'FPS: 60.0', got %s", out.Output)
	}
}

func TestRunConsoleCommand_FallbackToRCAPI(t *testing.T) {
	// RC API server responds to the /remote/object/call endpoint.
	rcServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/remote/object/call" {
			t.Errorf("expected /remote/object/call, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["functionName"] != "ExecuteConsoleCommand" {
			t.Errorf("expected ExecuteConsoleCommand, got %v", body["functionName"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer rcServer.Close()

	// Plugin is offline (port 1), RC API is online.
	h := &Handler{
		Client: newTestClient(rcServer.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, out, err := h.RunConsoleCommand(context.Background(), &mcp.CallToolRequest{}, RunConsoleCommandInput{
		Command: "stat unit",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success via RC API fallback")
	}
	if out.Command != "stat unit" {
		t.Errorf("expected command 'stat unit', got %s", out.Command)
	}
}

func TestRunConsoleCommand_BothOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.RunConsoleCommand(context.Background(), &mcp.CallToolRequest{}, RunConsoleCommandInput{
		Command: "stat fps",
	})
	if err == nil {
		t.Fatal("expected error when both plugin and RC API are offline")
	}
}

func TestRunConsoleCommand_MissingCommand(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.RunConsoleCommand(context.Background(), &mcp.CallToolRequest{}, RunConsoleCommandInput{})
	if err == nil {
		t.Fatal("expected error for missing command")
	}
}
