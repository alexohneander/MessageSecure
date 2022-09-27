package main

import (
	"MessageSecure/config"
	"os"
)

func main() {

	// Configure Settings
	app := config.ConfigureSettings()

	// Configure Routes
	config.ConfigureRoutes(app)

	// Start server
	app.Listen(":" + os.Getenv("PORT"))
}
