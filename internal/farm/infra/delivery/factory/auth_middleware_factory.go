package factory

import (
	"os"

	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/middleware"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/service"
)

func AuthMiddlewareFactory() middleware.Middleware {
	tokenService := service.NewJwtTokenService(os.Getenv("JWT_SECRET"))
	middleware := middleware.NewAuthMiddleware(tokenService)

	return middleware
}
