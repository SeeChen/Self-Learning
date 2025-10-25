<div align=center>

# Golang Project: Web Crawler

[Overview](#1-summary)</br>
[Architecture](#2-architecture)</br>
[Environment Setup](#3-environmental-setup)</br>
[Configuration](#4-configuration)</br>
[Execution](#5-execution)</br>
[Future Work](#6-future-work)</br>

</div>

## 1. Overview
This project is a concurrent, modular web crawler built in Go, designed to scrape quotes, author, and tags from [Quotes to Scrape](https://quotes.toscrape.com/).

It demonstrates best practices in Go programming -- including structured configuratoin, concurrency control, context-based cancellation, retry logic, and clean arhitecture.


## 2. Architecture
This system is devided into clear, independent modules:
|PACKAGE|DESCRIPTION|
|---:|:---:|
|[main](./cmd/main.go)|Main entry point that initializes and starts the crawler.|
|[Crawler](./internal/crawler/crawler.go)|Core logic for URL scheduling, parsing, and coordination|
|[Fetcher](./internal/crawler/fetch.go)|Handles HTTP requests with retry and timeout management|
|[Parser](./internal/crawler/parser.go)|Extracts data using HTML tokenization|
|[Config](./pkg/config/config.go)|Loads and manages configuration from JSON files|
|[Logger](./pkg/logger/logger.go)|Handles structured, leveled logging|
|[Models](./internal/models/quote.go)|Defines shared data models (e.g. `Quote`)|

## 3. Environment Setup
|CATEGORY|DETAILS|
|:---:|:---:|
|OS|Windows 11/Linux/WSL|
|Go Version|go1.25.0|
|Editor|Visual Studio Code|

## 4. Configuration
All configuration is nanaged through `config/config.json` or automatically loaded via defaults defined in `pkg/config/config.go`.

### Example JSON file:
```json
{
  "crawler": {
    "base_url": "https://quotes.toscrape.com",
    "max_pages": 10,
    "concurrency": 5,
    "request_timeout": 10,
    "retry_attempts": 3,
    "rate_limit_per_sec": 10,
    "user_agent": "WebCrawler/1.0"
  },
  "logger": {
    "level": "info",
    "output_path": "logs/crawler.log",
    "max_size_mb": 100,
    "max_backups": 3,
    "max_age_days": 28
  }
}
```

## 5. Execution
### 5.1 Build the Crawler
```bash
go build -o crawler ./cmd/main
```

### 5.2 Run without build
```bash
go run ./cmd/main.go
```

## 6. Future Work
- Add persitent database integration.
- Implement distributed crawling via gRPC
- Add caching layer with Redis
- Integrate Prmetheus for metrics and health monitoring.

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 25-OCT-2025 18:37 UTC +08:00*
</div>