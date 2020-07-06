// Product model

package models

// Product : Product structure
type Product struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	UserID        uint    `json:"userID"`
	URL           string  `json:"url"`
	ProductName   string  `json:"productName"`
	Store         string  `json:"store"`
	RegularPrice  float64 `json:"regularPrice"`
	DiscountPrice float64 `json:"discountPrice"`
	Periodicity   int     `json:"periodicity"`
	UserEmail     string  `json:"userEmail"`
}

// CreateProduct : Basic structure for a new product
type CreateProduct struct {
	URL         string `json:"url" binding:"required"`
	Store       string `json:"store" binding:"required"`
	UserID      uint   `json:"userID"`
	Periodicity int    `json:"periodicity" binding:"required"`
	UserEmail   string `json:"userEmail"`
}
