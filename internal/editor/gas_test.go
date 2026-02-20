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

func TestGASOps_GrantAbility(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/gas/ops" {
			t.Errorf("expected /api/gas/ops, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "grant_ability" {
			t.Errorf("expected grant_ability, got %v", body["operation"])
		}
		if body["actor_path"] != "/Game/Map:PersistentLevel.Hero_0" {
			t.Errorf("unexpected actor_path: %v", body["actor_path"])
		}
		if body["ability_class"] != "/Game/Abilities/GA_FireBlast.GA_FireBlast_C" {
			t.Errorf("unexpected ability_class: %v", body["ability_class"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"ability_spec_handle":42}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation:    "grant_ability",
		ActorPath:    "/Game/Map:PersistentLevel.Hero_0",
		AbilityClass: "/Game/Abilities/GA_FireBlast.GA_FireBlast_C",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestGASOps_RevokeAbility(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "revoke_ability" {
			t.Errorf("expected revoke_ability, got %v", body["operation"])
		}
		if body["ability_tag"] != "Ability.FireBlast" {
			t.Errorf("unexpected ability_tag: %v", body["ability_tag"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"revoked_count":1}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation:  "revoke_ability",
		ActorPath:  "/Game/Map:PersistentLevel.Hero_0",
		AbilityTag: "Ability.FireBlast",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGASOps_ListAbilities(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_abilities" {
			t.Errorf("expected list_abilities, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"abilities":[{"class":"GA_FireBlast_C","level":1,"active":false}],"count":1}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation: "list_abilities",
		ActorPath: "/Game/Map:PersistentLevel.Hero_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestGASOps_ApplyEffect(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "apply_effect" {
			t.Errorf("expected apply_effect, got %v", body["operation"])
		}
		if body["effect_class"] != "/Game/Effects/GE_Heal.GE_Heal_C" {
			t.Errorf("unexpected effect_class: %v", body["effect_class"])
		}
		if body["effect_level"] != float64(3) {
			t.Errorf("expected effect_level 3, got %v", body["effect_level"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	effectLevel := float64(3)
	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation:   "apply_effect",
		ActorPath:   "/Game/Map:PersistentLevel.Hero_0",
		EffectClass: "/Game/Effects/GE_Heal.GE_Heal_C",
		EffectLevel: &effectLevel,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGASOps_GetAttributes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "get_attributes" {
			t.Errorf("expected get_attributes, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"attribute_sets":[{"name":"UHealthSet","attributes":[{"name":"Health","base":100,"current":85}]}]}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation: "get_attributes",
		ActorPath: "/Game/Map:PersistentLevel.Hero_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestGASOps_SetAttribute(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "set_attribute" {
			t.Errorf("expected set_attribute, got %v", body["operation"])
		}
		if body["attribute_name"] != "Health" {
			t.Errorf("expected attribute_name Health, got %v", body["attribute_name"])
		}
		if body["attribute_value"] != float64(100) {
			t.Errorf("expected attribute_value 100, got %v", body["attribute_value"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	attrValue := float64(100)
	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation:      "set_attribute",
		ActorPath:      "/Game/Map:PersistentLevel.Hero_0",
		AttributeName:  "Health",
		AttributeValue: &attrValue,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGASOps_SetAttributeZeroValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["attribute_value"] != float64(0) {
			t.Errorf("expected attribute_value 0, got %v", body["attribute_value"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	zeroValue := float64(0)
	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation:      "set_attribute",
		ActorPath:      "/Game/Map:PersistentLevel.Hero_0",
		AttributeName:  "Health",
		AttributeValue: &zeroValue,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGASOps_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["instigator_path"] != "/Game/Map:PersistentLevel.Enemy_0" {
			t.Errorf("expected instigator_path from Params merge, got %v", body["instigator_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation: "apply_effect",
		ActorPath: "/Game/Map:PersistentLevel.Hero_0",
		Params:    json.RawMessage(`{"instigator_path":"/Game/Map:PersistentLevel.Enemy_0"}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGASOps_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestGASOps_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.GASOps(context.Background(), &mcp.CallToolRequest{}, GASOpsInput{
		Operation: "list_abilities",
		ActorPath: "/Game/Map:PersistentLevel.Hero_0",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
