package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB : DB object
var DB *gorm.DB

// ConnectDataBase : Starting the DB
func ConnectDataBase() {

	dbData := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PSW"))
	fmt.Println(dbData)
	database, err := gorm.Open("postgres", dbData)

	if err != nil {
		log.Println(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Book{})
	database.AutoMigrate(&Product{})

	DB = database
}
