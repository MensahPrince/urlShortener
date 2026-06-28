package handlers

import (
	"github.com/MensahPrince/urlShortener/auth/db"
	"github.com/gofiber/fiber/v3"
)

func GetOriginal(c fiber.Ctx) error {

	//Database set up
	database := db.DB

	//Get shortcode and assign it to code
	code := c.Params("shorten")

	var link string

	if code == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}

	err := database.QueryRow("SELECT original_url FROM urls WHERE short_code = ?",
		code).Scan(&link)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Shortcode not found",
		})
	}

	return c.Redirect().To(link)

}
