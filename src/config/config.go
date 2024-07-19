package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPass := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errorDb != nil {
		panic("Failed to connect mysql database")
	}

	return db
}

func DisconnectDb(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}

	dbSQL.Close()
}
