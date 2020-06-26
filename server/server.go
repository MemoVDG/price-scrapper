package main

import (
	"./models"

	"./controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBookByID)
	r.POST("/books", controllers.CreateBook)
	// PATCH is like PUT but only update the filds that it gets
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
