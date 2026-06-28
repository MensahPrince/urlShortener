package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/MensahPrince/urlShortener/types"
	"github.com/MensahPrince/urlShortener/auth/db"
	"github.com/MensahPrince/urlShortener/utils"
)

func Shorten(c fiber.Ctx) error {
	
	database := db.DB

	var req types.ShortenRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id := utils.GenerateShortCode()

	_, err := database.Exec(
		"INSERT INTO urls (original_url, short_code) VALUES (?, ?)",
		req.URL,
		id,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	} else {
		c.SendString("Done Shortening URL.")
	}

	var shortened = fmt.Sprintf("http://127.0.0.1:3000/%s", id)
	return c.JSON(fiber.Map{
		"url": shortened,
	})
}