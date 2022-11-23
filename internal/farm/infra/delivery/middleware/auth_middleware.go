package middleware

import (
	"fmt"
	"net/http"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
)

type (
	authMiddleware struct {
		TokenService gateway.TokenService
	}
)

func NewAuthMiddleware(tokenService gateway.TokenService) Middleware {
	return authMiddleware{
		TokenService: tokenService,
	}
}

func (m authMiddleware) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": shared.ErrApiUnauthorized.Error()})
		c.Abort()
		return
	}
	userId, err := m.TokenService.ValidateToken(model.Token(token))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": shared.ErrApiUnauthorized.Error()})
		c.Abort()
		return
	}

	c.Request.Header.Set("X-User-Id", fmt.Sprint(userId))

	c.Next()
}
