// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks_usecase

import (
	context "context"

	model "github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
	usecase "github.com/brunobrolesi/open-garden-core/internal/sensor/business/usecase"
	mock "github.com/stretchr/testify/mock"
)

// GetFarmSensorUseCase is an autogenerated mock type for the GetFarmSensorUseCase type
type GetFarmSensorUseCase struct {
	mock.Mock
}

// Exec provides a mock function with given fields: ctx, input
func (_m *GetFarmSensorUseCase) Exec(ctx context.Context, input usecase.GetFarmSensorInputDto) (model.FarmSensor, error) {
	ret := _m.Called(ctx, input)

	var r0 model.FarmSensor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, usecase.GetFarmSensorInputDto) (model.FarmSensor, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, usecase.GetFarmSensorInputDto) model.FarmSensor); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(model.FarmSensor)
	}

	if rf, ok := ret.Get(1).(func(context.Context, usecase.GetFarmSensorInputDto) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGetFarmSensorUseCase creates a new instance of GetFarmSensorUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetFarmSensorUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetFarmSensorUseCase {
	mock := &GetFarmSensorUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}