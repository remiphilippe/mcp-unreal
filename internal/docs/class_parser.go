// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

package docs

import (
	"strings"
)

// ClassInfo holds structured information about a UE class.
// Parsed from markdown class reference documents.
type ClassInfo struct {
	Name        string   `json:"name"`
	Parent      string   `json:"parent,omitempty"`
	Module      string   `json:"module,omitempty"`
	Description string   `json:"description"`
	KeyProps    []string `json:"key_properties,omitempty"`
	KeyFuncs    []string `json:"key_functions,omitempty"`
	Source      string   `json:"source,omitempty"`
	URL         string   `json:"url,omitempty"`
}

// ParseClassDoc extracts structured class information from markdown content.
// Expected format:
//
//	# ClassName
//	**Parent**: ParentClass
//	**Module**: ModuleName
//	Description text...
//	## Key Properties
//	- `PropertyName` — description
//	## Key Functions
//	- `FunctionName(params)` — description
func ParseClassDoc(name, content string) ClassInfo {
	info := ClassInfo{
		Name: name,
	}

	lines := strings.Split(content, "\n")
	section := "" // current section being parsed

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Track sections.
		if strings.HasPrefix(trimmed, "## ") {
			sectionName := strings.ToLower(strings.TrimPrefix(trimmed, "## "))
			switch {
			case strings.Contains(sectionName, "propert"):
				section = "properties"
			case strings.Contains(sectionName, "function") || strings.Contains(sectionName, "method"):
				section = "functions"
			default:
				section = ""
			}
			continue
		}

		// Parse metadata.
		if strings.HasPrefix(trimmed, "**Parent**:") || strings.HasPrefix(trimmed, "**Parent Class**:") {
			info.Parent = extractMetaValue(trimmed)
			continue
		}
		if strings.HasPrefix(trimmed, "**Module**:") {
			info.Module = extractMetaValue(trimmed)
			continue
		}

		// Parse list items in property/function sections.
		if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
			item := strings.TrimLeft(trimmed, "-* ")
			item = strings.TrimSpace(item)
			if item == "" {
				continue
			}
			switch section {
			case "properties":
				info.KeyProps = append(info.KeyProps, item)
			case "functions":
				info.KeyFuncs = append(info.KeyFuncs, item)
			}
			continue
		}

		// Collect description from non-section, non-heading lines before first section.
		if section == "" && !strings.HasPrefix(trimmed, "#") && !strings.HasPrefix(trimmed, "**") && trimmed != "" {
			if info.Description == "" {
				info.Description = trimmed
			} else if len(info.Description) < 500 {
				info.Description += " " + trimmed
			}
		}
	}

	return info
}

// extractMetaValue pulls the value after "**Key**: value" markdown.
func extractMetaValue(line string) string {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) < 2 {
		return ""
	}
	val := strings.TrimSpace(parts[1])
	// Strip backticks and markdown formatting.
	val = strings.Trim(val, "`*")
	return val
}
