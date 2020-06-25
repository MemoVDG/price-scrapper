package main

import (
	"fmt"
	"github.com/gocolly/colly"
)


func main() {
	// Create a collector
	m := colly.NewCollector()
	var regularPrice, discountPrice string

	// Find the div section with the price
	m.OnHTML("#price", func(e *colly.HTMLElement) {
		priceSection := e.DOM
		
		// Case when the regular price un underline and there is a discount
		if (priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text() != "") {
			regularPrice = priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text()
			
			// TODO -- 
			// There are some products that you need to chose the size or some option
			// So amazon give you a range $2,814.19 - $3,574.81
			// Check if the price has more than one $ and the to the user tha select an option and give the new link
		}

		// Case when is the only prices and the product does not have discount
		if (priceSection.Find("#priceblock_ourprice").Text() != "") {
			regularPrice = priceSection.Find("#priceblock_ourprice").Text()
		}

		// Case when the element has a discount
		discountPrice = priceSection.Find("#priceblock_dealprice").Text()
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit("https://www.amazon.com.mx/dp/B07SVBBP2B?tag=jalme-20&th=1&psc=1")
	fmt.Print(regularPrice, discountPrice)
}