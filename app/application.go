package application

import (
	"log"

	"github.com/joho/godotenv"
	server "github.com/vkhoa145/facebook-mini-api/app/server"
	"github.com/vkhoa145/facebook-mini-api/config"
)

func StartApp() {
	// err := godotenv.Load("/Users/vodangkhoa/Documents/Projects/facebook-mini-api/.env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := server.NewFiber()
	config := config.LoadConfig()
	server := server.NewServer(app, config)
	error := server.Start()
	if error != nil {
		log.Fatal("Error starting server: ", error)
	}
}
