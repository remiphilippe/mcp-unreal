package editor

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- get_output_log ---

// GetOutputLogInput defines parameters for the get_output_log tool.
type GetOutputLogInput struct {
	Category     string  `json:"category,omitempty" jsonschema:"Filter by log category substring (e.g. LogTemp, LogBlueprintUserMessages, LogMCPUnreal)"`
	Verbosity    string  `json:"verbosity,omitempty" jsonschema:"Minimum verbosity: fatal, error, warning, display, log, verbose (default: all)"`
	Pattern      string  `json:"pattern,omitempty" jsonschema:"Regex pattern to match against log message text (e.g. 'terrain.*failed', 'spawn')"`
	MaxLines     int     `json:"max_lines,omitempty" jsonschema:"Maximum number of log entries to return (default 100)"`
	SinceSeconds float64 `json:"since_seconds,omitempty" jsonschema:"Only return log entries from the last N seconds (e.g. 30 for last 30 seconds)"`
}

// LogEntry represents a single output log entry.
type LogEntry struct {
	Category  string `json:"category" jsonschema:"log category"`
	Verbosity string `json:"verbosity" jsonschema:"log verbosity level"`
	Message   string `json:"message" jsonschema:"log message text"`
}

// GetOutputLogOutput is returned by the get_output_log tool.
type GetOutputLogOutput struct {
	Entries []LogEntry `json:"entries" jsonschema:"log entries"`
	Count   int        `json:"count" jsonschema:"number of entries returned"`
}

// --- capture_viewport ---

// CaptureViewportInput defines parameters for the capture_viewport tool.
type CaptureViewportInput struct {
	OutputPath string `json:"output_path,omitempty" jsonschema:"File path to save the PNG screenshot — if empty, returns base64"`
	World      string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
	IncludeUI  bool   `json:"include_ui,omitempty" jsonschema:"If true, capture includes Slate/UMG UI overlays (HUD, menus, debug text). Requires PIE. Uses async FScreenshotRequest — may take one extra frame"`
}

// CaptureViewportOutput is returned by the capture_viewport tool.
type CaptureViewportOutput struct {
	Success     bool   `json:"success" jsonschema:"whether the capture succeeded"`
	FilePath    string `json:"file_path,omitempty" jsonschema:"file path if saved to disk"`
	ImageBase64 string `json:"image_base64,omitempty" jsonschema:"base64-encoded PNG if no output_path"`
	Format      string `json:"format,omitempty" jsonschema:"image format (png)"`
	Width       int    `json:"width" jsonschema:"image width in pixels"`
	Height      int    `json:"height" jsonschema:"image height in pixels"`
}

// --- execute_script ---

// ExecuteScriptInput defines parameters for the execute_script tool.
type ExecuteScriptInput struct {
	Script string `json:"script" jsonschema:"required,Python script code to execute in the editor (requires Python Editor Script Plugin)"`
	World  string `json:"world,omitempty" jsonschema:"Target world: auto (default, PIE if active else editor), pie (error if not running), editor (always editor)"`
}

// ExecuteScriptOutput is returned by the execute_script tool.
type ExecuteScriptOutput struct {
	Success bool   `json:"success" jsonschema:"whether the script executed successfully"`
	Output  string `json:"output,omitempty" jsonschema:"script output if captured"`
}

// RegisterEditorUtils adds the output log, viewport capture, and script execution
// tools to the MCP server.
func (h *Handler) RegisterEditorUtils(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "get_output_log",
		Description: "Read recent entries from the UE editor output log. " +
			"Supports filtering by category (e.g. LogTemp, LogBlueprintUserMessages), " +
			"minimum verbosity level (fatal, error, warning, display, log, verbose), " +
			"regex pattern matching on message text, and time-based filtering (since_seconds). " +
			"Returns up to max_lines entries (default 100). " +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.GetOutputLog)

	mcp.AddTool(server, &mcp.Tool{
		Name: "capture_viewport",
		Description: "Capture a screenshot of the active viewport as PNG. " +
			"If output_path is provided, saves to disk. Otherwise returns base64-encoded image data. " +
			"Supports world parameter to capture PIE game viewport or editor viewport (default: auto). " +
			"Set include_ui=true to capture with Slate/UMG overlays (HUD, menus) — requires PIE. " +
			"Requires the editor running with MCPUnreal plugin (port 8090).",
	}, h.CaptureViewport)

	mcp.AddTool(server, &mcp.Tool{
		Name: "execute_script",
		Description: "Execute a Python script in the UE editor via the Python Editor Script Plugin. " +
			"WARNING: This can run arbitrary code in the editor — use with care. " +
			"The full script is logged before execution for security auditing. " +
			"Requires the editor running with MCPUnreal plugin (port 8090) and " +
			"the Python Editor Script Plugin enabled.",
	}, h.ExecuteScript)

	mcp.AddTool(server, &mcp.Tool{
		Name: "live_compile",
		Description: "Trigger Live Coding (hot reload) compilation in the running editor. " +
			"Recompiles modified C++ files without restarting the editor. " +
			"Returns whether compilation succeeded and any error details. " +
			"Requires the editor running with MCPUnreal plugin (port 8090) and " +
			"Live Coding enabled in Editor Preferences.",
	}, h.LiveCompile)

	mcp.AddTool(server, &mcp.Tool{
		Name: "pie_control",
		Description: "Control Play In Editor (PIE) sessions. " +
			"Operations: 'start' begins a PIE session (async — returns immediately, PIE starts next frame), " +
			"'stop' ends the current PIE session, 'status' checks if PIE is active. " +
			"After start/stop, call with operation 'status' to verify the state change completed. " +
			"Use 'simulate' option with start for Simulate In Editor mode.",
	}, h.PIEControl)

	mcp.AddTool(server, &mcp.Tool{
		Name: "player_control",
		Description: "Control the player pawn and editor viewport camera. " +
			"Operations: 'get_info' (get player controller, pawn, location, rotation, camera — requires PIE), " +
			"'teleport' (move pawn to location/rotation — requires PIE), " +
			"'set_rotation' (set player view direction — requires PIE), " +
			"'set_view_target' (change camera to follow another actor — requires PIE), " +
			"'get_camera' (get editor viewport camera position — works without PIE), " +
			"'set_camera' (move editor viewport camera — works without PIE). " +
			"Use after pie_control(start) to navigate the player in PIE, or without PIE for editor camera.",
	}, h.PlayerControl)
}

// GetOutputLog implements the get_output_log tool.
func (h *Handler) GetOutputLog(ctx context.Context, req *mcp.CallToolRequest, input GetOutputLogInput) (*mcp.CallToolResult, GetOutputLogOutput, error) {
	body := map[string]any{}
	if input.Category != "" {
		body["category"] = input.Category
	}
	if input.Verbosity != "" {
		body["verbosity"] = input.Verbosity
	}
	if input.Pattern != "" {
		body["pattern"] = input.Pattern
	}
	if input.MaxLines > 0 {
		body["max_lines"] = input.MaxLines
	}
	if input.SinceSeconds > 0 {
		body["since_seconds"] = input.SinceSeconds
	}

	resp, err := h.Client.PluginCall(ctx, "/api/editor/output_log", body)
	if err != nil {
		return nil, GetOutputLogOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out GetOutputLogOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, GetOutputLogOutput{}, fmt.Errorf("parsing output log response: %w", err)
	}

	return nil, out, nil
}

// CaptureViewport implements the capture_viewport tool.
// When no output_path is given, returns the image as MCP ImageContent so the
// LLM can actually see the screenshot (not just a base64 text blob).
func (h *Handler) CaptureViewport(ctx context.Context, req *mcp.CallToolRequest, input CaptureViewportInput) (*mcp.CallToolResult, CaptureViewportOutput, error) {
	body := map[string]any{}
	if input.OutputPath != "" {
		body["output_path"] = input.OutputPath
	}
	if input.World != "" {
		body["world"] = input.World
	}
	if input.IncludeUI {
		body["include_ui"] = true
	}

	resp, err := h.Client.PluginCall(ctx, "/api/editor/capture_viewport", body)
	if err != nil {
		return nil, CaptureViewportOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out CaptureViewportOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, CaptureViewportOutput{}, fmt.Errorf("parsing viewport capture response: %w", err)
	}

	// When the plugin returns base64 image data (no output_path), return it as
	// MCP ImageContent so the LLM receives a viewable image, not a text blob.
	if out.ImageBase64 != "" {
		imageBytes, err := base64.StdEncoding.DecodeString(out.ImageBase64)
		if err != nil {
			return nil, CaptureViewportOutput{}, fmt.Errorf("decoding base64 image data: %w", err)
		}

		result := &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.ImageContent{
					Data:     imageBytes,
					MIMEType: "image/png",
				},
			},
		}

		// Clear the base64 field from the structured output — image is in Content.
		out.ImageBase64 = ""
		return result, out, nil
	}

	return nil, out, nil
}

// ExecuteScript implements the execute_script tool.
func (h *Handler) ExecuteScript(ctx context.Context, req *mcp.CallToolRequest, input ExecuteScriptInput) (*mcp.CallToolResult, ExecuteScriptOutput, error) {
	if input.Script == "" {
		return nil, ExecuteScriptOutput{}, fmt.Errorf("script is required")
	}

	h.Logger.Warn("executing script via editor", "script_length", len(input.Script))

	body := map[string]any{
		"script": input.Script,
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/editor/execute_script", body)
	if err != nil {
		return nil, ExecuteScriptOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out ExecuteScriptOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, ExecuteScriptOutput{}, fmt.Errorf("parsing script execution response: %w", err)
	}

	return nil, out, nil
}

// --- pie_control ---

// PIEControlInput defines parameters for the pie_control tool.
type PIEControlInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: start (begin PIE session), stop (end PIE session), status (check PIE state)"`
	MapPath   string `json:"map_path,omitempty" jsonschema:"Map to play (e.g. /Game/Maps/MyLevel). Only used with start. Default: current editor map"`
	Simulate  bool   `json:"simulate,omitempty" jsonschema:"If true, start Simulate In Editor instead of Play In Editor. Only used with start"`
}

// PIEControlOutput is returned by the pie_control tool.
type PIEControlOutput struct {
	Success   bool   `json:"success" jsonschema:"whether the operation succeeded"`
	Message   string `json:"message,omitempty" jsonschema:"human-readable result message"`
	PIEActive bool   `json:"pie_active" jsonschema:"whether PIE is currently active"`
	PIEMap    string `json:"pie_map,omitempty" jsonschema:"map name if PIE is active"`
	Error     string `json:"error,omitempty" jsonschema:"error message if operation failed"`
}

// --- player_control ---

// PlayerControlInput defines parameters for the player_control tool.
type PlayerControlInput struct {
	Operation string      `json:"operation" jsonschema:"required,Operation: get_info (player state), teleport (move pawn), set_rotation (set view direction), set_view_target (change camera target), get_camera (editor viewport), set_camera (move editor viewport)"`
	Location  *[3]float64 `json:"location,omitempty" jsonschema:"[X,Y,Z] world position in centimeters (for teleport, set_camera)"`
	Rotation  *[3]float64 `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] in degrees (for teleport, set_rotation, set_camera)"`
	ActorPath string      `json:"actor_path,omitempty" jsonschema:"Actor path or label for set_view_target"`
	World     string      `json:"world,omitempty" jsonschema:"Target world: auto (default), pie, editor. Player ops default to PIE, camera ops default to editor"`
}

// PlayerControlOutput is returned by the player_control tool.
type PlayerControlOutput struct {
	ControllerPath  string      `json:"controller_path,omitempty" jsonschema:"path to the player controller"`
	PawnPath        string      `json:"pawn_path,omitempty" jsonschema:"path to the possessed pawn"`
	PawnClass       string      `json:"pawn_class,omitempty" jsonschema:"class name of the pawn"`
	Location        *[3]float64 `json:"location,omitempty" jsonschema:"[X,Y,Z] pawn world position"`
	Rotation        *[3]float64 `json:"rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] pawn rotation"`
	ControlRotation *[3]float64 `json:"control_rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] controller view rotation"`
	CameraLocation  *[3]float64 `json:"camera_location,omitempty" jsonschema:"[X,Y,Z] camera position"`
	CameraRotation  *[3]float64 `json:"camera_rotation,omitempty" jsonschema:"[Pitch,Yaw,Roll] camera rotation"`
	TargetPath      string      `json:"target_path,omitempty" jsonschema:"path to the view target actor"`
	Success         bool        `json:"success" jsonschema:"whether the operation succeeded"`
	Message         string      `json:"message,omitempty" jsonschema:"human-readable result message"`
	Error           string      `json:"error,omitempty" jsonschema:"error message if operation failed"`
}

// --- live_compile ---

// LiveCompileInput defines parameters for the live_compile tool.
type LiveCompileInput struct{}

// LiveCompileOutput is returned by the live_compile tool.
type LiveCompileOutput struct {
	Success bool   `json:"success" jsonschema:"whether compilation succeeded"`
	Status  string `json:"status,omitempty" jsonschema:"compilation status (e.g. Compiling, Succeeded, Failed)"`
	Errors  string `json:"errors,omitempty" jsonschema:"compilation error details if any"`
}

// LiveCompile implements the live_compile tool.
func (h *Handler) LiveCompile(ctx context.Context, req *mcp.CallToolRequest, input LiveCompileInput) (*mcp.CallToolResult, LiveCompileOutput, error) {
	resp, err := h.Client.PluginCall(ctx, "/api/editor/live_compile", map[string]any{})
	if err != nil {
		return nil, LiveCompileOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out LiveCompileOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, LiveCompileOutput{}, fmt.Errorf("parsing live compile response: %w", err)
	}

	return nil, out, nil
}

// PIEControl implements the pie_control tool.
func (h *Handler) PIEControl(ctx context.Context, req *mcp.CallToolRequest, input PIEControlInput) (*mcp.CallToolResult, PIEControlOutput, error) {
	if input.Operation == "" {
		return nil, PIEControlOutput{}, fmt.Errorf("operation is required (start, stop, status)")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.MapPath != "" {
		body["map_path"] = input.MapPath
	}
	if input.Simulate {
		body["simulate"] = true
	}

	resp, err := h.Client.PluginCall(ctx, "/api/editor/pie_control", body)
	if err != nil {
		return nil, PIEControlOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out PIEControlOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, PIEControlOutput{}, fmt.Errorf("parsing PIE control response: %w", err)
	}

	return nil, out, nil
}

// PlayerControl implements the player_control tool.
func (h *Handler) PlayerControl(ctx context.Context, req *mcp.CallToolRequest, input PlayerControlInput) (*mcp.CallToolResult, PlayerControlOutput, error) {
	if input.Operation == "" {
		return nil, PlayerControlOutput{}, fmt.Errorf("operation is required (get_info, teleport, set_rotation, set_view_target, get_camera, set_camera)")
	}

	body := map[string]any{
		"operation": input.Operation,
	}
	if input.Location != nil {
		body["location"] = input.Location
	}
	if input.Rotation != nil {
		body["rotation"] = input.Rotation
	}
	if input.ActorPath != "" {
		body["actor_path"] = input.ActorPath
	}
	if input.World != "" {
		body["world"] = input.World
	}

	resp, err := h.Client.PluginCall(ctx, "/api/editor/player_control", body)
	if err != nil {
		return nil, PlayerControlOutput{}, fmt.Errorf(
			"editor unreachable — ensure UE is running with the MCPUnreal plugin loaded: %w", err,
		)
	}

	var out PlayerControlOutput
	if err := json.Unmarshal(resp, &out); err != nil {
		return nil, PlayerControlOutput{}, fmt.Errorf("parsing player control response: %w", err)
	}

	return nil, out, nil
}
