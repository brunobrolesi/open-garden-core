package gateway

import "github.com/brunobrolesi/open-garden-core/internal/farm/business/model"

type TokenService interface {
	ValidateToken(token model.Token) (int, error)
}
