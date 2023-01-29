package util

import (
	"examples/microservices/config"
	"examples/microservices/models"
)

func CreateTables() {
	config.DB.AutoMigrate(&models.Todo{})
	config.DB.AutoMigrate(&models.User{})
}
