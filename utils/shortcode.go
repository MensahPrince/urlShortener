package utils

import (
	nanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateShortCode() string {
	id, _ := nanoid.New()

	return id
}