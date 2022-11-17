// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks_gateway

import mock "github.com/stretchr/testify/mock"

// HashService is an autogenerated mock type for the HashService type
type HashService struct {
	mock.Mock
}

// CompareStringAndHash provides a mock function with given fields: s, hash
func (_m *HashService) CompareStringAndHash(s string, hash string) error {
	ret := _m.Called(s, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(s, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateHash provides a mock function with given fields: s
func (_m *HashService) GenerateHash(s string) (string, error) {
	ret := _m.Called(s)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHashService interface {
	mock.TestingT
	Cleanup(func())
}

// NewHashService creates a new instance of HashService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHashService(t mockConstructorTestingTNewHashService) *HashService {
	mock := &HashService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
