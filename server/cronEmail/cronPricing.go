package cronEmail

import (
	"fmt"
	"log"
	"strconv"

	"../models"
	price "../utils"
	"github.com/robfig/cron"
)

// CronJobs : It'll run in paralell with the server
func CronJobs() {

	c := cron.New()
	c.AddFunc("@every 0h0m30s", func() {
		var products []models.Product
		if err := models.DB.Where("periodicity = ?", 5).Find(&products).Error; err != nil {
			log.Println("Error getting the data for this time periodicity")
		}

		fmt.Print(len(products))
		for _, product := range products {
			// Find user
			//var user models.User

			// Check the product
			regularPriceStr, discountPriceStr, _ := price.GetPriceByStore(product.URL, product.Store)
			regularPrice, errR := strconv.ParseFloat(regularPriceStr, 64)
			discountPrice, errD := strconv.ParseFloat(discountPriceStr, 64)
			if errR != nil || errD != nil {
				regularPrice = discountPrice
				log.Print(errR, errD)
			}

			// TODO -- Send email to the user, to notify that price's product change
			if regularPrice < product.RegularPrice || discountPrice < product.DiscountPrice {
				SendEmail("memovdg@gmail.com", product.ProductName)
			}

			SendEmail("memovdg@gmail.com", product.ProductName)

			// Update price
			product.DiscountPrice = discountPrice
			product.RegularPrice = regularPrice

			models.DB.Save(&product)

		}
	})
	c.Start()
	select {}

}
