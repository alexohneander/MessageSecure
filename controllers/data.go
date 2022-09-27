package controllers

import (
	mongodb "MessageSecure/database"
	"MessageSecure/models"

	"github.com/gofiber/fiber/v2"
)

func EncryptData(c *fiber.Ctx) error {

	formMessage := c.FormValue("message")
	formToken := c.FormValue("token")
	id, _ := GenerateRandomStringURLSafe(10)

	message := models.Message{
		Id:      id,
		Message: formMessage,
	}

	client := mongodb.CreateClient()
	mongodb.AddDataToCollection(client, message)

	return c.Render("encrypted", fiber.Map{
		"Title":     "Nachricht verschlüsselt",
		"MetaTitle": "Encrypted",
		"Message":   formMessage,
		"Id":        id,
		"Token":     formToken,
	})
}

func DecryptData(c *fiber.Ctx) error {

	id := c.Params("id")
	token := c.Params("token")

	client := mongodb.CreateClient()
	message := mongodb.GetDataFromCollection(client, id)

	return c.Render("decrypted", fiber.Map{
		"Title":     "Lesen und zerstören?",
		"MetaTitle": "Lesen und zerstören?",
		"Content":   "Du liest und zerstörst gerade die Nachricht mit der id " + id,
		"Token":     token,
		"Id":        id,
		"Message":   message.Message,
	})
}
