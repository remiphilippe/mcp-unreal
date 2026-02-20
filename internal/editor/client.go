// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

// Package editor implements MCP tools that communicate with the running
// UE 5.7 editor via the Remote Control API (port 30010) and the MCPUnreal
// editor plugin (port 8090).
//
// The Remote Control API is a built-in UE plugin that exposes UObject
// properties and UFUNCTIONs over HTTP. The MCPUnreal plugin extends this
// with actor management, Blueprint editing, and other editor internals.
//
// See IMPLEMENTATION.md §2 for communication paths and §3.3–§3.11 for
// the tool inventory.
package editor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/remiphilippe/mcp-unreal/internal/config"
)

const (
	// defaultRequestTimeout is the maximum time for a single HTTP request
	// to the editor. Most operations complete in under a second; builds
	// and tests use their own timeouts via headless tools.
	defaultRequestTimeout = 30 * time.Second

	// defaultConnectTimeout is used for ping/health checks.
	defaultConnectTimeout = 5 * time.Second
)

// Client provides HTTP communication with the UE editor via both
// the Remote Control API and the MCPUnreal editor plugin.
type Client struct {
	rcAPIBaseURL  string
	pluginBaseURL string
	httpClient    *http.Client
	logger        *slog.Logger
}

// Handler holds references needed by all editor tools.
type Handler struct {
	Client *Client
	Logger *slog.Logger
}

// NewClient creates an editor client from the server configuration.
func NewClient(cfg *config.Config, logger *slog.Logger) *Client {
	return &Client{
		rcAPIBaseURL:  cfg.RCAPIURL(),
		pluginBaseURL: cfg.PluginURL(),
		httpClient: &http.Client{
			Timeout: defaultRequestTimeout,
		},
		logger: logger,
	}
}

// RCAPIURL returns the base URL for the Remote Control API.
func (c *Client) RCAPIURL() string { return c.rcAPIBaseURL }

// PluginURL returns the base URL for the MCPUnreal plugin.
func (c *Client) PluginURL() string { return c.pluginBaseURL }

// RCAPICall sends an HTTP PUT request to the Remote Control API and
// returns the response body as raw JSON. The RC API uses PUT for all
// mutating operations (set property, call function, search assets).
func (c *Client) RCAPICall(ctx context.Context, endpoint string, body any) (json.RawMessage, error) {
	return c.doRequest(ctx, http.MethodPut, c.rcAPIBaseURL+endpoint, body, "RC API")
}

// PluginCall sends an HTTP POST request to the MCPUnreal editor plugin
// and returns the response body as raw JSON.
func (c *Client) PluginCall(ctx context.Context, endpoint string, body any) (json.RawMessage, error) {
	return c.doRequest(ctx, http.MethodPost, c.pluginBaseURL+endpoint, body, "MCPUnreal plugin")
}

// PingRCAPI checks whether the Remote Control API is reachable.
func (c *Client) PingRCAPI(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, defaultConnectTimeout)
	defer cancel()
	_, err := c.doRequest(ctx, http.MethodGet, c.rcAPIBaseURL+"/remote/info", nil, "")
	return err == nil
}

// PingPlugin checks whether the MCPUnreal editor plugin is reachable.
func (c *Client) PingPlugin(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, defaultConnectTimeout)
	defer cancel()
	_, err := c.doRequest(ctx, http.MethodPost, c.pluginBaseURL+"/api/status", nil, "")
	return err == nil
}

// doRequest performs an HTTP request and returns the response body.
// It handles JSON marshaling, content-type headers, and produces
// user-friendly errors when the editor is offline.
func (c *Client) doRequest(ctx context.Context, method, url string, body any, serviceName string) (json.RawMessage, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshaling request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	c.logger.Debug("editor HTTP request", "method", method, "url", url)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if serviceName != "" {
			return nil, fmt.Errorf(
				"%s unreachable at %s — ensure UE is running with %s enabled: %w",
				serviceName, url, serviceName, err,
			)
		}
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response from %s: %w", serviceName, err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("%s returned HTTP %d: %s", serviceName, resp.StatusCode, truncate(string(respBody), 500))
	}

	// The UE plugin may return HTTP 200 with an error payload.
	// Detect {"error":"..."} responses and surface the message.
	var errCheck struct {
		Error string `json:"error"`
	}
	if json.Unmarshal(respBody, &errCheck) == nil && errCheck.Error != "" {
		return nil, fmt.Errorf("%s: %s", serviceName, errCheck.Error)
	}

	return json.RawMessage(respBody), nil
}

// truncate shortens a string to maxLen, appending "..." if truncated.
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
