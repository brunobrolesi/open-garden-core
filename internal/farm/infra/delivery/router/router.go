package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/farm/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.RouterGroup {
	user := r.Group("/farm")
	{
		user.POST("/", factory.AuthMiddlewareFactory().Handle, factory.CreateFarmFactory().Handle)
	}

	return user
}
