package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestCharacterConfig_GetConfig(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/characters/config" {
			t.Errorf("expected /api/characters/config, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "get_config" {
			t.Errorf("expected get_config, got %v", body["operation"])
		}
		if body["blueprint_path"] != "/Game/Characters/BP_Player" {
			t.Errorf("unexpected blueprint_path: %v", body["blueprint_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"max_walk_speed":600,"jump_z_velocity":420}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.CharacterConfig(context.Background(), &mcp.CallToolRequest{}, CharacterConfigInput{
		Operation:     "get_config",
		BlueprintPath: "/Game/Characters/BP_Player",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestCharacterConfig_SetMovement(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "set_movement" {
			t.Errorf("expected set_movement, got %v", body["operation"])
		}
		if body["max_walk_speed"] != float64(800) {
			t.Errorf("expected max_walk_speed 800, got %v", body["max_walk_speed"])
		}
		if body["jump_z_velocity"] != float64(500) {
			t.Errorf("expected jump_z_velocity 500, got %v", body["jump_z_velocity"])
		}
		if body["gravity_scale"] != float64(1.5) {
			t.Errorf("expected gravity_scale 1.5, got %v", body["gravity_scale"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	maxWalk := 800.0
	jumpZ := 500.0
	gravity := 1.5
	_, _, err := h.CharacterConfig(context.Background(), &mcp.CallToolRequest{}, CharacterConfigInput{
		Operation:     "set_movement",
		BlueprintPath: "/Game/Characters/BP_Player",
		MaxWalkSpeed:  &maxWalk,
		JumpZVelocity: &jumpZ,
		GravityScale:  &gravity,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCharacterConfig_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["custom_key"] != "custom_value" {
			t.Errorf("expected custom_key=custom_value from Params merge, got %v", body["custom_key"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.CharacterConfig(context.Background(), &mcp.CallToolRequest{}, CharacterConfigInput{
		Operation:     "get_config",
		BlueprintPath: "/Game/Characters/BP_Player",
		Params:        json.RawMessage(`{"custom_key":"custom_value"}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCharacterConfig_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CharacterConfig(context.Background(), &mcp.CallToolRequest{}, CharacterConfigInput{
		BlueprintPath: "/Game/Characters/BP_Player",
	})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestCharacterConfig_MissingBlueprintPath(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CharacterConfig(context.Background(), &mcp.CallToolRequest{}, CharacterConfigInput{
		Operation: "get_config",
	})
	if err == nil {
		t.Fatal("expected error for missing blueprint_path")
	}
}

func TestCharacterConfig_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CharacterConfig(context.Background(), &mcp.CallToolRequest{}, CharacterConfigInput{
		Operation:     "get_config",
		BlueprintPath: "/Game/Characters/BP_Player",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
