package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
)

type (
	CreateUserInputDto struct {
		CompanyName string
		Email       string
		Password    string
	}

	CreateUserUseCase interface {
		Exec(ctx context.Context, user CreateUserInputDto) (model.Token, error)
	}

	createUser struct {
		HashService    gateway.HashService
		UserRepository gateway.UserRepository
		TokenService   gateway.TokenService
	}
)

func NewCreateUserUseCase(hashService gateway.HashService, userRepository gateway.UserRepository, tokenService gateway.TokenService) CreateUserUseCase {
	return createUser{
		HashService:    hashService,
		UserRepository: userRepository,
		TokenService:   tokenService,
	}
}

func (c createUser) Exec(ctx context.Context, user CreateUserInputDto) (model.Token, error) {
	isEmailInUse, err := c.UserRepository.IsEmailInUse(ctx, user.Email)

	if err != nil {
		return "", err
	}

	if isEmailInUse {
		return "", model.ErrEmailInUse
	}

	hashedPassword, err := c.HashService.GenerateHash(user.Password)

	if err != nil {
		return "", err
	}

	u := model.User{
		CompanyName: user.CompanyName,
		Email:       user.Email,
		Password:    hashedPassword,
		Active:      true,
	}
	newUser, err := c.UserRepository.CreateUser(ctx, u)

	if err != nil {
		return "", err
	}

	token, err := c.TokenService.GenerateToken(newUser.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
