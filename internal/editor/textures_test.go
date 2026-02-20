// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

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

func newTextureTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestTextureOps_Import(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/textures/ops" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "import" {
			t.Errorf("expected import, got %v", body["operation"])
		}
		if body["source_path"] != "/tmp/terrain.png" {
			t.Errorf("expected source_path, got %v", body["source_path"])
		}
		if body["compression"] != "TC_Default" {
			t.Errorf("expected TC_Default, got %v", body["compression"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(TextureOpsOutput{
			Success: true,
			Asset:   "/Game/Textures/T_Terrain",
			Message: "Imported texture",
		})
	}))
	defer ts.Close()

	h := newTextureTestHandler(ts)
	_, out, err := h.TextureOps(context.Background(), &mcp.CallToolRequest{}, TextureOpsInput{
		Operation:   "import",
		SourcePath:  "/tmp/terrain.png",
		Destination: "/Game/Textures/T_Terrain",
		Compression: "TC_Default",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Asset != "/Game/Textures/T_Terrain" {
		t.Errorf("expected asset path, got %s", out.Asset)
	}
}

func TestTextureOps_GetInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(TextureOpsOutput{
			Success: true,
			Info: &TextureInfo{
				Asset:       "/Game/Textures/T_Terrain",
				Name:        "T_Terrain",
				Width:       2048,
				Height:      2048,
				Format:      "PF_B8G8R8A8",
				MipCount:    11,
				Compression: "TC_Default",
				SizeKB:      16384,
			},
		})
	}))
	defer ts.Close()

	h := newTextureTestHandler(ts)
	_, out, err := h.TextureOps(context.Background(), &mcp.CallToolRequest{}, TextureOpsInput{
		Operation: "get_info",
		Asset:     "/Game/Textures/T_Terrain",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Info == nil {
		t.Fatal("expected info to be present")
	}
	if out.Info.Width != 2048 || out.Info.Height != 2048 {
		t.Errorf("unexpected dimensions: %dx%d", out.Info.Width, out.Info.Height)
	}
	if out.Info.MipCount != 11 {
		t.Errorf("expected 11 mips, got %d", out.Info.MipCount)
	}
}

func TestTextureOps_SetMaterialTexture(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["material_instance"] != "/Game/Materials/MI_Terrain" {
			t.Errorf("expected material_instance, got %v", body["material_instance"])
		}
		if body["param_name"] != "BaseColor" {
			t.Errorf("expected param_name BaseColor, got %v", body["param_name"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(TextureOpsOutput{
			Success: true,
			Message: "Texture parameter set",
		})
	}))
	defer ts.Close()

	h := newTextureTestHandler(ts)
	_, out, err := h.TextureOps(context.Background(), &mcp.CallToolRequest{}, TextureOpsInput{
		Operation:        "set_material_texture",
		MaterialInstance: "/Game/Materials/MI_Terrain",
		ParamName:        "BaseColor",
		Texture:          "/Game/Textures/T_Terrain",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
}

func TestTextureOps_List(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(TextureOpsOutput{
			Success: true,
			Count:   2,
			Textures: []TextureInfo{
				{Asset: "/Game/Textures/T_Terrain", Name: "T_Terrain", Width: 2048, Height: 2048},
				{Asset: "/Game/Textures/T_Normal", Name: "T_Normal", Width: 1024, Height: 1024},
			},
		})
	}))
	defer ts.Close()

	h := newTextureTestHandler(ts)
	_, out, err := h.TextureOps(context.Background(), &mcp.CallToolRequest{}, TextureOpsInput{
		Operation: "list",
		Path:      "/Game/Textures/",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 2 {
		t.Errorf("expected 2, got %d", out.Count)
	}
}

func TestTextureOps_MissingOperation(t *testing.T) {
	h := newTextureTestHandler(nil)
	_, _, err := h.TextureOps(context.Background(), &mcp.CallToolRequest{}, TextureOpsInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required', got: %v", err)
	}
}

func TestTextureOps_PluginOffline(t *testing.T) {
	h := newTextureTestHandler(nil)
	_, _, err := h.TextureOps(context.Background(), &mcp.CallToolRequest{}, TextureOpsInput{
		Operation: "list",
		Path:      "/Game/Textures/",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
