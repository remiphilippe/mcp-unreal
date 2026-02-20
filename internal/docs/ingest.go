package docs

import (
	"crypto/sha256"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// skipFiles contains filenames that are not documentation content and
// should never be indexed (project meta-files, changelogs, etc.).
var skipFiles = map[string]bool{
	"README.md":       true,
	"CHANGELOG.md":    true,
	"CONTRIBUTING.md": true,
	"LICENSE.md":      true,
}

// IngestDirectory reads all markdown files from a directory tree and
// indexes them into the given Index. Files in skipFiles are excluded.
// The source parameter identifies where the docs came from (e.g.
// "ue5.7", "realtimemesh", "project").
func IngestDirectory(idx *Index, dir, source string, logger *slog.Logger) (int, error) {
	count := 0

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !isMarkdown(path) {
			return nil
		}
		if skipFiles[filepath.Base(path)] {
			return nil
		}

		data, err := os.ReadFile(path) //nolint:gosec // path from filepath.Walk within trusted dir
		if err != nil {
			logger.Warn("skipping unreadable file", "path", path, "error", err)
			return nil
		}

		content := string(data)
		if strings.TrimSpace(content) == "" {
			return nil
		}

		entry := parseMarkdownDoc(path, content, source)
		if err := idx.IndexDoc(entry); err != nil {
			logger.Warn("failed to index doc", "path", path, "error", err)
			return nil
		}

		count++
		logger.Debug("indexed doc", "path", path, "title", entry.Title, "category", entry.Category)
		return nil
	})

	return count, err
}

// IngestFile indexes a single file (e.g. a project's CLAUDE.md).
func IngestFile(idx *Index, path, source string) error {
	cleanPath := filepath.Clean(path)
	data, err := os.ReadFile(cleanPath) //nolint:gosec // caller provides trusted path
	if err != nil {
		return fmt.Errorf("reading %s: %w", cleanPath, err)
	}

	content := string(data)
	if strings.TrimSpace(content) == "" {
		return nil
	}

	entry := parseMarkdownDoc(path, content, source)
	return idx.IndexDoc(entry)
}

// parseMarkdownDoc extracts metadata from a markdown file's content
// and path to build a DocEntry.
func parseMarkdownDoc(path, content, source string) DocEntry {
	title := extractTitle(content)
	if title == "" {
		// Use filename without extension as fallback.
		title = strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	}

	category := inferCategory(path, content)
	classes := extractClassNames(content)

	// Generate a stable ID from the file path.
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(path)))[:16]

	return DocEntry{
		ID:       id,
		Title:    title,
		Category: category,
		Source:   source,
		Content:  content,
		Classes:  classes,
	}
}

// extractTitle pulls the first # heading from markdown content.
func extractTitle(content string) string {
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "# "))
		}
	}
	return ""
}

var categoryKeywords = map[string][]string{
	"actor":        {"actor", "spawn", "pawn", "character", "controller", "gamemode"},
	"blueprint":    {"blueprint", "graph", "node", "blueprint pin", "compile"},
	"material":     {"material", "shader", "texture", "rendering"},
	"animation":    {"animation", "anim", "skeleton", "montage", "state machine"},
	"input":        {"input", "enhanced input", "action mapping", "input action"},
	"realtimemesh": {"realtimemesh", "proceduralmesh", "mesh generation", "section group"},
	"gameplay":     {"gameplay", "game mode", "game state", "player state", "ability"},
	"rendering":    {"rendering", "viewport", "camera", "light", "post process"},
	"networking":   {"networking", "replication", "net driver", "net multicast"},
}

// inferCategory guesses the doc category from the file path and content.
func inferCategory(path, content string) string {
	// First check path components.
	lowerPath := strings.ToLower(path)
	for cat, keywords := range categoryKeywords {
		for _, kw := range keywords {
			if strings.Contains(lowerPath, kw) {
				return cat
			}
		}
	}

	// Then check content (first 500 chars for performance).
	sample := strings.ToLower(content)
	if len(sample) > 500 {
		sample = sample[:500]
	}
	for cat, keywords := range categoryKeywords {
		for _, kw := range keywords {
			if strings.Contains(sample, kw) {
				return cat
			}
		}
	}

	return "general"
}

// ueClassRe matches UE class name patterns: AActor, UObject, FStruct, etc.
var ueClassRe = regexp.MustCompile(`\b([AUFE][A-Z][a-zA-Z0-9]{2,})\b`)

// extractClassNames finds UE class name references in the content.
func extractClassNames(content string) []string {
	matches := ueClassRe.FindAllString(content, -1)
	seen := make(map[string]bool)
	var classes []string
	for _, m := range matches {
		if !seen[m] && isLikelyClassName(m) {
			seen[m] = true
			classes = append(classes, m)
		}
	}
	return classes
}

// isLikelyClassName filters out common false positives from the class regex.
func isLikelyClassName(name string) bool {
	// Too short to be a class name.
	if len(name) < 4 {
		return false
	}
	// Common false positives.
	falsePositives := map[string]bool{
		"FNAME": true, "FTEXT": true, "FSTRING": true,
		"UINT8": true, "UINT16": true, "UINT32": true, "UINT64": true,
		"FALSE": true, "FLOAT": true, "UFUNCTION": true, "UPROPERTY": true,
		"UCLASS": true, "USTRUCT": true, "UENUM": true, "UINTERFACE": true,
		"ENGINE": true, "EDITOR": true, "ENSURE": true,
	}
	return !falsePositives[strings.ToUpper(name)]
}

func isMarkdown(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".md" || ext == ".markdown"
}
