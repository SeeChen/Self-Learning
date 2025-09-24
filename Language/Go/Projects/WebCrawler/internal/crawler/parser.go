package crawler

import (
	"WebCrawler/internal/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(html string) ([]models.Quote, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var results []models.Quote
	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		var content string = s.Find(".text").Text()
		var author string = s.Find(".author").Text()

		var listTag []string

		var tags = s.Find(".tag")
		tags.Each(func(i int, tag *goquery.Selection) {
			listTag = append(listTag, tag.Text())
		})

		results = append(results, models.Quote{
			Content: content,
			Author:  author,
			Tags:    listTag,
		})
	})

	return results, nil
}
