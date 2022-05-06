// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	storage "github.com/onflow/flow-go/storage"
)

// ExecutionReceipts is an autogenerated mock type for the ExecutionReceipts type
type ExecutionReceipts struct {
	mock.Mock
}

// BatchStore provides a mock function with given fields: receipt, batch
func (_m *ExecutionReceipts) BatchStore(receipt *flow.ExecutionReceipt, batch storage.BatchStorage) error {
	ret := _m.Called(receipt, batch)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.ExecutionReceipt, storage.BatchStorage) error); ok {
		r0 = rf(receipt, batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ByBlockID provides a mock function with given fields: blockID
func (_m *ExecutionReceipts) ByBlockID(blockID flow.Identifier) (flow.ExecutionReceiptList, error) {
	ret := _m.Called(blockID)

	var r0 flow.ExecutionReceiptList
	if rf, ok := ret.Get(0).(func(flow.Identifier) flow.ExecutionReceiptList); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.ExecutionReceiptList)
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

// ByID provides a mock function with given fields: receiptID
func (_m *ExecutionReceipts) ByID(receiptID flow.Identifier) (*flow.ExecutionReceipt, error) {
	ret := _m.Called(receiptID)

	var r0 *flow.ExecutionReceipt
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.ExecutionReceipt); ok {
		r0 = rf(receiptID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.ExecutionReceipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(receiptID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: receipt
func (_m *ExecutionReceipts) Store(receipt *flow.ExecutionReceipt) error {
	ret := _m.Called(receipt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.ExecutionReceipt) error); ok {
		r0 = rf(receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}