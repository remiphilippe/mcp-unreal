package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestPCGOps_Execute(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/pcg/ops" {
			t.Errorf("expected /api/pcg/ops, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "execute" {
			t.Errorf("expected execute, got %v", body["operation"])
		}
		if body["actor_path"] != "/Game/Map:PersistentLevel.PCGActor_0" {
			t.Errorf("unexpected actor_path: %v", body["actor_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"generated_count":42}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation: "execute",
		ActorPath: "/Game/Map:PersistentLevel.PCGActor_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestPCGOps_Cleanup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "cleanup" {
			t.Errorf("expected cleanup, got %v", body["operation"])
		}
		if body["actor_path"] != "/Game/Map:PersistentLevel.PCGActor_0" {
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

	_, _, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation: "cleanup",
		ActorPath: "/Game/Map:PersistentLevel.PCGActor_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPCGOps_GetGraphInfo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "get_graph_info" {
			t.Errorf("expected get_graph_info, got %v", body["operation"])
		}
		if body["graph_path"] != "/Game/PCG/PCG_Foliage" {
			t.Errorf("unexpected graph_path: %v", body["graph_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"nodes":[{"id":"Node_0","title":"Surface Sampler"}],"edges":[]}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation: "get_graph_info",
		GraphPath: "/Game/PCG/PCG_Foliage",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestPCGOps_SetParameter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "set_parameter" {
			t.Errorf("expected set_parameter, got %v", body["operation"])
		}
		if body["parameter_name"] != "Density" {
			t.Errorf("expected parameter_name Density, got %v", body["parameter_name"])
		}
		if body["parameter_value"] != float64(0.5) {
			t.Errorf("expected parameter_value 0.5, got %v", body["parameter_value"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation:      "set_parameter",
		ActorPath:      "/Game/Map:PersistentLevel.PCGActor_0",
		ParameterName:  "Density",
		ParameterValue: 0.5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPCGOps_AddNode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "add_node" {
			t.Errorf("expected add_node, got %v", body["operation"])
		}
		if body["graph_path"] != "/Game/PCG/PCG_Foliage" {
			t.Errorf("unexpected graph_path: %v", body["graph_path"])
		}
		if body["node_type"] != "PCGSurfaceSamplerSettings" {
			t.Errorf("unexpected node_type: %v", body["node_type"])
		}
		if body["node_label"] != "Ground Sampler" {
			t.Errorf("unexpected node_label: %v", body["node_label"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"node_id":"Node_1"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation: "add_node",
		GraphPath: "/Game/PCG/PCG_Foliage",
		NodeType:  "PCGSurfaceSamplerSettings",
		NodeLabel: "Ground Sampler",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestPCGOps_ConnectNodes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "connect_nodes" {
			t.Errorf("expected connect_nodes, got %v", body["operation"])
		}
		if body["node_id"] != "Node_0" {
			t.Errorf("unexpected node_id: %v", body["node_id"])
		}
		if body["target_node_id"] != "Node_1" {
			t.Errorf("unexpected target_node_id: %v", body["target_node_id"])
		}
		if body["source_pin_label"] != "Out" {
			t.Errorf("unexpected source_pin_label: %v", body["source_pin_label"])
		}
		if body["target_pin_label"] != "In" {
			t.Errorf("unexpected target_pin_label: %v", body["target_pin_label"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation:      "connect_nodes",
		GraphPath:      "/Game/PCG/PCG_Foliage",
		NodeID:         "Node_0",
		TargetNodeID:   "Node_1",
		SourcePinLabel: "Out",
		TargetPinLabel: "In",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPCGOps_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["seed"] != float64(12345) {
			t.Errorf("expected seed=12345 from Params merge, got %v", body["seed"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation: "execute",
		ActorPath: "/Game/Map:PersistentLevel.PCGActor_0",
		Params:    json.RawMessage(`{"seed":12345}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPCGOps_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestPCGOps_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.PCGOps(context.Background(), &mcp.CallToolRequest{}, PCGOpsInput{
		Operation: "execute",
		ActorPath: "/Game/Map:PersistentLevel.PCGActor_0",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
