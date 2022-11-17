package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
)

type UserRepository interface {
	CreateUser(model.User, context.Context) (model.User, error)
	GetUserByEmail(string, context.Context) (model.User, error)
}
