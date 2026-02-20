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

func TestProceduralMesh_CreateSection(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/mesh/procedural" {
			t.Errorf("expected /api/mesh/procedural, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create_section" {
			t.Errorf("expected create_section, got %v", body["operation"])
		}
		// Verify vertices are present.
		verts, ok := body["vertices"].([]any)
		if !ok || len(verts) != 3 {
			t.Errorf("expected 3 vertices, got %v", body["vertices"])
		}
		// Verify triangles.
		tris, ok := body["triangles"].([]any)
		if !ok || len(tris) != 3 {
			t.Errorf("expected 3 triangle indices, got %v", body["triangles"])
		}
		// Verify normals.
		normals, ok := body["normals"].([]any)
		if !ok || len(normals) != 3 {
			t.Errorf("expected 3 normals, got %v", body["normals"])
		}
		// Verify UVs.
		uvs, ok := body["uvs"].([]any)
		if !ok || len(uvs) != 3 {
			t.Errorf("expected 3 UVs, got %v", body["uvs"])
		}
		// Verify colors.
		colors, ok := body["colors"].([]any)
		if !ok || len(colors) != 3 {
			t.Errorf("expected 3 colors, got %v", body["colors"])
		}
		// Verify location.
		loc, ok := body["location"].([]any)
		if !ok || len(loc) != 3 {
			t.Errorf("expected location array, got %v", body["location"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"actor_path":"/Game/Test:PersistentLevel.ProcMesh_0"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.ProceduralMesh(context.Background(), &mcp.CallToolRequest{}, ProceduralMeshInput{
		Operation: "create_section",
		ActorName: "TestMesh",
		Vertices:  [][3]float64{{0, 0, 0}, {100, 0, 0}, {50, 100, 0}},
		Triangles: []int{0, 1, 2},
		Normals:   [][3]float64{{0, 0, 1}, {0, 0, 1}, {0, 0, 1}},
		UVs:       [][2]float64{{0, 0}, {1, 0}, {0.5, 1}},
		Colors:    [][4]float64{{1, 0, 0, 1}, {0, 1, 0, 1}, {0, 0, 1, 1}},
		Location:  [3]float64{100, 200, 300},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestProceduralMesh_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.ProceduralMesh(context.Background(), &mcp.CallToolRequest{}, ProceduralMeshInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestProceduralMesh_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.ProceduralMesh(context.Background(), &mcp.CallToolRequest{}, ProceduralMeshInput{
		Operation: "create_section",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestRealtimeMesh_CreateLod(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/mesh/realtime" {
			t.Errorf("expected /api/mesh/realtime, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "create_section" {
			t.Errorf("expected create_section, got %v", body["operation"])
		}
		if body["lod_index"] != float64(1) {
			t.Errorf("expected lod_index 1, got %v", body["lod_index"])
		}
		if body["screen_size"] != float64(0.5) {
			t.Errorf("expected screen_size 0.5, got %v", body["screen_size"])
		}
		if body["section_group_key"] != "group0" {
			t.Errorf("expected section_group_key group0, got %v", body["section_group_key"])
		}
		if body["section_key"] != "section0" {
			t.Errorf("expected section_key section0, got %v", body["section_key"])
		}
		// Verify mesh data.
		verts, ok := body["vertices"].([]any)
		if !ok || len(verts) != 3 {
			t.Errorf("expected 3 vertices, got %v", body["vertices"])
		}
		tris, ok := body["triangles"].([]any)
		if !ok || len(tris) != 3 {
			t.Errorf("expected 3 triangle indices, got %v", body["triangles"])
		}
		normals, ok := body["normals"].([]any)
		if !ok || len(normals) != 3 {
			t.Errorf("expected 3 normals, got %v", body["normals"])
		}
		tangents, ok := body["tangents"].([]any)
		if !ok || len(tangents) != 3 {
			t.Errorf("expected 3 tangents, got %v", body["tangents"])
		}
		uvs, ok := body["uvs"].([]any)
		if !ok || len(uvs) != 3 {
			t.Errorf("expected 3 UVs, got %v", body["uvs"])
		}
		colors, ok := body["colors"].([]any)
		if !ok || len(colors) != 3 {
			t.Errorf("expected 3 colors, got %v", body["colors"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"actor_path":"/Game/Test:PersistentLevel.RMC_0"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.RealtimeMesh(context.Background(), &mcp.CallToolRequest{}, RealtimeMeshInput{
		Operation:       "create_section",
		LODIndex:        1,
		ScreenSize:      0.5,
		SectionGroupKey: "group0",
		SectionKey:      "section0",
		Vertices:        [][3]float64{{0, 0, 0}, {100, 0, 0}, {50, 100, 0}},
		Triangles:       []int{0, 1, 2},
		Normals:         [][3]float64{{0, 0, 1}, {0, 0, 1}, {0, 0, 1}},
		Tangents:        [][3]float64{{1, 0, 0}, {1, 0, 0}, {1, 0, 0}},
		UVs:             [][2]float64{{0, 0}, {1, 0}, {0.5, 1}},
		Colors:          [][4]float64{{1, 0, 0, 1}, {0, 1, 0, 1}, {0, 0, 1, 1}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Result == nil {
		t.Error("expected non-empty result")
	}
}

func TestRealtimeMesh_SetupCollision(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "setup_collision" {
			t.Errorf("expected setup_collision, got %v", body["operation"])
		}
		collVerts, ok := body["collision_vertices"].([]any)
		if !ok || len(collVerts) != 4 {
			t.Errorf("expected 4 collision vertices, got %v", body["collision_vertices"])
		}
		collTris, ok := body["collision_triangles"].([]any)
		if !ok || len(collTris) != 6 {
			t.Errorf("expected 6 collision triangle indices, got %v", body["collision_triangles"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, _, err := h.RealtimeMesh(context.Background(), &mcp.CallToolRequest{}, RealtimeMeshInput{
		Operation:          "setup_collision",
		ActorPath:          "/Game/Test:PersistentLevel.RMC_0",
		CollisionVertices:  [][3]float64{{0, 0, 0}, {100, 0, 0}, {100, 100, 0}, {0, 100, 0}},
		CollisionTriangles: []int{0, 1, 2, 0, 2, 3},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRealtimeMesh_MissingOperation(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.RealtimeMesh(context.Background(), &mcp.CallToolRequest{}, RealtimeMeshInput{})
	if err == nil {
		t.Fatal("expected error for missing operation")
	}
}

func TestRealtimeMesh_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.RealtimeMesh(context.Background(), &mcp.CallToolRequest{}, RealtimeMeshInput{
		Operation: "create_lod",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
