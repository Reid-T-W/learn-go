package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	//Load env variables
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
		log.Fatal("Error loading .env file")
	}
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DATABASE")
	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_password, database))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
