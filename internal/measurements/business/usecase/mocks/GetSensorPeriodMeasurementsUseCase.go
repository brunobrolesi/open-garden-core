// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks_usecase

import (
	context "context"

	model "github.com/brunobrolesi/open-garden-core/internal/measurements/business/model"
	usecase "github.com/brunobrolesi/open-garden-core/internal/measurements/business/usecase"
	mock "github.com/stretchr/testify/mock"
)

// GetSensorPeriodMeasurementsUseCase is an autogenerated mock type for the GetSensorPeriodMeasurementsUseCase type
type GetSensorPeriodMeasurementsUseCase struct {
	mock.Mock
}

// Exec provides a mock function with given fields: ctx, input
func (_m *GetSensorPeriodMeasurementsUseCase) Exec(ctx context.Context, input usecase.GetSensorPeriodMeasurementsInputDto) (model.SensorMeasurements, error) {
	ret := _m.Called(ctx, input)

	var r0 model.SensorMeasurements
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, usecase.GetSensorPeriodMeasurementsInputDto) (model.SensorMeasurements, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, usecase.GetSensorPeriodMeasurementsInputDto) model.SensorMeasurements); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.SensorMeasurements)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, usecase.GetSensorPeriodMeasurementsInputDto) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewGetSensorPeriodMeasurementsUseCase creates a new instance of GetSensorPeriodMeasurementsUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGetSensorPeriodMeasurementsUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *GetSensorPeriodMeasurementsUseCase {
	mock := &GetSensorPeriodMeasurementsUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}