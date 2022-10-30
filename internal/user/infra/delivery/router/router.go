package router

import (
	"github.com/brunobrolesi/open-garden-core/internal/user/infra/delivery/factory"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.RouterGroup {
	user := r.Group("")
	{
		user.POST("/register", factory.CreateUserFactory().Handle)
	}

	return user
}
