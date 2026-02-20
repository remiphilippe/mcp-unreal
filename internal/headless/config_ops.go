// config_ops.go implements the config_ops MCP tool for reading and writing
// UE .ini config files (DefaultEngine.ini, DefaultGame.ini, etc.).
//
// This is a headless tool — it reads/writes files directly from the project's
// Config/ directory and does not require the editor to be running.
// See issue #39.
package headless

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- config_ops ---

// ConfigOpsInput defines parameters for the config_ops tool.
type ConfigOpsInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: get, set, delete, list, list_sections"`
	File      string `json:"file" jsonschema:"required,INI file name without extension (e.g. DefaultEngine, DefaultGame, DefaultInput, DefaultGameUserSettings)"`
	Section   string `json:"section,omitempty" jsonschema:"INI section name (e.g. /Script/Engine.RendererSettings). Required for get, set, delete, list."`
	Key       string `json:"key,omitempty" jsonschema:"Config key name. Required for get, set, delete."`
	Value     string `json:"value,omitempty" jsonschema:"Value to set. Required for set operation."`
}

// ConfigOpsOutput is returned by the config_ops tool.
type ConfigOpsOutput struct {
	Success  bool              `json:"success" jsonschema:"whether the operation succeeded"`
	File     string            `json:"file" jsonschema:"INI file that was operated on"`
	Section  string            `json:"section,omitempty" jsonschema:"section that was queried or modified"`
	Key      string            `json:"key,omitempty" jsonschema:"key that was queried or modified"`
	Value    string            `json:"value,omitempty" jsonschema:"value retrieved or set"`
	Values   map[string]string `json:"values,omitempty" jsonschema:"all key-value pairs in the section (for list operation)"`
	Sections []string          `json:"sections,omitempty" jsonschema:"all section names in the file (for list_sections operation)"`
}

// RegisterConfig adds the config_ops tool to the MCP server.
func (h *Handler) RegisterConfig(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "config_ops",
		Description: "Read and write UE project .ini config files (DefaultEngine.ini, DefaultGame.ini, etc.). " +
			"Operations: get (read a key), set (write a key), delete (remove a key), " +
			"list (all keys in a section), list_sections (all sections in a file). " +
			"Does not require the editor to be running. " +
			"File paths are resolved relative to the project Config/ directory.",
	}, h.ConfigOps)
}

// ConfigOps implements the config_ops tool.
func (h *Handler) ConfigOps(ctx context.Context, req *mcp.CallToolRequest, input ConfigOpsInput) (*mcp.CallToolResult, ConfigOpsOutput, error) {
	if input.Operation == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("operation is required")
	}
	if input.File == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("file is required")
	}

	// Resolve the INI file path within the project Config/ directory.
	iniPath, err := h.resolveINIPath(input.File)
	if err != nil {
		return nil, ConfigOpsOutput{}, err
	}

	switch input.Operation {
	case "get":
		return h.configGet(iniPath, input)
	case "set":
		return h.configSet(iniPath, input)
	case "delete":
		return h.configDelete(iniPath, input)
	case "list":
		return h.configList(iniPath, input)
	case "list_sections":
		return h.configListSections(iniPath, input)
	default:
		return nil, ConfigOpsOutput{}, fmt.Errorf("unknown operation %q — use get, set, delete, list, or list_sections", input.Operation)
	}
}

// resolveINIPath safely resolves an INI file name to a full path within Config/.
func (h *Handler) resolveINIPath(file string) (string, error) {
	if h.Config.ProjectRoot == "" {
		return "", fmt.Errorf("no UE project root detected — set MCP_UNREAL_PROJECT env var")
	}

	// Sanitize: strip any extension the user might have added.
	file = strings.TrimSuffix(file, ".ini")

	// Reject path traversal attempts.
	if strings.Contains(file, "..") || strings.Contains(file, "/") || strings.Contains(file, "\\") {
		return "", fmt.Errorf("invalid file name %q — use just the file name without path separators (e.g. DefaultEngine)", file)
	}

	configDir := filepath.Join(h.Config.ProjectRoot, "Config")
	iniPath := filepath.Join(configDir, file+".ini")

	// Verify the resolved path is still within Config/.
	absPath, err := filepath.Abs(iniPath)
	if err != nil {
		return "", fmt.Errorf("resolving path: %w", err)
	}
	absConfigDir, err := filepath.Abs(configDir)
	if err != nil {
		return "", fmt.Errorf("resolving config dir: %w", err)
	}
	if !strings.HasPrefix(absPath, absConfigDir) {
		return "", fmt.Errorf("path traversal blocked: %s is outside Config/", file)
	}

	return iniPath, nil
}

func (h *Handler) configGet(iniPath string, input ConfigOpsInput) (*mcp.CallToolResult, ConfigOpsOutput, error) {
	if input.Section == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("section is required for get operation")
	}
	if input.Key == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("key is required for get operation")
	}

	sections, err := parseINI(iniPath)
	if err != nil {
		return nil, ConfigOpsOutput{}, err
	}

	keys, ok := sections[input.Section]
	if !ok {
		return nil, ConfigOpsOutput{}, fmt.Errorf("section [%s] not found in %s", input.Section, filepath.Base(iniPath))
	}

	value, ok := keys[input.Key]
	if !ok {
		return nil, ConfigOpsOutput{}, fmt.Errorf("key %q not found in section [%s]", input.Key, input.Section)
	}

	return nil, ConfigOpsOutput{
		Success: true,
		File:    filepath.Base(iniPath),
		Section: input.Section,
		Key:     input.Key,
		Value:   value,
	}, nil
}

func (h *Handler) configSet(iniPath string, input ConfigOpsInput) (*mcp.CallToolResult, ConfigOpsOutput, error) {
	if input.Section == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("section is required for set operation")
	}
	if input.Key == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("key is required for set operation")
	}

	if err := setINIValue(iniPath, input.Section, input.Key, input.Value); err != nil {
		return nil, ConfigOpsOutput{}, err
	}

	return nil, ConfigOpsOutput{
		Success: true,
		File:    filepath.Base(iniPath),
		Section: input.Section,
		Key:     input.Key,
		Value:   input.Value,
	}, nil
}

func (h *Handler) configDelete(iniPath string, input ConfigOpsInput) (*mcp.CallToolResult, ConfigOpsOutput, error) {
	if input.Section == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("section is required for delete operation")
	}
	if input.Key == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("key is required for delete operation")
	}

	if err := deleteINIValue(iniPath, input.Section, input.Key); err != nil {
		return nil, ConfigOpsOutput{}, err
	}

	return nil, ConfigOpsOutput{
		Success: true,
		File:    filepath.Base(iniPath),
		Section: input.Section,
		Key:     input.Key,
	}, nil
}

func (h *Handler) configList(iniPath string, input ConfigOpsInput) (*mcp.CallToolResult, ConfigOpsOutput, error) {
	if input.Section == "" {
		return nil, ConfigOpsOutput{}, fmt.Errorf("section is required for list operation")
	}

	sections, err := parseINI(iniPath)
	if err != nil {
		return nil, ConfigOpsOutput{}, err
	}

	keys, ok := sections[input.Section]
	if !ok {
		return nil, ConfigOpsOutput{}, fmt.Errorf("section [%s] not found in %s", input.Section, filepath.Base(iniPath))
	}

	return nil, ConfigOpsOutput{
		Success: true,
		File:    filepath.Base(iniPath),
		Section: input.Section,
		Values:  keys,
	}, nil
}

func (h *Handler) configListSections(iniPath string, _ ConfigOpsInput) (*mcp.CallToolResult, ConfigOpsOutput, error) {
	sections, err := parseINI(iniPath)
	if err != nil {
		return nil, ConfigOpsOutput{}, err
	}

	names := make([]string, 0, len(sections))
	for name := range sections {
		names = append(names, name)
	}
	sort.Strings(names)

	return nil, ConfigOpsOutput{
		Success:  true,
		File:     filepath.Base(iniPath),
		Sections: names,
	}, nil
}

// ---------------------------------------------------------------------------
// UE INI parser
// ---------------------------------------------------------------------------
//
// UE INI files follow a simple format:
//   [SectionName]
//   Key=Value
//   +Key=ArrayValue   (append to array)
//   ;comment
//
// We preserve the +/- prefixes on keys as-is to maintain UE semantics.

// parseINI reads a UE INI file and returns a map of section → key → value.
// For duplicate keys (array values with + prefix), only the last value is kept
// in the map; the full file is preserved during writes.
func parseINI(path string) (map[string]map[string]string, error) {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file not found: %s", filepath.Base(path))
		}
		return nil, fmt.Errorf("opening config file: %w", err)
	}
	defer func() { _ = f.Close() }()

	sections := make(map[string]map[string]string)
	currentSection := ""

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Skip empty lines and comments.
		if trimmed == "" || strings.HasPrefix(trimmed, ";") || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Section header.
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			currentSection = trimmed[1 : len(trimmed)-1]
			if _, ok := sections[currentSection]; !ok {
				sections[currentSection] = make(map[string]string)
			}
			continue
		}

		// Key=Value pair.
		if idx := strings.IndexByte(trimmed, '='); idx > 0 && currentSection != "" {
			key := trimmed[:idx]
			value := trimmed[idx+1:]
			sections[currentSection][key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	return sections, nil
}

// setINIValue sets a key in a section, creating the section if needed.
// It preserves all other content in the file (comments, ordering, etc.).
func setINIValue(path, section, key, value string) error {
	lines, err := readLines(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("reading config file: %w", err)
	}

	newLine := key + "=" + value
	sectionHeader := "[" + section + "]"

	inSection := false
	keySet := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Track current section.
		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			if inSection && !keySet {
				// We were in the target section but didn't find the key.
				// Insert the new key before this next section header.
				lines = insertLine(lines, i, newLine)
				keySet = true
				break
			}
			inSection = (trimmed == sectionHeader)
			continue
		}

		if inSection {
			if idx := strings.IndexByte(trimmed, '='); idx > 0 {
				existingKey := trimmed[:idx]
				if existingKey == key {
					lines[i] = newLine
					keySet = true
					break
				}
			}
		}
	}

	// If we were in the section at EOF without finding the key, append.
	if inSection && !keySet {
		lines = append(lines, newLine)
		keySet = true
	}

	// If the section doesn't exist at all, add it at the end.
	if !keySet {
		if len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) != "" {
			lines = append(lines, "")
		}
		lines = append(lines, sectionHeader, newLine)
	}

	return writeLines(path, lines)
}

// deleteINIValue removes a key from a section.
func deleteINIValue(path, section, key string) error {
	lines, err := readLines(path)
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	sectionHeader := "[" + section + "]"
	inSection := false
	deleted := false

	var result []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			inSection = (trimmed == sectionHeader)
		}

		if inSection {
			if idx := strings.IndexByte(trimmed, '='); idx > 0 {
				existingKey := trimmed[:idx]
				if existingKey == key {
					deleted = true
					continue // skip this line
				}
			}
		}

		result = append(result, line)
	}

	if !deleted {
		return fmt.Errorf("key %q not found in section [%s]", key, section)
	}

	return writeLines(path, result)
}

// readLines reads a file into a slice of lines.
func readLines(path string) ([]string, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	// Split preserving empty lines. TrimRight to avoid trailing empty element
	// from a final newline.
	content := strings.TrimRight(string(data), "\n\r")
	if content == "" {
		return nil, nil
	}
	return strings.Split(content, "\n"), nil
}

// writeLines writes a slice of lines back to a file with a trailing newline.
func writeLines(path string, lines []string) error {
	// Ensure the directory exists.
	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}
	content := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(path, []byte(content), 0o600)
}

// insertLine inserts a line at position i, shifting everything after.
func insertLine(lines []string, i int, line string) []string {
	lines = append(lines, "")
	copy(lines[i+1:], lines[i:])
	lines[i] = line
	return lines
}
