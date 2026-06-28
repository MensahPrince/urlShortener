package routes

import (
	auth "github.com/MensahPrince/urlShortener/auth/handlers"
	sh "github.com/MensahPrince/urlShortener/handlers"
	"github.com/MensahPrince/urlShortener/auth/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/", auth.Base)
	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)
	app.Get("/profile", middleware.JWTMiddleware, auth.FetchProfile)
	app.Post("/edit/:param", middleware.JWTMiddleware, auth.EditHandler)
	app.Get("/request-otp", middleware.JWTMiddleware, auth.RequestOTP)
	app.Post("/reset", middleware.JWTMiddleware, auth.ResetPassword)
	app.Post("/delete", middleware.JWTMiddleware, auth.DeleteAccount)

	app.Get("/", sh.Base)
	app.Post("/redirect", sh.GetOriginal)
	app.Post("/:shorten", sh.Shorten)
}
