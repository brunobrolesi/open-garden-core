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
		Exec(user CreateUserInputDto, ctx context.Context) (model.Token, error)
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

func (c createUser) Exec(user CreateUserInputDto, ctx context.Context) (model.Token, error) {
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
	newUser, err := c.UserRepository.CreateUser(u, ctx)

	if err != nil {
		return "", err
	}

	token, err := c.TokenService.GenerateToken(newUser.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
