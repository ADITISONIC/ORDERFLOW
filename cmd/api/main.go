package main

import (
	

	"orderflow/config"
	"orderflow/database"
	"orderflow/routes"

)

func main() {

	config.LoadEnv()
	config.ConnectDB()
	database.Migrate()
	router := routes.SetupRouter()

	

	router.Run(":8080")
}
