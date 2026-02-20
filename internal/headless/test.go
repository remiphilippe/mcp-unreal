package headless

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- run_tests ---

// RunTestsInput defines parameters for the run_tests tool.
type RunTestsInput struct {
	Filter   string `json:"filter,omitempty" jsonschema:"Test name filter pattern (e.g. 'MyProject.' for all project tests). Runs all tests if empty."`
	Config   string `json:"config,omitempty" jsonschema:"Build configuration: Development, DebugGame, Shipping. Default Development."`
	Platform string `json:"platform,omitempty" jsonschema:"Target platform. Defaults to current platform."`
}

// TestResult represents a single test's outcome.
type TestResult struct {
	Name     string   `json:"name" jsonschema:"fully qualified test name"`
	Status   string   `json:"status" jsonschema:"pass, fail, or skip"`
	Duration string   `json:"duration,omitempty" jsonschema:"test duration if available"`
	Events   []string `json:"events,omitempty" jsonschema:"failure or warning event messages"`
}

// RunTestsOutput is returned by the run_tests tool.
type RunTestsOutput struct {
	Success    bool         `json:"success" jsonschema:"true if all tests passed"`
	TotalTests int          `json:"total_tests" jsonschema:"total number of tests run"`
	Passed     int          `json:"passed" jsonschema:"number of passing tests"`
	Failed     int          `json:"failed" jsonschema:"number of failing tests"`
	Skipped    int          `json:"skipped" jsonschema:"number of skipped tests"`
	Duration   string       `json:"duration" jsonschema:"total run duration"`
	Results    []TestResult `json:"results" jsonschema:"per-test results"`
	LogPath    string       `json:"log_path,omitempty" jsonschema:"path to the full UE log file"`
	ExitCode   int          `json:"exit_code" jsonschema:"process exit code"`
}

// RegisterTests adds the test automation tools to the MCP server.
func (h *Handler) RegisterTests(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "run_tests",
		Description: "Run headless UE automation tests using UnrealEditor-Cmd with -nullrhi (no GPU). " +
			"Returns structured JSON with per-test pass/fail results and failure event details. " +
			"Does not require the editor to be running. " +
			"Call build_project first if you have edited C++ files.",
	}, h.RunTests)

	mcp.AddTool(server, &mcp.Tool{
		Name: "list_tests",
		Description: "List available UE automation test names matching a filter pattern. " +
			"Use this to discover test names before calling run_tests. " +
			"Does not require the editor to be running.",
	}, h.ListTests)

	mcp.AddTool(server, &mcp.Tool{
		Name: "run_visual_tests",
		Description: "Run UE automation tests WITH GPU rendering (no -nullrhi flag). " +
			"Use this instead of run_tests when tests require rendering or visual validation. " +
			"Slower than run_tests but supports screenshot comparison and visual regression testing. " +
			"Does not require the editor to be running (uses headless editor with GPU).",
	}, h.RunVisualTests)
}

// RunTests implements the run_tests tool.
func (h *Handler) RunTests(ctx context.Context, req *mcp.CallToolRequest, input RunTestsInput) (*mcp.CallToolResult, RunTestsOutput, error) {
	editorPath := h.Config.UEEditorPath
	if _, err := os.Stat(editorPath); err != nil {
		return nil, RunTestsOutput{}, fmt.Errorf(
			"UnrealEditor-Cmd not found at %s — set UE_EDITOR_PATH env var or install UE 5.7",
			editorPath,
		)
	}

	projectFile := h.Config.UProjectFile
	if projectFile == "" {
		return nil, RunTestsOutput{}, fmt.Errorf(
			"no .uproject file found — set MCP_UNREAL_PROJECT or run from inside a UE project directory",
		)
	}

	filter := input.Filter
	if filter == "" {
		filter = "." // Match all tests.
	}

	// Build the command: UnrealEditor-Cmd <project> -ExecCmds="Automation RunTests <filter>" -nullrhi -unattended -nopause
	execCmd := fmt.Sprintf("Automation RunTests %s;Quit", filter)
	args := []string{
		projectFile,
		"-ExecCmds=" + execCmd,
		"-nullrhi",
		"-unattended",
		"-nopause",
		"-nosplash",
		"-nosound",
		"-NoPCH",
		"-NoSharedPCH",
	}

	start := time.Now()
	stdout, stderr, exitCode, err := h.runCommand(ctx, editorPath, args, 30*time.Minute)
	duration := time.Since(start)

	if err != nil && exitCode == -1 {
		return nil, RunTestsOutput{}, fmt.Errorf("failed to run tests: %w", err)
	}

	combined := stdout + "\n" + stderr
	results := parseTestResults(combined)

	passed := 0
	failed := 0
	skipped := 0
	for _, r := range results {
		switch r.Status {
		case "pass":
			passed++
		case "fail":
			failed++
		case "skip":
			skipped++
		}
	}

	return nil, RunTestsOutput{
		Success:    failed == 0 && len(results) > 0,
		TotalTests: len(results),
		Passed:     passed,
		Failed:     failed,
		Skipped:    skipped,
		Duration:   duration.Round(time.Second).String(),
		Results:    results,
		ExitCode:   exitCode,
	}, nil
}

// --- run_visual_tests ---

// RunVisualTests implements the run_visual_tests tool. It runs tests with GPU
// rendering enabled (no -nullrhi), which is needed for visual regression tests
// and screenshot comparison.
func (h *Handler) RunVisualTests(ctx context.Context, req *mcp.CallToolRequest, input RunTestsInput) (*mcp.CallToolResult, RunTestsOutput, error) {
	editorPath := h.Config.UEEditorPath
	if _, err := os.Stat(editorPath); err != nil {
		return nil, RunTestsOutput{}, fmt.Errorf(
			"UnrealEditor-Cmd not found at %s — set UE_EDITOR_PATH env var or install UE 5.7",
			editorPath,
		)
	}

	projectFile := h.Config.UProjectFile
	if projectFile == "" {
		return nil, RunTestsOutput{}, fmt.Errorf(
			"no .uproject file found — set MCP_UNREAL_PROJECT or run from inside a UE project directory",
		)
	}

	filter := input.Filter
	if filter == "" {
		filter = "."
	}

	// Same as RunTests but WITHOUT -nullrhi to enable GPU rendering.
	execCmd := fmt.Sprintf("Automation RunTests %s;Quit", filter)
	args := []string{
		projectFile,
		"-ExecCmds=" + execCmd,
		"-unattended",
		"-nopause",
		"-nosplash",
		"-nosound",
	}

	start := time.Now()
	stdout, stderr, exitCode, err := h.runCommand(ctx, editorPath, args, 30*time.Minute)
	duration := time.Since(start)

	if err != nil && exitCode == -1 {
		return nil, RunTestsOutput{}, fmt.Errorf("failed to run visual tests: %w", err)
	}

	combined := stdout + "\n" + stderr
	results := parseTestResults(combined)

	passed := 0
	failed := 0
	skipped := 0
	for _, r := range results {
		switch r.Status {
		case "pass":
			passed++
		case "fail":
			failed++
		case "skip":
			skipped++
		}
	}

	return nil, RunTestsOutput{
		Success:    failed == 0 && len(results) > 0,
		TotalTests: len(results),
		Passed:     passed,
		Failed:     failed,
		Skipped:    skipped,
		Duration:   duration.Round(time.Second).String(),
		Results:    results,
		ExitCode:   exitCode,
	}, nil
}

// --- list_tests ---

// ListTestsInput defines parameters for the list_tests tool.
type ListTestsInput struct {
	Filter string `json:"filter,omitempty" jsonschema:"Test name filter pattern. Lists all if empty."`
}

// ListTestsOutput is returned by the list_tests tool.
type ListTestsOutput struct {
	Tests []string `json:"tests" jsonschema:"list of available test names"`
	Total int      `json:"total" jsonschema:"total number of matching tests"`
}

// ListTests implements the list_tests tool.
func (h *Handler) ListTests(ctx context.Context, req *mcp.CallToolRequest, input ListTestsInput) (*mcp.CallToolResult, ListTestsOutput, error) {
	editorPath := h.Config.UEEditorPath
	if _, err := os.Stat(editorPath); err != nil {
		return nil, ListTestsOutput{}, fmt.Errorf(
			"UnrealEditor-Cmd not found at %s — set UE_EDITOR_PATH env var or install UE 5.7",
			editorPath,
		)
	}

	projectFile := h.Config.UProjectFile
	if projectFile == "" {
		return nil, ListTestsOutput{}, fmt.Errorf(
			"no .uproject file found — set MCP_UNREAL_PROJECT or run from inside a UE project directory",
		)
	}

	execCmd := "Automation List;Quit"
	args := []string{
		projectFile,
		"-ExecCmds=" + execCmd,
		"-nullrhi",
		"-unattended",
		"-nopause",
		"-nosplash",
		"-nosound",
	}

	stdout, stderr, _, err := h.runCommand(ctx, editorPath, args, 10*time.Minute)
	if err != nil {
		return nil, ListTestsOutput{}, fmt.Errorf("failed to list tests: %w", err)
	}

	combined := stdout + "\n" + stderr
	tests := parseTestList(combined)

	// Apply filter if provided.
	if input.Filter != "" {
		var filtered []string
		lowerFilter := strings.ToLower(input.Filter)
		for _, t := range tests {
			if strings.Contains(strings.ToLower(t), lowerFilter) {
				filtered = append(filtered, t)
			}
		}
		tests = filtered
	}

	return nil, ListTestsOutput{
		Tests: tests,
		Total: len(tests),
	}, nil
}

// --- log parsing ---

// UE test log line patterns:
//
//	LogAutomationController: Test Completed. Result={Passed|Failed} Test={TestName}
//	LogAutomationController: Error: ...
//	LogAutomationController: Warning: ...
//	LogAutomationController: BeginEvents: ...
//	LogAutomationController:   Event: ...
var (
	testResultRe = regexp.MustCompile(
		`LogAutomationController.*Test Completed\.\s+Result=\{(\w+)\}\s+Test=\{([^}]+)\}`,
	)
	testEventRe = regexp.MustCompile(
		`LogAutomationController.*(?:Error|Warning):\s+(.+)`,
	)
	testListRe = regexp.MustCompile(
		`LogAutomationController.*\]\s+([\w.]+)`,
	)
	testDurationRe = regexp.MustCompile(
		`Test Completed.*Duration=\{([^}]+)\}`,
	)
)

// parseTestResults extracts per-test results from UE log output.
func parseTestResults(output string) []TestResult {
	lines := strings.Split(output, "\n")

	var results []TestResult
	currentEvents := []string{}

	for _, line := range lines {
		// Collect error/warning events.
		if eventMatch := testEventRe.FindStringSubmatch(line); eventMatch != nil {
			currentEvents = append(currentEvents, strings.TrimSpace(eventMatch[1]))
			continue
		}

		// Match test completion lines.
		if resultMatch := testResultRe.FindStringSubmatch(line); resultMatch != nil {
			status := "pass"
			switch strings.ToLower(resultMatch[1]) {
			case "failed", "fail":
				status = "fail"
			case "skipped", "skip", "notrun":
				status = "skip"
			}

			duration := ""
			if durMatch := testDurationRe.FindStringSubmatch(line); durMatch != nil {
				duration = durMatch[1]
			}

			result := TestResult{
				Name:     resultMatch[2],
				Status:   status,
				Duration: duration,
			}

			// Attach accumulated events to failing tests.
			if status == "fail" && len(currentEvents) > 0 {
				result.Events = currentEvents
			}
			currentEvents = nil

			results = append(results, result)
		}
	}

	return results
}

// parseTestList extracts test names from "Automation List" output.
func parseTestList(output string) []string {
	lines := strings.Split(output, "\n")
	seen := make(map[string]bool)
	var tests []string

	for _, line := range lines {
		if !strings.Contains(line, "LogAutomationController") {
			continue
		}
		if matches := testListRe.FindStringSubmatch(line); matches != nil {
			name := strings.TrimSpace(matches[1])
			if name != "" && !seen[name] && strings.Contains(name, ".") {
				seen[name] = true
				tests = append(tests, name)
			}
		}
	}

	return tests
}
