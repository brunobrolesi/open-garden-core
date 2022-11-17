// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks_usecase

import (
	context "context"

	model "github.com/brunobrolesi/open-garden-core/internal/user/business/model"
	mock "github.com/stretchr/testify/mock"

	usecase "github.com/brunobrolesi/open-garden-core/internal/user/business/usecase"
)

// AuthenticateUserUseCase is an autogenerated mock type for the AuthenticateUserUseCase type
type AuthenticateUserUseCase struct {
	mock.Mock
}

// Exec provides a mock function with given fields: credentials, ctx
func (_m *AuthenticateUserUseCase) Exec(credentials usecase.AuthenticateUserInputDto, ctx context.Context) (model.Token, error) {
	ret := _m.Called(credentials, ctx)

	var r0 model.Token
	if rf, ok := ret.Get(0).(func(usecase.AuthenticateUserInputDto, context.Context) model.Token); ok {
		r0 = rf(credentials, ctx)
	} else {
		r0 = ret.Get(0).(model.Token)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(usecase.AuthenticateUserInputDto, context.Context) error); ok {
		r1 = rf(credentials, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthenticateUserUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthenticateUserUseCase creates a new instance of AuthenticateUserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthenticateUserUseCase(t mockConstructorTestingTNewAuthenticateUserUseCase) *AuthenticateUserUseCase {
	mock := &AuthenticateUserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
