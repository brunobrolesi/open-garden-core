package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/sensor/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func SensorRouter(r *gin.Engine) *gin.RouterGroup {
	sensor := r.Group("/")
	{
		sensor.POST("/farms/:farm_id/sensors", factory.AuthMiddlewareFactory().Handle, factory.AddFarmSensorFactory().Handle)
	}

	return sensor
}
