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

func TestGetTestLog_ReadsFile(t *testing.T) {
	absPath, err := filepath.Abs("testdata/passing_tests.log")
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath:  absPath,
		MaxLines: 100,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	if out.TotalLines == 0 {
		t.Error("expected non-zero total lines")
	}
	if out.Returned == 0 {
		t.Error("expected non-zero returned lines")
	}
	if out.LogPath != absPath {
		t.Errorf("LogPath = %q, want %q", out.LogPath, absPath)
	}
}

func TestGetTestLog_Filter(t *testing.T) {
	absPath, err := filepath.Abs("testdata/failing_tests.log")
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath: absPath,
		Filter:  "error",
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	if out.Returned == 0 {
		t.Error("expected filtered results with 'error' lines")
	}

	for _, line := range strings.Split(out.Content, "\n") {
		if line == "" {
			continue
		}
		if !strings.Contains(strings.ToLower(line), "error") {
			t.Errorf("filtered line does not contain 'error': %q", line)
		}
	}
}

func TestGetTestLog_MaxLines(t *testing.T) {
	absPath, err := filepath.Abs("testdata/passing_tests.log")
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath:  absPath,
		MaxLines: 2,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	if out.Returned > 2 {
		t.Errorf("returned %d lines, want <= 2", out.Returned)
	}
}

func TestGetTestLog_MaxLinesCapped(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "big.log")

	var lines []string
	for i := 0; i < 600; i++ {
		lines = append(lines, "log line")
	}
	if err := os.WriteFile(logFile, []byte(strings.Join(lines, "\n")), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath:  logFile,
		MaxLines: 999,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	if out.Returned > 500 {
		t.Errorf("returned %d lines, want <= 500 (capped)", out.Returned)
	}
}

func TestGetTestLog_DefaultMaxLines(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "big.log")

	var lines []string
	for i := 0; i < 300; i++ {
		lines = append(lines, "log line")
	}
	if err := os.WriteFile(logFile, []byte(strings.Join(lines, "\n")), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath:  logFile,
		MaxLines: 0,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	if out.Returned > 200 {
		t.Errorf("returned %d lines, want <= 200 (default)", out.Returned)
	}
}

func TestGetTestLog_Offset(t *testing.T) {
	absPath, err := filepath.Abs("testdata/passing_tests.log")
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, full, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath:  absPath,
		MaxLines: 500,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath:  absPath,
		Offset:   3,
		MaxLines: 500,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	expectedReturned := full.Returned - 3
	if expectedReturned < 0 {
		expectedReturned = 0
	}
	if out.Returned != expectedReturned {
		t.Errorf("returned %d lines with offset 3, want %d", out.Returned, expectedReturned)
	}
}

func TestGetTestLog_OffsetBeyondEnd(t *testing.T) {
	absPath, err := filepath.Abs("testdata/passing_tests.log")
	if err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath: absPath,
		Offset:  99999,
	})
	if err != nil {
		t.Fatalf("GetTestLog returned error: %v", err)
	}

	if out.Returned != 0 {
		t.Errorf("returned %d lines for offset beyond end, want 0", out.Returned)
	}
}

func TestGetTestLog_PathTraversal(t *testing.T) {
	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, _, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath: "../../../etc/passwd",
	})
	if err == nil {
		t.Error("expected error for path traversal attempt")
	}
	if !strings.Contains(err.Error(), "path traversal not allowed") {
		t.Errorf("error = %q, want to contain 'path traversal not allowed'", err.Error())
	}
}

func TestGetTestLog_PathTraversalRelative(t *testing.T) {
	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, _, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath: "logs/../../etc/passwd",
	})
	if err == nil {
		t.Error("expected error for relative path traversal")
	}
	if !strings.Contains(err.Error(), "path traversal not allowed") {
		t.Errorf("error = %q, want to contain 'path traversal not allowed'", err.Error())
	}
}

func TestGetTestLog_MissingFile(t *testing.T) {
	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, _, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{
		LogPath: "/nonexistent/path/test.log",
	})
	if err == nil {
		t.Error("expected error for missing file")
	}
}

//nolint:gosec // test: temp directories
func TestGetTestLog_AutoDetect(t *testing.T) {
	dir := t.TempDir()
	logsDir := filepath.Join(dir, "Saved", "Logs")
	if err := os.MkdirAll(logsDir, 0750); err != nil {
		t.Fatal(err)
	}

	logContent := "Line 1\nLine 2\nLine 3\n"
	logFile := filepath.Join(logsDir, "latest.log")
	if err := os.WriteFile(logFile, []byte(logContent), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{ProjectRoot: dir},
		Logger: testLogger(),
	}

	_, out, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{})
	if err != nil {
		t.Fatalf("GetTestLog auto-detect returned error: %v", err)
	}
	if out.LogPath == "" {
		t.Error("expected auto-detected LogPath")
	}
	if out.Returned == 0 {
		t.Error("expected non-zero returned lines from auto-detected log")
	}
}

func TestGetTestLog_AutoDetect_NoLogs(t *testing.T) {
	h := &Handler{
		Config: &config.Config{ProjectRoot: t.TempDir()},
		Logger: testLogger(),
	}

	_, _, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{})
	if err == nil {
		t.Error("expected error when no log files found")
	}
	if !strings.Contains(err.Error(), "no UE log file found") {
		t.Errorf("error = %q, want to contain 'no UE log file found'", err.Error())
	}
}

func TestGetTestLog_AutoDetect_NoProjectRoot(t *testing.T) {
	h := &Handler{
		Config: &config.Config{ProjectRoot: ""},
		Logger: testLogger(),
	}

	_, _, err := h.GetTestLog(context.Background(), nil, GetTestLogInput{})
	if err == nil {
		t.Error("expected error when no project root")
	}
}

func TestFindLatestLog(t *testing.T) {
	dir := t.TempDir()
	logsDir := filepath.Join(dir, "Saved", "Logs")
	if err := os.MkdirAll(logsDir, 0700); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(logsDir, "old.log"), []byte("old"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(logsDir, "new.log"), []byte("new"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{ProjectRoot: dir},
		Logger: testLogger(),
	}

	got := h.findLatestLog()
	if got == "" {
		t.Fatal("findLatestLog returned empty")
	}

	if filepath.Base(got) != "new.log" && filepath.Base(got) != "old.log" {
		t.Errorf("unexpected log file: %s", got)
	}
}

func TestFindLatestLog_EmptyProjectRoot(t *testing.T) {
	h := &Handler{
		Config: &config.Config{ProjectRoot: ""},
		Logger: testLogger(),
	}

	got := h.findLatestLog()
	if got != "" {
		t.Errorf("findLatestLog with empty ProjectRoot = %q, want empty", got)
	}
}

func TestFindLatestLog_NoLogsDir(t *testing.T) {
	h := &Handler{
		Config: &config.Config{ProjectRoot: t.TempDir()},
		Logger: testLogger(),
	}

	got := h.findLatestLog()
	if got != "" {
		t.Errorf("findLatestLog with no Saved/Logs/ = %q, want empty", got)
	}
}

//nolint:gosec // test: temp directories
func TestFindLatestLog_IgnoresDirectories(t *testing.T) {
	dir := t.TempDir()
	logsDir := filepath.Join(dir, "Saved", "Logs")
	if err := os.MkdirAll(logsDir, 0750); err != nil {
		t.Fatal(err)
	}

	// Create a subdirectory named "foo.log" -- should be ignored.
	if err := os.MkdirAll(filepath.Join(logsDir, "foo.log"), 0750); err != nil {
		t.Fatal(err)
	}

	// Create a non-log file -- should also be ignored.
	if err := os.WriteFile(filepath.Join(logsDir, "notes.txt"), []byte("x"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{ProjectRoot: dir},
		Logger: testLogger(),
	}

	got := h.findLatestLog()
	if got != "" {
		t.Errorf("findLatestLog should ignore dirs and non-log files, got %q", got)
	}
}

//nolint:gosec // test: temp directories
func TestFindLatestLog_PicksNewest(t *testing.T) {
	dir := t.TempDir()
	logsDir := filepath.Join(dir, "Saved", "Logs")
	if err := os.MkdirAll(logsDir, 0750); err != nil {
		t.Fatal(err)
	}

	oldFile := filepath.Join(logsDir, "old.log")
	if err := os.WriteFile(oldFile, []byte("old content"), 0600); err != nil {
		t.Fatal(err)
	}

	newFile := filepath.Join(logsDir, "new.log")
	if err := os.WriteFile(newFile, []byte("new content"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{ProjectRoot: dir},
		Logger: testLogger(),
	}

	got := h.findLatestLog()
	if got == "" {
		t.Fatal("findLatestLog returned empty")
	}
	if filepath.Base(got) != "new.log" {
		t.Errorf("findLatestLog = %q, want new.log (latest)", filepath.Base(got))
	}
}
