// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/onflow/flow-go/storage"
)

// ServiceEvents is an autogenerated mock type for the ServiceEvents type
type ServiceEvents struct {
	mock.Mock
}

// BatchStore provides a mock function with given fields: blockID, events, batch
func (_m *ServiceEvents) BatchStore(blockID flow.Identifier, events []flow.Event, batch storage.BatchStorage) error {
	ret := _m.Called(blockID, events, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, []flow.Event, storage.BatchStorage) error); ok {
		r0 = rf(blockID, events, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ByBlockID provides a mock function with given fields: blockID
func (_m *ServiceEvents) ByBlockID(blockID flow.Identifier) ([]flow.Event, error) {
	ret := _m.Called(blockID)

	var r0 []flow.Event
	if rf, ok := ret.Get(0).(func(flow.Identifier) []flow.Event); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]flow.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}