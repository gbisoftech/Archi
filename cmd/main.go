package main

import (
	"main/config"
	"main/internal/models"
	"main/internal/routers"
)

func main() {

	config.ReadEnv()
	config.ConnectDB()
	defer config.CloseDB()

	models.Init()

	router := routers.SetupRouter()
	router.Run(":" + config.ServerPort)
}
