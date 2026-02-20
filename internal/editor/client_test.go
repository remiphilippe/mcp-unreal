package editor

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

func newTestClient(rcAPIURL, pluginURL string) *Client {
	return &Client{
		rcAPIBaseURL:  rcAPIURL,
		pluginBaseURL: pluginURL,
		httpClient:    &http.Client{},
		logger:        testLogger(),
	}
}

func TestClient_RCAPICall_PUT(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT method, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected application/json content type, got %s", r.Header.Get("Content-Type"))
		}
		if r.URL.Path != "/remote/object/property" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"RelativeLocation": map[string]float64{"X": 100, "Y": 200, "Z": 300},
		})
	}))
	defer server.Close()

	client := newTestClient(server.URL, "http://127.0.0.1:1")

	resp, err := client.RCAPICall(context.Background(), "/remote/object/property", map[string]string{"test": "value"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response, got nil")
	}
}

func TestClient_PluginCall_POST(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		if r.URL.Path != "/api/actors/list" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode([]ActorInfo{
			{Name: "TestActor", Class: "StaticMeshActor", Path: "/Game/Test.TestActor"},
		})
	}))
	defer server.Close()

	client := newTestClient("http://127.0.0.1:1", server.URL)

	resp, err := client.PluginCall(context.Background(), "/api/actors/list", map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response, got nil")
	}
}

func TestClient_RCAPICall_NilBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "" {
			t.Errorf("expected no content-type for nil body, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL, "http://127.0.0.1:1")

	resp, err := client.RCAPICall(context.Background(), "/remote/info", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(resp) != "{}" {
		t.Errorf("expected {}, got %s", string(resp))
	}
}

func TestClient_RCAPICall_Offline(t *testing.T) {
	client := newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1")

	_, err := client.RCAPICall(context.Background(), "/remote/object/property", nil)
	if err == nil {
		t.Fatal("expected error for offline RC API")
	}
}

func TestClient_PluginCall_Offline(t *testing.T) {
	client := newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1")

	_, err := client.PluginCall(context.Background(), "/api/status", nil)
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestClient_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid object path"}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL, "http://127.0.0.1:1")

	_, err := client.RCAPICall(context.Background(), "/remote/object/property", map[string]string{"objectPath": "bad"})
	if err == nil {
		t.Fatal("expected error for HTTP 400")
	}
}

func TestClient_PingRCAPI_Online(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL, "http://127.0.0.1:1")

	if !client.PingRCAPI(context.Background()) {
		t.Error("expected PingRCAPI to return true for online server")
	}
}

func TestClient_PingRCAPI_Offline(t *testing.T) {
	client := newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1")

	if client.PingRCAPI(context.Background()) {
		t.Error("expected PingRCAPI to return false for offline server")
	}
}

func TestClient_PingPlugin_Online(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"name":"MCPUnreal"}`))
	}))
	defer server.Close()

	client := newTestClient("http://127.0.0.1:1", server.URL)

	if !client.PingPlugin(context.Background()) {
		t.Error("expected PingPlugin to return true for online server")
	}
}

func TestClient_PingPlugin_Offline(t *testing.T) {
	client := newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1")

	if client.PingPlugin(context.Background()) {
		t.Error("expected PingPlugin to return false for offline server")
	}
}

func TestClient_RCAPIURL(t *testing.T) {
	client := newTestClient("http://127.0.0.1:30010", "http://127.0.0.1:8090")

	if client.RCAPIURL() != "http://127.0.0.1:30010" {
		t.Errorf("unexpected RCAPIURL: %s", client.RCAPIURL())
	}
}

func TestClient_PluginURL(t *testing.T) {
	client := newTestClient("http://127.0.0.1:30010", "http://127.0.0.1:8090")

	if client.PluginURL() != "http://127.0.0.1:8090" {
		t.Errorf("unexpected PluginURL: %s", client.PluginURL())
	}
}

func TestClient_PluginCall_ErrorPayload(t *testing.T) {
	// Plugin returns HTTP 200 with {"error":"..."} body — client must detect this.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"error":"Actor not found: MissingActor"}`))
	}))
	defer server.Close()

	client := newTestClient("http://127.0.0.1:1", server.URL)

	_, err := client.PluginCall(context.Background(), "/api/actors/delete", map[string]any{
		"actor_names": []string{"MissingActor"},
	})
	if err == nil {
		t.Fatal("expected error for HTTP 200 with error payload")
	}
	if !strings.Contains(err.Error(), "Actor not found") {
		t.Errorf("expected error message to contain 'Actor not found', got: %s", err.Error())
	}
}

func TestClient_RCAPICall_ErrorPayload(t *testing.T) {
	// RC API can also return 200 with error payload in some cases.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"error":"Object not found"}`))
	}))
	defer server.Close()

	client := newTestClient(server.URL, "http://127.0.0.1:1")

	_, err := client.RCAPICall(context.Background(), "/remote/object/call", map[string]any{
		"objectPath": "/Game/Missing",
	})
	if err == nil {
		t.Fatal("expected error for RC API 200 with error payload")
	}
	if !strings.Contains(err.Error(), "Object not found") {
		t.Errorf("expected 'Object not found' in error, got: %s", err.Error())
	}
}

func TestClient_PluginCall_ValidErrorFreePayload(t *testing.T) {
	// Normal response with no "error" field should succeed.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success":true,"actor_path":"/Game/Test.Actor"}`))
	}))
	defer server.Close()

	client := newTestClient("http://127.0.0.1:1", server.URL)

	resp, err := client.PluginCall(context.Background(), "/api/actors/spawn", map[string]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response, got nil")
	}
}

func TestClient_PluginCall_EmptyErrorField(t *testing.T) {
	// Response with {"error":""} should NOT be treated as an error.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"error":"","success":true}`))
	}))
	defer server.Close()

	client := newTestClient("http://127.0.0.1:1", server.URL)

	_, err := client.PluginCall(context.Background(), "/api/actors/list", nil)
	if err != nil {
		t.Fatalf("empty error field should not fail: %v", err)
	}
}

func TestClient_PluginCall_ArrayResponse(t *testing.T) {
	// Array response like actor list should not trigger error detection.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"name":"Actor1"},{"name":"Actor2"}]`))
	}))
	defer server.Close()

	client := newTestClient("http://127.0.0.1:1", server.URL)

	resp, err := client.PluginCall(context.Background(), "/api/actors/list", nil)
	if err != nil {
		t.Fatalf("array response should not fail: %v", err)
	}
	if resp == nil {
		t.Fatal("expected response, got nil")
	}
}
