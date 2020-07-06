package main

import (
	"log"

	"./controllers"
	"./cronEmail"
	"./models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init gets called before the main function
func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	r := gin.Default()
	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBookByID)
	r.POST("/books", controllers.CreateBook)
	// PATCH is like PUT but only update the filds that it gets
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// API Products
	r.GET("/products", controllers.FindProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.FindProductByID)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	go func() {
		cronEmail.CronJobs()
	}()
	r.Run()
}
