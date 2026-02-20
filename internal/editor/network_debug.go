package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- network_debug ---

// HTTPRequestInfo describes a recent HTTP request.
type HTTPRequestInfo struct {
	URL        string  `json:"url" jsonschema:"request URL"`
	Method     string  `json:"method" jsonschema:"HTTP method (GET, POST, etc.)"`
	StatusCode int     `json:"status_code,omitempty" jsonschema:"response status code"`
	DurationMs float64 `json:"duration_ms,omitempty" jsonschema:"request duration in milliseconds"`
	Error      string  `json:"error,omitempty" jsonschema:"error message if request failed"`
	Active     bool    `json:"active,omitempty" jsonschema:"whether the request is still in-flight"`
	Timestamp  string  `json:"timestamp,omitempty" jsonschema:"ISO 8601 timestamp"`
}

// WebSocketInfo describes a WebSocket connection.
type WebSocketInfo struct {
	URL       string `json:"url" jsonschema:"WebSocket endpoint URL"`
	State     string `json:"state" jsonschema:"connection state: connected, disconnected, connecting, error"`
	Protocol  string `json:"protocol,omitempty" jsonschema:"WebSocket protocol"`
	Error     string `json:"error,omitempty" jsonschema:"error message if disconnected due to error"`
	Timestamp string `json:"timestamp,omitempty" jsonschema:"last state change timestamp"`
}

// NetworkDebugInput defines parameters for the network_debug tool.
type NetworkDebugInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: list_active, recent_requests, websocket_status, summary"`
	LastN     int    `json:"last_n,omitempty" jsonschema:"Number of recent requests to return (default 20). For recent_requests."`
}

// NetworkDebugOutput is returned by the network_debug tool.
type NetworkDebugOutput struct {
	ActiveRequests      []HTTPRequestInfo `json:"active_requests,omitempty" jsonschema:"currently in-flight HTTP requests"`
	RecentRequests      []HTTPRequestInfo `json:"recent_requests,omitempty" jsonschema:"recent HTTP request history"`
	WebSockets          []WebSocketInfo   `json:"websockets,omitempty" jsonschema:"WebSocket connection states"`
	Count               int               `json:"count,omitempty" jsonschema:"number of results"`
	Note                string            `json:"note,omitempty" jsonschema:"additional information (e.g. for websocket_status)"`
	TotalTracked        int               `json:"total_tracked,omitempty" jsonschema:"total HTTP requests ever tracked (for summary)"`
	ActiveCount         int               `json:"active_count,omitempty" jsonschema:"currently active request count (for summary)"`
	AutoTrackingEnabled bool              `json:"auto_tracking_enabled,omitempty" jsonschema:"whether automatic HTTP tracking is enabled (for summary)"`
}

// RegisterNetworkDebug adds the network_debug tool to the MCP server.
func (h *Handler) RegisterNetworkDebug(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "network_debug",
		Description: "Introspect active HTTP requests and WebSocket connections in the editor. " +
			"Operations: list_active (in-flight HTTP requests and active WebSocket connections), " +
			"recent_requests (last N HTTP request/response log with URLs, status codes, durations), " +
			"websocket_status (WebSocket connection states), " +
			"summary (total tracked requests, active count, auto-tracking status). " +
			"Useful for debugging live data feeds, API connectivity, and network errors. " +
			"Requires the editor running with the MCPUnreal plugin loaded (port 8090).",
	}, h.NetworkDebug)
}

// NetworkDebug implements the network_debug tool.
func (h *Handler) NetworkDebug(ctx context.Context, req *mcp.CallToolRequest, input NetworkDebugInput) (*mcp.CallToolResult, NetworkDebugOutput, error) {
	if input.Operation == "" {
		return nil, NetworkDebugOutput{}, fmt.Errorf("operation is required")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.LastN > 0 {
		body["last_n"] = input.LastN
	}

	resp, err := h.Client.PluginCall(ctx, "/api/network/debug", body)
	if err != nil {
		return nil, NetworkDebugOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out NetworkDebugOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, NetworkDebugOutput{}, fmt.Errorf("parsing network debug response: %w", err)
	}

	return nil, out, nil
}
