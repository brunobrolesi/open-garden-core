package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/api/middleware"
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func SensorRouter(r *gin.Engine, authMiddleware middleware.Middleware) *gin.RouterGroup {
	sensor := r.Group("/")
	{
		sensor.POST("/farms/:farm_id/sensors", authMiddleware.Handle, factory.AddFarmSensorFactory().Handle)
		sensor.GET("/farms/:farm_id/sensors", authMiddleware.Handle, factory.GetFarmSensorsFactory().Handle)
		sensor.POST("/farms/:farm_id/sensors/:farm_sensor_id", authMiddleware.Handle, factory.GetFarmSensorFactory().Handle)
	}

	return sensor
}
