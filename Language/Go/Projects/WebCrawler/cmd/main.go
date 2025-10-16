package main

import (
	"WebCrawler/internal/crawler"
	"WebCrawler/internal/exporter"
	"WebCrawler/internal/models"
	"WebCrawler/pkg/config"
	"WebCrawler/pkg/logger"
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Command-line flags used to configure the Web Crawler application.
var (
	configPath   = flag.String("config", "configs/config.json", "Path to the configuration file")
	startPage    = flag.Int("start", 1, "Starting page number for crawling")
	endPage      = flag.Int("end", 10, "Ending page number for crawling")
	outputFile   = flag.String("output", "output/quotes", "Output file path (without extension)")
	outputFormat = flag.String("format", "json", "Output format. Options: json, csv, markdown")
	verbose      = flag.Bool("verbose", false, "Enable verbose (debug) logging output")
)

// main serves as the entry point for the Web Crawler application.
//
// It performs the following steps:
//  1. Parses CLI flags.
//  2. Loads configuration from a JSON file or defaults.
//  3. Initializes the logger based on configuration and flags.
//  4. Sets up graceful shutdown via OS signal interception.
//  5. Creates and runs the crawler.
//  6. Exports the scraped data.
//  7. Displays a preview of sample records.
//
// Exits with a non-zero status on fatal errors.
func main() {
	flag.Parse()

	cfg, err := loadConfig(*configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	logLevel := cfg.Logger.Level
	if *verbose {
		logLevel = "debug"
	}
	if err := logger.Init(logLevel, cfg.Logger.OutputPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	printBanner()
	logger.Info("Web Crawler started")
	logger.Info("Version: 1.0.0")

	// Create cancellable context for graceful shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt and termination signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		logger.Warn("Interrupt signal received. Gracefully shutting down...")
		cancel()
	}()

	// Initialize crawler.
	c := crawler.NewCrawler(&cfg.Crawler)

	logger.Info("Crawling pages from %d to %d", *startPage, *endPage)
	quotes, err := c.Start(ctx, *startPage, *endPage)
	if err != nil {
		logger.Fatal("Crawling failed: %v", err)
	}

	if len(quotes) == 0 {
		logger.Warn("No data retrieved from crawling")
		return
	}

	if err := exportData(quotes, *outputFile, *outputFormat); err != nil {
		logger.Fatal("Failed to export data: %v", err)
	}

	displaySamples(quotes, 5)
	logger.Info("Crawling completed successfully!")
}

// loadConfig reads the application configuration from a specified file.
//
// If the configuration file does not exist, it falls back to the default configuration.
// The configuration defines parameters for logger, crawler, and exporter.
//
// Args:
//
//	path: The filesystem path to the configuration JSON file.
//
// Returns:
//
//	A pointer to a config.Config struct and an error if loading fails.
func loadConfig(path string) (*config.Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		logger.Warn("Configuration file not found. Using default configuration.")
		return config.Default(), nil
	}

	cfg, err := config.Load(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return cfg, nil
}

// exportData writes the crawled quotes to a file using the specified export format.
//
// Supported formats include JSON, CSV, and Markdown.
// The exporter is determined dynamically using an exporter factory.
//
// Args:
//
//	quotes: A slice of Quote structs containing the scraped data.
//	baseFilename: The file name (without extension) to which data will be written.
//	format: The output format ("json", "csv", or "markdown").
//
// Returns:
//
//	An error if the export operation fails.
func exportData(quotes []models.Quote, baseFilename, format string) error {
	factory := &exporter.ExportFactory{}
	exp, err := factory.GetExporter(format)
	if err != nil {
		return err
	}

	// Add the appropriate file extension.
	filename := fmt.Sprintf("%s.%s", baseFilename, format)
	if format == "markdown" {
		filename = fmt.Sprintf("%s.md", baseFilename)
	}

	logger.Info("Exporting data to: %s", filename)

	if err := exp.Export(quotes, filename); err != nil {
		return err
	}

	logger.Info("Data exported successfully: %s (%d records)", filename, len(quotes))
	return nil
}

// displaySamples prints a limited number of sample quotes for preview.
//
// This function is primarily used for quick verification of crawled data
// without requiring the user to open the exported file.
//
// Args:
//
//	quotes: A slice of Quote structs containing the crawled data.
//	limit: The maximum number of sample quotes to display.
func displaySamples(quotes []models.Quote, limit int) {
	logger.Info("\n========== Sample Data (Top %d) ==========", limit)

	count := limit
	if len(quotes) < limit {
		count = len(quotes)
	}

	for i := 0; i < count; i++ {
		q := quotes[i]
		logger.Info("\n--- Quote %d ---", i+1)
		logger.Info("Content: %v", q.Content)
		logger.Info("Author: %v", q.Author)
		logger.Info("Tags: %v", q.Tags)
	}

	logger.Info("\n========================================")
}

// printBanner prints a stylized ASCII banner for the application.
//
// The banner is displayed at startup to provide visual identification
// and basic version information. This is purely aesthetic.
func printBanner() {
	banner := `
╔══════════════════════════════════════╗
║        Web Crawler v1.0.0            ║
║   High-performance Web Crawler       ║
╚══════════════════════════════════════╝
`
	fmt.Println(banner)
}
