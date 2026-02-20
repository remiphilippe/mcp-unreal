package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGetLevelActors_Success(t *testing.T) {
	actors := []ActorInfo{
		{Name: "Cube", Class: "StaticMeshActor", Path: "/Game/Test:PersistentLevel.Cube",
			Location: [3]float64{100, 200, 300}, Rotation: [3]float64{0, 45, 0}, Scale: [3]float64{1, 1, 1}},
		{Name: "Light", Class: "PointLight", Path: "/Game/Test:PersistentLevel.Light",
			Location: [3]float64{0, 0, 500}, Scale: [3]float64{1, 1, 1}},
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(actors)
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GetLevelActors(context.Background(), &mcp.CallToolRequest{}, GetLevelActorsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Total != 2 {
		t.Errorf("expected 2 actors, got %d", out.Total)
	}
	if out.Actors[0].Name != "Cube" {
		t.Errorf("expected first actor to be Cube, got %s", out.Actors[0].Name)
	}
}

func TestGetLevelActors_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.GetLevelActors(context.Background(), &mcp.CallToolRequest{}, GetLevelActorsInput{})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestSpawnActor_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)

		// Verify default scale is [1,1,1].
		scale, ok := body["scale"].([]interface{})
		if !ok || len(scale) != 3 {
			t.Error("expected scale array of length 3")
		} else if scale[0].(float64) != 1 || scale[1].(float64) != 1 || scale[2].(float64) != 1 {
			t.Errorf("expected default scale [1,1,1], got %v", scale)
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(SpawnActorOutput{
			ActorPath: "/Game/Test:PersistentLevel.PointLight_0",
			ActorName: "PointLight_0",
			Class:     "PointLight",
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.SpawnActor(context.Background(), &mcp.CallToolRequest{}, SpawnActorInput{
		ClassName: "PointLight",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Class != "PointLight" {
		t.Errorf("expected PointLight class, got %s", out.Class)
	}
}

func TestSpawnActor_MissingClassName(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.SpawnActor(context.Background(), &mcp.CallToolRequest{}, SpawnActorInput{})
	if err == nil {
		t.Fatal("expected error for missing class_name")
	}
}

func TestDeleteActors_MissingInput(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.DeleteActors(context.Background(), &mcp.CallToolRequest{}, DeleteActorsInput{})
	if err == nil {
		t.Fatal("expected error when no paths or names provided")
	}
}

func TestMoveActor_Location(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT for RC API, got %s", r.Method)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)

		if body["functionName"] != "K2_SetActorLocation" {
			t.Errorf("expected K2_SetActorLocation, got %v", body["functionName"])
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ReturnValue":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	loc := [3]float64{100, 200, 300}
	_, out, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Cube",
		Location:   &loc,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if calls != 1 {
		t.Errorf("expected 1 RC API call, got %d", calls)
	}
}

func TestMoveActor_AllComponents(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	loc := [3]float64{100, 200, 300}
	rot := [3]float64{0, 90, 0}
	scale := [3]float64{2, 2, 2}
	_, _, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Cube",
		Location:   &loc,
		Rotation:   &rot,
		Scale:      &scale,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if calls != 3 {
		t.Errorf("expected 3 RC API calls (location + rotation + scale), got %d", calls)
	}
}

func TestMoveActor_MissingObjectPath(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{})
	if err == nil {
		t.Fatal("expected error for missing object_path")
	}
}

func TestMoveActor_NoTransformComponents(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Cube",
	})
	if err == nil {
		t.Fatal("expected error when no transform components provided")
	}
}

func TestDeleteActors_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/actors/delete" {
			t.Errorf("expected /api/actors/delete, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		paths, ok := body["actor_paths"].([]any)
		if !ok || len(paths) != 1 {
			t.Errorf("expected 1 actor path, got %v", body["actor_paths"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(DeleteActorsOutput{
			DeletedCount: 1,
			Deleted:      []string{"Cube"},
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.DeleteActors(context.Background(), &mcp.CallToolRequest{}, DeleteActorsInput{
		ActorPaths: []string{"/Game/Test:PersistentLevel.Cube"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.DeletedCount != 1 {
		t.Errorf("expected 1 deleted, got %d", out.DeletedCount)
	}
	if len(out.Deleted) != 1 || out.Deleted[0] != "Cube" {
		t.Errorf("unexpected deleted actors: %v", out.Deleted)
	}
}

func TestDeleteActors_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.DeleteActors(context.Background(), &mcp.CallToolRequest{}, DeleteActorsInput{
		ActorPaths: []string{"/Game/Test:PersistentLevel.Cube"},
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestMoveActor_Rotation(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)

		if body["functionName"] != "K2_SetActorRotation" {
			t.Errorf("expected K2_SetActorRotation, got %v", body["functionName"])
		}

		params, ok := body["parameters"].(map[string]any)
		if !ok {
			t.Fatal("expected parameters map")
		}
		newRot, ok := params["NewRotation"].(map[string]any)
		if !ok {
			t.Fatal("expected NewRotation map")
		}
		if newRot["Yaw"].(float64) != 90 {
			t.Errorf("expected Yaw=90, got %v", newRot["Yaw"])
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	rot := [3]float64{0, 90, 0}
	_, out, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Cube",
		Rotation:   &rot,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Rotation != [3]float64{0, 90, 0} {
		t.Errorf("expected rotation [0,90,0], got %v", out.Rotation)
	}
	if calls != 1 {
		t.Errorf("expected 1 RC API call, got %d", calls)
	}
}

func TestMoveActor_Scale(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)

		if body["functionName"] != "SetActorScale3D" {
			t.Errorf("expected SetActorScale3D, got %v", body["functionName"])
		}

		params, ok := body["parameters"].(map[string]any)
		if !ok {
			t.Fatal("expected parameters map")
		}
		newScale, ok := params["NewScale3D"].(map[string]any)
		if !ok {
			t.Fatal("expected NewScale3D map")
		}
		if newScale["X"].(float64) != 2 || newScale["Y"].(float64) != 2 || newScale["Z"].(float64) != 2 {
			t.Errorf("expected scale [2,2,2], got %v", newScale)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	scale := [3]float64{2, 2, 2}
	_, out, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Cube",
		Scale:      &scale,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Scale != [3]float64{2, 2, 2} {
		t.Errorf("expected scale [2,2,2], got %v", out.Scale)
	}
	if calls != 1 {
		t.Errorf("expected 1 RC API call, got %d", calls)
	}
}

func TestMoveActor_LocationAndRotation(t *testing.T) {
	functionNames := []string{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		functionNames = append(functionNames, body["functionName"].(string))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ReturnValue":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	loc := [3]float64{100, 200, 300}
	rot := [3]float64{0, 45, 0}
	_, _, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Cube",
		Location:   &loc,
		Rotation:   &rot,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(functionNames) != 2 {
		t.Fatalf("expected 2 RC API calls, got %d", len(functionNames))
	}
	if functionNames[0] != "K2_SetActorLocation" {
		t.Errorf("expected first call K2_SetActorLocation, got %s", functionNames[0])
	}
	if functionNames[1] != "K2_SetActorRotation" {
		t.Errorf("expected second call K2_SetActorRotation, got %s", functionNames[1])
	}
}

func TestMoveActor_RCAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"Object not found"}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	loc := [3]float64{100, 200, 300}
	_, _, err := h.MoveActor(context.Background(), &mcp.CallToolRequest{}, MoveActorInput{
		ObjectPath: "/Game/Test:PersistentLevel.Missing",
		Location:   &loc,
	})
	if err == nil {
		t.Fatal("expected error for RC API failure")
	}
}

func TestSpawnActor_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.SpawnActor(context.Background(), &mcp.CallToolRequest{}, SpawnActorInput{
		ClassName: "PointLight",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
