package main

import (
	"MessageSecure/config"
)

func main() {

	// Configure Settings
	app := config.ConfigureSettings()

	// Configure Routes
	config.ConfigureRoutes(app)

	// Start server
	app.Listen(":3000")
}
