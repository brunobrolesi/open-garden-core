package service

import (
	"strconv"

	"github.com/brunobrolesi/open-garden-core/internal/farm/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/shared"
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

func (s jwtTokenService) ValidateToken(t model.Token) (int, error) {
	token, err := jwt.ParseWithClaims(string(t), &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Secret), nil
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
