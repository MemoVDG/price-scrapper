// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
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
	err := chromedp.Run(ctx, getProductInformation(&discountPrice, &regularPrice, &productName))
	if err != nil {
		log.Fatal(err)
	}

	if regularPrice == "False" {
		regularPrice = discountPrice
	}

	fmt.Print(regularPrice, discountPrice, productName)

}

func getProductInformation(discountPrice, regularPrice, productName *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("https://www.liverpool.com.mx/tienda/pdp/bolsa-crossbody-steve-madden-amarilla-capitonada/1080852653"),
		//chromedp.Navigate("https://www.liverpool.com.mx/tienda/pdp/horno-de-microondas-ge-profile-2.2-pies-cúbicos-acero-peb7227andd/1088752577"),
		chromedp.WaitVisible("//div[@class='m-product__price--collection']", chromedp.BySearch),
		chromedp.Evaluate("document.querySelector('p.a-product__paragraphRegularPrice.m-0.d-inline') ? document.querySelector('p.a-product__paragraphRegularPrice.m-0.d-inline').innerText : 'False'", regularPrice),
		chromedp.Evaluate("document.querySelector('p.a-product__paragraphDiscountPrice.m-0.d-inline').innerText", discountPrice),
		chromedp.Evaluate("document.querySelector('h1.a-product__information--title').innerText", productName),
	}
}
