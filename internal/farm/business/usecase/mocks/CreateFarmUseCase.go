// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks_usecase

import (
	context "context"

	model "github.com/brunobrolesi/open-garden-core/internal/farm/business/model"
	mock "github.com/stretchr/testify/mock"

	usecase "github.com/brunobrolesi/open-garden-core/internal/farm/business/usecase"
)

// CreateFarmUseCase is an autogenerated mock type for the CreateFarmUseCase type
type CreateFarmUseCase struct {
	mock.Mock
}

// Exec provides a mock function with given fields: ctx, farm
func (_m *CreateFarmUseCase) Exec(ctx context.Context, farm usecase.CreateFarmInputDto) (model.Farm, error) {
	ret := _m.Called(ctx, farm)

	var r0 model.Farm
	if rf, ok := ret.Get(0).(func(context.Context, usecase.CreateFarmInputDto) model.Farm); ok {
		r0 = rf(ctx, farm)
	} else {
		r0 = ret.Get(0).(model.Farm)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, usecase.CreateFarmInputDto) error); ok {
		r1 = rf(ctx, farm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCreateFarmUseCase interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateFarmUseCase creates a new instance of CreateFarmUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateFarmUseCase(t mockConstructorTestingTNewCreateFarmUseCase) *CreateFarmUseCase {
	mock := &CreateFarmUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
