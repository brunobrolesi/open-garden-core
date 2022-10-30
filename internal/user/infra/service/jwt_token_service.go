package service

import (
	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/golang-jwt/jwt/v4"
)

type jwtTokenService struct {
	Secret string
}

func NewJwtTokenService(secret string) gateway.TokenService {
	return jwtTokenService{
		Secret: secret,
	}
}

func (s jwtTokenService) GenerateToken(userId int) (model.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
	})

	tokenString, err := token.SignedString([]byte(s.Secret))

	if err != nil {
		return "", err
	}

	return model.Token(tokenString), nil
}
