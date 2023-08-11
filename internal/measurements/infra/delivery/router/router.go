package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/api/middleware"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func MeasurementRouter(r *gin.Engine, authMiddleware middleware.Middleware) *gin.RouterGroup {
	user := r.Group("/measurements")
	{
		user.GET("/sensors/:id", authMiddleware.Handle, factory.GetSensorMeasurementsFactory().Handle)
	}

	return user
}
