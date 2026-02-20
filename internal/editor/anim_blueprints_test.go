package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestAnimBlueprintQuery_ListStateMachines(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "list_state_machines" {
			t.Errorf("expected list_state_machines, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"state_machines":[{"name":"Locomotion"}]}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.AnimBlueprintQuery(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintQueryInput{
		Operation:     "list_state_machines",
		BlueprintPath: "/Game/ABP_Test",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestAnimBlueprintQuery_MissingPath(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.AnimBlueprintQuery(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintQueryInput{
		Operation: "list_state_machines",
	})
	if err == nil {
		t.Fatal("expected error for missing blueprint_path")
	}
}

func TestAnimBlueprintModify_CreateState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create_state" {
			t.Errorf("expected create_state, got %v", body["operation"])
		}
		if body["state_name"] != "Idle" {
			t.Errorf("expected state_name Idle, got %v", body["state_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation:        "create_state",
		BlueprintPath:    "/Game/ABP_Test",
		StateMachineName: "Locomotion",
		StateName:        "Idle",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestAnimBlueprintModify_MissingOperation(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		BlueprintPath: "/Game/ABP_Test",
	})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestAnimBlueprintModify_MissingBlueprintPath(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation: "create_state",
	})
	if err == nil {
		t.Fatal("expected error for missing blueprint_path")
	}
}

func TestAnimBlueprintQuery_MissingOperation(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.AnimBlueprintQuery(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintQueryInput{
		BlueprintPath: "/Game/ABP_Test",
	})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestAnimBlueprintQuery_PluginOffline(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.AnimBlueprintQuery(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintQueryInput{
		Operation:     "list_state_machines",
		BlueprintPath: "/Game/ABP_Test",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestAnimBlueprintModify_CreateTransition(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create_transition" {
			t.Errorf("expected create_transition, got %v", body["operation"])
		}
		if body["from_state"] != "Idle" {
			t.Errorf("expected from_state Idle, got %v", body["from_state"])
		}
		if body["to_state"] != "Running" {
			t.Errorf("expected to_state Running, got %v", body["to_state"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation:        "create_transition",
		BlueprintPath:    "/Game/ABP_Test",
		StateMachineName: "Locomotion",
		FromState:        "Idle",
		ToState:          "Running",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestAnimBlueprintModify_RenameState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["old_name"] != "OldState" {
			t.Errorf("expected old_name OldState, got %v", body["old_name"])
		}
		if body["new_name"] != "NewState" {
			t.Errorf("expected new_name NewState, got %v", body["new_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation:        "rename_state",
		BlueprintPath:    "/Game/ABP_Test",
		StateMachineName: "Locomotion",
		OldName:          "OldState",
		NewName:          "NewState",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAnimBlueprintModify_DeleteTransition(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["transition_id"] != "trans-abc-123" {
			t.Errorf("expected transition_id trans-abc-123, got %v", body["transition_id"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation:        "delete_transition",
		BlueprintPath:    "/Game/ABP_Test",
		StateMachineName: "Locomotion",
		TransitionID:     "trans-abc-123",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAnimBlueprintModify_AddAnimNode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["node_class"] != "AnimNodeSequencePlayer" {
			t.Errorf("expected node_class AnimNodeSequencePlayer, got %v", body["node_class"])
		}
		if body["node_id"] != "node-to-delete" {
			t.Errorf("expected node_id node-to-delete, got %v", body["node_id"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation:        "add_anim_node",
		BlueprintPath:    "/Game/ABP_Test",
		StateMachineName: "Locomotion",
		NodeClass:        "AnimNodeSequencePlayer",
		NodeID:           "node-to-delete",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestAnimBlueprintModify_PluginOffline(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.AnimBlueprintModify(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintModifyInput{
		Operation:     "compile",
		BlueprintPath: "/Game/ABP_Test",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestAnimBlueprintQuery_InspectStateMachine(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["state_machine_name"] != "Locomotion" {
			t.Errorf("expected state_machine_name Locomotion, got %v", body["state_machine_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"states":["Idle","Running"]}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.AnimBlueprintQuery(context.Background(), &mcp.CallToolRequest{}, AnimBlueprintQueryInput{
		Operation:        "inspect_state_machine",
		BlueprintPath:    "/Game/ABP_Test",
		StateMachineName: "Locomotion",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}
