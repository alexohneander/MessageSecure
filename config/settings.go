package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/template/html"
)

func ConfigureSettings() *fiber.App {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		// Views Layout is the global layout for all template render until override on Render function.
		ViewsLayout: "layouts/main",
	})

	// Set Public Folder
	app.Static("/", "./public")

	// Set Favicon
	app.Use(favicon.New())

	// Provide a custom compression level
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))

	return app
}
