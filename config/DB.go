package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	ServerPort string
	dsn        string
)

func ReadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	ServerPort = os.Getenv("SERVER_PORT")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("PASSWORD")

	dsn = user + "@" + password + "tcp(" + host + ":" + port + ")/mysql?charset=utf8mb4&parseTime=True&loc=Local"
}

func ConnectDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func CloseDB() {
	db, err := DB.DB()
	if err != nil {
		panic("failed to close database")
	}
	db.Close()
}
