package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

// ConnectDataBase : Starting the DB
func ConnectDataBase(dbPath string) {
	database, err := gorm.Open("sqlite3", dbPath)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})
	database.AutoMigrate(&Product{})

	DB = database
}
