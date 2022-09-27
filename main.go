package main

import (
	"MessageSecure/config"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// get Env File
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file")
	}

	// Configure Settings
	app := config.ConfigureSettings()

	// Configure Routes
	config.ConfigureRoutes(app)

	// Start server
	startErr := app.Listen(":" + os.Getenv("PORT"))
	if startErr != nil {
		log.Fatal(startErr)
	}
}
