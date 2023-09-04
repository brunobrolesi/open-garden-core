package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/api/middleware"
	farmRouter "github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/router"
	healthRouter "github.com/brunobrolesi/open-garden-core/internal/health/infra/delivery/router"
	measurementsRouter "github.com/brunobrolesi/open-garden-core/internal/measurements/infra/delivery/router"
	sensorRouter "github.com/brunobrolesi/open-garden-core/internal/sensor/infra/delivery/router"
	userRouter "github.com/brunobrolesi/open-garden-core/internal/user/infra/delivery/router"
	"github.com/gin-gonic/gin"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})

	router.Group("/api")
	{
		healthRouter.HealthRouter(router)
		userRouter.UserRouter(router)
		farmRouter.FarmRouter(router, middleware.AuthMiddlewareFactory())
		sensorRouter.SensorRouter(router, middleware.AuthMiddlewareFactory())
		measurementsRouter.MeasurementRouter(router, middleware.AuthMiddlewareFactory())
	}

	return router
}
