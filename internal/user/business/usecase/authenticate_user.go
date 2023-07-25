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
		Exec(ctx context.Context, credentials AuthenticateUserInputDto) (model.Token, error)
	}

	authenticateUser struct {
		hashService    gateway.HashService
		userRepository gateway.UserRepository
		tokenService   gateway.TokenService
	}
)

func NewAuthenticateUserUseCase(hashService gateway.HashService, userRepository gateway.UserRepository, tokenService gateway.TokenService) AuthenticateUserUseCase {
	return authenticateUser{
		hashService:    hashService,
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (a authenticateUser) Exec(ctx context.Context, credentials AuthenticateUserInputDto) (model.Token, error) {
	user, err := a.userRepository.GetUserByEmail(ctx, credentials.Email)

	if err != nil {
		return "", err
	}

	if user.EmailIsEmpty() {
		return "", model.ErrAuthentication
	}

	if err := a.hashService.CompareStringAndHash(credentials.Password, user.Password); err != nil {
		return "", model.ErrAuthentication
	}

	token, err := a.tokenService.GenerateToken(user.Id)

	if err != nil {
		return "", err
	}

	return token, nil
}
