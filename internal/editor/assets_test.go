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

func TestGetAssetInfo_Success(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.WriteHeader(http.StatusOK)
		switch r.URL.Path {
		case "/api/assets/info":
			_, _ = w.Write([]byte(`{"name":"BP_Player","path":"/Game/BP_Player","class":"Blueprint","package":"/Game","package_flags":0,"tags":{"NativeParentClass":"Actor"}}`))
		case "/api/assets/dependencies":
			_, _ = w.Write([]byte(`{"dependencies":["/Engine/BasicShapes"]}`))
		case "/api/assets/referencers":
			_, _ = w.Write([]byte(`{"referencers":["/Game/Maps/MainLevel"]}`))
		default:
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.GetAssetInfo(context.Background(), &mcp.CallToolRequest{}, GetAssetInfoInput{
		AssetPath: "/Game/BP_Player.BP_Player",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Name != "BP_Player" {
		t.Errorf("expected BP_Player, got %s", out.Name)
	}
	if len(out.Dependencies) != 1 {
		t.Errorf("expected 1 dependency, got %d", len(out.Dependencies))
	}
	if len(out.Referencers) != 1 {
		t.Errorf("expected 1 referencer, got %d", len(out.Referencers))
	}
	if callCount != 3 {
		t.Errorf("expected 3 plugin calls (info + deps + refs), got %d", callCount)
	}
}

func TestGetAssetInfo_MissingPath(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.GetAssetInfo(context.Background(), &mcp.CallToolRequest{}, GetAssetInfoInput{})
	if err == nil {
		t.Fatal("expected error for missing asset_path")
	}
}

func TestSearchAssets_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["name_filter"] != "BP_Test" {
			t.Errorf("expected name_filter BP_Test, got %v", body["name_filter"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"name":"BP_Test","path":"/Game/BP_Test","class":"Blueprint","package":"/Game"}]`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, out, err := h.SearchAssets(context.Background(), &mcp.CallToolRequest{}, SearchAssetsInput{
		NameFilter: "BP_Test",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Total != 1 {
		t.Errorf("expected 1 asset, got %d", out.Total)
	}
}

func TestSearchAssets_WithFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["class_filter"] != "StaticMesh" {
			t.Errorf("expected class_filter StaticMesh, got %v", body["class_filter"])
		}
		if body["path_filter"] != "/Game/Meshes" {
			t.Errorf("expected path_filter /Game/Meshes, got %v", body["path_filter"])
		}
		if body["recursive_path"] != true {
			t.Errorf("expected recursive_path true, got %v", body["recursive_path"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[]`))
	}))
	defer server.Close()

	h := &Handler{Client: newTestClient("http://127.0.0.1:1", server.URL), Logger: testLogger()}
	_, _, err := h.SearchAssets(context.Background(), &mcp.CallToolRequest{}, SearchAssetsInput{
		ClassFilter:   "StaticMesh",
		PathFilter:    "/Game/Meshes",
		RecursivePath: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSearchAssets_PluginOffline(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.SearchAssets(context.Background(), &mcp.CallToolRequest{}, SearchAssetsInput{})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestGetAssetInfo_PluginOffline(t *testing.T) {
	h := &Handler{Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"), Logger: testLogger()}
	_, _, err := h.GetAssetInfo(context.Background(), &mcp.CallToolRequest{}, GetAssetInfoInput{
		AssetPath: "/Game/BP_Test.BP_Test",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
