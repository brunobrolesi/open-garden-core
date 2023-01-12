// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks_gateway

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/brunobrolesi/open-garden-core/internal/sensor/business/model"
)

// FarmRepository is an autogenerated mock type for the FarmRepository type
type FarmRepository struct {
	mock.Mock
}

// GetFarmByIdAndUserId provides a mock function with given fields: ctx, id, userId
func (_m *FarmRepository) GetFarmByIdAndUserId(ctx context.Context, id int, userId int) (model.Farm, error) {
	ret := _m.Called(ctx, id, userId)

	var r0 model.Farm
	if rf, ok := ret.Get(0).(func(context.Context, int, int) model.Farm); ok {
		r0 = rf(ctx, id, userId)
	} else {
		r0 = ret.Get(0).(model.Farm)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, id, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewFarmRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewFarmRepository creates a new instance of FarmRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFarmRepository(t mockConstructorTestingTNewFarmRepository) *FarmRepository {
	mock := &FarmRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}