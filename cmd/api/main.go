// @title OrderFlow API
// @version 1.0
// @description Event-driven Order Processing System using Go, Kafka, Redis & MySQL.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"orderflow/cache"
	"orderflow/config"
	"orderflow/consumer"
	"orderflow/database"
	"orderflow/routes"
	_ "orderflow/docs"
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
