// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package headless

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/remiphilippe/mcp-unreal/internal/config"
)

// testLogger returns a debug-level logger writing to stderr (safe for tests).
func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
}

// createFakeEditorWithFixture creates a fake UnrealEditor-Cmd script in a temp dir
// that outputs the given fixture content and exits with the given code.
//
//nolint:gosec // test helper: scripts need 0750 to be executable, temp files are safe
func createFakeEditorWithFixture(t *testing.T, fixture string, exitCode int) (editorPath, projectFile string) {
	t.Helper()
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}

	// Write fixture to a temp file so the script can cat it.
	fixtureFile := filepath.Join(dir, "output.txt")
	if err := os.WriteFile(fixtureFile, []byte(fixture), 0600); err != nil {
		t.Fatal(err)
	}

	editorPath = filepath.Join(binDir, "UnrealEditor-Cmd")
	script := fmt.Sprintf("#!/bin/sh\ncat '%s'\nexit %d\n", fixtureFile, exitCode)
	if err := os.WriteFile(editorPath, []byte(script), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile = filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	return editorPath, projectFile
}

// createFakeEditor creates a minimal fake editor that outputs a short string.
func createFakeEditor(t *testing.T, stdout string, exitCode int) (editorPath, projectFile string) {
	t.Helper()
	return createFakeEditorWithFixture(t, stdout, exitCode)
}

func TestParseBuildErrors(t *testing.T) {
	tests := []struct {
		name    string
		output  string
		wantN   int
		wantMsg string
	}{
		{
			name:    "MSVC style error",
			output:  `D:\Project\Source\MyClass.cpp(42): error C2065: 'foo': undeclared identifier`,
			wantN:   1,
			wantMsg: `D:\Project\Source\MyClass.cpp(42): error C2065: 'foo': undeclared identifier`,
		},
		{
			name:    "clang style error",
			output:  `/Users/dev/Source/MyClass.cpp(42): error : use of undeclared identifier 'foo'`,
			wantN:   1,
			wantMsg: `/Users/dev/Source/MyClass.cpp(42): error : use of undeclared identifier 'foo'`,
		},
		{
			name:   "no errors",
			output: "Build succeeded.\nTotal time: 30.5s",
			wantN:  0,
		},
		{
			name: "multiple errors deduped",
			output: `Source/A.cpp(10): error C2065: 'x': undeclared
Source/A.cpp(10): error C2065: 'x': undeclared
Source/B.cpp(20): error C2065: 'y': undeclared`,
			wantN: 2,
		},
		{
			name:   "empty output",
			output: "",
			wantN:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := parseBuildErrors(tt.output)
			if len(errors) != tt.wantN {
				t.Errorf("got %d errors, want %d: %v", len(errors), tt.wantN, errors)
			}
			if tt.wantMsg != "" && len(errors) > 0 && errors[0] != tt.wantMsg {
				t.Errorf("error[0] = %q, want %q", errors[0], tt.wantMsg)
			}
		})
	}
}

func TestParseBuildWarnings(t *testing.T) {
	output := `Source/MyClass.cpp(10): warning C4267: conversion from 'size_t' to 'int'
Source/MyClass.cpp(20): warning C4996: 'sprintf' deprecated
Build succeeded with 2 warnings.`

	warnings := parseBuildWarnings(output)
	if len(warnings) != 2 {
		t.Errorf("got %d warnings, want 2: %v", len(warnings), warnings)
	}
}

func TestLastNLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		n     int
		want  string
	}{
		{"fewer lines than n", "a\nb\nc", 5, "a\nb\nc"},
		{"exact n", "a\nb\nc", 3, "a\nb\nc"},
		{"more lines than n", "a\nb\nc\nd\ne", 2, "d\ne"},
		{"empty", "", 5, ""},
		{"single line", "hello", 1, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lastNLines(tt.input, tt.n)
			if got != tt.want {
				t.Errorf("lastNLines(%q, %d) = %q, want %q", tt.input, tt.n, got, tt.want)
			}
		})
	}
}

func TestDedup(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b", ""}
	got := dedup(input)
	want := []string{"a", "b", "c"}

	if len(got) != len(want) {
		t.Fatalf("dedup returned %d items, want %d: %v", len(got), len(want), got)
	}
	for i, v := range want {
		if got[i] != v {
			t.Errorf("dedup[%d] = %q, want %q", i, got[i], v)
		}
	}
}

func TestDedup_Nil(t *testing.T) {
	got := dedup(nil)
	if got != nil {
		t.Errorf("dedup(nil) = %v, want nil", got)
	}
}

func TestDefaultPlatform(t *testing.T) {
	p := defaultPlatform()
	switch runtime.GOOS {
	case "darwin":
		if p != "Mac" {
			t.Errorf("defaultPlatform() on darwin = %q, want Mac", p)
		}
	case "windows":
		if p != "Win64" {
			t.Errorf("defaultPlatform() on windows = %q, want Win64", p)
		}
	case "linux":
		if p != "Linux" {
			t.Errorf("defaultPlatform() on linux = %q, want Linux", p)
		}
	default:
		if p != "Mac" && p != "Win64" && p != "Linux" {
			t.Errorf("defaultPlatform() = %q, want one of Mac/Win64/Linux", p)
		}
	}
}

func TestFileExists_True(t *testing.T) {
	f := filepath.Join(t.TempDir(), "exists.txt")
	if err := os.WriteFile(f, []byte("x"), 0600); err != nil {
		t.Fatal(err)
	}
	if !fileExists(f) {
		t.Errorf("fileExists(%q) = false, want true", f)
	}
}

func TestFileExists_False(t *testing.T) {
	if fileExists(filepath.Join(t.TempDir(), "nope.txt")) {
		t.Error("fileExists returned true for nonexistent file")
	}
}

func TestRunCommand_Success(t *testing.T) {
	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	stdout, stderr, exitCode, err := h.runCommand(context.Background(), "sh", []string{"-c", "echo hello"}, 5*1e9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 0 {
		t.Errorf("exit code = %d, want 0", exitCode)
	}
	if !strings.Contains(stdout, "hello") {
		t.Errorf("stdout = %q, want to contain 'hello'", stdout)
	}
	_ = stderr
}

func TestRunCommand_ExitError(t *testing.T) {
	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, _, exitCode, err := h.runCommand(context.Background(), "sh", []string{"-c", "exit 42"}, 5*1e9)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if exitCode != 42 {
		t.Errorf("exit code = %d, want 42", exitCode)
	}
}

func TestRunCommand_NotFound(t *testing.T) {
	h := &Handler{
		Config: &config.Config{},
		Logger: testLogger(),
	}

	_, _, exitCode, err := h.runCommand(context.Background(), "/nonexistent/binary/xyz", nil, 5*1e9)
	if err == nil {
		t.Fatal("expected error for nonexistent binary")
	}
	if exitCode != -1 {
		t.Errorf("exit code = %d, want -1", exitCode)
	}
}

func TestBuildProject_Success(t *testing.T) {
	editorPath, projectFile := createFakeEditor(t, "Build succeeded.\n", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.BuildProject(context.Background(), nil, BuildInput{})
	if err != nil {
		t.Fatalf("BuildProject returned error: %v", err)
	}
	if !out.Success {
		t.Errorf("Success = false, want true")
	}
	if out.ExitCode != 0 {
		t.Errorf("ExitCode = %d, want 0", out.ExitCode)
	}
	if out.ErrorCount != 0 {
		t.Errorf("ErrorCount = %d, want 0", out.ErrorCount)
	}
}

func TestBuildProject_WithErrors(t *testing.T) {
	fixture := `Source/MyClass.cpp(42): error C2065: 'foo': undeclared identifier
Source/MyClass.cpp(99): error C2065: 'bar': undeclared identifier
Source/MyClass.cpp(10): warning C4996: 'sprintf' deprecated
Build failed.
`
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 1)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.BuildProject(context.Background(), nil, BuildInput{})
	if err != nil {
		t.Fatalf("BuildProject returned error: %v", err)
	}
	if out.Success {
		t.Error("Success = true, want false")
	}
	if out.ExitCode != 1 {
		t.Errorf("ExitCode = %d, want 1", out.ExitCode)
	}
	if out.ErrorCount != 2 {
		t.Errorf("ErrorCount = %d, want 2", out.ErrorCount)
	}
	if out.WarningCount != 1 {
		t.Errorf("WarningCount = %d, want 1", out.WarningCount)
	}
}

func TestBuildProject_WithManyWarnings(t *testing.T) {
	var lines []string
	for i := 0; i < 25; i++ {
		lines = append(lines, fmt.Sprintf("Source/File.cpp(%d): warning C4996: warning %d", i+1, i+1))
	}
	fixture := strings.Join(lines, "\n") + "\n"
	editorPath, projectFile := createFakeEditorWithFixture(t, fixture, 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.BuildProject(context.Background(), nil, BuildInput{})
	if err != nil {
		t.Fatalf("BuildProject returned error: %v", err)
	}
	if out.WarningCount != 25 {
		t.Errorf("WarningCount = %d, want 25", out.WarningCount)
	}
	if len(out.Warnings) != 20 {
		t.Errorf("len(Warnings) = %d, want 20 (capped)", len(out.Warnings))
	}
}

func TestBuildProject_NoEditor(t *testing.T) {
	h := &Handler{
		Config: &config.Config{
			UEEditorPath: "/nonexistent/UnrealEditor-Cmd",
			UProjectFile: "/some/Test.uproject",
		},
		Logger: testLogger(),
	}

	_, _, err := h.BuildProject(context.Background(), nil, BuildInput{})
	if err == nil {
		t.Fatal("expected error for missing editor")
	}
	if !strings.Contains(err.Error(), "UnrealEditor-Cmd not found") {
		t.Errorf("error = %q, want to contain 'UnrealEditor-Cmd not found'", err.Error())
	}
}

func TestBuildProject_NoProject(t *testing.T) {
	editorPath, _ := createFakeEditor(t, "", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: "",
		},
		Logger: testLogger(),
	}

	_, _, err := h.BuildProject(context.Background(), nil, BuildInput{})
	if err == nil {
		t.Fatal("expected error for missing project")
	}
	if !strings.Contains(err.Error(), "no .uproject file found") {
		t.Errorf("error = %q, want to contain 'no .uproject file found'", err.Error())
	}
}

func TestBuildProject_CustomInputs(t *testing.T) {
	editorPath, projectFile := createFakeEditor(t, "Build succeeded.\n", 0)

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.BuildProject(context.Background(), nil, BuildInput{
		Target:   "MyGameEditor",
		Config:   "Shipping",
		Platform: "Win64",
		Clean:    true,
	})
	if err != nil {
		t.Fatalf("BuildProject returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestBuildProject_DefaultTarget(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	fixtureFile := filepath.Join(dir, "output.txt")
	if err := os.WriteFile(fixtureFile, []byte("Build succeeded.\n"), 0600); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	script := fmt.Sprintf("#!/bin/sh\ncat '%s'\nexit 0\n", fixtureFile)
	if err := os.WriteFile(editorPath, []byte(script), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "MyAwesomeGame.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.BuildProject(context.Background(), nil, BuildInput{})
	if err != nil {
		t.Fatalf("BuildProject returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestCookProject_Success(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	batchDir := filepath.Join(dir, "Engine", "Build", "BatchFiles")
	if err := os.MkdirAll(batchDir, 0750); err != nil {
		t.Fatal(err)
	}

	fixtureFile := filepath.Join(dir, "cook_output.txt")
	if err := os.WriteFile(fixtureFile, []byte("Cook complete.\n"), 0600); err != nil {
		t.Fatal(err)
	}

	runUAT := filepath.Join(batchDir, "RunUAT.sh")
	script := fmt.Sprintf("#!/bin/sh\ncat '%s'\nexit 0\n", fixtureFile)
	if err := os.WriteFile(runUAT, []byte(script), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.CookProject(context.Background(), nil, CookInput{})
	if err != nil {
		t.Fatalf("CookProject returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
	if out.ExitCode != 0 {
		t.Errorf("ExitCode = %d, want 0", out.ExitCode)
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestCookProject_Iterative(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	batchDir := filepath.Join(dir, "Engine", "Build", "BatchFiles")
	if err := os.MkdirAll(batchDir, 0750); err != nil {
		t.Fatal(err)
	}
	runUAT := filepath.Join(batchDir, "RunUAT.sh")
	if err := os.WriteFile(runUAT, []byte("#!/bin/sh\necho iterative cook\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.CookProject(context.Background(), nil, CookInput{
		Platform:  "Mac",
		Config:    "Shipping",
		Iterative: true,
	})
	if err != nil {
		t.Fatalf("CookProject returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestCookProject_Failure(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	batchDir := filepath.Join(dir, "Engine", "Build", "BatchFiles")
	if err := os.MkdirAll(batchDir, 0750); err != nil {
		t.Fatal(err)
	}
	runUAT := filepath.Join(batchDir, "RunUAT.sh")
	if err := os.WriteFile(runUAT, []byte("#!/bin/sh\necho cook failed\nexit 1\n"), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.CookProject(context.Background(), nil, CookInput{})
	if err != nil {
		t.Fatalf("CookProject returned error: %v", err)
	}
	if out.Success {
		t.Error("Success = true, want false for exit code 1")
	}
	if out.ExitCode != 1 {
		t.Errorf("ExitCode = %d, want 1", out.ExitCode)
	}
}

func TestCookProject_NoProject(t *testing.T) {
	h := &Handler{
		Config: &config.Config{
			UEEditorPath: "/some/path/UnrealEditor-Cmd",
			UProjectFile: "",
		},
		Logger: testLogger(),
	}

	_, _, err := h.CookProject(context.Background(), nil, CookInput{})
	if err == nil {
		t.Fatal("expected error for missing project")
	}
	if !strings.Contains(err.Error(), "no .uproject file found") {
		t.Errorf("error = %q, want to contain 'no .uproject file found'", err.Error())
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestCookProject_NoRunUAT(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, _, err := h.CookProject(context.Background(), nil, CookInput{})
	if err == nil {
		t.Fatal("expected error for missing RunUAT")
	}
	if !strings.Contains(err.Error(), "RunUAT script not found") {
		t.Errorf("error = %q, want to contain 'RunUAT script not found'", err.Error())
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestGenerateProjectFiles_Success(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	genScriptDir := filepath.Join(dir, "Engine", "Build", "BatchFiles", "Mac")
	if err := os.MkdirAll(genScriptDir, 0750); err != nil {
		t.Fatal(err)
	}
	genScript := filepath.Join(genScriptDir, "GenerateProjectFiles.sh")
	if err := os.WriteFile(genScript, []byte("#!/bin/sh\necho Generated.\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.GenerateProjectFiles(context.Background(), nil, GenerateProjectFilesInput{})
	if err != nil {
		t.Fatalf("GenerateProjectFiles returned error: %v", err)
	}
	if !out.Success {
		t.Error("Success = false, want true")
	}
	if out.ExitCode != 0 {
		t.Errorf("ExitCode = %d, want 0", out.ExitCode)
	}
}

func TestGenerateProjectFiles_NoProject(t *testing.T) {
	h := &Handler{
		Config: &config.Config{
			UEEditorPath: "/some/path/UnrealEditor-Cmd",
			UProjectFile: "",
		},
		Logger: testLogger(),
	}

	_, _, err := h.GenerateProjectFiles(context.Background(), nil, GenerateProjectFilesInput{})
	if err == nil {
		t.Fatal("expected error for missing project")
	}
	if !strings.Contains(err.Error(), "no .uproject file found") {
		t.Errorf("error = %q, want to contain 'no .uproject file found'", err.Error())
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestGenerateProjectFiles_NoScript(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, _, err := h.GenerateProjectFiles(context.Background(), nil, GenerateProjectFilesInput{})
	if err == nil {
		t.Fatal("expected error for missing GenerateProjectFiles script")
	}
	if !strings.Contains(err.Error(), "GenerateProjectFiles script not found") {
		t.Errorf("error = %q, want to contain 'GenerateProjectFiles script not found'", err.Error())
	}
}

//nolint:gosec // test: temp files with executable scripts
func TestGenerateProjectFiles_Failure(t *testing.T) {
	dir := t.TempDir()
	binDir := filepath.Join(dir, "Engine", "Binaries", "Mac")
	if err := os.MkdirAll(binDir, 0750); err != nil {
		t.Fatal(err)
	}
	editorPath := filepath.Join(binDir, "UnrealEditor-Cmd")
	if err := os.WriteFile(editorPath, []byte("#!/bin/sh\nexit 0\n"), 0750); err != nil {
		t.Fatal(err)
	}

	genScriptDir := filepath.Join(dir, "Engine", "Build", "BatchFiles", "Mac")
	if err := os.MkdirAll(genScriptDir, 0750); err != nil {
		t.Fatal(err)
	}
	genScript := filepath.Join(genScriptDir, "GenerateProjectFiles.sh")
	if err := os.WriteFile(genScript, []byte("#!/bin/sh\necho failed\nexit 1\n"), 0750); err != nil {
		t.Fatal(err)
	}

	projectFile := filepath.Join(dir, "Test.uproject")
	if err := os.WriteFile(projectFile, []byte("{}"), 0600); err != nil {
		t.Fatal(err)
	}

	h := &Handler{
		Config: &config.Config{
			UEEditorPath: editorPath,
			UProjectFile: projectFile,
		},
		Logger: testLogger(),
	}

	_, out, err := h.GenerateProjectFiles(context.Background(), nil, GenerateProjectFilesInput{})
	if err != nil {
		t.Fatalf("GenerateProjectFiles returned error: %v", err)
	}
	if out.Success {
		t.Error("Success = true, want false")
	}
	if out.ExitCode != 1 {
		t.Errorf("ExitCode = %d, want 1", out.ExitCode)
	}
}

//nolint:gosec // test: temp files
func TestFindRunUATScript_Found(t *testing.T) {
	dir := t.TempDir()
	batchDir := filepath.Join(dir, "Engine", "Build", "BatchFiles")
	if err := os.MkdirAll(batchDir, 0750); err != nil {
		t.Fatal(err)
	}
	runUAT := filepath.Join(batchDir, "RunUAT.sh")
	if err := os.WriteFile(runUAT, []byte("#!/bin/sh\n"), 0750); err != nil {
		t.Fatal(err)
	}

	editorPath := filepath.Join(dir, "Engine", "Binaries", "Mac", "UnrealEditor-Cmd")

	got := findRunUATScript(editorPath)
	if got == "" {
		t.Fatal("findRunUATScript returned empty, want the script path")
	}
	if got != runUAT {
		t.Errorf("findRunUATScript = %q, want %q", got, runUAT)
	}
}

func TestFindRunUATScript_NotFound(t *testing.T) {
	dir := t.TempDir()
	editorPath := filepath.Join(dir, "Engine", "Binaries", "Mac", "UnrealEditor-Cmd")

	got := findRunUATScript(editorPath)
	if got != "" {
		t.Errorf("findRunUATScript = %q, want empty", got)
	}
}

//nolint:gosec // test: temp files
func TestFindGenerateProjectFilesScript_Found(t *testing.T) {
	dir := t.TempDir()
	macDir := filepath.Join(dir, "Engine", "Build", "BatchFiles", "Mac")
	if err := os.MkdirAll(macDir, 0750); err != nil {
		t.Fatal(err)
	}
	genScript := filepath.Join(macDir, "GenerateProjectFiles.sh")
	if err := os.WriteFile(genScript, []byte("#!/bin/sh\n"), 0750); err != nil {
		t.Fatal(err)
	}

	editorPath := filepath.Join(dir, "Engine", "Binaries", "Mac", "UnrealEditor-Cmd")

	got := findGenerateProjectFilesScript(editorPath)
	if got == "" {
		t.Fatal("findGenerateProjectFilesScript returned empty, want the script path")
	}
	if got != genScript {
		t.Errorf("findGenerateProjectFilesScript = %q, want %q", got, genScript)
	}
}

func TestFindGenerateProjectFilesScript_NotFound(t *testing.T) {
	dir := t.TempDir()
	editorPath := filepath.Join(dir, "Engine", "Binaries", "Mac", "UnrealEditor-Cmd")

	got := findGenerateProjectFilesScript(editorPath)
	if got != "" {
		t.Errorf("findGenerateProjectFilesScript = %q, want empty", got)
	}
}
