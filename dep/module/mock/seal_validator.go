// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// SealValidator is an autogenerated mock type for the SealValidator type
type SealValidator struct {
	mock.Mock
}

// Validate provides a mock function with given fields: candidate
func (_m *SealValidator) Validate(candidate *flow.Block) (*flow.Seal, error) {
	ret := _m.Called(candidate)

	var r0 *flow.Seal
	if rf, ok := ret.Get(0).(func(*flow.Block) *flow.Seal); ok {
		r0 = rf(candidate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Seal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*flow.Block) error); ok {
		r1 = rf(candidate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}