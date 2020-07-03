package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	// Create a collector
	m := colly.NewCollector()
	var regularPrice, discountPrice string
	var re = regexp.MustCompile("[$,]")

	// Find the div section with the price
	m.OnHTML("#price", func(e *colly.HTMLElement) {
		priceSection := e.DOM
		// Case when the regular price un underline and there is a discount
		if priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text() != "" {
			regularPrice = re.ReplaceAllString(priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text(), "")

			// TODO --
			// There are some products that you need to chose the size or some option
			// So amazon give you a range $2,814.19 - $3,574.81
			// Check if the price has more than one $ and the to the user tha select an option and give the new link
		}

		// Case when is the only prices and the product does not have discount
		if priceSection.Find("#priceblock_ourprice").Text() != "" {
			regularPrice = re.ReplaceAllString(priceSection.Find("#priceblock_ourprice").Text(), "")
		}

		// Case when the element has a discount
		discountPrice = re.ReplaceAllString(priceSection.Find("#priceblock_dealprice").Text(), "")
	})

	m.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(e.Text))
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit("https://www.amazon.com.mx/TOPLIVING-Ejecutiva-Operativa-Ergonomica-Giratoria/dp/B089Y73LWP/ref=sr_1_1?__mk_es_MX=ÅMÅŽÕÑ&dchild=1&pf_rd_i=gb_main&pf_rd_m=AVDBXBAVVSXLQ&pf_rd_p=d7c672db-636e-444e-9402-8d9f57de08af&pf_rd_r=H8VRM0NC9J1ZMJ0BTYNF&pf_rd_s=slot-5&pf_rd_t=701&qid=1593530739&smid=A3QC1CWBMD7YSQ&sr=8-1")
	fmt.Print(regularPrice, discountPrice)
}
