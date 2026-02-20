package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newComponentTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestGetActorComponents_Success(t *testing.T) {
	response := GetActorComponentsOutput{
		Actor: "BaseMapActor_0",
		Class: "PelorusBaseMapActor",
		Path:  "PersistentLevel.BaseMapActor_0",
		Components: []ComponentInfo{
			{
				Name:    "DefaultSceneRoot",
				Class:   "USceneComponent",
				Visible: true,
				Children: []any{
					ComponentInfo{
						Name:    "TerrainMesh",
						Class:   "URealtimeMeshComponent",
						Visible: true,
					},
					ComponentInfo{
						Name:          "PointInstances",
						Class:         "UInstancedStaticMeshComponent",
						Visible:       true,
						InstanceCount: intPtr(4200),
					},
				},
			},
		},
		NonSceneComponents: []NonSceneComponentInfo{
			{
				Name:     "AisClient",
				Class:    "UActorComponent",
				IsActive: true,
			},
		},
		TotalComponents: 4,
	}

	respBytes, _ := json.Marshal(response)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/actors/components" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(respBytes)
	}))
	defer ts.Close()

	h := newComponentTestHandler(ts)
	ctx := context.Background()

	_, out, err := h.GetActorComponents(ctx, nil, GetActorComponentsInput{
		ActorName:         "BaseMapActor_0",
		IncludeTransforms: false,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Actor != "BaseMapActor_0" {
		t.Errorf("expected BaseMapActor_0, got %s", out.Actor)
	}
	if out.TotalComponents != 4 {
		t.Errorf("expected 4 total components, got %d", out.TotalComponents)
	}
	if len(out.Components) != 1 {
		t.Fatalf("expected 1 root component, got %d", len(out.Components))
	}
	root := out.Components[0]
	if len(root.Children) != 2 {
		t.Errorf("expected 2 children, got %d", len(root.Children))
	}
	// After JSON round-trip, children are map[string]any.
	ismMap, ok := root.Children[1].(map[string]any)
	if !ok {
		t.Fatalf("expected child to be map[string]any, got %T", root.Children[1])
	}
	if count, ok := ismMap["instance_count"].(float64); !ok || int(count) != 4200 {
		t.Errorf("expected ISM instance_count 4200, got %v", ismMap["instance_count"])
	}
	// Check non-scene components.
	if len(out.NonSceneComponents) != 1 {
		t.Fatalf("expected 1 non-scene component, got %d", len(out.NonSceneComponents))
	}
	if out.NonSceneComponents[0].Name != "AisClient" {
		t.Errorf("expected AisClient, got %s", out.NonSceneComponents[0].Name)
	}
}

func TestGetActorComponents_WithTransforms(t *testing.T) {
	response := GetActorComponentsOutput{
		Actor: "PointLight_0",
		Class: "APointLight",
		Path:  "PersistentLevel.PointLight_0",
		Components: []ComponentInfo{
			{
				Name:    "LightComponent0",
				Class:   "UPointLightComponent",
				Visible: true,
				Transform: &ComponentTransform{
					Location: [3]float64{100, 200, 300},
					Rotation: [3]float64{0, 45, 0},
					Scale:    [3]float64{1, 1, 1},
				},
			},
		},
		TotalComponents: 1,
	}

	respBytes, _ := json.Marshal(response)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(respBytes)
	}))
	defer ts.Close()

	h := newComponentTestHandler(ts)
	ctx := context.Background()

	_, out, err := h.GetActorComponents(ctx, nil, GetActorComponentsInput{
		ActorPath:         "PersistentLevel.PointLight_0",
		IncludeTransforms: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(out.Components))
	}
	comp := out.Components[0]
	if comp.Transform == nil {
		t.Fatal("expected transform to be present")
	}
	if comp.Transform.Location[0] != 100 {
		t.Errorf("expected X=100, got %f", comp.Transform.Location[0])
	}
}

func TestGetActorComponents_MissingInput(t *testing.T) {
	h := newComponentTestHandler(nil)
	ctx := context.Background()

	_, _, err := h.GetActorComponents(ctx, nil, GetActorComponentsInput{})
	if err == nil {
		t.Fatal("expected error for missing actor_path and actor_name")
	}
	if !strings.Contains(err.Error(), "required") {
		t.Errorf("expected 'required' in error, got: %v", err)
	}
}

func TestGetActorComponents_ActorNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"error":"Actor not found: NonExistent"}`))
	}))
	defer ts.Close()

	h := newComponentTestHandler(ts)
	ctx := context.Background()

	_, _, err := h.GetActorComponents(ctx, nil, GetActorComponentsInput{
		ActorName: "NonExistent",
	})
	if err == nil {
		t.Fatal("expected error for missing actor")
	}
}

func intPtr(n int) *int { return &n }
