package gateway

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
)

type UserRepository interface {
	CreateUser(context.Context, model.User) (model.User, error)
	GetUserByEmail(context.Context, string) (model.User, error)
	IsEmailInUse(context.Context, string) (bool, error)
}
