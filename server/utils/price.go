package price

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

// AmazonProduct : Get the prices in the Amazon url
func AmazonProduct(url string) (string, string, string) {
	// Create a collector
	m := colly.NewCollector()
	// Regex to remove not number characteres
	var re = regexp.MustCompile("[$,]")
	var regularPrice, discountPrice, productName string

	// Find the div section with the price
	m.OnHTML("#price", func(e *colly.HTMLElement) {
		priceSection := e.DOM

		// Case when the regular price un underline and there is a discount
		if priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text() != "" {
			// Remove characteres that not are a number
			regularPrice = re.ReplaceAllString(strings.TrimSpace(priceSection.Find(".priceBlockStrikePriceString.a-text-strike").Text()), "")
			// TODO --
			// There are some products that you need to chose the size or some option
			// So amazon give you a range $2,814.19 - $3,574.81
			// Check if the price has more than one $ and the to the user tha select an option and give the new link
		}

		// Case when is the only prices and the product does not have discount
		if priceSection.Find("#priceblock_saleprice").Text() != "" {
			discountPrice = re.ReplaceAllString(strings.TrimSpace(priceSection.Find("#priceblock_saleprice").Text()), "")
		}
		if priceSection.Find("#priceblock_dealprice").Text() != "" {
			// Case when the element has a discount
			discountPrice = re.ReplaceAllString(strings.TrimSpace(priceSection.Find("#priceblock_dealprice").Text()), "")
		}

		if priceSection.Find("#priceblock_ourprice").Text() != "" {
			discountPrice = re.ReplaceAllString(strings.TrimSpace(priceSection.Find("#priceblock_ourprice").Text()), "")
		}

		fmt.Println(regularPrice, discountPrice)
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
	return regularPrice, discountPrice, productName
}

// MercadoLibreProduct : Get the prices in the ML url
func MercadoLibreProduct(url string) (string, string, string) {
	// Create a collector
	m := colly.NewCollector()
	// Regex to remove not number characteres
	var re = regexp.MustCompile("[$,]")
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
	// Remove characteres that not are a number
	regularPrice = re.ReplaceAllString(pricing[0], "")
	if len(pricing) > 1 {
		discountPrice = re.ReplaceAllString(pricing[1], "")
	} else {
		discountPrice = regularPrice
	}

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

	return regularPrice, discountPrice, productName
}

func getProductInformation(discountPrice, regularPrice, productName *string, url string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("//div[@class='m-product__price--collection']", chromedp.BySearch),
		chromedp.Evaluate("document.querySelector('p.a-product__paragraphRegularPrice.m-0.d-inline') ? document.querySelector('p.a-product__paragraphRegularPrice.m-0.d-inline').childNodes[1].nodeValue.replace(',','') + '' : 'False'", regularPrice),
		chromedp.Evaluate("document.querySelector('p.a-product__paragraphDiscountPrice.m-0.d-inline').childNodes[1].nodeValue.replace(',','') + '' ", discountPrice),
		chromedp.Evaluate("document.querySelector('h1.a-product__information--title').innerText", productName),
	}
}

// GetPriceByStore : Function used in the cron job to check the url based on the store
func GetPriceByStore(url, store string) (string, string, string) {

	switch store {
	case "Amazon":
		regularPrice, discountPrice, productName := AmazonProduct(url)
		return regularPrice, discountPrice, productName
	case "MercadoLibre":
		regularPrice, discountPrice, productName := MercadoLibreProduct(url)
		return regularPrice, discountPrice, productName
	case "Liverpool":
		regularPrice, discountPrice, productName := LiverpoolProduct(url)
		return regularPrice, discountPrice, productName
	default:
		return "", "", ""
	}
}
