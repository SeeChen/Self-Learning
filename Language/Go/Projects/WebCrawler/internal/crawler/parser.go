package crawler

import (
	"WebCrawler/internal/models"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//
// ======== HTML Parser ========
//

// Parser is responsible for extracting structured quote data from HTML pages.
//
// It uses CSS selectors and the goquery library to traverse the DOM
// and extract content, author names, and tags for each quote element.
type Parser struct {
	selector string
}

//
// ======== Constructor ========
//

// NewParser creates a new Parser instance with a default CSS selector.
//
// The default selector targets elements with the `.quote` class,
// which are used on `quotes.toscrape.com` to represent individual quotes.
//
// Returns:
//
//	A pointer to an initialized Parser.
func NewParser() *Parser {
	return &Parser{
		selector: ".quote",
	}
}

//
// ======== Core Parsing Logic ========
//

// Parse extracts all quotes from an HTML document string.
//
// It parses the input HTML using goquery, identifies quote elements
// based on the configured CSS selector, and returns a slice of Quote structs.
//
// Args:
//
//	html: Raw HTML content to be parsed.
//
// Returns:
//
//	A slice of Quote objects parsed from the HTML, or an error if parsing fails.
//
// Errors:
//   - Returns an error if the HTML is empty or invalid.
//   - Returns an error if no quote elements are found.
func (p *Parser) Parse(html string) ([]models.Quote, error) {
	if html == "" {
		return nil, fmt.Errorf("HTML content is empty")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTML document: %w", err)
	}

	var quotes []models.Quote

	doc.Find(p.selector).Each(func(i int, s *goquery.Selection) {
		quote := p.parseQuote(s)
		if quote != nil {
			quotes = append(quotes, *quote)
		}
	})

	if len(quotes) == 0 {
		return nil, fmt.Errorf("no quotes found in the HTML")
	}

	return quotes, nil
}

//
// ======== Helper Methods ========
//

// parseQuote extracts structured data from a single quote element.
//
// It reads the text, author name, and associated tags from the DOM element.
// If required fields (content or author) are missing, the function returns nil.
//
// Args:
//
//	s: goquery Selection representing a single `.quote` element.
//
// Returns:
//
//	A pointer to a Quote struct, or nil if parsing fails.
func (p *Parser) parseQuote(s *goquery.Selection) *models.Quote {
	content := strings.TrimSpace(s.Find(".text").Text())
	author := strings.TrimSpace(s.Find(".author").Text())

	if content == "" || author == "" {
		return nil
	}

	// Extract tags associated with the quote.
	var tags []string
	s.Find(".tag").Each(func(i int, tag *goquery.Selection) {
		tagText := strings.TrimSpace(tag.Text())
		if tagText != "" {
			tags = append(tags, tagText)
		}
	})

	return &models.Quote{
		Content: content,
		Author:  author,
		Tags:    tags,
	}
}

//
// ======== Configuration ========
//

// SetSelector updates the CSS selector used to identify quote elements.
//
// This allows users to customize the parser for different HTML structures.
//
// Args:
//
//	selector: A CSS selector string used to locate quote elements.
func (p *Parser) SetSelector(selector string) {
	p.selector = selector
}

// ParseWithSelector temporarily overrides the parser’s selector
// to extract quotes using a custom CSS selector.
//
// It restores the original selector after parsing completes.
//
// Args:
//
//	html: HTML content to parse.
//	selector: CSS selector to use for this parsing operation.
//
// Returns:
//
//	A slice of Quote objects extracted with the custom selector, or an error if parsing fails.
func (p *Parser) ParseWithSelector(html, selector string) ([]models.Quote, error) {
	originalSelector := p.selector
	p.selector = selector
	defer func() { p.selector = originalSelector }()

	return p.Parse(html)
}
