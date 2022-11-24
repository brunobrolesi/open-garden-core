package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.RouterGroup {
	user := r.Group("/farms")
	{
		user.POST("/", factory.AuthMiddlewareFactory().Handle, factory.CreateFarmFactory().Handle)
		user.GET("/", factory.AuthMiddlewareFactory().Handle, factory.GetUserFarmsFactory().Handle)
		user.GET("/:id", factory.AuthMiddlewareFactory().Handle, factory.GetUserFarmFactory().Handle)
	}

	return user
}
