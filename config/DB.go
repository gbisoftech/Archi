package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dsn = "root:@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

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
