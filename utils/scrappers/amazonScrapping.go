package scrappers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func AmazonScrapping(url string) (int, string) {
	// Create a collector
	collector := colly.NewCollector()
	var stringPrice, productTitle string

	// Search tag where the price is
	collector.OnHTML("span[class=a-price-whole]", func(e *colly.HTMLElement) {
		if stringPrice == "" {
			stringPrice = e.Text
		}
	})
	// Search tag where the product name is
	collector.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		productTitle = strings.TrimSpace(e.Text)
	})

	// Set error handler
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	collector.Visit(url)

	// Format price
	replacer := strings.NewReplacer(".", "", ",", "")
	priceFormated := replacer.Replace(stringPrice)
	// TODO: Handle error
	price, _ := strconv.Atoi(priceFormated)
	return price, productTitle
}
