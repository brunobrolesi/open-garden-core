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

func TestCreateUserUseCase(t *testing.T) {
	type TestSuite struct {
		Sut                usecase.CreateUserUseCase
		HashServiceMock    *mocks_gateway.HashService
		UserRepositoryMock *mocks_gateway.UserRepository
		TokenServiceMock   *mocks_gateway.TokenService
	}
	makeTestSuite := func() TestSuite {
		hashServiceMock := mocks_gateway.NewHashService(t)
		userRepositoryMock := mocks_gateway.NewUserRepository(t)
		tokenServiceMock := mocks_gateway.NewTokenService(t)
		sut := usecase.NewCreateUserUseCase(hashServiceMock, userRepositoryMock, tokenServiceMock)
		return TestSuite{
			Sut:                sut,
			HashServiceMock:    hashServiceMock,
			UserRepositoryMock: userRepositoryMock,
			TokenServiceMock:   tokenServiceMock,
		}
	}
	makeCreateUserModel := func() usecase.CreateUserModel {
		return usecase.CreateUserModel{
			CompanyName: "valid_name",
			Email:       "valid@mail.com",
			Password:    "valid_password",
		}
	}
	t.Run("Should call GenerateHash from HashService with correct value", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("", errors.New("any_error"))
		user := makeCreateUserModel()
		testSuite.Sut.CreateUser(user, context.Background())

		testSuite.HashServiceMock.AssertCalled(t, "GenerateHash", user.Password)
	})
	t.Run("Should return an error if GenerateHash from HashService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("", errors.New("hash_error"))
		user := makeCreateUserModel()
		token, err := testSuite.Sut.CreateUser(user, context.Background())
		assert.Empty(t, token)
		assert.Error(t, err, "hash_error")
	})
	t.Run("Should call CreateUser from UserRepository with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		user := makeCreateUserModel()
		ctx := context.Background()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("any_hash", nil)
		testSuite.UserRepositoryMock.On("CreateUser", mock.Anything, mock.Anything).Return(model.User{}, errors.New("any_error"))
		testSuite.Sut.CreateUser(user, ctx)
		expectedUserCall := model.User{
			CompanyName: user.CompanyName,
			Email:       user.Email,
			Password:    "any_hash",
			Active:      true,
		}
		testSuite.UserRepositoryMock.AssertCalled(t, "CreateUser", expectedUserCall, ctx)
	})
	t.Run("Should return an error if CreateUser from UserRepository returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		user := makeCreateUserModel()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("any_hash", nil)
		testSuite.UserRepositoryMock.On("CreateUser", mock.Anything, mock.Anything).Return(model.User{}, errors.New("user_repository_error"))
		token, err := testSuite.Sut.CreateUser(user, context.Background())
		assert.Empty(t, token)
		assert.Error(t, err, "user_repository_error")
	})
	t.Run("Should call GenerateToken from TokenService with correct values", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("any_hash", nil)
		user := model.User{
			Id:          1,
			CompanyName: "valid_company_name",
			Email:       "valid@mail.com",
			Password:    "any_hash",
			Active:      true,
		}
		testSuite.UserRepositoryMock.On("CreateUser", mock.Anything, mock.Anything).Return(user, nil)
		testSuite.TokenServiceMock.On("GenerateToken", mock.Anything).Return(model.Token(""), errors.New("token_error"))
		testSuite.Sut.CreateUser(makeCreateUserModel(), context.Background())
		testSuite.TokenServiceMock.AssertCalled(t, "GenerateToken", user.Id)
	})
	t.Run("Should return an error if GenerateToken from TokenService returns an error", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("any_hash", nil)
		user := model.User{
			Id:          1,
			CompanyName: "valid_company_name",
			Email:       "valid@mail.com",
			Password:    "any_hash",
			Active:      true,
		}
		testSuite.UserRepositoryMock.On("CreateUser", mock.Anything, mock.Anything).Return(user, nil)
		testSuite.TokenServiceMock.On("GenerateToken", mock.Anything).Return(model.Token(""), errors.New("token_error"))
		token, err := testSuite.Sut.CreateUser(makeCreateUserModel(), context.Background())
		assert.Empty(t, token)
		assert.Error(t, err, "token_error")
	})
	t.Run("Should return a token on success", func(t *testing.T) {
		testSuite := makeTestSuite()
		testSuite.HashServiceMock.On("GenerateHash", mock.Anything).Return("any_hash", nil)
		user := model.User{
			Id:          1,
			CompanyName: "valid_company_name",
			Email:       "valid@mail.com",
			Password:    "any_hash",
			Active:      true,
		}
		testSuite.UserRepositoryMock.On("CreateUser", mock.Anything, mock.Anything).Return(user, nil)
		testSuite.TokenServiceMock.On("GenerateToken", mock.Anything).Return(model.Token("token"), nil)
		token, err := testSuite.Sut.CreateUser(makeCreateUserModel(), context.Background())
		assert.Equal(t, model.Token("token"), token)
		assert.Nil(t, err)
	})
}
