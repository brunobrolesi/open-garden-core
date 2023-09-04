package api

import (
	"os"

	"github.com/brunobrolesi/open-garden-core/internal/api/router"
)

func StartApp() {
	router := router.ApiRouter()

	port := os.Getenv("APP_PORT")
	router.Run(":" + port)
}
