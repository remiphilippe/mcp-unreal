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

func newUIQueryTestHandler(ts *httptest.Server) *Handler {
	pluginURL := "http://127.0.0.1:1"
	if ts != nil {
		pluginURL = ts.URL
	}
	return &Handler{
		Client: newTestClient("http://127.0.0.1:1", pluginURL),
		Logger: testLogger(),
	}
}

func TestUIQuery_Tree(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/ui/query" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["operation"] != "tree" {
			t.Errorf("expected tree, got %v", body["operation"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(UIQueryOutput{
			Count: 3,
			Widgets: []WidgetInfo{
				{Type: "SWindow", Name: "MainWindow", Visible: true, Enabled: true, Children: []any{
					WidgetInfo{Type: "SPanel", Name: "Root", Visible: true, Enabled: true},
					WidgetInfo{Type: "STextBlock", Name: "Title", Visible: true, Enabled: true},
				}},
			},
		})
	}))
	defer ts.Close()

	h := newUIQueryTestHandler(ts)
	_, out, err := h.UIQuery(context.Background(), &mcp.CallToolRequest{}, UIQueryInput{
		Operation: "tree",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 3 {
		t.Errorf("expected 3, got %d", out.Count)
	}
	if len(out.Widgets) != 1 {
		t.Fatalf("expected 1 root widget, got %d", len(out.Widgets))
	}
	if out.Widgets[0].Type != "SWindow" {
		t.Errorf("expected SWindow, got %s", out.Widgets[0].Type)
	}
}

func TestUIQuery_Find(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body["class"] != "SCheckBox" {
			t.Errorf("expected SCheckBox, got %v", body["class"])
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(UIQueryOutput{
			Count: 2,
			Widgets: []WidgetInfo{
				{Type: "SCheckBox", Name: "EncCheckbox", Visible: true, Enabled: true},
				{Type: "SCheckBox", Name: "AisCheckbox", Visible: true, Enabled: true},
			},
		})
	}))
	defer ts.Close()

	h := newUIQueryTestHandler(ts)
	_, out, err := h.UIQuery(context.Background(), &mcp.CallToolRequest{}, UIQueryInput{
		Operation: "find",
		Class:     "SCheckBox",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Count != 2 {
		t.Errorf("expected 2, got %d", out.Count)
	}
}

func TestUIQuery_MissingOperation(t *testing.T) {
	h := newUIQueryTestHandler(nil)
	_, _, err := h.UIQuery(context.Background(), &mcp.CallToolRequest{}, UIQueryInput{})
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "operation is required") {
		t.Errorf("expected 'operation is required', got: %v", err)
	}
}

func TestUIQuery_PluginOffline(t *testing.T) {
	h := newUIQueryTestHandler(nil)
	_, _, err := h.UIQuery(context.Background(), &mcp.CallToolRequest{}, UIQueryInput{
		Operation: "tree",
	})
	if err == nil {
		t.Fatal("expected error for offline plugin")
	}
}
