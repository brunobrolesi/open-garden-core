package router

import (
	healthRouter "github.com/brunobrolesi/open-garden-core/internal/health/infra/delivery/router"
	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()

	router.Group("/api")
	{
		healthRouter.HealthRouter(router)
	}

	return router
}
