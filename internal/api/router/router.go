package router

import (
	farmRouter "github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/router"
	healthRouter "github.com/brunobrolesi/open-garden-core/internal/health/infra/delivery/router"
	userRouter "github.com/brunobrolesi/open-garden-core/internal/user/infra/delivery/router"
	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()

	router.Group("/api")
	{
		healthRouter.HealthRouter(router)
		userRouter.UserRouter(router)
		farmRouter.UserRouter(router)
	}

	return router
}
