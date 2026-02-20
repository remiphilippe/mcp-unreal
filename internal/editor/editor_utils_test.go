package editor

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestGetOutputLog_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/editor/output_log" {
			t.Errorf("expected /api/editor/output_log, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["category"] != "LogTemp" {
			t.Errorf("expected category LogTemp, got %v", body["category"])
		}
		if body["verbosity"] != "warning" {
			t.Errorf("expected verbosity warning, got %v", body["verbosity"])
		}
		if body["max_lines"] != float64(50) {
			t.Errorf("expected max_lines 50, got %v", body["max_lines"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(GetOutputLogOutput{
			Entries: []LogEntry{
				{Category: "LogTemp", Verbosity: "Warning", Message: "test warning"},
			},
			Count: 1,
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GetOutputLog(context.Background(), &mcp.CallToolRequest{}, GetOutputLogInput{
		Category:  "LogTemp",
		Verbosity: "warning",
		MaxLines:  50,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 1 {
		t.Errorf("expected 1 entry, got %d", out.Count)
	}
	if out.Entries[0].Message != "test warning" {
		t.Errorf("unexpected message: %s", out.Entries[0].Message)
	}
}

func TestGetOutputLog_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.GetOutputLog(context.Background(), &mcp.CallToolRequest{}, GetOutputLogInput{})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestCaptureViewport_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/editor/capture_viewport" {
			t.Errorf("expected /api/editor/capture_viewport, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["output_path"] != "/tmp/screenshot.png" {
			t.Errorf("expected output_path /tmp/screenshot.png, got %v", body["output_path"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(CaptureViewportOutput{
			Success:  true,
			FilePath: "/tmp/screenshot.png",
			Format:   "png",
			Width:    1920,
			Height:   1080,
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.CaptureViewport(context.Background(), &mcp.CallToolRequest{}, CaptureViewportInput{
		OutputPath: "/tmp/screenshot.png",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Width != 1920 || out.Height != 1080 {
		t.Errorf("unexpected dimensions: %dx%d", out.Width, out.Height)
	}
}

func TestCaptureViewport_Base64ReturnsImageContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(CaptureViewportOutput{
			Success:     true,
			ImageBase64: "iVBORw0KGgo=", // tiny valid base64
			Format:      "png",
			Width:       64,
			Height:      64,
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	result, out, err := h.CaptureViewport(context.Background(), &mcp.CallToolRequest{}, CaptureViewportInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	// ImageBase64 should be cleared from structured output (image is in Content).
	if out.ImageBase64 != "" {
		t.Error("expected ImageBase64 to be cleared from structured output")
	}
	// Result should contain MCP ImageContent.
	if result == nil {
		t.Fatal("expected non-nil CallToolResult with ImageContent")
	}
	if len(result.Content) != 1 {
		t.Fatalf("expected 1 content block, got %d", len(result.Content))
	}
	imgContent, ok := result.Content[0].(*mcp.ImageContent)
	if !ok {
		t.Fatalf("expected ImageContent, got %T", result.Content[0])
	}
	if imgContent.MIMEType != "image/png" {
		t.Errorf("expected image/png, got %s", imgContent.MIMEType)
	}
	if len(imgContent.Data) == 0 {
		t.Error("expected non-empty image data")
	}
}

func TestCaptureViewport_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.CaptureViewport(context.Background(), &mcp.CallToolRequest{}, CaptureViewportInput{})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestExecuteScript_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/editor/execute_script" {
			t.Errorf("expected /api/editor/execute_script, got %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["script"] != "print('hello')" {
			t.Errorf("unexpected script: %v", body["script"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ExecuteScriptOutput{
			Success: true,
			Output:  "hello",
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.ExecuteScript(context.Background(), &mcp.CallToolRequest{}, ExecuteScriptInput{
		Script: "print('hello')",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Output != "hello" {
		t.Errorf("expected output 'hello', got %s", out.Output)
	}
}

func TestExecuteScript_MissingScript(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.ExecuteScript(context.Background(), &mcp.CallToolRequest{}, ExecuteScriptInput{})
	if err == nil {
		t.Fatal("expected error for missing script")
	}
}

func TestExecuteScript_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.ExecuteScript(context.Background(), &mcp.CallToolRequest{}, ExecuteScriptInput{
		Script: "print('hello')",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestGetOutputLog_PatternFilter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["pattern"] != "terrain.*failed" {
			t.Errorf("expected pattern 'terrain.*failed', got %v", body["pattern"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(GetOutputLogOutput{
			Entries: []LogEntry{
				{Category: "LogTemp", Verbosity: "Error", Message: "terrain mesh failed to generate"},
			},
			Count: 1,
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GetOutputLog(context.Background(), &mcp.CallToolRequest{}, GetOutputLogInput{
		Pattern: "terrain.*failed",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 1 {
		t.Errorf("expected 1 entry, got %d", out.Count)
	}
}

func TestGetOutputLog_SinceSecondsFilter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["since_seconds"] != float64(30) {
			t.Errorf("expected since_seconds 30, got %v", body["since_seconds"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(GetOutputLogOutput{
			Entries: []LogEntry{
				{Category: "LogTemp", Verbosity: "Log", Message: "recent message"},
			},
			Count: 1,
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GetOutputLog(context.Background(), &mcp.CallToolRequest{}, GetOutputLogInput{
		SinceSeconds: 30,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 1 {
		t.Errorf("expected 1 entry, got %d", out.Count)
	}
}

func TestLiveCompile_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/editor/live_compile" {
			t.Errorf("expected /api/editor/live_compile, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(LiveCompileOutput{
			Success: true,
			Status:  "Compiling",
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.LiveCompile(context.Background(), &mcp.CallToolRequest{}, LiveCompileInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !out.Success {
		t.Error("expected success")
	}
	if out.Status != "Compiling" {
		t.Errorf("expected status Compiling, got %s", out.Status)
	}
}

func TestLiveCompile_Disabled(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(LiveCompileOutput{
			Success: false,
			Status:  "Disabled",
			Errors:  "Live Coding is disabled in Editor Preferences.",
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.LiveCompile(context.Background(), &mcp.CallToolRequest{}, LiveCompileInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Success {
		t.Error("expected failure")
	}
	if out.Errors == "" {
		t.Error("expected error details")
	}
}

func TestLiveCompile_PluginOffline(t *testing.T) {
	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", "http://127.0.0.1:1"),
		Logger: testLogger(),
	}

	_, _, err := h.LiveCompile(context.Background(), &mcp.CallToolRequest{}, LiveCompileInput{})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}

func TestGetOutputLog_CombinedFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["category"] != "LogBlueprintUserMessages" {
			t.Errorf("expected category LogBlueprintUserMessages, got %v", body["category"])
		}
		if body["pattern"] != "spawn" {
			t.Errorf("expected pattern 'spawn', got %v", body["pattern"])
		}
		if body["since_seconds"] != float64(60) {
			t.Errorf("expected since_seconds 60, got %v", body["since_seconds"])
		}
		if body["verbosity"] != "warning" {
			t.Errorf("expected verbosity warning, got %v", body["verbosity"])
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(GetOutputLogOutput{
			Entries: []LogEntry{},
			Count:   0,
		})
	}))
	defer server.Close()

	h := &Handler{
		Client: newTestClient("http://127.0.0.1:1", server.URL),
		Logger: testLogger(),
	}

	_, out, err := h.GetOutputLog(context.Background(), &mcp.CallToolRequest{}, GetOutputLogInput{
		Category:     "LogBlueprintUserMessages",
		Verbosity:    "warning",
		Pattern:      "spawn",
		SinceSeconds: 60,
		MaxLines:     50,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 0 {
		t.Errorf("expected 0 entries, got %d", out.Count)
	}
}
