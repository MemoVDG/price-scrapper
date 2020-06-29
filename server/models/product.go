// Product model

package models

// Product : Product structure
type Product struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	URL           string `json:"url"`
	Store         string `json:"store"`
	RegularPrice  string `json:"regularPrice"`
	DiscountPrice string `json:"oldPrice"`
	User          string `json:"user"`
}

// CreateProduct : Basic structure for a new product
type CreateProduct struct {
	URL   string `json:"url" binding:"required"`
	Store string `json:"store" binding:"required"`
	User  string `json:"user"`
}
