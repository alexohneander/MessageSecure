package main

import (
	"MessageSecure/config"
	"log"
	"os"
)

func main() {

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
