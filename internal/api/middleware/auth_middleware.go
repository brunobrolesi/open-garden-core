package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/brunobrolesi/open-garden-core/internal/shared"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type (
	authMiddleware struct {
		tokenService TokenService
	}

	Token string

	TokenService interface {
		ValidateToken(token Token) (int, error)
	}

	jwtTokenService struct {
		secret string
	}
)

func AuthMiddlewareFactory() Middleware {
	tokenService := NewJwtTokenService(os.Getenv("JWT_SECRET"))

	return NewAuthMiddleware(tokenService)
}

func NewAuthMiddleware(tokenService TokenService) Middleware {
	return authMiddleware{
		tokenService: tokenService,
	}
}

func (m authMiddleware) Handle(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": shared.ErrApiUnauthorized.Error()})
		c.Abort()
		return
	}
	userId, err := m.tokenService.ValidateToken(Token(token))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": shared.ErrApiUnauthorized.Error()})
		c.Abort()
		return
	}

	c.Request.Header.Set("X-User-Id", fmt.Sprint(userId))

	c.Next()
}

func NewJwtTokenService(secret string) TokenService {
	return jwtTokenService{
		secret: secret,
	}
}

func (s jwtTokenService) ValidateToken(t Token) (int, error) {
	token, err := jwt.ParseWithClaims(string(t), &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, shared.ErrApiUnauthorized
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	if !ok {
		return 0, shared.ErrApiUnauthorized
	}

	userId, err := strconv.Atoi(claims.Subject)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
