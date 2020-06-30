package price

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

// AmazonProduct : Get the prices in the Amazon url
func AmazonProduct(url string) (string, string, string) {
	// Create a collector
	m := colly.NewCollector()
	var regularPrice, discountPrice, productName string

	// Find the div section with the price
	m.OnHTML("#price", func(e *colly.HTMLElement) {
		priceSection := e.DOM

		// Case when the regular price un underline and there is a discount
		if priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text() != "" {
			regularPrice = priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text()

			// TODO --
			// There are some products that you need to chose the size or some option
			// So amazon give you a range $2,814.19 - $3,574.81
			// Check if the price has more than one $ and the to the user tha select an option and give the new link
		}

		// Case when is the only prices and the product does not have discount
		if priceSection.Find("#priceblock_ourprice").Text() != "" {
			regularPrice = priceSection.Find("#priceblock_ourprice").Text()
		}

		// Case when the element has a discount
		discountPrice = priceSection.Find("#priceblock_dealprice").Text()
	})

	m.OnHTML("#productTitle", func(e *colly.HTMLElement) {
		productName = strings.TrimSpace(e.Text)
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit(url)
	fmt.Println(regularPrice, discountPrice)
	return regularPrice, discountPrice, productName
}

// MercadoLibreProduct : Get the prices in the ML url
func MercadoLibreProduct(url string) (string, string, string) {
	// Create a collector
	m := colly.NewCollector()
	var regularPrice, discountPrice string
	var pricing []string
	var productName string

	// Set HTML callback
	// Won't be called if error occurs
	m.OnHTML("span.price-tag-fraction", func(e *colly.HTMLElement) {
		pricing = append(pricing, e.Text)
	})

	// Get the name of the product
	m.OnHTML("h1.item-title__primary ", func(e *colly.HTMLElement) {
		text := strings.TrimSpace(e.Text)
		productName = text
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit(url)
	regularPrice = pricing[0]
	discountPrice = pricing[1]

	return regularPrice, discountPrice, productName

}

// LiverpoolProduct : Get the prices in the Liverpool url
func LiverpoolProduct(url string) (string, string, string) {

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", false),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// run task list
	var discountPrice, regularPrice, productName string
	err := chromedp.Run(ctx, getProductInformation(&discountPrice, &regularPrice, &productName, url))
	if err != nil {
		log.Fatal(err)
	}

	if regularPrice == "False" {
		regularPrice = discountPrice
	}

	fmt.Print(regularPrice, discountPrice)
	return regularPrice, discountPrice, productName
}

func getProductInformation(discountPrice, regularPrice, productName *string, url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("//div[@class='m-product__price--collection']", chromedp.BySearch),
		chromedp.Evaluate("document.querySelector('p.a-product__paragraphRegularPrice.m-0.d-inline') ? document.querySelector('p.a-product__paragraphRegularPrice.m-0.d-inline').innerText : 'False'", regularPrice),
		chromedp.Evaluate("document.querySelector('p.a-product__paragraphDiscountPrice.m-0.d-inline').innerText", discountPrice),
		chromedp.Evaluate("document.querySelector('h1.a-product__information--title').innerText", productName),
	}
}
