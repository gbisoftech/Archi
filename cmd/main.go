package main

import (
	"main/config"
	"main/internal/models"
	"main/internal/routers"
)

func main() {

	config.ConnectDB()
	defer config.CloseDB()

	models.Init()

	router := routers.SetupRouter()
	router.Run(":8080")
}
