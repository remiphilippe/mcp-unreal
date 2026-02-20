package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestInputOps_ListActions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/input/ops" {
			t.Errorf("expected /api/input/ops, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_actions" {
			t.Errorf("expected list_actions, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"actions":[{"name":"IA_Move","path":"/Game/Input/IA_Move"}]}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation: "list_actions",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestInputOps_BindAction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "bind_action" {
			t.Errorf("expected bind_action, got %v", body["operation"])
		}
		if body["key"] != "W" {
			t.Errorf("expected key W, got %v", body["key"])
		}
		mods, ok := body["modifiers"].([]any)
		if !ok || len(mods) != 1 || mods[0] != "Negate" {
			t.Errorf("expected modifiers [Negate], got %v", body["modifiers"])
		}
		triggers, ok := body["triggers"].([]any)
		if !ok || len(triggers) != 1 || triggers[0] != "Pressed" {
			t.Errorf("expected triggers [Pressed], got %v", body["triggers"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation:  "bind_action",
		AssetPath:  "/Game/Input/IMC_Default",
		ActionName: "IA_MoveForward",
		Key:        "W",
		Modifiers:  []string{"Negate"},
		Triggers:   []string{"Pressed"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInputOps_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["extra_field"] != "extra_value" {
			t.Errorf("expected extra_field=extra_value from Params merge, got %v", body["extra_field"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation: "list_actions",
		Params:    json.RawMessage(`{"extra_field":"extra_value"}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInputOps_AddAction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "add_action" {
			t.Errorf("expected add_action, got %v", body["operation"])
		}
		if body["action_name"] != "IA_Sprint" {
			t.Errorf("expected IA_Sprint, got %v", body["action_name"])
		}
		if body["value_type"] != "bool" {
			t.Errorf("expected bool, got %v", body["value_type"])
		}
		if body["package_path"] != "/Game/Input" {
			t.Errorf("expected /Game/Input, got %v", body["package_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"path":"/Game/Input/IA_Sprint","name":"IA_Sprint"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation:   "add_action",
		ActionName:  "IA_Sprint",
		ValueType:   "bool",
		PackagePath: "/Game/Input",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestInputOps_RemoveAction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "remove_action" {
			t.Errorf("expected remove_action, got %v", body["operation"])
		}
		if body["asset_path"] != "/Game/Input/IA_Sprint" {
			t.Errorf("expected /Game/Input/IA_Sprint, got %v", body["asset_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"deleted_count":1}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation: "remove_action",
		AssetPath: "/Game/Input/IA_Sprint",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInputOps_AddContext(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "add_context" {
			t.Errorf("expected add_context, got %v", body["operation"])
		}
		if body["context_name"] != "IMC_Combat" {
			t.Errorf("expected IMC_Combat, got %v", body["context_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"path":"/Game/Input/IMC_Combat","name":"IMC_Combat"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation:   "add_context",
		ContextName: "IMC_Combat",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInputOps_UnbindAction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "unbind_action" {
			t.Errorf("expected unbind_action, got %v", body["operation"])
		}
		if body["asset_path"] != "/Game/Input/IMC_Default" {
			t.Errorf("expected /Game/Input/IMC_Default, got %v", body["asset_path"])
		}
		if body["action_name"] != "IA_MoveForward" {
			t.Errorf("expected IA_MoveForward, got %v", body["action_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"action":"IA_MoveForward","context":"IMC_Default"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation:  "unbind_action",
		AssetPath:  "/Game/Input/IMC_Default",
		ActionName: "IA_MoveForward",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestInputOps_GetBindings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "get_bindings" {
			t.Errorf("expected get_bindings, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"bindings":[{"action":"IA_Move","key":"W"}],"count":1}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation: "get_bindings",
		AssetPath: "/Game/Input/IMC_Default",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestInputOps_ListContexts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_contexts" {
			t.Errorf("expected list_contexts, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"contexts":[{"name":"IMC_Default","path":"/Game/Input/IMC_Default"}],"count":1}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation: "list_contexts",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestInputOps_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestInputOps_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.InputOps(context.Background(), &mcp.CallToolRequest{}, InputOpsInput{
		Operation: "list_actions",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
