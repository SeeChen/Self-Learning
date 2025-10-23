package crawler

import (
	"WebCrawler/internal/models"
	"WebCrawler/pkg/config"
	"WebCrawler/pkg/logger"
	"context"
	"fmt"
	"sync"
	"time"
)

//
// ======== Core Structures ========
//

// Crawler represents the main crawling engine.
//
// It coordinates page fetching, HTML parsing, rate-limiting,
// concurrency control, and statistical tracking. A Crawler
// instance is typically configured through a CrawlerConfig
// object loaded from application settings.
type Crawler struct {
	config      *config.CrawlerConfig
	fetcher     *Fetcher
	parser      *Parser
	rateLimiter *RateLimiter
	stats       *Statistics
}

// Statistics holds runtime metrics collected during a crawl session.
//
// These include page success/failure counts, quote totals,
// latency metrics, and start/end timestamps.
type Statistics struct {
	mu             sync.RWMutex
	TotalPages     int
	SuccessPages   int
	FailedPages    int
	TotalQuotes    int
	StartTime      time.Time
	EndTime        time.Time
	AverageLatency time.Duration
	totalLatency   time.Duration
}

//
// ======== Crawler Lifecycle ========
//

// NewCrawler creates and initializes a new Crawler instance.
//
// It constructs internal components such as the Fetcher,
// Parser, RateLimiter, and Statistics tracker.
func NewCrawler(cfg *config.CrawlerConfig) *Crawler {
	return &Crawler{
		config:      cfg,
		fetcher:     NewFetcher(cfg),
		parser:      NewParser(),
		rateLimiter: NewRateLimiter(cfg.RateLimitPerSec),
		stats: &Statistics{
			StartTime: time.Now(),
		},
	}
}

// Start launches the crawler for a specified page range.
//
// This method spawns a pool of worker goroutines to process
// multiple pages concurrently, while respecting rate limits
// and graceful cancellation through the provided context.
func (c *Crawler) Start(ctx context.Context, startPage, endPage int) ([]models.Quote, error) {
	logger.Info("Crawler engine started, page range: %d-%d", startPage, endPage)
	c.stats.StartTime = time.Now()

	defer c.rateLimiter.Stop()
	defer c.fetcher.Close()

	var (
		allQuotes []models.Quote
		mu        sync.Mutex
		wg        sync.WaitGroup
	)

	// Create communication channels.
	tasks := make(chan int, endPage-startPage+1)
	results := make(chan []models.Quote, endPage-startPage+1)
	errors := make(chan error, endPage-startPage+1)

	// Launch worker goroutines.
	for i := 0; i < c.config.Concurrency; i++ {
		wg.Add(1)
		go c.worker(ctx, &wg, tasks, results, errors)
	}

	// Feed page numbers into the task channel.
	go func() {
		defer close(tasks)
		for page := startPage; page <= endPage; page++ {
			select {
			case <-ctx.Done():
				return
			case tasks <- page:
			}
		}
	}()

	// Wait for workers to finish and close result channels.
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	// Collect results until both channels are closed.
	for results != nil || errors != nil {
		select {
		case quotes, ok := <-results:
			if !ok {
				results = nil
				continue
			}
			mu.Lock()
			allQuotes = append(allQuotes, quotes...)
			c.stats.mu.Lock()
			c.stats.TotalQuotes += len(quotes)
			c.stats.mu.Unlock()
			mu.Unlock()

		case err, ok := <-errors:
			if !ok {
				errors = nil
				continue
			}
			logger.Error("Crawling error: %v", err)
		}
	}

	c.stats.EndTime = time.Now()
	c.printStats()

	return allQuotes, nil
}

//
// ======== Internal Helpers ========
//

// worker is a goroutine responsible for processing individual pages.
//
// It waits for rate-limiter tokens, performs the fetch-and-parse
// sequence, records success/failure metrics, and forwards results
// or errors through dedicated channels.
func (c *Crawler) worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	tasks <-chan int,
	results chan<- []models.Quote,
	errors chan<- error,
) {
	defer wg.Done()

	for page := range tasks {
		select {
		case <-ctx.Done():
			return
		default:
			// Wait for rate limiter token with context awareness.
			if err := c.rateLimiter.Wait(ctx); err != nil {
				errors <- fmt.Errorf("rate limiter canceled: %w", err)
				return
			}

			quotes, err := c.crawlPage(ctx, page)
			if err != nil {
				errors <- fmt.Errorf("page %d crawl failed: %w", page, err)
				c.stats.mu.Lock()
				c.stats.FailedPages++
				c.stats.mu.Unlock()
				continue
			}

			results <- quotes
			c.stats.mu.Lock()
			c.stats.SuccessPages++
			c.stats.mu.Unlock()
		}
	}
}

// crawlPage fetches and parses a single page from the target site.
//
// It constructs the URL, fetches HTML content (with retry),
// parses quotes, and records latency statistics.
func (c *Crawler) crawlPage(ctx context.Context, page int) ([]models.Quote, error) {
	url := fmt.Sprintf("%s/page/%d/", c.config.BaseURL, page)
	start := time.Now()
	logger.Debug("Fetching: %s", url)

	html, err := c.fetcher.FetchWithRetry(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch HTML: %w", err)
	}

	quotes, err := c.parser.Parse(html)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	latency := time.Since(start)
	c.stats.mu.Lock()
	c.stats.TotalPages++
	c.stats.totalLatency += latency
	c.stats.AverageLatency = c.stats.totalLatency / time.Duration(c.stats.TotalPages)
	c.stats.mu.Unlock()

	logger.Info("Page %d crawled successfully: %d quotes, duration %v", page, len(quotes), latency)
	return quotes, nil
}

// GetStatistics returns a snapshot copy of the current statistics.
//
// The returned value is a deep copy to prevent concurrent data races.
func (c *Crawler) GetStatistics() *Statistics {
	c.stats.mu.RLock()
	defer c.stats.mu.RUnlock()

	statsCopy := *c.stats
	return &statsCopy
}

// printStats logs a summary of crawl statistics.
//
// This includes total pages, successes/failures, quote count,
// total runtime, and average latency.
func (c *Crawler) printStats() {
	c.stats.mu.RLock()
	defer c.stats.mu.RUnlock()

	duration := c.stats.EndTime.Sub(c.stats.StartTime)
	logger.Info("=== Crawl Statistics ===")
	logger.Info("Total pages: %d", c.stats.TotalPages)
	logger.Info("Succeeded: %d", c.stats.SuccessPages)
	logger.Info("Failed: %d", c.stats.FailedPages)
	logger.Info("Total quotes: %d", c.stats.TotalQuotes)
	logger.Info("Duration: %v", duration)
	logger.Info("Avg latency: %v", c.stats.AverageLatency)
	if duration > 0 {
		logger.Info("Crawl rate: %.2f pages/sec", float64(c.stats.TotalPages)/duration.Seconds())
	}
}

//
// ======== Rate Limiter ========
//

// RateLimiter enforces a maximum request rate using a token bucket.
//
// Tokens are replenished at a fixed interval determined by maxRate.
type RateLimiter struct {
	ticker   *time.Ticker
	tokens   chan struct{}
	maxRate  int
	stopChan chan struct{}
}

// NewRateLimiter creates a new RateLimiter.
func NewRateLimiter(ratePerSec int) *RateLimiter {
	if ratePerSec <= 0 {
		ratePerSec = 1
	}

	rl := &RateLimiter{
		ticker:   time.NewTicker(time.Second / time.Duration(ratePerSec)),
		tokens:   make(chan struct{}, ratePerSec),
		maxRate:  ratePerSec,
		stopChan: make(chan struct{}),
	}

	// Pre-fill the token bucket.
	for i := 0; i < ratePerSec; i++ {
		rl.tokens <- struct{}{}
	}

	// Refill tokens periodically.
	go func() {
		for {
			select {
			case <-rl.ticker.C:
				select {
				case rl.tokens <- struct{}{}:
				default:
				}
			case <-rl.stopChan:
				rl.ticker.Stop()
				return
			}
		}
	}()

	return rl
}

// Wait blocks until a token is available or context is canceled.
func (rl *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-rl.tokens:
		return nil
	}
}

// Stop terminates the rate limiter and stops token replenishment.
//
// This should be called during shutdown to release resources.
func (rl *RateLimiter) Stop() {
	select {
	case <-rl.stopChan:
		// already closed
	default:
		close(rl.stopChan)
	}
}
