package exporter

import (
	"WebCrawler/internal/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//
// ======== Interfaces & Factories ========
//

// Exporter defines a common interface for exporting quote data.
//
// Each implementation (JSON, CSV, Markdown) provides its own
// format-specific logic for serializing and writing the data.
type Exporter interface {
	Export(quotes []models.Quote, filename string) error
}

// ExportFactory is responsible for creating exporters
// based on the desired output format.
type ExportFactory struct{}

// GetExporter returns an appropriate Exporter implementation
// based on the given format string. Supported formats: json, csv, markdown.
//
// Example:
//
//	exporter, _ := factory.GetExporter("csv")
//	exporter.Export(data, "output.csv")
func (f *ExportFactory) GetExporter(format string) (Exporter, error) {
	switch strings.ToLower(format) {
	case "json":
		return NewJSONExporter(true), nil
	case "csv":
		return NewCSVExporter(), nil
	case "markdown", "md":
		return NewMarkdownExporter(), nil
	default:
		return nil, fmt.Errorf("unsupported export format: %s", format)
	}
}

//
// ======== JSON Exporter ========
//

// JSONExporter exports quote data to a JSON file.
type JSONExporter struct {
	indent bool
}

// NewJSONExporter creates a new JSONExporter.
//
// If indent is true, the output will be formatted with indentation
// for human readability.
func NewJSONExporter(indent bool) *JSONExporter {
	return &JSONExporter{indent: indent}
}

// Export writes the provided quotes to a JSON file at the given path.
//
// The function ensures that the destination directory exists,
// encodes the data as JSON, and handles any I/O errors gracefully.
func (e *JSONExporter) Export(quotes []models.Quote, filename string) error {
	if err := ensureDir(filename); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if e.indent {
		encoder.SetIndent("", "  ")
	}

	if err := encoder.Encode(quotes); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}

//
// ======== CSV Exporter ========
//

// CSVExporter exports quote data to a CSV file.
type CSVExporter struct{}

// NewCSVExporter creates a new CSVExporter instance.
func NewCSVExporter() *CSVExporter {
	return &CSVExporter{}
}

// Export writes the provided quotes to a CSV file.
//
// The exported file includes headers: "Content", "Author", and "Tags".
// Tags are concatenated with semicolons when multiple tags exist.
func (e *CSVExporter) Export(quotes []models.Quote, filename string) error {
	if err := ensureDir(filename); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"Content", "Author", "Tags"}); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	for _, quote := range quotes {
		tags := strings.Join(quote.Tags, ";")
		record := []string{quote.Content, quote.Author, tags}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}
	}

	return nil
}

//
// ======== Markdown Exporter ========
//

// MarkdownExporter exports quote data to a Markdown file.
type MarkdownExporter struct{}

// NewMarkdownExporter creates a new MarkdownExporter instance.
func NewMarkdownExporter() *MarkdownExporter {
	return &MarkdownExporter{}
}

// Export writes the provided quotes to a Markdown file.
//
// Each quote is formatted as a Markdown section containing the
// quote text, author, and associated tags.
func (e *MarkdownExporter) Export(quotes []models.Quote, filename string) error {
	if err := ensureDir(filename); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	fmt.Fprintf(file, "# Quotes Collection\n\n")
	fmt.Fprintf(file, "Total quotes: %d\n\n", len(quotes))
	fmt.Fprintf(file, "---\n\n")

	for i, quote := range quotes {
		fmt.Fprintf(file, "## Quote %d\n\n", i+1)
		fmt.Fprintf(file, "> %s\n\n", quote.Content)
		fmt.Fprintf(file, "**Author:** %s\n\n", quote.Author)

		if len(quote.Tags) > 0 {
			tagList := make([]string, len(quote.Tags))
			for j, tag := range quote.Tags {
				tagList[j] = fmt.Sprintf("`%s`", tag)
			}
			fmt.Fprintf(file, "**Tags:** %s\n\n", strings.Join(tagList, ", "))
		}

		fmt.Fprintf(file, "---\n\n")
	}

	return nil
}

//
// ======== Utility Functions ========
//

// ensureDir ensures that the directory for a given filename exists.
//
// It recursively creates all necessary parent directories with 0755 permissions.
func ensureDir(filename string) error {
	dir := filepath.Dir(filename)
	return os.MkdirAll(dir, 0755)
}
