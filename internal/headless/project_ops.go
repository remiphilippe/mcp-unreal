package headless

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- project_ops ---

// UProjectPlugin describes a plugin entry in the .uproject file.
type UProjectPlugin struct {
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

// UProjectModule describes a module entry in the .uproject file.
type UProjectModule struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// UProjectFile represents the parsed .uproject JSON structure.
type UProjectFile struct {
	FileVersion     int              `json:"file_version,omitempty"`
	EngineVersion   string           `json:"engine_version,omitempty"`
	Category        string           `json:"category,omitempty"`
	Description     string           `json:"description,omitempty"`
	Modules         []UProjectModule `json:"modules,omitempty"`
	Plugins         []UProjectPlugin `json:"plugins,omitempty"`
	TargetPlatforms []string         `json:"target_platforms,omitempty"`
}

// ProjectOpsInput defines parameters for the project_ops tool.
type ProjectOpsInput struct {
	Operation string `json:"operation" jsonschema:"required,Operation: get_info, list_plugins, enable_plugin, disable_plugin, add_module, set_target_platforms"`
	// For enable_plugin, disable_plugin.
	Name string `json:"name,omitempty" jsonschema:"Plugin or module name. For enable_plugin, disable_plugin, add_module."`
	// For add_module.
	Type string `json:"type,omitempty" jsonschema:"Module type: Runtime, Editor, Developer, Program. For add_module."`
	// For set_target_platforms.
	Platforms []string `json:"platforms,omitempty" jsonschema:"Target platform list (e.g. ['Mac', 'Win64']). For set_target_platforms."`
}

// ProjectOpsOutput is returned by the project_ops tool.
type ProjectOpsOutput struct {
	Success         bool             `json:"success" jsonschema:"whether the operation succeeded"`
	ProjectName     string           `json:"project_name,omitempty" jsonschema:"project name"`
	EngineVersion   string           `json:"engine_version,omitempty" jsonschema:"engine version association"`
	Modules         []UProjectModule `json:"modules,omitempty" jsonschema:"project modules"`
	Plugins         []UProjectPlugin `json:"plugins,omitempty" jsonschema:"project plugins"`
	TargetPlatforms []string         `json:"target_platforms,omitempty" jsonschema:"target platforms"`
	Message         string           `json:"message,omitempty" jsonschema:"status message"`
}

// RegisterProject adds the project_ops tool to the MCP server.
func (h *Handler) RegisterProject(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "project_ops",
		Description: "Read and modify the .uproject file: get project info, list/enable/disable plugins, " +
			"add modules, and set target platforms. " +
			"Operations: get_info, list_plugins, enable_plugin, disable_plugin, add_module, set_target_platforms. " +
			"Does not require the editor to be running — reads/writes the .uproject file directly. " +
			"Creates a .uproject.bak backup before writing.",
	}, h.ProjectOps)
}

// findUProject locates the .uproject file in the project root.
func (h *Handler) findUProject() (string, error) {
	root := h.Config.ProjectRoot
	if root == "" {
		return "", fmt.Errorf("project root not configured — set MCP_UNREAL_PROJECT env var")
	}

	// Check if root is directly a .uproject file.
	if filepath.Ext(root) == ".uproject" {
		if _, err := os.Stat(root); err == nil {
			return root, nil
		}
	}

	// Search for .uproject in the root directory.
	entries, err := os.ReadDir(root)
	if err != nil {
		return "", fmt.Errorf("reading project root %s: %w", root, err)
	}
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".uproject" {
			return filepath.Join(root, entry.Name()), nil
		}
	}

	return "", fmt.Errorf("no .uproject file found in %s", root)
}

// readUProject reads and parses the .uproject file.
func readUProject(path string) (*uprojectRaw, error) {
	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	var proj uprojectRaw
	if err := json.Unmarshal(data, &proj); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return &proj, nil
}

// writeUProject writes the .uproject file, creating a backup first.
func writeUProject(path string, proj *uprojectRaw) error {
	// Backup.
	bakPath := path + ".bak"
	if data, err := os.ReadFile(filepath.Clean(path)); err == nil {
		_ = os.WriteFile(bakPath, data, 0o600)
	}

	data, err := json.MarshalIndent(proj, "", "\t")
	if err != nil {
		return fmt.Errorf("marshaling .uproject: %w", err)
	}
	return os.WriteFile(filepath.Clean(path), data, 0o600)
}

// uprojectRaw is the raw JSON structure for round-trip serialization.
type uprojectRaw struct {
	FileVersion       int                        `json:"FileVersion,omitempty"`
	EngineAssociation string                     `json:"EngineAssociation,omitempty"`
	Category          string                     `json:"Category,omitempty"`
	Description       string                     `json:"Description,omitempty"`
	Modules           []uprojectModule           `json:"Modules,omitempty"`
	Plugins           []uprojectPlugin           `json:"Plugins,omitempty"`
	TargetPlatforms   []string                   `json:"TargetPlatforms,omitempty"`
	Extra             map[string]json.RawMessage `json:"-"`
}

type uprojectModule struct {
	Name                 string `json:"Name"`
	Type                 string `json:"Type"`
	LoadingPhase         string `json:"LoadingPhase,omitempty"`
	AdditionalDependency string `json:"AdditionalDependencies,omitempty"`
}

type uprojectPlugin struct {
	Name    string `json:"Name"`
	Enabled bool   `json:"Enabled"`
}

// ProjectOps implements the project_ops tool.
func (h *Handler) ProjectOps(ctx context.Context, req *mcp.CallToolRequest, input ProjectOpsInput) (*mcp.CallToolResult, ProjectOpsOutput, error) {
	if input.Operation == "" {
		return nil, ProjectOpsOutput{}, fmt.Errorf("operation is required")
	}

	uprojectPath, err := h.findUProject()
	if err != nil {
		return nil, ProjectOpsOutput{}, err
	}

	proj, err := readUProject(uprojectPath)
	if err != nil {
		return nil, ProjectOpsOutput{}, err
	}

	projectName := filepath.Base(filepath.Dir(uprojectPath))
	if filepath.Ext(uprojectPath) == ".uproject" {
		projectName = filepath.Base(uprojectPath[:len(uprojectPath)-len(".uproject")])
	}

	switch input.Operation {
	case "get_info":
		out := ProjectOpsOutput{
			Success:         true,
			ProjectName:     projectName,
			EngineVersion:   proj.EngineAssociation,
			TargetPlatforms: proj.TargetPlatforms,
		}
		for _, m := range proj.Modules {
			out.Modules = append(out.Modules, UProjectModule{Name: m.Name, Type: m.Type})
		}
		for _, p := range proj.Plugins {
			out.Plugins = append(out.Plugins, UProjectPlugin(p))
		}
		return nil, out, nil

	case "list_plugins":
		out := ProjectOpsOutput{Success: true}
		for _, p := range proj.Plugins {
			out.Plugins = append(out.Plugins, UProjectPlugin(p))
		}
		return nil, out, nil

	case "enable_plugin":
		if input.Name == "" {
			return nil, ProjectOpsOutput{}, fmt.Errorf("name is required for enable_plugin")
		}
		found := false
		for i := range proj.Plugins {
			if proj.Plugins[i].Name == input.Name {
				proj.Plugins[i].Enabled = true
				found = true
				break
			}
		}
		if !found {
			proj.Plugins = append(proj.Plugins, uprojectPlugin{Name: input.Name, Enabled: true})
		}
		if err := writeUProject(uprojectPath, proj); err != nil {
			return nil, ProjectOpsOutput{}, err
		}
		return nil, ProjectOpsOutput{
			Success: true,
			Message: fmt.Sprintf("Plugin '%s' enabled", input.Name),
		}, nil

	case "disable_plugin":
		if input.Name == "" {
			return nil, ProjectOpsOutput{}, fmt.Errorf("name is required for disable_plugin")
		}
		found := false
		for i := range proj.Plugins {
			if proj.Plugins[i].Name == input.Name {
				proj.Plugins[i].Enabled = false
				found = true
				break
			}
		}
		if !found {
			proj.Plugins = append(proj.Plugins, uprojectPlugin{Name: input.Name, Enabled: false})
		}
		if err := writeUProject(uprojectPath, proj); err != nil {
			return nil, ProjectOpsOutput{}, err
		}
		return nil, ProjectOpsOutput{
			Success: true,
			Message: fmt.Sprintf("Plugin '%s' disabled", input.Name),
		}, nil

	case "add_module":
		if input.Name == "" {
			return nil, ProjectOpsOutput{}, fmt.Errorf("name is required for add_module")
		}
		modType := input.Type
		if modType == "" {
			modType = "Runtime"
		}
		// Check if module already exists.
		for _, m := range proj.Modules {
			if m.Name == input.Name {
				return nil, ProjectOpsOutput{
					Success: true,
					Message: fmt.Sprintf("Module '%s' already exists", input.Name),
				}, nil
			}
		}
		proj.Modules = append(proj.Modules, uprojectModule{
			Name:         input.Name,
			Type:         modType,
			LoadingPhase: "Default",
		})
		if err := writeUProject(uprojectPath, proj); err != nil {
			return nil, ProjectOpsOutput{}, err
		}
		return nil, ProjectOpsOutput{
			Success: true,
			Message: fmt.Sprintf("Module '%s' (%s) added", input.Name, modType),
		}, nil

	case "set_target_platforms":
		if len(input.Platforms) == 0 {
			return nil, ProjectOpsOutput{}, fmt.Errorf("platforms array is required for set_target_platforms")
		}
		proj.TargetPlatforms = input.Platforms
		if err := writeUProject(uprojectPath, proj); err != nil {
			return nil, ProjectOpsOutput{}, err
		}
		return nil, ProjectOpsOutput{
			Success:         true,
			TargetPlatforms: input.Platforms,
			Message:         "Target platforms updated",
		}, nil

	default:
		return nil, ProjectOpsOutput{}, fmt.Errorf("unknown operation: %s", input.Operation)
	}
}
