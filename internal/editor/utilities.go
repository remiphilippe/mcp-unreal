package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- run_console_command ---

// RunConsoleCommandInput defines parameters for the run_console_command tool.
type RunConsoleCommandInput struct {
	Command string `json:"command" jsonschema:"required,UE console command to execute (e.g. stat fps, obj list, ShowFlag.Collision 1)"`
	World   string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// RunConsoleCommandOutput is returned by the run_console_command tool.
type RunConsoleCommandOutput struct {
	Success bool   `json:"success" jsonschema:"whether the command was sent successfully"`
	Command string `json:"command" jsonschema:"the command that was executed"`
	Output  string `json:"output,omitempty" jsonschema:"command output if captured"`
}

// RegisterUtilities adds the run_console_command tool to the MCP server.
func (h *Handler) RegisterUtilities(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "run_console_command",
		Description: "Execute a UE console command in the running editor. " +
			"WARNING: Can execute arbitrary editor commands — use with care. " +
			"Requires the editor to be running. Tries the MCPUnreal plugin first, " +
			"falls back to the Remote Control API. " +
			"Common commands: stat fps, stat unit, obj list, ShowFlag.*, log list, r.SetRes 1920x1080",
	}, h.RunConsoleCommand)
}

// RunConsoleCommand implements the run_console_command tool.
// It tries the plugin endpoint first (which captures output), then
// falls back to the RC API via KismetSystemLibrary::ExecuteConsoleCommand.
func (h *Handler) RunConsoleCommand(ctx context.Context, req *mcp.CallToolRequest, input RunConsoleCommandInput) (*mcp.CallToolResult, RunConsoleCommandOutput, error) {
	if input.Command == "" {
		return nil, RunConsoleCommandOutput{}, fmt.Errorf("command is required")
	}

	h.Logger.Info("executing console command", "command", input.Command)

	// Try plugin endpoint first — it captures command output.
	body := map[string]any{
		"command": input.Command,
	}
	if input.World != "" {
		body["world"] = input.World
	}
	resp, err := h.Client.PluginCall(ctx, "/api/editor/console_command", body)
	if err == nil {
		var result struct {
			Output string `json:"output"`
		}
		_ = json.Unmarshal(resp, &result)
		return nil, RunConsoleCommandOutput{
			Success: true,
			Command: input.Command,
			Output:  result.Output,
		}, nil
	}

	// Fallback: RC API via KismetSystemLibrary::ExecuteConsoleCommand.
	// This sends the command but cannot capture its output.
	rcBody := map[string]any{
		"objectPath":   "/Script/Engine.Default__KismetSystemLibrary",
		"functionName": "ExecuteConsoleCommand",
		"parameters": map[string]any{
			"WorldContextObject": "",
			"Command":            input.Command,
		},
	}
	_, rcErr := h.Client.RCAPICall(ctx, "/remote/object/call", rcBody)
	if rcErr != nil {
		return nil, RunConsoleCommandOutput{}, fmt.Errorf(
			"could not execute console command — ensure the UE editor is running: "+
				"plugin error: %v, RC API error: %w", err, rcErr,
		)
	}

	return nil, RunConsoleCommandOutput{
		Success: true,
		Command: input.Command,
		Output:  "command sent via RC API (output not captured — use get_output_log to see results)",
	}, nil
}
