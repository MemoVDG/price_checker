package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Create a collector
	m := colly.NewCollector()
	var price string
	//var re = regexp.MustCompile("[$,]")
	m.OnHTML("span[class=a-price-whole]", func(e *colly.HTMLElement) {
		if price == "" {
			price = e.Text
		}
	})

	m.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(e.Text))
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit("https://www.amazon.com.mx/Bluelander-Escritorio-Enfriamiento-Alfombrillas-Antideslizantes/dp/B09LC61DJJ/ref=pd_sbs_5/140-4683112-3107017?pd_rd_w=FqdUv&pf_rd_p=b7b7182e-de22-409a-8f7d-0db823b6f8b4&pf_rd_r=VSME89BC8H0XJ6GX5TK5&pd_rd_r=7c1cc6f2-6010-4948-bbaf-d3203061115e&pd_rd_wg=waQks&pd_rd_i=B09LC61DJJ&psc=1")
	replacer := strings.NewReplacer(".", "", ",", "")
	t := replacer.Replace(price)
	intval, _ := strconv.Atoi(t)
	fmt.Print(intval)
}
