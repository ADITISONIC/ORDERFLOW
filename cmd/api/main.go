package main

import (
	"orderflow/cache"
	"orderflow/config"
	"orderflow/consumer"
	"orderflow/database"
	"orderflow/routes"
)

func main() {

	config.LoadEnv()
	config.ConnectDB()
	cache.ConnectRedis()
	database.Migrate()
	go consumer.StartConsumer()
	router := routes.SetupRouter()


	router.Run(":8080")
}
