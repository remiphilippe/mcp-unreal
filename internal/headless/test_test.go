// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package headless

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/remiphilippe/mcp-unreal/internal/config"
)

func readTestdata(t *testing.T, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("testdata", name)) //nolint:gosec // test fixtures only
	if err != nil {
		t.Fatalf("failed to read testdata/%s: %v", name, err)
	}
	return string(data)
}

func TestParseTestResults(t *testing.T) {
	tests := []struct {
		name       string
		fixture    string
		wantTotal  int
		wantPassed int
		wantFailed int
	}{
		{
			name:       "all passing",
			fixture:    "passing_tests.log",
			wantTotal:  3,
			wantPassed: 3,
			wantFailed: 0,
		},
		{
			name:       "some failing",
			fixture:    "failing_tests.log",
			wantTotal:  4,
			wantPassed: 2,
			wantFailed: 2,
		},
		{
			name:       "empty log",
			fixture:    "empty.log",
			wantTotal:  0,
			wantPassed: 0,
			wantFailed: 0,
		},
		{
			name:       "no tests matched filter",
			fixture:    "no_tests_found.log",
			wantTotal:  0,
			wantPassed: 0,
			wantFailed: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := readTestdata(t, tt.fixture)
			results := parseTestResults(output)

			if len(results) != tt.wantTotal {
				t.Errorf("total tests = %d, want %d", len(results), tt.wantTotal)
			}

			passed := 0
			failed := 0
			for _, r := range results {
				switch r.Status {
				case "pass":
					passed++
				case "fail":
					failed++
				}
			}

			if passed != tt.wantPassed {
				t.Errorf("passed = %d, want %d", passed, tt.wantPassed)
			}
			if failed != tt.wantFailed {
				t.Errorf("failed = %d, want %d", failed, tt.wantFailed)
			}
		})
	}
}

func TestParseTestResults_FailureEvents(t *testing.T) {
	output := readTestdata(t, "failing_tests.log")
	results := parseTestResults(output)

	// Find the CalculationTest failure — it should have error events.
	var calcTest *TestResult
	for i, r := range results {
		if r.Name == "MyProject.Unit.CalculationTest" {
			calcTest = &results[i]
			break
		}
	}

	if calcTest == nil {
		t.Fatal("MyProject.Unit.CalculationTest not found in results")
	}

	if calcTest.Status != "fail" {
		t.Errorf("CalculationTest status = %q, want fail", calcTest.Status)
	}

	if len(calcTest.Events) == 0 {
		t.Error("CalculationTest should have failure events")
	}

	// Check that error messages were captured.
	foundExpectedError := false
	for _, event := range calcTest.Events {
		if event == "Expected value 42 but got 0" {
			foundExpectedError = true
		}
	}
	if !foundExpectedError {
		t.Errorf("expected 'Expected value 42 but got 0' in events, got: %v", calcTest.Events)
	}
}

func TestParseTestResults_Duration(t *testing.T) {
	output := readTestdata(t, "passing_tests.log")
	results := parseTestResults(output)

	if len(results) == 0 {
		t.Fatal("expected results")
	}

	// First test should have a duration.
	if results[0].Duration == "" {
		t.Error("expected duration on first test result")
	}
	if results[0].Duration != "0.012s" {
		t.Errorf("duration = %q, want 0.012s", results[0].Duration)
	}
}

func TestParseTestResults_SkipStatus(t *testing.T) {
	// Test that skip/notrun statuses are parsed correctly.
	output := `[2025.01.01-00.00.00:000][  0]LogAutomationController: Display: Test Completed. Result={Skipped} Test={MyProject.Unit.SkippedTest}`
	results := parseTestResults(output)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Status != "skip" {
		t.Errorf("status = %q, want skip", results[0].Status)
	}
}

func TestParseTestResults_PassingHasNoEvents(t *testing.T) {
	output := readTestdata(t, "passing_tests.log")
	results := parseTestResults(output)
	for _, r := range results {
		if len(r.Events) > 0 {
			t.Errorf("passing test %q should not have events, got: %v", r.Name, r.Events)
		}
	}
}

func TestParseTestList(t *testing.T) {
	output := readTestdata(t, "test_list.log")
	tests := parseTestList(output)

	if len(tests) != 5 {
		t.Errorf("expected 5 tests, got %d: %v", len(tests), tests)
	}

	expected := []string{
		"MyProject.Unit.MathUtils",
		"MyProject.Unit.StringUtils",
		"MyProject.Unit.CalculationTest",
		"MyProject.Integration.GameMode",
		"MyProject.Visual.ScreenshotTest",
	}

	for _, want := range expected {
		found := false
		for _, got := range tests {
			if got == want {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected test %q not found in list: %v", want, tests)
		}
	}
}

func TestParseTestList_Empty(t *testing.T) {
	tests := parseTestList("")
	if len(tests) != 0 {
		t.Errorf("expected 0 tests from empty input, got %d", len(tests))
	}
}

func TestParseTestList_Dedup(t *testing.T) {
	// Duplicate test names should be deduplicated.
	output := `[2025.01.01-00.00.00:000][  0]LogAutomationController: Display: ] MyProject.Unit.MathUtils
[2025.01.01-00.00.00:000][  0]LogAutomationController: Display: ] MyProject.Unit.MathUtils`
	tests := parseTestList(output)
	if len(tests) != 1 {
		t.Errorf("expected 1 test (deduped), got %d: %v", len(tests), tests)
	}
}

func TestParseTestList_IgnoresNonDotted(t *testing.T) {
	// Names without dots are ignored (not proper test names).
	output := `[2025.01.01-00.00.00:000][  0]LogAutomationController: Display: ] SomeNonTestName`
	tests := parseTestList(output)
	if len(tests) != 0 {
		t.Errorf("expected 0 tests for non-dotted name, got %d: %v", len(tests), tests)
	}
}

// --- Handler-level tests for run_tests, run_visual_tests, list_tests ---

func TestRunTests_AllPass(t *testing.T) {
	fixture := readTestdata(t, "passing_tests.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.RunTests(context.Background(), nil, RunTestsInput{
		Filter: "MyProject.",
	})
	if err != nil {
		t.Fatalf("RunTests returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
	if out.TotalTests != 3 {
		t.Errorf("TotalTests = %d, want 3", out.TotalTests)
	}
	if out.Passed != 3 {
		t.Errorf("Passed = %d, want 3", out.Passed)
	}
	if out.Failed != 0 {
		t.Errorf("Failed = %d, want 0", out.Failed)
	}
	if out.ExitCode != 0 {
		t.Errorf("ExitCode = %d, want 0", out.ExitCode)
	}
}

func TestRunTests_SomeFail(t *testing.T) {
	fixture := readTestdata(t, "failing_tests.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 1)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.RunTests(context.Background(), nil, RunTestsInput{})
	if err != nil {
		t.Fatalf("RunTests returned error: %v", err)
	}
	if out.Success {
		t.Error("Success = true, want false (there are failures)")
	}
	if out.TotalTests != 4 {
		t.Errorf("TotalTests = %d, want 4", out.TotalTests)
	}
	if out.Passed != 2 {
		t.Errorf("Passed = %d, want 2", out.Passed)
	}
	if out.Failed != 2 {
		t.Errorf("Failed = %d, want 2", out.Failed)
	}

	// Verify failure events were captured.
	for _, r := range out.Results {
		if r.Status == "fail" && r.Name == "MyProject.Unit.CalculationTest" {
			if len(r.Events) == 0 {
				t.Error("expected events on CalculationTest failure")
			}
		}
	}
}

func TestRunTests_EmptyLog(t *testing.T) {
	editorPath, projectFile := createFakeEditor(t, "", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.RunTests(context.Background(), nil, RunTestsInput{})
	if err != nil {
		t.Fatalf("RunTests returned error: %v", err)
	}
	// No results parsed from empty output means not successful (len(results) == 0).
	if out.Success {
		t.Error("Success = true, want false for empty log (no tests run)")
	}
	if out.TotalTests != 0 {
		t.Errorf("TotalTests = %d, want 0", out.TotalTests)
	}
}

func TestRunTests_NoEditor(t *testing.T) {
	h := &Handler{
		Config: &config.Config{
			UEEditorPath: "/nonexistent/UnrealEditor-Cmd",
			UProjectFile: "/some/Test.uproject",
		},
		Logger: testLogger(),
	}

	_, _, err := h.RunTests(context.Background(), nil, RunTestsInput{})
	if err == nil {
		t.Fatal("expected error for missing editor")
	}
	if !strings.Contains(err.Error(), "UnrealEditor-Cmd not found") {
		t.Errorf("error = %q, want to contain 'UnrealEditor-Cmd not found'", err.Error())
	}
}

func TestRunTests_NoProject(t *testing.T) {
	editorPath, _ := createFakeEditor(t, "", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: "",
		},
		Logger: testLogger(),
	}

	_, _, err := h.RunTests(context.Background(), nil, RunTestsInput{})
	if err == nil {
		t.Fatal("expected error for missing project")
	}
	if !strings.Contains(err.Error(), "no .uproject file found") {
		t.Errorf("error = %q, want to contain 'no .uproject file found'", err.Error())
	}
}

func TestRunTests_DefaultFilter(t *testing.T) {
	// Empty filter should default to "." (match all).
	fixture := readTestdata(t, "passing_tests.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.RunTests(context.Background(), nil, RunTestsInput{
		Filter: "", // Should default to ".".
	})
	if err != nil {
		t.Fatalf("RunTests returned error: %v", err)
	}
	if out.TotalTests != 3 {
		t.Errorf("TotalTests = %d, want 3", out.TotalTests)
	}
}

func TestRunVisualTests_Success(t *testing.T) {
	fixture := readTestdata(t, "passing_tests.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.RunVisualTests(context.Background(), nil, RunTestsInput{
		Filter: "MyProject.",
	})
	if err != nil {
		t.Fatalf("RunVisualTests returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
	if out.TotalTests != 3 {
		t.Errorf("TotalTests = %d, want 3", out.TotalTests)
	}
	if out.Passed != 3 {
		t.Errorf("Passed = %d, want 3", out.Passed)
	}
}

func TestRunVisualTests_NoEditor(t *testing.T) {
	h := &Handler{
		Config: &config.Config{
			UEEditorPath: "/nonexistent/UnrealEditor-Cmd",
			UProjectFile: "/some/Test.uproject",
		},
		Logger: testLogger(),
	}

	_, _, err := h.RunVisualTests(context.Background(), nil, RunTestsInput{})
	if err == nil {
		t.Fatal("expected error for missing editor")
	}
	if !strings.Contains(err.Error(), "UnrealEditor-Cmd not found") {
		t.Errorf("error = %q, want to contain 'UnrealEditor-Cmd not found'", err.Error())
	}
}

func TestRunVisualTests_NoProject(t *testing.T) {
	editorPath, _ := createFakeEditor(t, "", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: "",
		},
		Logger: testLogger(),
	}

	_, _, err := h.RunVisualTests(context.Background(), nil, RunTestsInput{})
	if err == nil {
		t.Fatal("expected error for missing project")
	}
	if !strings.Contains(err.Error(), "no .uproject file found") {
		t.Errorf("error = %q, want to contain 'no .uproject file found'", err.Error())
	}
}

func TestRunVisualTests_Failures(t *testing.T) {
	fixture := readTestdata(t, "failing_tests.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 1)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.RunVisualTests(context.Background(), nil, RunTestsInput{})
	if err != nil {
		t.Fatalf("RunVisualTests returned error: %v", err)
	}
	if out.Success {
		t.Error("Success = true, want false")
	}
	if out.Failed != 2 {
		t.Errorf("Failed = %d, want 2", out.Failed)
	}
}

func TestListTests_Success(t *testing.T) {
	fixture := readTestdata(t, "test_list.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.ListTests(context.Background(), nil, ListTestsInput{})
	if err != nil {
		t.Fatalf("ListTests returned error: %v", err)
	}
	if out.Total != 5 {
		t.Errorf("Total = %d, want 5", out.Total)
	}
	if len(out.Tests) != 5 {
		t.Errorf("len(Tests) = %d, want 5", len(out.Tests))
	}
}

func TestListTests_WithFilter(t *testing.T) {
	fixture := readTestdata(t, "test_list.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.ListTests(context.Background(), nil, ListTestsInput{
		Filter: "Unit",
	})
	if err != nil {
		t.Fatalf("ListTests returned error: %v", err)
	}
	// Should match 3 tests: MathUtils, StringUtils, CalculationTest (all under Unit).
	if out.Total != 3 {
		t.Errorf("Total = %d, want 3 (Unit tests)", out.Total)
	}
	for _, test := range out.Tests {
		if !strings.Contains(strings.ToLower(test), "unit") {
			t.Errorf("test %q does not match filter 'Unit'", test)
		}
	}
}

func TestListTests_FilterCaseInsensitive(t *testing.T) {
	fixture := readTestdata(t, "test_list.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.ListTests(context.Background(), nil, ListTestsInput{
		Filter: "visual",
	})
	if err != nil {
		t.Fatalf("ListTests returned error: %v", err)
	}
	if out.Total != 1 {
		t.Errorf("Total = %d, want 1 (Visual.ScreenshotTest)", out.Total)
	}
}

func TestListTests_FilterNoMatch(t *testing.T) {
	fixture := readTestdata(t, "test_list.log")
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.ListTests(context.Background(), nil, ListTestsInput{
		Filter: "NonExistentTestCategory",
	})
	if err != nil {
		t.Fatalf("ListTests returned error: %v", err)
	}
	if out.Total != 0 {
		t.Errorf("Total = %d, want 0 for non-matching filter", out.Total)
	}
}

func TestListTests_NoEditor(t *testing.T) {
	h := &Handler{
		Config: &config.Config{
			UEEditorPath: "/nonexistent/UnrealEditor-Cmd",
			UProjectFile: "/some/Test.uproject",
		},
		Logger: testLogger(),
	}

	_, _, err := h.ListTests(context.Background(), nil, ListTestsInput{})
	if err == nil {
		t.Fatal("expected error for missing editor")
	}
	if !strings.Contains(err.Error(), "UnrealEditor-Cmd not found") {
		t.Errorf("error = %q, want to contain 'UnrealEditor-Cmd not found'", err.Error())
	}
}

func TestListTests_NoProject(t *testing.T) {
	editorPath, _ := createFakeEditor(t, "", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: "",
		},
		Logger: testLogger(),
	}

	_, _, err := h.ListTests(context.Background(), nil, ListTestsInput{})
	if err == nil {
		t.Fatal("expected error for missing project")
	}
	if !strings.Contains(err.Error(), "no .uproject file found") {
		t.Errorf("error = %q, want to contain 'no .uproject file found'", err.Error())
	}
}
