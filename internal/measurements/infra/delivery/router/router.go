package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/api/middleware"
	"github.com/brunobrolesi/open-garden-core/internal/measurements/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func MeasurementRouter(r *gin.Engine, authMiddleware middleware.Middleware) *gin.RouterGroup {
	measurement := r.Group("/measurements")
	{
		measurement.GET("/sensors/:id", authMiddleware.Handle, factory.GetSensorMeasurementsFactory().Handle)
		measurement.POST("/sensor", factory.AddSensorMeasurementFactory().Handle)
	}

	return measurement
}
