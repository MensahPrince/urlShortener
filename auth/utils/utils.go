package utils

import (
	"github.com/MensahPrince/urlShortener/auth/db"
	"github.com/MensahPrince/urlShortener/auth/types"
)

func CheckDB() types.DBSTATUS {
	if db.DB == nil {
		return types.DBSTATUS{
			Success: false,
			Message: "Database Connection Failed",
		}
	}

	return types.DBSTATUS{
		Success: true,
		Message: "Connected",
	}
}
