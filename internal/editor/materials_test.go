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

func TestMaterialOps_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/materials/ops" {
			t.Errorf("expected /api/materials/ops, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create" {
			t.Errorf("expected create, got %v", body["operation"])
		}
		if body["package_path"] != "/Game/Materials" {
			t.Errorf("expected package_path /Game/Materials, got %v", body["package_path"])
		}
		if body["material_name"] != "M_TestMaterial" {
			t.Errorf("expected material_name M_TestMaterial, got %v", body["material_name"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"path":"/Game/Materials/M_TestMaterial"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.MaterialOps(context.Background(), &mcp.CallToolRequest{}, MaterialOpsInput{
		Operation:    "create",
		PackagePath:  "/Game/Materials",
		MaterialName: "M_TestMaterial",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-nil result")
	}
}

func TestMaterialOps_SetParameter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "set_parameter" {
			t.Errorf("expected set_parameter, got %v", body["operation"])
		}
		if body["material_path"] != "/Game/Materials/M_Test" {
			t.Errorf("unexpected material_path: %v", body["material_path"])
		}
		if body["parameter_name"] != "Roughness" {
			t.Errorf("expected parameter_name Roughness, got %v", body["parameter_name"])
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

	_, _, err := h.MaterialOps(context.Background(), &mcp.CallToolRequest{}, MaterialOpsInput{
		Operation:      "set_parameter",
		MaterialPath:   "/Game/Materials/M_Test",
		ParameterName:  "Roughness",
		ParameterValue: 0.5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMaterialOps_SetParameterWithColor(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		color, ok := body["color"].([]any)
		if !ok {
			t.Fatal("expected color array in body")
		}
		if len(color) != 4 {
			t.Fatalf("expected 4 color components, got %d", len(color))
		}
		if color[0].(float64) != 1.0 || color[1].(float64) != 0.0 || color[2].(float64) != 0.0 || color[3].(float64) != 1.0 {
			t.Errorf("unexpected color values: %v", color)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.MaterialOps(context.Background(), &mcp.CallToolRequest{}, MaterialOpsInput{
		Operation:     "set_parameter",
		MaterialPath:  "/Game/Materials/M_Test",
		ParameterName: "BaseColor",
		Color:         [4]float64{1.0, 0.0, 0.0, 1.0},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMaterialOps_WithExtraParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["blend_mode"] != "Translucent" {
			t.Errorf("expected blend_mode=Translucent from Params merge, got %v", body["blend_mode"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.MaterialOps(context.Background(), &mcp.CallToolRequest{}, MaterialOpsInput{
		Operation: "create",
		Params:    json.RawMessage(`{"blend_mode":"Translucent"}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMaterialOps_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.MaterialOps(context.Background(), &mcp.CallToolRequest{}, MaterialOpsInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestMaterialOps_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.MaterialOps(context.Background(), &mcp.CallToolRequest{}, MaterialOpsInput{
		Operation: "create",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
