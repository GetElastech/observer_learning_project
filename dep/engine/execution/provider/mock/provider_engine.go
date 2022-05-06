// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	context "context"

	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	network "github.com/onflow/flow-go/network"
)

// ProviderEngine is an autogenerated mock type for the ProviderEngine type
type ProviderEngine struct {
	mock.Mock
}

// BroadcastExecutionReceipt provides a mock function with given fields: _a0, _a1
func (_m *ProviderEngine) BroadcastExecutionReceipt(_a0 context.Context, _a1 *flow.ExecutionReceipt) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *flow.ExecutionReceipt) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Done provides a mock function with given fields:
func (_m *ProviderEngine) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Process provides a mock function with given fields: channel, originID, event
func (_m *ProviderEngine) Process(channel network.Channel, originID flow.Identifier, event interface{}) error {
	ret := _m.Called(channel, originID, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(network.Channel, flow.Identifier, interface{}) error); ok {
		r0 = rf(channel, originID, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessLocal provides a mock function with given fields: event
func (_m *ProviderEngine) ProcessLocal(event interface{}) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *ProviderEngine) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Submit provides a mock function with given fields: channel, originID, event
func (_m *ProviderEngine) Submit(channel network.Channel, originID flow.Identifier, event interface{}) {
	_m.Called(channel, originID, event)
}

// SubmitLocal provides a mock function with given fields: event
func (_m *ProviderEngine) SubmitLocal(event interface{}) {
	_m.Called(event)
}