package config

import (
	"encoding/json"
	"os"
	"time"
)

// Config defines the top-level application configuration.
//
// It aggregates settings for the crawler, logger, and optional database.
// The configuration is typically loaded from a JSON file.
type Config struct {
	Crawler  CrawlerConfig  `json:"crawler"`
	Logger   LoggerConfig   `json:"logger"`
	Database DatabaseConfig `json:"database,omitempty"`
}

// CrawlerConfig defines runtime parameters for the web crawler.
//
// These settings control crawling limits, concurrency, timeouts,
// retries, and user agent information.
type CrawlerConfig struct {
	BaseURL         string        `json:"base_url"`           // Base URL of the target site to crawl.
	MaxPages        int           `json:"max_pages"`          // Maximum number of pages to crawl.
	Concurrency     int           `json:"concurrency"`        // Number of concurrent crawling workers.
	RequestTimeout  time.Duration `json:"request_timeout"`    // Timeout per HTTP request.
	RetryAttempts   int           `json:"retry_attempts"`     // Number of retry attempts on failure.
	RateLimitPerSec int           `json:"rate_limit_per_sec"` // Maximum allowed requests per second.
	UserAgent       string        `json:"user_agent"`         // Custom user agent for requests.
}

// LoggerConfig defines configuration parameters for the logging system.
//
// It specifies the log level, output path, rotation policy,
// and log file retention limits.
type LoggerConfig struct {
	Level      string `json:"level"`        // Logging level (e.g., debug, info, warn, error).
	OutputPath string `json:"output_path"`  // File path for log output.
	MaxSize    int    `json:"max_size_mb"`  // Maximum log file size in megabytes before rotation.
	MaxBackups int    `json:"max_backups"`  // Number of old log files to retain.
	MaxAge     int    `json:"max_age_days"` // Maximum age (in days) of retained log files.
}

// DatabaseConfig defines optional parameters for database connectivity.
//
// If the database feature is disabled, these settings are ignored.
type DatabaseConfig struct {
	Enabled  bool   `json:"enabled"`  // Whether database storage is enabled.
	Driver   string `json:"driver"`   // Database driver (e.g., mysql, postgres, sqlite).
	Host     string `json:"host"`     // Database host address.
	Port     int    `json:"port"`     // Database port number.
	Database string `json:"database"` // Database name.
	Username string `json:"username"` // Username for authentication.
	Password string `json:"password"` // Password for authentication.
}

// Load reads configuration from a JSON file and applies default values.
//
// Args:
//
//	path: Path to the configuration JSON file.
//
// Returns:
//
//	A pointer to a Config struct containing merged settings,
//	or an error if the file cannot be opened or parsed.
//
// Example:
//
//	cfg, err := config.Load("configs/config.json")
//	if err != nil {
//	    log.Fatal(err)
//	}
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	setDefaults(&cfg)
	return &cfg, nil
}

// setDefaults assigns default values to configuration fields that are unset.
//
// This ensures that minimal configuration files still produce a functional setup.
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

// Default returns a pre-configured default Config instance.
//
// This function is used when no external configuration file is available.
// The defaults are tuned for a typical single-machine crawler deployment.
//
// Returns:
//
//	A pointer to a Config struct populated with default settings.
func Default() *Config {
	cfg := &Config{
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
	return cfg
}
