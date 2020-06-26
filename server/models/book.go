// models/book.go

package models

// Book : Basic book structure
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// CreateBookInput : Validate the new new input
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// UpdateBookInput : Validate the update data
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
