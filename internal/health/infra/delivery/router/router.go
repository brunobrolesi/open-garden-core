package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/health/infra/delivery/handler"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.Engine) *gin.RouterGroup {
	health := r.Group("")
	{
		health.GET("/ping", handler.HealthCheckHandler)
	}

	return health
}
