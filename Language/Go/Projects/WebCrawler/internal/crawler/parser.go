package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseHTML(html string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var results []string
	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		results = append(results, s.Text())
	})

	return results, nil
}
