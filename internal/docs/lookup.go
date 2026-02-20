package docs

import (
	"context"
	"fmt"
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// --- lookup_docs tool ---

// LookupDocsInput defines parameters for the lookup_docs tool.
type LookupDocsInput struct {
	Query     string `json:"query" jsonschema:"Natural language query about UE APIs, classes, or patterns"`
	Category  string `json:"category,omitempty" jsonschema:"Optional filter: actor, blueprint, material, animation, input, realtimemesh, gameplay, rendering, networking"`
	MaxTokens int    `json:"max_tokens,omitempty" jsonschema:"Max approximate tokens to return. Default 3000."`
}

// DocResult represents a single search result.
type DocResult struct {
	Title    string  `json:"title" jsonschema:"document title"`
	Source   string  `json:"source" jsonschema:"document source (ue5.7, realtimemesh, project)"`
	Category string  `json:"category" jsonschema:"document category"`
	Snippet  string  `json:"snippet" jsonschema:"relevant content snippet"`
	URL      string  `json:"url,omitempty" jsonschema:"original documentation URL if available"`
	Score    float64 `json:"score" jsonschema:"relevance score"`
}

// LookupDocsOutput is returned by the lookup_docs tool.
type LookupDocsOutput struct {
	Results []DocResult `json:"results" jsonschema:"search results ordered by relevance"`
	Total   int         `json:"total" jsonschema:"number of results returned"`
}

// --- lookup_class tool ---

// LookupClassInput defines parameters for the lookup_class tool.
type LookupClassInput struct {
	ClassName string `json:"class_name" jsonschema:"UE class name, e.g. AActor, UCharacterMovementComponent, URealtimeMeshSimple"`
}

// LookupClassOutput is returned by the lookup_class tool.
type LookupClassOutput struct {
	Found bool      `json:"found" jsonschema:"whether the class was found in the index"`
	Class ClassInfo `json:"class,omitempty" jsonschema:"structured class reference if found"`
}

// Register adds the lookup tools to the MCP server.
func (d *Index) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name: "lookup_docs",
		Description: "Search UE 5.7 API docs, RealtimeMesh docs, and project-specific docs. " +
			"Returns relevant snippets with source references. " +
			"Use this before writing code that uses UE APIs you are unsure about. " +
			"Always available — does not require the editor to be running.",
	}, d.LookupDocs)

	mcp.AddTool(server, &mcp.Tool{
		Name: "lookup_class",
		Description: "Get the full class reference for a specific UE class: inheritance chain, " +
			"key properties, key functions, and usage notes. " +
			"Use this when you need detailed API information for a specific class. " +
			"Always available — does not require the editor to be running.",
	}, d.LookupClass)
}

// LookupDocs implements the lookup_docs tool.
// See IMPLEMENTATION.md §4.4 for the design.
func (d *Index) LookupDocs(ctx context.Context, req *mcp.CallToolRequest, input LookupDocsInput) (*mcp.CallToolResult, LookupDocsOutput, error) {
	if input.Query == "" {
		return nil, LookupDocsOutput{}, fmt.Errorf("query is required")
	}

	maxTokens := input.MaxTokens
	if maxTokens <= 0 {
		maxTokens = 3000
	}

	queryStr := input.Query
	if input.Category != "" {
		queryStr = fmt.Sprintf("+category:%s %s", input.Category, queryStr)
	}

	searchReq := bleve.NewSearchRequest(bleve.NewQueryStringQuery(queryStr))
	searchReq.Size = 10
	searchReq.Fields = []string{"title", "source", "category", "content", "url"}

	result, err := d.index.Search(searchReq)
	if err != nil {
		return nil, LookupDocsOutput{}, fmt.Errorf("search failed: %w", err)
	}

	var results []DocResult
	totalTokens := 0
	for _, hit := range result.Hits {
		content := strField(hit.Fields, "content")

		// Rough token estimate: 1 token ~ 4 chars (IMPLEMENTATION.md §10: token budget).
		tokens := len(content) / 4
		if totalTokens+tokens > maxTokens {
			remaining := (maxTokens - totalTokens) * 4
			if remaining > 0 && len(content) > remaining {
				content = content[:remaining] + "..."
			} else {
				break
			}
		}
		totalTokens += tokens

		results = append(results, DocResult{
			Title:    strField(hit.Fields, "title"),
			Source:   strField(hit.Fields, "source"),
			Category: strField(hit.Fields, "category"),
			Snippet:  content,
			URL:      strField(hit.Fields, "url"),
			Score:    hit.Score,
		})
	}

	return nil, LookupDocsOutput{Results: results, Total: len(results)}, nil
}

// LookupClass implements the lookup_class tool.
func (d *Index) LookupClass(ctx context.Context, req *mcp.CallToolRequest, input LookupClassInput) (*mcp.CallToolResult, LookupClassOutput, error) {
	if input.ClassName == "" {
		return nil, LookupClassOutput{}, fmt.Errorf("class_name is required")
	}

	// First try: exact title match — the most authoritative doc for a class
	// is the one titled with the class name itself.
	titleQuery := fmt.Sprintf("+title:%s", input.ClassName)
	titleReq := bleve.NewSearchRequest(bleve.NewQueryStringQuery(titleQuery))
	titleReq.Size = 5
	titleReq.Fields = []string{"title", "source", "content", "url", "classes"}

	result, err := d.index.Search(titleReq)
	if err == nil && result.Total > 0 {
		// Prefer the hit whose title exactly matches the class name.
		hit := pickBestHit(result, input.ClassName)
		content := strField(hit.Fields, "content")
		info := ParseClassDoc(input.ClassName, content)
		info.Source = strField(hit.Fields, "source")
		info.URL = strField(hit.Fields, "url")
		return nil, LookupClassOutput{Found: true, Class: info}, nil
	}

	// Second try: search by classes field — finds docs that reference this class.
	query := fmt.Sprintf("+classes:%s", input.ClassName)
	searchReq := bleve.NewSearchRequest(bleve.NewQueryStringQuery(query))
	searchReq.Size = 10
	searchReq.Fields = []string{"title", "source", "content", "url", "classes"}

	result, err = d.index.Search(searchReq)
	if err == nil && result.Total > 0 {
		hit := pickBestHit(result, input.ClassName)
		content := strField(hit.Fields, "content")
		info := ParseClassDoc(input.ClassName, content)
		info.Source = strField(hit.Fields, "source")
		info.URL = strField(hit.Fields, "url")
		return nil, LookupClassOutput{Found: true, Class: info}, nil
	}

	// Fallback: try a content search for the class name.
	fallbackQuery := bleve.NewSearchRequest(bleve.NewQueryStringQuery(input.ClassName))
	fallbackQuery.Size = 5
	fallbackQuery.Fields = []string{"title", "source", "content", "url"}

	result, err = d.index.Search(fallbackQuery)
	if err != nil || result.Total == 0 {
		return nil, LookupClassOutput{Found: false}, nil
	}

	hit := pickBestHit(result, input.ClassName)
	content := strField(hit.Fields, "content")
	info := ParseClassDoc(input.ClassName, content)
	info.Source = strField(hit.Fields, "source")
	info.URL = strField(hit.Fields, "url")

	return nil, LookupClassOutput{Found: true, Class: info}, nil
}

// pickBestHit selects the most relevant hit for a class lookup.
// It prefers the hit whose title exactly matches the class name,
// falling back to the highest-scored hit.
func pickBestHit(result *bleve.SearchResult, className string) *search.DocumentMatch {
	for _, hit := range result.Hits {
		title := strField(hit.Fields, "title")
		if strings.EqualFold(title, className) {
			return hit
		}
	}
	return result.Hits[0]
}

func strField(fields map[string]interface{}, key string) string {
	if v, ok := fields[key].(string); ok {
		return v
	}
	return ""
}
