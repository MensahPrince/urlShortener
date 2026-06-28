package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func Base(c fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"server": "urlShortener",
		"status": "200ok",
		"health": "Healthy",
	})
}
