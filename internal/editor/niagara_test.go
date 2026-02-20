package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestNiagaraOps_SpawnSystem(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/niagara/ops" {
			t.Errorf("expected /api/niagara/ops, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "spawn_system" {
			t.Errorf("expected spawn_system, got %v", body["operation"])
		}
		if body["system_path"] != "/Game/VFX/NS_Fire" {
			t.Errorf("unexpected system_path: %v", body["system_path"])
		}
		// Verify location.
		loc, ok := body["location"].([]any)
		if !ok || len(loc) != 3 {
			t.Errorf("expected location array, got %v", body["location"])
		}
		if body["actor_name"] != "FireVFX" {
			t.Errorf("unexpected actor_name: %v", body["actor_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"actor_path":"/Game/Map:PersistentLevel.NiagaraActor_0"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:  "spawn_system",
		SystemPath: "/Game/VFX/NS_Fire",
		Location:   [3]float64{100, 200, 300},
		ActorName:  "FireVFX",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestNiagaraOps_SetParameter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "set_parameter" {
			t.Errorf("expected set_parameter, got %v", body["operation"])
		}
		if body["parameter_name"] != "SpawnRate" {
			t.Errorf("expected parameter_name SpawnRate, got %v", body["parameter_name"])
		}
		if body["parameter_value"] != float64(500) {
			t.Errorf("expected parameter_value 500, got %v", body["parameter_value"])
		}
		if body["parameter_type"] != "float" {
			t.Errorf("expected parameter_type float, got %v", body["parameter_type"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:      "set_parameter",
		ActorPath:      "/Game/Map:PersistentLevel.NiagaraActor_0",
		ParameterName:  "SpawnRate",
		ParameterValue: float64(500),
		ParameterType:  "float",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNiagaraOps_GetSystemInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "get_system_info" {
			t.Errorf("expected get_system_info, got %v", body["operation"])
		}
		if body["system_path"] != "/Game/VFX/NS_Fire" {
			t.Errorf("unexpected system_path: %v", body["system_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"emitters":[{"name":"Sparks","enabled":true}],"parameters":[{"name":"SpawnRate","type":"float"}]}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:  "get_system_info",
		SystemPath: "/Game/VFX/NS_Fire",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestNiagaraOps_AddEmitter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "add_emitter" {
			t.Errorf("expected add_emitter, got %v", body["operation"])
		}
		if body["system_path"] != "/Game/VFX/NS_Fire" {
			t.Errorf("unexpected system_path: %v", body["system_path"])
		}
		if body["emitter_path"] != "/Game/VFX/Emitters/NE_Sparks" {
			t.Errorf("unexpected emitter_path: %v", body["emitter_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"emitter_name":"Sparks"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:   "add_emitter",
		SystemPath:  "/Game/VFX/NS_Fire",
		EmitterPath: "/Game/VFX/Emitters/NE_Sparks",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestNiagaraOps_RemoveEmitter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "remove_emitter" {
			t.Errorf("expected remove_emitter, got %v", body["operation"])
		}
		if body["emitter_name"] != "Sparks" {
			t.Errorf("unexpected emitter_name: %v", body["emitter_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:   "remove_emitter",
		SystemPath:  "/Game/VFX/NS_Fire",
		EmitterName: "Sparks",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNiagaraOps_Activate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "activate" {
			t.Errorf("expected activate, got %v", body["operation"])
		}
		if body["actor_path"] != "/Game/Map:PersistentLevel.NiagaraActor_0" {
			t.Errorf("unexpected actor_path: %v", body["actor_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation: "activate",
		ActorPath: "/Game/Map:PersistentLevel.NiagaraActor_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNiagaraOps_Deactivate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "deactivate" {
			t.Errorf("expected deactivate, got %v", body["operation"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation: "deactivate",
		ActorPath: "/Game/Map:PersistentLevel.NiagaraActor_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNiagaraOps_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["warmup_time"] != float64(2.5) {
			t.Errorf("expected warmup_time=2.5 from Params merge, got %v", body["warmup_time"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:  "spawn_system",
		SystemPath: "/Game/VFX/NS_Fire",
		Params:     json.RawMessage(`{"warmup_time":2.5}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNiagaraOps_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestNiagaraOps_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.NiagaraOps(context.Background(), &mcp.CallToolRequest{}, NiagaraOpsInput{
		Operation:  "spawn_system",
		SystemPath: "/Game/VFX/NS_Fire",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
