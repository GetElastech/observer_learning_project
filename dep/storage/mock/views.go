// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// Views is an autogenerated mock type for the Views type
type Views struct {
	mock.Mock
}

// Retrieve provides a mock function with given fields: action
func (_m *Views) Retrieve(action uint8) (uint64, error) {
	ret := _m.Called(action)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint8) uint64); ok {
		r0 = rf(action)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint8) error); ok {
		r1 = rf(action)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: action, view
func (_m *Views) Store(action uint8, view uint64) error {
	ret := _m.Called(action, view)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8, uint64) error); ok {
		r0 = rf(action, view)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}