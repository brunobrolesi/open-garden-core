package gateway

import "github.com/brunobrolesi/open-garden-core/internal/user/business/model"

type TokenService interface {
	GenerateToken(userId int) (model.Token, error)
}
