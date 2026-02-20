// Package status implements the "status" MCP tool which reports server
// health, editor connectivity, project info, and feature availability.
//
// See IMPLEMENTATION.md §3.11 (Editor Utilities — status tool) and
// §10 (Graceful degradation principle).
package status

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/remiphilippe/mcp-unreal/internal/config"
)

// Handler holds references needed by the status tool.
type Handler struct {
	Config  *config.Config
	Version string
}

// Input defines parameters for the status tool.
type Input struct{}

// Output is returned by the status tool.
type Output struct {
	ServerVersion string   `json:"server_version" jsonschema:"mcp-unreal server version"`
	Platform      string   `json:"platform" jsonschema:"OS and architecture"`
	GoVersion     string   `json:"go_version" jsonschema:"Go runtime version"`
	ProjectRoot   string   `json:"project_root" jsonschema:"detected UE project root directory"`
	UProjectFile  string   `json:"uproject_file,omitempty" jsonschema:"path to .uproject file if found"`
	UEEditorPath  string   `json:"ue_editor_path" jsonschema:"configured path to UnrealEditor-Cmd"`
	UEInstalled   bool     `json:"ue_installed" jsonschema:"whether UnrealEditor-Cmd exists on disk"`
	EditorOnline  bool     `json:"editor_online" jsonschema:"whether the UE editor Remote Control API is reachable"`
	PluginOnline  bool     `json:"plugin_online" jsonschema:"whether the MCPUnreal editor plugin is reachable"`
	PIEActive     bool     `json:"pie_active" jsonschema:"whether Play In Editor is currently active"`
	PIEMap        string   `json:"pie_map,omitempty" jsonschema:"map name of the PIE world if active"`
	RCAPIPort     int      `json:"rc_api_port" jsonschema:"Remote Control API port"`
	PluginPort    int      `json:"plugin_port" jsonschema:"MCPUnreal plugin port"`
	Features      []string `json:"features" jsonschema:"list of available feature categories"`
}

// Register adds the status tool to the MCP server.
func (h *Handler) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "status",
		Description: "Check mcp-unreal server health, UE installation, and editor connectivity. " +
			"Call this first to verify your environment is set up correctly. " +
			"Returns project info, editor online status, and available features.",
	}, h.Status)
}

// Status implements the status tool handler.
func (h *Handler) Status(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	cfg := h.Config

	out := Output{
		ServerVersion: h.Version,
		Platform:      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		GoVersion:     runtime.Version(),
		ProjectRoot:   cfg.ProjectRoot,
		UProjectFile:  cfg.UProjectFile,
		UEEditorPath:  cfg.UEEditorPath,
		RCAPIPort:     cfg.RCAPIPort,
		PluginPort:    cfg.PluginPort,
	}

	// Check if UE editor binary exists on disk.
	if _, err := os.Stat(cfg.UEEditorPath); err == nil {
		out.UEInstalled = true
	}

	// Ping RC API (localhost only, per CLAUDE.md Security §3).
	out.EditorOnline = pingHTTP(ctx, cfg.RCAPIURL())

	// Ping plugin.
	ps := pingPlugin(ctx, cfg.PluginURL()+"/api/status")
	out.PluginOnline = ps.Online
	out.PIEActive = ps.PIEActive
	out.PIEMap = ps.PIEMap

	// Determine available features based on what's reachable.
	out.Features = h.availableFeatures(out)

	return nil, out, nil
}

// availableFeatures returns feature categories based on current state.
func (h *Handler) availableFeatures(out Output) []string {
	features := []string{}

	if out.UEInstalled {
		features = append(features, "headless_build", "headless_test")
	}

	// Doc search is always available (local bleve index).
	features = append(features, "doc_search")

	if out.EditorOnline {
		features = append(features, "rc_api_properties", "rc_api_functions", "console_commands")
	}

	if out.PluginOnline {
		features = append(features, "actors", "blueprints", "anim_blueprints", "assets",
			"materials", "characters", "input", "levels", "mesh", "output_log",
			"viewport_capture", "script_execution")
	}

	if out.PIEActive {
		features = append(features, "pie_world")
	}

	return features
}

// pluginStatus holds the result of pinging the MCPUnreal plugin endpoint.
type pluginStatus struct {
	Online    bool
	PIEActive bool
	PIEMap    string
}

// pingPlugin sends a GET to the plugin status endpoint and parses PIE fields.
func pingPlugin(ctx context.Context, url string) pluginStatus {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return pluginStatus{}
	}

	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return pluginStatus{}
	}
	defer func() { _ = resp.Body.Close() }()

	var body struct {
		PIEActive bool   `json:"pie_active"`
		PIEMap    string `json:"pie_map"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		// Reachable but couldn't parse — still online.
		return pluginStatus{Online: true}
	}

	return pluginStatus{
		Online:    true,
		PIEActive: body.PIEActive,
		PIEMap:    body.PIEMap,
	}
}

// pingHTTP sends a quick HTTP request to check if a service is reachable.
func pingHTTP(ctx context.Context, url string) bool {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false
	}

	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	_ = resp.Body.Close()
	return true
}
