package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/MensahPrince/urlShortener/utils"
)

func IsAuthenticated(c fiber.Ctx) error {
	ip := c.IP()
	const LIMIT int = 10
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		allowed, err := utils.Allow(c.Context(), ip, LIMIT)
		if err != nil{
			return fiber.NewError(fiber.StatusInternalServerError, "rate limit check failed")
		}
		if !allowed{
			return fiber.NewError(fiber.StatusUnauthorized, "Please log in to continue.")
		}
		return c.Next()
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	jwtKey := os.Getenv("JWT_KEY")

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil || !token.Valid {
		
		return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired token")
	}

	return c.Next()
}
