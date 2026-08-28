package main

import (
	"backend-api/config"
	"backend-api/database"
	"backend-api/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
