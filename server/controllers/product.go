package controllers

import (
	"fmt"
	"net/http"

	"../models"
	price "../utils"
	"github.com/gin-gonic/gin"
)

// FindProducts : List all the products
func FindProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// CreateProduct : Add a new product to the DB
func CreateProduct(c *gin.Context) {
	var input models.CreateProduct
	var regularPrice, discountPrice string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Analize product
	if input.Store == "Amazon" {
		regularPrice, discountPrice = price.AmazonProduct(input.URL)
	} else if input.Store == "MercadoLibre" {
		regularPrice, discountPrice = price.MercadoLibreProduct(input.URL)
	} else {
		regularPrice, discountPrice = price.LiverpoolProduct(input.URL)
	}

	fmt.Println(regularPrice, discountPrice)
	// Create Product
	product := models.Product{URL: input.URL, User: input.User, RegularPrice: regularPrice, DiscountPrice: discountPrice, Store: input.Store}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// FindProductByID : Find product using the ID
func FindProductByID(c *gin.Context) {
	var product models.Product

	if err := models.DB.Where("id= ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not find"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product, "message": "success"})
}

// DeleteProduct : Remove product record from the DB
func DeleteProduct(c *gin.Context) {
	// Check if element exist
	var product models.Product
	if err := models.DB.Where("id= ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
