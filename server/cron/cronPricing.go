package main

import (
	"log"

	"../models"
	price "../utils"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {

	c := cron.New()
	models.ConnectDataBase("./test.db")
	c.AddFunc("@every 0h0m10s", func() {
		var products []models.Product
		if err := models.DB.Where("periodicity = ?", 5).Find(&products).Error; err != nil {
			log.Println("Error getting the data for this time periodicity")

		}
		log.Println("Products found")
		for _, product := range products {

			regularPrice, discountPrice, _ := price.GetPriceByStore(product.URL, product.Store)

			// TODO -- Send email to the user, to notify that price's product change
			if regularPrice < product.RegularPrice || discountPrice < product.DiscountPrice {
				log.Println("Price Change")
			}

			// Update price

		}
	})
	c.Start()
	select {}

}
