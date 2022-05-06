// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	protocol "github.com/onflow/flow-go/state/protocol"
)

// Epoch is an autogenerated mock type for the Epoch type
type Epoch struct {
	mock.Mock
}

// Cluster provides a mock function with given fields: index
func (_m *Epoch) Cluster(index uint) (protocol.Cluster, error) {
	ret := _m.Called(index)

	var r0 protocol.Cluster
	if rf, ok := ret.Get(0).(func(uint) protocol.Cluster); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.Cluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Clustering provides a mock function with given fields:
func (_m *Epoch) Clustering() (flow.ClusterList, error) {
	ret := _m.Called()

	var r0 flow.ClusterList
	if rf, ok := ret.Get(0).(func() flow.ClusterList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.ClusterList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Counter provides a mock function with given fields:
func (_m *Epoch) Counter() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DKG provides a mock function with given fields:
func (_m *Epoch) DKG() (protocol.DKG, error) {
	ret := _m.Called()

	var r0 protocol.DKG
	if rf, ok := ret.Get(0).(func() protocol.DKG); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.DKG)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DKGPhase1FinalView provides a mock function with given fields:
func (_m *Epoch) DKGPhase1FinalView() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DKGPhase2FinalView provides a mock function with given fields:
func (_m *Epoch) DKGPhase2FinalView() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DKGPhase3FinalView provides a mock function with given fields:
func (_m *Epoch) DKGPhase3FinalView() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalView provides a mock function with given fields:
func (_m *Epoch) FinalView() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstView provides a mock function with given fields:
func (_m *Epoch) FirstView() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitialIdentities provides a mock function with given fields:
func (_m *Epoch) InitialIdentities() (flow.IdentityList, error) {
	ret := _m.Called()

	var r0 flow.IdentityList
	if rf, ok := ret.Get(0).(func() flow.IdentityList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RandomSource provides a mock function with given fields:
func (_m *Epoch) RandomSource() ([]byte, error) {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}