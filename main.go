package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v3"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type App struct {
	DB *sql.DB
}

type ShortenRequest struct {
	URL string `json:"url"`
}

func main() {

	db, err := sql.Open(
		"mysql",
		"root:password@tcp(localhost:3306)/urlshortener")

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"server": "urlShortener",
			"status": "200ok",
			"health": "Healthy",
		})
	})

	app.Post("/shorten", func(c fiber.Ctx) error {
		var req ShortenRequest

		if err := c.Bind().Body(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid request",
			})
		}
		id, _ := nanoid.New()

		_, err := db.Exec(
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
	})

	app.Get("/:shorten", func(c fiber.Ctx) error {
		shortCode := c.Params("shorten")

		var originalURL string

		err := db.QueryRow(
			"SELECT original_url FROM urls WHERE short_code = ?",
			shortCode,
		).Scan(&originalURL)

		if err != nil {
			if err == sql.ErrNoRows {
				return c.Status(404).JSON(fiber.Map{
					"error": "URL not found",
				})
			}

			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Redirect().To(originalURL)
	})

	log.Fatal(app.Listen(":3000"))
}
