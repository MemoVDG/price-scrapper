package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	var regularPriceStr, discountPriceStr, productName string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	regularPriceStr, discountPriceStr, productName = price.GetPriceByStore(input.URL, input.Store)
	fmt.Println(regularPriceStr)

	regularPrice, errR := strconv.ParseFloat(regularPriceStr, 64)
	discountPrice, errD := strconv.ParseFloat(discountPriceStr, 64)
	if errR != nil || errD != nil {
		regularPrice = discountPrice
		log.Print(errR, errD)
	}

	// Create Product
	product := models.Product{
		URL:           input.URL,
		User:          input.User,
		RegularPrice:  regularPrice,
		DiscountPrice: discountPrice,
		Store:         input.Store,
		ProductName:   productName,
		Periodicity:   input.Periodicity,
	}
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
