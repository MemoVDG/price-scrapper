package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Create a collector
	m := colly.NewCollector()

	// Set HTML callback
	// Won't be called if error occurs
	m.OnHTML("span.price-tag-fraction", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit("https://articulo.mercadolibre.com.mx/MLM-775959840-cubrebocas-kn95-certificado-mascarilla-tapabocas-1pz-n95-_JM")
}
