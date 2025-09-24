package crawler

import (
	"WebCrawler/internal/models"
	"WebCrawler/pkg/logger"
	"fmt"
)

func Start(url string) []models.Quote {
	logger.Info(fmt.Sprintf("Fetch: %s.", url))

	html, err := Fetch(url)
	if err != nil {
		logger.Error(fmt.Sprintf("Fetch Error: %v.", err.Error()))
		return nil
	}

	results, err := ParseHTML(html)
	if err != nil {
		logger.Error(fmt.Sprintf("Parse Error: %v.", err.Error()))
		return nil
	}

	return results
}
