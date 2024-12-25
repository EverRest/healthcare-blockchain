package main

import (
    "healthcare-blockchain/models"
	"healthcare-blockchain/api"
	"healthcare-blockchain/config"
	"healthcare-blockchain/database"
)

func main() {
	config.LoadConfig()
	database.Connect()
	database.Migrate(
		&models.User{},
		&models.BlockMetadata{},
	)
	router := api.SetupRouter()
    if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
        panic(err)
    }
	router.Run(":" + config.AppConfig.AppPort)
}
