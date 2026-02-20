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

func newISMTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestISMOps_Create(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/ism/ops" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create" {
			t.Errorf("expected operation create, got %v", body["operation"])
		}
		if body["actor_name"] != "MyActor" {
			t.Errorf("expected actor_name MyActor, got %v", body["actor_name"])
		}
		if body["mesh"] != "/Engine/BasicShapes/Cube" {
			t.Errorf("expected mesh path, got %v", body["mesh"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "InstancedStaticMeshComponent_0",
			InstanceCount: 0,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation: "create",
		ActorName: "MyActor",
		Mesh:      "/Engine/BasicShapes/Cube",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.ComponentName != "InstancedStaticMeshComponent_0" {
		t.Errorf("expected component name, got %s", out.ComponentName)
	}
}

func TestISMOps_AddInstances(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "add_instances" {
			t.Errorf("expected add_instances, got %v", body["operation"])
		}
		transforms, ok := body["transforms"].([]any)
		if !ok || len(transforms) != 2 {
			t.Errorf("expected 2 transforms, got %v", body["transforms"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "ISM_0",
			InstanceCount: 2,
			AddedCount:    2,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "add_instances",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
		Transforms: []ISMTransform{
			{Location: [3]float64{100, 200, 300}},
			{Location: [3]float64{400, 500, 600}, Scale: [3]float64{2, 2, 2}},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.AddedCount != 2 {
		t.Errorf("expected 2 added, got %d", out.AddedCount)
	}
	if out.InstanceCount != 2 {
		t.Errorf("expected 2 instances, got %d", out.InstanceCount)
	}
}

func TestISMOps_UpdateInstance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "update_instance" {
			t.Errorf("expected update_instance, got %v", body["operation"])
		}
		if body["instance_index"] != float64(3) {
			t.Errorf("expected instance_index 3, got %v", body["instance_index"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "ISM_0",
			InstanceCount: 10,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	idx := 3
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "update_instance",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
		InstanceIndex: &idx,
		Transform:     &ISMTransform{Location: [3]float64{999, 888, 777}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.InstanceCount != 10 {
		t.Errorf("expected 10 instances, got %d", out.InstanceCount)
	}
}

func TestISMOps_ClearInstances(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "clear_instances" {
			t.Errorf("expected clear_instances, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "ISM_0",
			InstanceCount: 0,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "clear_instances",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.InstanceCount != 0 {
		t.Errorf("expected 0 instances, got %d", out.InstanceCount)
	}
}

func TestISMOps_GetInstanceCount(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "ISM_0",
			InstanceCount: 4200,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "get_instance_count",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.InstanceCount != 4200 {
		t.Errorf("expected 4200, got %d", out.InstanceCount)
	}
}

func TestISMOps_RemoveInstance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["instance_index"] != float64(5) {
			t.Errorf("expected instance_index 5, got %v", body["instance_index"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "ISM_0",
			InstanceCount: 9,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	idx := 5
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "remove_instance",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
		InstanceIndex: &idx,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.InstanceCount != 9 {
		t.Errorf("expected 9 instances, got %d", out.InstanceCount)
	}
}

func TestISMOps_SetMaterial(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["material"] != "/Game/Materials/M_Red" {
			t.Errorf("expected material path, got %v", body["material"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "ISM_0",
			InstanceCount: 10,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "set_material",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
		Material:      "/Game/Materials/M_Red",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestISMOps_UseHISM(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["use_hism"] != true {
			t.Errorf("expected use_hism true, got %v", body["use_hism"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(ISMOpsOutput{
			Success:       true,
			ComponentName: "HISM_0",
			InstanceCount: 0,
		})
	}))
	defer ts.Close()

	h := newISMTestHandler(ts)
	_, out, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation: "create",
		ActorName: "MyActor",
		UseHISM:   true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ComponentName != "HISM_0" {
		t.Errorf("expected HISM_0, got %s", out.ComponentName)
	}
}

func TestISMOps_MissingOperation(t *testing.T) {
	h := newISMTestHandler(nil)
	_, _, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required' in error, got: %v", err)
	}
}

func TestISMOps_PluginOffline(t *testing.T) {
	h := newISMTestHandler(nil)
	_, _, err := h.ISMOps(context.Background(), &mcp.CallToolRequest{}, ISMOpsInput{
		Operation:     "get_instance_count",
		ActorName:     "MyActor",
		ComponentName: "ISM_0",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
	if !strings.Contains(err.Error(), "editor unreachable") {
		t.Errorf("expected 'editor unreachable' in error, got: %v", err)
	}
}
