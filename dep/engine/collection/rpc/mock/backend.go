// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Backend is an autogenerated mock type for the Backend type
type Backend struct {
	mock.Mock
}

// ProcessTransaction provides a mock function with given fields: _a0
func (_m *Backend) ProcessTransaction(_a0 *flow.TransactionBody) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.TransactionBody) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
