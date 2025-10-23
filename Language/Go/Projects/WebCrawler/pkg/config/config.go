package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

//
// ======== Configuration Structures ========
//

// Config represents the root application configuration.
//
// It groups together settings for the web crawler, logging system,
// and (optionally) database connections.
type Config struct {
	Crawler  CrawlerConfig  `json:"crawler"`            // Crawler-related settings.
	Logger   LoggerConfig   `json:"logger"`             // Logging configuration.
	Database DatabaseConfig `json:"database,omitempty"` // Optional database connection parameters.
}

//
// ======== Crawler Configuration ========
//

// CrawlerConfig defines parameters for controlling crawling behavior.
//
// These include concurrency, request retry policy, rate limits,
// and HTTP client settings.
type CrawlerConfig struct {
	BaseURL         string        `json:"base_url"`           // Base URL of the target website.
	MaxPages        int           `json:"max_pages"`          // Maximum number of pages to crawl.
	Concurrency     int           `json:"concurrency"`        // Number of concurrent worker goroutines.
	RequestTimeout  time.Duration `json:"request_timeout"`    // HTTP request timeout duration.
	RetryAttempts   int           `json:"retry_attempts"`     // Maximum number of retry attempts per request.
	RateLimitPerSec int           `json:"rate_limit_per_sec"` // Maximum requests per second.
	UserAgent       string        `json:"user_agent"`         // Custom User-Agent header for requests.
}

//
// ======== Logger Configuration ========
//

// LoggerConfig defines file rotation and verbosity settings
// for the application logger.
type LoggerConfig struct {
	Level      string `json:"level"`        // Log level (e.g., "debug", "info", "error").
	OutputPath string `json:"output_path"`  // Path to the log output file.
	MaxSize    int    `json:"max_size_mb"`  // Maximum log file size before rotation.
	MaxBackups int    `json:"max_backups"`  // Number of rotated log files to keep.
	MaxAge     int    `json:"max_age_days"` // Maximum number of days to retain old log files.
}

//
// ======== Database Configuration ========
//

// DatabaseConfig defines optional database connection settings.
//
// If `Enabled` is false, the crawler operates without persistence.
type DatabaseConfig struct {
	Enabled  bool   `json:"enabled"`  // Whether database integration is enabled.
	Driver   string `json:"driver"`   // Database driver name (e.g., "mysql", "postgres").
	Host     string `json:"host"`     // Database server hostname or IP address.
	Port     int    `json:"port"`     // Database server port number.
	Database string `json:"database"` // Database name.
	Username string `json:"username"` // Username for authentication.
	Password string `json:"password"` // Password for authentication.
}

//
// ======== Configuration Loading ========
//

// Load reads configuration settings from a JSON file.
//
// It automatically fills in missing fields with safe defaults
// defined in `setDefaults()`.
//
// Args:
//
//	path: The file path to the JSON configuration file.
//
// Returns:
//
//	A pointer to a fully initialized Config instance, or an error if reading fails.
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config JSON: %w", err)
	}

	setDefaults(&cfg)
	return &cfg, nil
}

//
// ======== Default Handling ========
//

// setDefaults assigns default values for missing configuration fields.
//
// This ensures the application runs reliably even with minimal config input.
func setDefaults(cfg *Config) {
	if cfg.Crawler.Concurrency == 0 {
		cfg.Crawler.Concurrency = 5
	}
	if cfg.Crawler.RequestTimeout == 0 {
		cfg.Crawler.RequestTimeout = 10 * time.Second
	}
	if cfg.Crawler.RetryAttempts == 0 {
		cfg.Crawler.RetryAttempts = 3
	}
	if cfg.Crawler.RateLimitPerSec == 0 {
		cfg.Crawler.RateLimitPerSec = 10
	}
	if cfg.Crawler.UserAgent == "" {
		cfg.Crawler.UserAgent = "WebCrawler/1.0"
	}
	if cfg.Logger.Level == "" {
		cfg.Logger.Level = "info"
	}
}

//
// ======== Default Configuration Generator ========
//

// Default returns a pre-configured Config object
// with sensible defaults suitable for development or testing.
//
// Returns:
//
//	A pointer to a Config struct populated with standard defaults.
func Default() *Config {
	return &Config{
		Crawler: CrawlerConfig{
			BaseURL:         "https://quotes.toscrape.com",
			MaxPages:        10,
			Concurrency:     5,
			RequestTimeout:  10 * time.Second,
			RetryAttempts:   3,
			RateLimitPerSec: 10,
			UserAgent:       "WebCrawler/1.0",
		},
		Logger: LoggerConfig{
			Level:      "info",
			OutputPath: "logs/crawler.log",
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     28,
		},
	}
}
