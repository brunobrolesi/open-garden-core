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
		hashService    gateway.HashService
		userRepository gateway.UserRepository
		tokenService   gateway.TokenService
	}
)

func NewCreateUserUseCase(hashService gateway.HashService, userRepository gateway.UserRepository, tokenService gateway.TokenService) CreateUserUseCase {
	return createUser{
		hashService:    hashService,
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (c createUser) Exec(ctx context.Context, user CreateUserInputDto) (model.Token, error) {
	isEmailInUse, err := c.userRepository.IsEmailInUse(ctx, user.Email)

	if err != nil {
		return "", err
	}

	if isEmailInUse {
		return "", model.ErrEmailInUse
	}

	hashedPassword, err := c.hashService.GenerateHash(user.Password)

	if err != nil {
		return "", err
	}

	u := model.User{
		CompanyName: user.CompanyName,
		Email:       user.Email,
		Password:    hashedPassword,
		Active:      true,
	}
	newUser, err := c.userRepository.CreateUser(ctx, u)

	if err != nil {
		return "", err
	}

	token, err := c.tokenService.GenerateToken(newUser.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
