// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks_gateway

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/brunobrolesi/open-garden-core/internal/user/business/model"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) CreateUser(_a0 model.User, _a1 context.Context) (model.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(model.User, context.Context) model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.User, context.Context) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) GetUserByEmail(_a0 string, _a1 context.Context) (model.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.User
	if rf, ok := ret.Get(0).(func(string, context.Context) model.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsEmailInUse provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) IsEmailInUse(_a0 string, _a1 context.Context) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, context.Context) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, context.Context) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
