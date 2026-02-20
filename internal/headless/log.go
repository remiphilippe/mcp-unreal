// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package headless

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- get_test_log ---

// GetTestLogInput defines parameters for the get_test_log tool.
type GetTestLogInput struct {
	LogPath  string `json:"log_path,omitempty" jsonschema:"Path to the UE log file. Auto-detected from project Saved/Logs/ if empty."`
	MaxLines int    `json:"max_lines,omitempty" jsonschema:"Maximum number of lines to return. Default 200."`
	Offset   int    `json:"offset,omitempty" jsonschema:"Line offset to start reading from (0-based). Default 0."`
	Filter   string `json:"filter,omitempty" jsonschema:"Only return lines containing this substring (case-insensitive)."`
}

// GetTestLogOutput is returned by the get_test_log tool.
type GetTestLogOutput struct {
	LogPath    string `json:"log_path" jsonschema:"path to the log file read"`
	Content    string `json:"content" jsonschema:"log content (filtered and truncated)"`
	TotalLines int    `json:"total_lines" jsonschema:"total lines in the file"`
	Returned   int    `json:"returned" jsonschema:"number of lines returned"`
}

// RegisterLog adds the get_test_log tool to the MCP server.
func (h *Handler) RegisterLog(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "get_test_log",
		Description: "Read the raw UE log from the most recent test run. " +
			"Supports line limits, offsets, and substring filtering. " +
			"Use this to inspect detailed output when run_tests reports failures. " +
			"Does not require the editor to be running.",
	}, h.GetTestLog)
}

// GetTestLog implements the get_test_log tool.
func (h *Handler) GetTestLog(ctx context.Context, req *mcp.CallToolRequest, input GetTestLogInput) (*mcp.CallToolResult, GetTestLogOutput, error) {
	logPath := input.LogPath
	if logPath == "" {
		logPath = h.findLatestLog()
		if logPath == "" {
			return nil, GetTestLogOutput{}, fmt.Errorf(
				"no UE log file found in project Saved/Logs/ — specify log_path explicitly",
			)
		}
	}

	// Path validation: reject traversal attempts (CLAUDE.md Security §4).
	cleaned := filepath.Clean(logPath)
	if strings.Contains(cleaned, "..") {
		return nil, GetTestLogOutput{}, fmt.Errorf(
			"invalid log path: path traversal not allowed",
		)
	}

	data, err := os.ReadFile(cleaned)
	if err != nil {
		return nil, GetTestLogOutput{}, fmt.Errorf("reading log file %s: %w", cleaned, err)
	}

	lines := strings.Split(string(data), "\n")
	totalLines := len(lines)

	// Apply filter.
	if input.Filter != "" {
		lowerFilter := strings.ToLower(input.Filter)
		var filtered []string
		for _, line := range lines {
			if strings.Contains(strings.ToLower(line), lowerFilter) {
				filtered = append(filtered, line)
			}
		}
		lines = filtered
	}

	// Apply offset.
	offset := input.Offset
	if offset > 0 && offset < len(lines) {
		lines = lines[offset:]
	} else if offset >= len(lines) {
		lines = nil
	}

	// Apply max lines (default 200, cap at 500 for token budget).
	maxLines := input.MaxLines
	if maxLines <= 0 {
		maxLines = 200
	}
	if maxLines > 500 {
		maxLines = 500
	}
	if len(lines) > maxLines {
		lines = lines[:maxLines]
	}

	return nil, GetTestLogOutput{
		LogPath:    cleaned,
		Content:    strings.Join(lines, "\n"),
		TotalLines: totalLines,
		Returned:   len(lines),
	}, nil
}

// findLatestLog looks for the most recent .log file in the project's
// Saved/Logs/ directory.
func (h *Handler) findLatestLog() string {
	if h.Config.ProjectRoot == "" {
		return ""
	}

	logsDir := filepath.Join(h.Config.ProjectRoot, "Saved", "Logs")
	entries, err := os.ReadDir(logsDir)
	if err != nil {
		return ""
	}

	var latestPath string
	var latestTime int64

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".log") {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		if info.ModTime().UnixNano() > latestTime {
			latestTime = info.ModTime().UnixNano()
			latestPath = filepath.Join(logsDir, e.Name())
		}
	}

	return latestPath
}
