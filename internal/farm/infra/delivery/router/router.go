package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/api/middleware"
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func FarmRouter(r *gin.Engine, authMiddleware middleware.Middleware) *gin.RouterGroup {
	user := r.Group("/farms")
	{
		user.POST("/", authMiddleware.Handle, factory.CreateFarmFactory().Handle)
		user.GET("/", authMiddleware.Handle, factory.GetUserFarmsFactory().Handle)
		user.GET("/:id", authMiddleware.Handle, factory.GetUserFarmFactory().Handle)
	}

	return user
}
