package crawler

import (
	"WebCrawler/pkg/config"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

//
// ======== HTTP Fetcher ========
//

// Fetcher is responsible for making HTTP requests and retrieving HTML content.
//
// It supports configurable timeouts, retry mechanisms with exponential backoff,
// and customizable User-Agent headers. Fetcher should be reused across multiple
// requests for connection pooling efficiency.
type Fetcher struct {
	client     *http.Client
	maxRetries int
	retryDelay time.Duration
	userAgent  string
}

//
// ======== Constructor ========
//

// NewFetcher creates and configures a new Fetcher instance.
//
// It initializes an HTTP client with connection pooling,
// timeouts, and retry configuration derived from the crawler configuration.
//
// Args:
//
//	cfg: Pointer to a CrawlerConfig struct providing timeout and retry settings.
//
// Returns:
//
//	A pointer to an initialized Fetcher ready for HTTP requests.
func NewFetcher(cfg *config.CrawlerConfig) *Fetcher {
	return &Fetcher{
		client: &http.Client{
			Timeout: cfg.RequestTimeout,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		maxRetries: cfg.RetryAttempts,
		retryDelay: 2 * time.Second,
		userAgent:  cfg.UserAgent,
	}
}

//
// ======== Core Fetch Methods ========
//

// Fetch retrieves the raw HTML content of the specified URL.
//
// It builds an HTTP GET request with a context, applies standard headers,
// and reads the response body into a string. The caller is responsible
// for providing a cancellable context to prevent long-running requests.
//
// Args:
//
//	ctx: Context for timeout and cancellation control.
//	url: Target URL to fetch.
//
// Returns:
//
//	A string containing the HTML content, or an error if the request fails.
func (f *Fetcher) Fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set request headers for standard web crawling compatibility.
	req.Header.Set("User-Agent", f.userAgent)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := f.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

// FetchWithRetry performs an HTTP fetch operation with automatic retries.
//
// It uses an exponential backoff strategy between retry attempts
// and respects the provided context for cancellation.
//
// Args:
//
//	ctx: Context for timeout and cancellation control.
//	url: Target URL to fetch.
//
// Returns:
//
//	The HTML content string, or an error if all retry attempts fail.
func (f *Fetcher) FetchWithRetry(ctx context.Context, url string) (string, error) {
	var lastErr error

	for attempt := 0; attempt <= f.maxRetries; attempt++ {
		if attempt > 0 {
			// Apply exponential backoff delay.
			delay := f.retryDelay * time.Duration(1<<uint(attempt-1))
			select {
			case <-ctx.Done():
				return "", ctx.Err()
			case <-time.After(delay):
			}
		}

		html, err := f.Fetch(ctx, url)
		if err == nil {
			return html, nil
		}

		lastErr = err

		if !shouldRetry(err) {
			break
		}
	}

	return "", fmt.Errorf("fetch failed after %d retries: %w", f.maxRetries, lastErr)
}

//
// ======== Retry Logic ========
//

// shouldRetry determines whether an error is eligible for retry.
//
// In this implementation, all non-nil errors trigger retries,
// but this function can be extended to inspect error types,
// status codes, or network conditions.
//
// Args:
//
//	err: The error encountered during an HTTP request.
//
// Returns:
//
//	True if the error should trigger a retry; false otherwise.
func shouldRetry(err error) bool {
	if err == nil {
		return false
	}
	// Example: only retry on transient network errors
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return false
	}
	return true
}

//
// ======== Cleanup ========
//

// Close releases idle HTTP connections.
//
// It should be called when the Fetcher is no longer needed to
// properly close connection pools and free resources.
func (f *Fetcher) Close() {
	f.client.CloseIdleConnections()
}
