// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockinsecure

import (
	insecure "github.com/onflow/flow-go/insecure"
	mock "github.com/stretchr/testify/mock"
)

// AttackerServer is an autogenerated mock type for the AttackerServer type
type AttackerServer struct {
	mock.Mock
}

// Observe provides a mock function with given fields: _a0
func (_m *AttackerServer) Observe(_a0 insecure.Attacker_ObserveServer) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(insecure.Attacker_ObserveServer) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
