// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package status

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/remiphilippe/mcp-unreal/internal/config"
)

func TestStatusWithEditorOffline(t *testing.T) {
	cfg := &config.Config{
		UEEditorPath:  "/nonexistent/UnrealEditor-Cmd",
		ProjectRoot:   "/tmp",
		RCAPIPort:     39999, // nothing listening
		PluginPort:    39998, // nothing listening
		DocsIndexPath: "./test-index.bleve",
	}

	h := &Handler{Config: cfg, Version: "0.1.0-test"}
	_, out, err := h.Status(context.Background(), nil, Input{})
	if err != nil {
		t.Fatalf("Status returned error: %v", err)
	}

	if out.ServerVersion != "0.1.0-test" {
		t.Errorf("ServerVersion = %q, want 0.1.0-test", out.ServerVersion)
	}
	if out.UEInstalled {
		t.Error("UEInstalled should be false for nonexistent path")
	}
	if out.EditorOnline {
		t.Error("EditorOnline should be false when nothing is listening")
	}
	if out.PluginOnline {
		t.Error("PluginOnline should be false when nothing is listening")
	}

	// Doc search should always be available.
	if !containsFeature(out.Features, "doc_search") {
		t.Error("doc_search feature should always be present")
	}

	// Headless features should NOT be present when UE not installed.
	if containsFeature(out.Features, "headless_build") {
		t.Error("headless_build should not be present when UE not installed")
	}
}

func TestStatusWithEditorOnline(t *testing.T) {
	// Start mock RC API server.
	rcAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer rcAPI.Close()

	// Start mock plugin server.
	plugin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	}))
	defer plugin.Close()

	rcPort := portFromURL(t, rcAPI.URL)
	pluginPort := portFromURL(t, plugin.URL)

	cfg := &config.Config{
		UEEditorPath:  "/nonexistent/UnrealEditor-Cmd",
		ProjectRoot:   "/tmp",
		RCAPIPort:     rcPort,
		PluginPort:    pluginPort,
		DocsIndexPath: "./test-index.bleve",
	}

	h := &Handler{Config: cfg, Version: "0.1.0-test"}
	_, out, err := h.Status(context.Background(), nil, Input{})
	if err != nil {
		t.Fatalf("Status returned error: %v", err)
	}

	if !out.EditorOnline {
		t.Error("EditorOnline should be true when RC API is reachable")
	}
	if !out.PluginOnline {
		t.Error("PluginOnline should be true when plugin is reachable")
	}

	// Editor features should be present.
	if !containsFeature(out.Features, "rc_api_properties") {
		t.Error("rc_api_properties feature should be present when editor online")
	}
	if !containsFeature(out.Features, "blueprints") {
		t.Error("blueprints feature should be present when plugin online")
	}
}

func TestAvailableFeaturesDegradation(t *testing.T) {
	h := &Handler{Config: &config.Config{}, Version: "test"}

	// All offline: only doc_search.
	features := h.availableFeatures(Output{
		UEInstalled:  false,
		EditorOnline: false,
		PluginOnline: false,
	})
	if len(features) != 1 || features[0] != "doc_search" {
		t.Errorf("all offline: features = %v, want [doc_search]", features)
	}

	// UE installed but editor offline: headless + doc_search.
	features = h.availableFeatures(Output{
		UEInstalled:  true,
		EditorOnline: false,
		PluginOnline: false,
	})
	if !containsFeature(features, "headless_build") {
		t.Error("headless_build should be present when UE installed")
	}
	if !containsFeature(features, "headless_test") {
		t.Error("headless_test should be present when UE installed")
	}
	if containsFeature(features, "blueprints") {
		t.Error("blueprints should not be present when plugin offline")
	}
}

func containsFeature(features []string, name string) bool {
	for _, f := range features {
		if f == name {
			return true
		}
	}
	return false
}

func portFromURL(t *testing.T, url string) int {
	t.Helper()
	parts := strings.Split(url, ":")
	port, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		t.Fatalf("failed to parse port from URL %q: %v", url, err)
	}
	return port
}
