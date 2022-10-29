package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
)

type CreateUserModel struct {
	CompanyName string
	Email       string
	Password    string
}

type CreateUserUseCase interface {
	CreateUser(u CreateUserModel, c context.Context) (model.Token, error)
}
