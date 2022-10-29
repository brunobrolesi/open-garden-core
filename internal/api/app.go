package api

import (
	"github.com/brunobrolesi/open-garden-core/internal/api/router"
)

func StartApp() {
	router := router.ApiRouter()

	router.Run(":8080")
}
