// Code generated by mockery v1.0.0. DO NOT EDIT.

package mempool

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// Identifiers is an autogenerated mock type for the Identifiers type
type Identifiers struct {
	mock.Mock
}

// Add provides a mock function with given fields: id
func (_m *Identifiers) Add(id flow.Identifier) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// All provides a mock function with given fields:
func (_m *Identifiers) All() flow.IdentifierList {
	ret := _m.Called()

	var r0 flow.IdentifierList
	if rf, ok := ret.Get(0).(func() flow.IdentifierList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentifierList)
		}
	}

	return r0
}

// Has provides a mock function with given fields: id
func (_m *Identifiers) Has(id flow.Identifier) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Rem provides a mock function with given fields: id
func (_m *Identifiers) Rem(id flow.Identifier) bool {
	ret := _m.Called(id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) bool); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Size provides a mock function with given fields:
func (_m *Identifiers) Size() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}