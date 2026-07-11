package database

import (
	"orderflow/config"
	"orderflow/models"
)

func Migrate() {

	config.DB.AutoMigrate(
		&models.User{},
		&models.Order{},
	)
}