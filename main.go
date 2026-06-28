package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/MensahPrince/urlShortener/routes"
	"github.com/joho/godotenv"
	"github.com/MensahPrince/urlShortener/auth/db"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file:", err)
	}

	db.Connect()

	app := fiber.New()

	routes.SetupAuthRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
