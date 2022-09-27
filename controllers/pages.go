package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func RenderIndex(c *fiber.Ctx) error {

	token, err := GenerateRandomStringURLSafe(32)

	if err != nil {
		token = ""
	}

	return c.Render("index", fiber.Map{
		"Title":     "New message",
		"MetaTitle": "New message",
		"Content":   "Send encrypted messages that self-destruct after being read.",
		"Token":     token,
		"Domain":    os.Getenv("DOMAIN"),
	})
}
