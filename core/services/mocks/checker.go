// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	services "github.com/smartcontractkit/chainlink/core/services"
	mock "github.com/stretchr/testify/mock"
)

// Checker is an autogenerated mock type for the Checker type
type Checker struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Checker) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsHealthy provides a mock function with given fields:
func (_m *Checker) IsHealthy() (bool, map[string]error) {
	ret := _m.Called()

	var r0 bool
	var r1 map[string]error
	if rf, ok := ret.Get(0).(func() (bool, map[string]error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() map[string]error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]error)
		}
	}

	return r0, r1
}

// IsReady provides a mock function with given fields:
func (_m *Checker) IsReady() (bool, map[string]error) {
	ret := _m.Called()

	var r0 bool
	var r1 map[string]error
	if rf, ok := ret.Get(0).(func() (bool, map[string]error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func() map[string]error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]error)
		}
	}

	return r0, r1
}

// Register provides a mock function with given fields: name, service
func (_m *Checker) Register(name string, service services.Checkable) error {
	ret := _m.Called(name, service)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, services.Checkable) error); ok {
		r0 = rf(name, service)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Checker) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields: name
func (_m *Checker) Unregister(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewChecker interface {
	mock.TestingT
	Cleanup(func())
}

// NewChecker creates a new instance of Checker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewChecker(t mockConstructorTestingTNewChecker) *Checker {
	mock := &Checker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
