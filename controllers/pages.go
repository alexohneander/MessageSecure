package controllers

import "github.com/gofiber/fiber/v2"

func RenderIndex(c *fiber.Ctx) error {

	token, err := GenerateRandomStringURLSafe(32)

	if err != nil {
		token = ""
	}

	return c.Render("index", fiber.Map{
		"Title":     "Neue Nachricht",
		"MetaTitle": "Neue Nachricht",
		"Content":   "Sende Verschlüsselte Nachrichten, die sich selbst zerstören, nachdem sie gelesen wurden.",
		"Token":     token,
	})
}
