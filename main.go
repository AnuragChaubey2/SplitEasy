package main

import (
	"log"

	"spliteasy/internal/config"
	"spliteasy/internal/database"
	"spliteasy/internal/router"
)

func main() {
	config.LoadEnv()
	database.InitDB()
	log.Println("App bootstrapped.")

	r := router.InitRouter()
	r.Run(":8080")

}
