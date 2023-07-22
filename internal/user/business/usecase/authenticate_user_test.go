package usecase_test

import (
	"context"
	"errors"
	"testing"

	mocks_gateway "github.com/brunobrolesi/open-garden-core/internal/user/business/gateway/mocks"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	"github.com/brunobrolesi/open-garden-core/internal/user/business/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAuthenticateUserUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                usecase.AuthenticateUserUseCase
		HashServiceMock    *mocks_gateway.HashService
		UserRepositoryMock *mocks_gateway.UserRepository
		TokenServiceMock   *mocks_gateway.TokenService
	}
	makeTestSuite := func() TestSuite {
		hashServiceMock := mocks_gateway.NewHashService(t)
		userRepositoryMock := mocks_gateway.NewUserRepository(t)
		tokenServiceMock := mocks_gateway.NewTokenService(t)
		sut := usecase.NewAuthenticateUserUseCase(hashServiceMock, userRepositoryMock, tokenServiceMock)
		return TestSuite{
			Sut:                sut,
			HashServiceMock:    hashServiceMock,
			UserRepositoryMock: userRepositoryMock,
			TokenServiceMock:   tokenServiceMock,
		}
	}
	makeAuthenticateUserInputDto := func() usecase.AuthenticateUserInputDto {
		return usecase.AuthenticateUserInputDto{
			Email:    "valid@mail.com",
			Password: "valid_password",
		}
	}
	makeUser := func() model.User {
		return model.User{
			Id:          1,
			CompanyName: "valid_name",
			Email:       "valid@mail.com",
			Password:    "hashed_password",
			Active:      true,
		}
	}
	t.Run("Should call GetUserByEmail from UserRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.User{}, errors.New("any_error"))
		credentials := makeAuthenticateUserInputDto()
		ctx := context.Background()
		testSuite.Sut.Exec(ctx, credentials)

		testSuite.UserRepositoryMock.AssertCalled(t, "GetUserByEmail", ctx, credentials.Email)
	})
	t.Run("Should return an error if GetUserByEmail from UserRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.User{}, errors.New("user_repository_error"))
		credentials := makeAuthenticateUserInputDto()
		token, err := testSuite.Sut.Exec(context.Background(), credentials)
		assert.Empty(t, token)
		assert.EqualError(t, err, "user_repository_error")
	})
	t.Run("Should return ErrAuthentication if user not exists", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(model.User{}, nil)
		credentials := makeAuthenticateUserInputDto()
		token, err := testSuite.Sut.Exec(context.Background(), credentials)
		assert.Empty(t, token)
		assert.EqualError(t, err, model.ErrAuthentication.Error())
	})
	t.Run("Should call CompareStringAndHash from HashService with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		credentials := makeAuthenticateUserInputDto()
		ctx := context.Background()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(makeUser(), nil)
		testSuite.HashServiceMock.On("CompareStringAndHash", mock.Anything, mock.Anything).Return(errors.New("any_error"))
		testSuite.Sut.Exec(ctx, credentials)
		testSuite.HashServiceMock.AssertCalled(t, "CompareStringAndHash", credentials.Password, makeUser().Password)
	})
	t.Run("Should return an error if CompareStringAndHash from HashService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		credentials := makeAuthenticateUserInputDto()
		ctx := context.Background()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(makeUser(), nil)
		testSuite.HashServiceMock.On("CompareStringAndHash", mock.Anything, mock.Anything).Return(errors.New("hash_service_error"))
		token, err := testSuite.Sut.Exec(ctx, credentials)
		assert.Empty(t, token)
		assert.EqualError(t, err, "invalid email or password")
	})
	t.Run("Should call GenerateToken from TokenService with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		credentials := makeAuthenticateUserInputDto()
		ctx := context.Background()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(makeUser(), nil)
		testSuite.HashServiceMock.On("CompareStringAndHash", mock.Anything, mock.Anything).Return(nil)
		testSuite.TokenServiceMock.On("GenerateToken", mock.Anything).Return(model.Token(""), errors.New("any_error"))
		testSuite.Sut.Exec(ctx, credentials)
		testSuite.TokenServiceMock.AssertCalled(t, "GenerateToken", makeUser().Id)
	})
	t.Run("Should return an error if GenerateToken from TokenService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		credentials := makeAuthenticateUserInputDto()
		ctx := context.Background()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(makeUser(), nil)
		testSuite.HashServiceMock.On("CompareStringAndHash", mock.Anything, mock.Anything).Return(nil)
		testSuite.TokenServiceMock.On("GenerateToken", mock.Anything).Return(model.Token(""), errors.New("token_service_error"))
		token, err := testSuite.Sut.Exec(ctx, credentials)
		assert.Empty(t, token)
		assert.EqualError(t, err, "token_service_error")

	})
	t.Run("Should return a token on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		credentials := makeAuthenticateUserInputDto()
		ctx := context.Background()
		testSuite.UserRepositoryMock.On("GetUserByEmail", mock.Anything, mock.Anything).Return(makeUser(), nil)
		testSuite.HashServiceMock.On("CompareStringAndHash", mock.Anything, mock.Anything).Return(nil)
		testSuite.TokenServiceMock.On("GenerateToken", mock.Anything).Return(model.Token("valid_token"), nil)
		token, err := testSuite.Sut.Exec(ctx, credentials)
		assert.Equal(t, model.Token("valid_token"), token)
		assert.Nil(t, err)
	})
}
