package models

import (
	"main/config"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	DB = config.DB
	DB.AutoMigrate(&User{}, &Book{})
}
