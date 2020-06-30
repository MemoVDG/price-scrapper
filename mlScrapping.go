package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Create a collector
	m := colly.NewCollector()
	var pricing []string
	var productName string
	// Set HTML callback
	// Won't be called if error occurs
	m.OnHTML("span.price-tag-fraction", func(e *colly.HTMLElement) {
		pricing = append(pricing, e.Text)

	})

	m.OnHTML("h1.item-title__primary ", func(e *colly.HTMLElement) {
		productName = e.Text
	})

	// Set error handler
	m.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping
	m.Visit("https://articulo.mercadolibre.com.mx/MLM-626659518-nuevo-reloj-platinum-bloom-elegant-contra-agua-inoxidable-_JM?quantity=1&variation=32629264855&onAttributesExp=true#reco_item_pos=2&reco_backend=machinalis-pads-boost-pdp&reco_backend_type=low_level&reco_client=vip-pads-right&reco_id=43fb173d-e287-4ff7-be6d-0fa0d489a38b&is_advertising=true&ad_domain=VIPCORE_RIGHT&ad_position=3&ad_click_id=ZTQ2YWYzNGMtZWEzZi00NDM3LWEyOTUtMTNhOTkxOGNmYjMy")

	fmt.Println(pricing)
	fmt.Println(productName)

}
