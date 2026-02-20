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

func TestBlueprintQuery_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/blueprints/list" {
			t.Errorf("expected /api/blueprints/list, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"blueprints":[{"name":"BP_Player","path":"/Game/BP_Player"}]}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "list",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestBlueprintQuery_InspectMissingPath(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "inspect",
	})
	if err == nil {
		t.Fatal("expected error for missing path")
	}
}

func TestBlueprintQuery_GetGraph(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/blueprints/get_graph" {
			t.Errorf("expected /api/blueprints/get_graph, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["graph_name"] != "EventGraph" {
			t.Errorf("expected graph_name EventGraph, got %v", body["graph_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"nodes":[]}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "get_graph",
		Path:      "/Game/BP_Test",
		GraphName: "EventGraph",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlueprintQuery_UnknownOperation(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "invalid",
	})
	if err == nil {
		t.Fatal("expected error for unknown operation")
	}
}

func TestBlueprintModify_AddVariable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "add_variable" {
			t.Errorf("expected add_variable operation, got %v", body["operation"])
		}
		if body["variable_name"] != "Health" {
			t.Errorf("expected variable_name Health, got %v", body["variable_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "add_variable",
		BlueprintPath: "/Game/BP_Test",
		VariableName:  "Health",
		VariableType:  "float",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if !out.Compiled {
		t.Error("expected compiled")
	}
}

func TestBlueprintModify_MissingOperation(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestBlueprintModify_ConnectPins(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "connect_pins" {
			t.Errorf("expected connect_pins, got %v", body["operation"])
		}
		if body["source_node_id"] != "node-1" {
			t.Errorf("expected source_node_id node-1, got %v", body["source_node_id"])
		}
		if body["source_pin"] != "ReturnValue" {
			t.Errorf("expected source_pin ReturnValue, got %v", body["source_pin"])
		}
		if body["target_node_id"] != "node-2" {
			t.Errorf("expected target_node_id node-2, got %v", body["target_node_id"])
		}
		if body["target_pin"] != "Input" {
			t.Errorf("expected target_pin Input, got %v", body["target_pin"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "connect_pins",
		BlueprintPath: "/Game/BP_Test",
		SourceNodeID:  "node-1",
		SourcePinName: "ReturnValue",
		TargetNodeID:  "node-2",
		TargetPinName: "Input",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestBlueprintQuery_GetGraph_MissingGraphName(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "get_graph",
		Path:      "/Game/BP_Test",
		// GraphName deliberately omitted.
	})
	if err == nil {
		t.Fatal("expected error for missing graph_name")
	}
}

func TestBlueprintModify_PluginOffline(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "compile",
		BlueprintPath: "/Game/BP_Test",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestBlueprintModify_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create" {
			t.Errorf("expected create, got %v", body["operation"])
		}
		if body["parent_class"] != "Actor" {
			t.Errorf("expected parent_class Actor, got %v", body["parent_class"])
		}
		if body["package_path"] != "/Game/Blueprints" {
			t.Errorf("expected package_path /Game/Blueprints, got %v", body["package_path"])
		}
		if body["name"] != "BP_NewActor" {
			t.Errorf("expected name BP_NewActor, got %v", body["name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true,"message":"created"}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "create",
		ClassName:     "Actor",
		PackagePath:   "/Game/Blueprints",
		BlueprintName: "BP_NewActor",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestBlueprintModify_AddNode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["node_class"] != "K2Node_CallFunction" {
			t.Errorf("expected node_class K2Node_CallFunction, got %v", body["node_class"])
		}
		if body["graph_name"] != "EventGraph" {
			t.Errorf("expected graph_name EventGraph, got %v", body["graph_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "add_node",
		BlueprintPath: "/Game/BP_Test",
		NodeClass:     "K2Node_CallFunction",
		GraphName:     "EventGraph",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlueprintModify_SetPinValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["node_id"] != "node-abc" {
			t.Errorf("expected node_id node-abc, got %v", body["node_id"])
		}
		if body["pin_name"] != "Value" {
			t.Errorf("expected pin_name Value, got %v", body["pin_name"])
		}
		if body["value"] != "42" {
			t.Errorf("expected value 42, got %v", body["value"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "set_pin_value",
		BlueprintPath: "/Game/BP_Test",
		NodeID:        "node-abc",
		PinName:       "Value",
		PinValue:      "42",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlueprintModify_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["custom_param"] != "custom_value" {
			t.Errorf("expected custom_param=custom_value from Params merge, got %v", body["custom_param"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":false}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "compile",
		BlueprintPath: "/Game/BP_Test",
		Params:        json.RawMessage(`{"custom_param":"custom_value"}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlueprintQuery_Inspect(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/blueprints/inspect" {
			t.Errorf("expected /api/blueprints/inspect, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["blueprint_path"] != "/Game/BP_Test" {
			t.Errorf("unexpected blueprint_path: %v", body["blueprint_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"variables":[],"functions":[]}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "inspect",
		Path:      "/Game/BP_Test",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestBlueprintQuery_PluginOffline(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{
		Operation: "list",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestBlueprintQuery_MissingOperation(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.BlueprintQuery(context.Background(), &mcp.CallToolRequest{}, BlueprintQueryInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestBlueprintModify_RemoveFunction(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["function_name"] != "MyFunction" {
			t.Errorf("expected function_name MyFunction, got %v", body["function_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"compiled":true}`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.BlueprintModify(context.Background(), &mcp.CallToolRequest{}, BlueprintModifyInput{
		Operation:     "remove_function",
		BlueprintPath: "/Game/BP_Test",
		FunctionName:  "MyFunction",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
