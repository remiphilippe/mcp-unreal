package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestSetProperty_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/remote/object/property" {
			t.Errorf("expected /remote/object/property, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["objectPath"] != "/Game/Maps/Main.Main:PersistentLevel.Cube" {
			t.Errorf("unexpected objectPath: %v", body["objectPath"])
		}
		if body["propertyName"] != "RelativeLocation" {
			t.Errorf("unexpected propertyName: %v", body["propertyName"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, out, err := h.SetProperty(context.Background(), &mcp.CallToolRequest{}, SetPropertyInput{
		ObjectPath:    "/Game/Maps/Main.Main:PersistentLevel.Cube",
		PropertyName:  "RelativeLocation",
		PropertyValue: json.RawMessage(`{"X":100,"Y":200,"Z":300}`),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Message == "" {
		t.Error("expected non-empty message")
	}
}

func TestSetProperty_MissingObjectPath(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.SetProperty(context.Background(), &mcp.CallToolRequest{}, SetPropertyInput{
		PropertyName:  "RelativeLocation",
		PropertyValue: json.RawMessage(`100`),
	})
	if err == nil {
		t.Fatal("expected error for missing object_path")
	}
}

func TestSetProperty_MissingPropertyName(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.SetProperty(context.Background(), &mcp.CallToolRequest{}, SetPropertyInput{
		ObjectPath:    "/Game/Maps/Main.Main:PersistentLevel.Cube",
		PropertyValue: json.RawMessage(`100`),
	})
	if err == nil {
		t.Fatal("expected error for missing property_name")
	}
}

func TestSetProperty_RCAPIOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.SetProperty(context.Background(), &mcp.CallToolRequest{}, SetPropertyInput{
		ObjectPath:    "/Game/Maps/Main.Main:PersistentLevel.Cube",
		PropertyName:  "Intensity",
		PropertyValue: json.RawMessage(`5000`),
	})
	if err == nil {
		t.Fatal("expected error for offline RC API")
	}
}

func TestGetProperty_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		if r.URL.Path != "/remote/object/property" {
			t.Errorf("expected /remote/object/property, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["access"] != "READ_ACCESS" {
			t.Errorf("expected READ_ACCESS, got %v", body["access"])
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"X":100,"Y":200,"Z":300}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, out, err := h.GetProperty(context.Background(), &mcp.CallToolRequest{}, GetPropertyInput{
		ObjectPath:   "/Game/Maps/Main.Main:PersistentLevel.Cube",
		PropertyName: "RelativeLocation",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ObjectPath != "/Game/Maps/Main.Main:PersistentLevel.Cube" {
		t.Errorf("unexpected object_path: %s", out.ObjectPath)
	}
	if out.PropertyName != "RelativeLocation" {
		t.Errorf("unexpected property_name: %s", out.PropertyName)
	}
	if string(out.PropertyValue) != `{"X":100,"Y":200,"Z":300}` {
		t.Errorf("unexpected property_value: %s", string(out.PropertyValue))
	}
}

func TestGetProperty_MissingObjectPath(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.GetProperty(context.Background(), &mcp.CallToolRequest{}, GetPropertyInput{
		PropertyName: "RelativeLocation",
	})
	if err == nil {
		t.Fatal("expected error for missing object_path")
	}
}

func TestGetProperty_MissingPropertyName(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.GetProperty(context.Background(), &mcp.CallToolRequest{}, GetPropertyInput{
		ObjectPath: "/Game/Maps/Main.Main:PersistentLevel.Cube",
	})
	if err == nil {
		t.Fatal("expected error for missing property_name")
	}
}

func TestCallFunction_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/remote/object/call" {
			t.Errorf("expected /remote/object/call, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["objectPath"] != "/Game/Maps/Main.Main:PersistentLevel.Cube" {
			t.Errorf("unexpected objectPath: %v", body["objectPath"])
		}
		if body["functionName"] != "K2_SetActorLocation" {
			t.Errorf("unexpected functionName: %v", body["functionName"])
		}
		// Verify no parameters key when nil.
		if _, ok := body["parameters"]; ok {
			t.Error("expected no parameters key when nil")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ReturnValue":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, out, err := h.CallFunction(context.Background(), &mcp.CallToolRequest{}, CallFunctionInput{
		ObjectPath:   "/Game/Maps/Main.Main:PersistentLevel.Cube",
		FunctionName: "K2_SetActorLocation",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(out.ReturnValue) != `{"ReturnValue":true}` {
		t.Errorf("unexpected return value: %s", string(out.ReturnValue))
	}
}

func TestCallFunction_WithParameters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		params, ok := body["parameters"].(map[string]any)
		if !ok {
			t.Fatal("expected parameters map in body")
		}
		if params["NewLocation"] == nil {
			t.Error("expected NewLocation parameter")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ReturnValue":true}`))
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient(server.URL, "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CallFunction(context.Background(), &mcp.CallToolRequest{}, CallFunctionInput{
		ObjectPath:   "/Game/Maps/Main.Main:PersistentLevel.Cube",
		FunctionName: "K2_SetActorLocation",
		Parameters: map[string]any{
			"NewLocation": map[string]float64{"X": 100, "Y": 200, "Z": 300},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCallFunction_MissingObjectPath(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CallFunction(context.Background(), &mcp.CallToolRequest{}, CallFunctionInput{
		FunctionName: "K2_SetActorLocation",
	})
	if err == nil {
		t.Fatal("expected error for missing object_path")
	}
}

func TestCallFunction_MissingFunctionName(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CallFunction(context.Background(), &mcp.CallToolRequest{}, CallFunctionInput{
		ObjectPath: "/Game/Maps/Main.Main:PersistentLevel.Cube",
	})
	if err == nil {
		t.Fatal("expected error for missing function_name")
	}
}
