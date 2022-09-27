package config

import (
	"MessageSecure/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func ConfigureRoutes(app *fiber.App) {

	app.Get("/", controllers.RenderIndex)
	app.Post("/message/encrypt", controllers.EncryptData)
	app.Get("/message/decrypt/:id/:token", controllers.DecryptData)

	// Provide a minimal config
	// app.Use(basicauth.New(basicauth.Config{
	// 	Users: map[string]string{
	// 		"admin": "Monitoring2442!*", // demo!
	// 	},
	// }))

	// Monitoring Dashboard
	app.Get("/dashboard", monitor.New())
}
