package main

import (
	"WebCrawler/internal/crawler"
	"WebCrawler/pkg/logger"
	"fmt"
)

func main() {
	logger.Info("Starting Web Crawler.")

	for i := 1; i <= 1; i++ {

		var url string = fmt.Sprintf("https://quotes.toscrape.com/page/%d/", i)
		logger.Info(url)

		results := crawler.Start(url)

		for _, item := range results {
			fmt.Print("=======")
			fmt.Println(item)
		}
	}
}
