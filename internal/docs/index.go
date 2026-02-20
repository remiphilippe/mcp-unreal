// Copyright (c) mcp-unreal project contributors. Apache-2.0 license.

// Package docs implements the documentation index and lookup tools
// for the mcp-unreal server.
//
// It uses Bleve (Go-native full-text search) to provide sub-millisecond
// UE API documentation lookups. See IMPLEMENTATION.md §4 for the full
// architecture.
package docs

import (
	"fmt"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/standard"
	"github.com/blevesearch/bleve/v2/mapping"
)

// DocEntry represents a single indexed document.
// See IMPLEMENTATION.md §4.1 for field descriptions.
type DocEntry struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Category string   `json:"category"` // actor, blueprint, material, animation, input, realtimemesh, gameplay, rendering, networking
	Source   string   `json:"source"`   // ue5.7, realtimemesh, project
	Content  string   `json:"content"`
	Classes  []string `json:"classes"` // related UE class names for cross-referencing
	URL      string   `json:"url"`
}

// Index wraps a Bleve index for documentation search.
type Index struct {
	index bleve.Index
}

// CreateIndex creates a new Bleve index at the given path with custom
// field mappings optimized for UE documentation search.
func CreateIndex(path string) (*Index, error) {
	indexMapping := buildIndexMapping()

	idx, err := bleve.New(path, indexMapping)
	if err != nil {
		return nil, fmt.Errorf("creating doc index at %s: %w", path, err)
	}

	return &Index{index: idx}, nil
}

// OpenIndex opens an existing Bleve index from disk.
func OpenIndex(path string) (*Index, error) {
	idx, err := bleve.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening doc index at %s: %w", path, err)
	}
	return &Index{index: idx}, nil
}

// OpenOrCreate opens an existing index or creates a new one if it doesn't exist.
func OpenOrCreate(path string) (*Index, error) {
	idx, err := OpenIndex(path)
	if err == nil {
		return idx, nil
	}
	return CreateIndex(path)
}

// Close closes the underlying Bleve index.
func (d *Index) Close() error {
	return d.index.Close()
}

// IndexDoc adds or updates a single document in the index.
func (d *Index) IndexDoc(entry DocEntry) error {
	return d.index.Index(entry.ID, entry)
}

// IndexBatch adds multiple documents in a single batch operation.
func (d *Index) IndexBatch(entries []DocEntry) error {
	batch := d.index.NewBatch()
	for _, entry := range entries {
		if err := batch.Index(entry.ID, entry); err != nil {
			return fmt.Errorf("batching doc %s: %w", entry.ID, err)
		}
	}
	return d.index.Batch(batch)
}

// DocCount returns the number of documents in the index.
func (d *Index) DocCount() (uint64, error) {
	return d.index.DocCount()
}

// buildIndexMapping creates the Bleve index mapping for DocEntry.
// Text fields (title, content) use the standard analyzer for full-text search.
// Keyword fields (category, source, classes) use exact-match for filtering.
func buildIndexMapping() mapping.IndexMapping {
	docMapping := bleve.NewDocumentMapping()

	// Text fields — full-text analyzed.
	titleField := bleve.NewTextFieldMapping()
	titleField.Analyzer = standard.Name
	titleField.Store = true
	docMapping.AddFieldMappingsAt("title", titleField)

	contentField := bleve.NewTextFieldMapping()
	contentField.Analyzer = standard.Name
	contentField.Store = true
	docMapping.AddFieldMappingsAt("content", contentField)

	// Keyword fields — exact match, stored for retrieval.
	categoryField := bleve.NewTextFieldMapping()
	categoryField.Analyzer = keyword.Name
	categoryField.Store = true
	docMapping.AddFieldMappingsAt("category", categoryField)

	sourceField := bleve.NewTextFieldMapping()
	sourceField.Analyzer = keyword.Name
	sourceField.Store = true
	docMapping.AddFieldMappingsAt("source", sourceField)

	urlField := bleve.NewTextFieldMapping()
	urlField.Analyzer = keyword.Name
	urlField.Store = true
	docMapping.AddFieldMappingsAt("url", urlField)

	// Classes — keyword array for cross-referencing by class name.
	classesField := bleve.NewTextFieldMapping()
	classesField.Analyzer = keyword.Name
	classesField.Store = true
	docMapping.AddFieldMappingsAt("classes", classesField)

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultMapping = docMapping

	return indexMapping
}
