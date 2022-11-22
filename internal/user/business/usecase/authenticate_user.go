package usecase

import (
	"context"

	"github.com/brunobrolesi/open-garden-core/internal/user/business/gateway"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
)

type (
	AuthenticateUserInputDto struct {
		Email    string
		Password string
	}

	AuthenticateUserUseCase interface {
		Exec(credentials AuthenticateUserInputDto, ctx context.Context) (model.Token, error)
	}

	authenticateUser struct {
		HashService    gateway.HashService
		UserRepository gateway.UserRepository
		TokenService   gateway.TokenService
	}
)

func NewAuthenticateUserUseCase(hashService gateway.HashService, userRepository gateway.UserRepository, tokenService gateway.TokenService) AuthenticateUserUseCase {
	return authenticateUser{
		HashService:    hashService,
		UserRepository: userRepository,
		TokenService:   tokenService,
	}
}

func (a authenticateUser) Exec(credentials AuthenticateUserInputDto, ctx context.Context) (model.Token, error) {
	user, err := a.UserRepository.GetUserByEmail(credentials.Email, ctx)

	if err != nil {
		return "", err
	}

	if user.EmailIsEmpty() {
		return "", model.ErrAuthentication
	}

	if err := a.HashService.CompareStringAndHash(credentials.Password, user.Password); err != nil {
		return "", model.ErrAuthentication
	}

	token, err := a.TokenService.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
