// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks_usecase

import (
	context "context"

	model "github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	mock "github.com/stretchr/testify/mock"
)

// GetUserFarmsUseCase is an autogenerated mock type for the GetUserFarmsUseCase type
type GetUserFarmsUseCase struct {
	mock.Mock
}

// Exec provides a mock function with given fields: ctx, userId
func (_m *GetUserFarmsUseCase) Exec(ctx context.Context, userId int) (model.Farms, error) {
	ret := _m.Called(ctx, userId)

	var r0 model.Farms
	if rf, ok := ret.Get(0).(func(context.Context, int) model.Farms); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Farms)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGetUserFarmsUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetUserFarmsUseCase creates a new instance of GetUserFarmsUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetUserFarmsUseCase(t mockConstructorTestingTNewGetUserFarmsUseCase) *GetUserFarmsUseCase {
	mock := &GetUserFarmsUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}