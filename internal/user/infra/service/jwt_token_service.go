package service

import (
	"fmt"
	"time"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/golang-jwt/jwt/v4"
)

type jwtTokenService struct {
	secret string
}

func NewJwtTokenService(secret string) gateway.TokenService {
	return jwtTokenService{
		secret: secret,
	}
}

func (s jwtTokenService) GenerateToken(userId int) (model.Token, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprint(userId),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.secret))

	if err != nil {
		return "", err
	}

	return model.Token(tokenString), nil
}
