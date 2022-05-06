// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onflow/flow-go/module (interfaces: Local,Requester)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	crypto "github.com/onflow/flow-go/crypto"
	hash "github.com/onflow/flow-go/crypto/hash"
	flow "github.com/onflow/flow-go/model/flow"
	reflect "reflect"
)

// MockLocal is a mock of Local interface
type MockLocal struct {
	ctrl     *gomock.Controller
	recorder *MockLocalMockRecorder
}

// MockLocalMockRecorder is the mock recorder for MockLocal
type MockLocalMockRecorder struct {
	mock *MockLocal
}

// NewMockLocal creates a new mock instance
func NewMockLocal(ctrl *gomock.Controller) *MockLocal {
	mock := &MockLocal{ctrl: ctrl}
	mock.recorder = &MockLocalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocal) EXPECT() *MockLocalMockRecorder {
	return m.recorder
}

// Address mocks base method
func (m *MockLocal) Address() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Address")
	ret0, _ := ret[0].(string)
	return ret0
}

// Address indicates an expected call of Address
func (mr *MockLocalMockRecorder) Address() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Address", reflect.TypeOf((*MockLocal)(nil).Address))
}

// NodeID mocks base method
func (m *MockLocal) NodeID() flow.Identifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeID")
	ret0, _ := ret[0].(flow.Identifier)
	return ret0
}

// NodeID indicates an expected call of NodeID
func (mr *MockLocalMockRecorder) NodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockLocal)(nil).NodeID))
}

// NotMeFilter mocks base method
func (m *MockLocal) NotMeFilter() flow.IdentityFilter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotMeFilter")
	ret0, _ := ret[0].(flow.IdentityFilter)
	return ret0
}

// NotMeFilter indicates an expected call of NotMeFilter
func (mr *MockLocalMockRecorder) NotMeFilter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotMeFilter", reflect.TypeOf((*MockLocal)(nil).NotMeFilter))
}

// Sign mocks base method
func (m *MockLocal) Sign(arg0 []byte, arg1 hash.Hasher) (crypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", arg0, arg1)
	ret0, _ := ret[0].(crypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign
func (mr *MockLocalMockRecorder) Sign(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockLocal)(nil).Sign), arg0, arg1)
}

// SignFunc mocks base method
func (m *MockLocal) SignFunc(arg0 []byte, arg1 hash.Hasher, arg2 func(crypto.PrivateKey, []byte, hash.Hasher) (crypto.Signature, error)) (crypto.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignFunc", arg0, arg1, arg2)
	ret0, _ := ret[0].(crypto.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignFunc indicates an expected call of SignFunc
func (mr *MockLocalMockRecorder) SignFunc(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignFunc", reflect.TypeOf((*MockLocal)(nil).SignFunc), arg0, arg1, arg2)
}

// MockRequester is a mock of Requester interface
type MockRequester struct {
	ctrl     *gomock.Controller
	recorder *MockRequesterMockRecorder
}

// MockRequesterMockRecorder is the mock recorder for MockRequester
type MockRequesterMockRecorder struct {
	mock *MockRequester
}

// NewMockRequester creates a new mock instance
func NewMockRequester(ctrl *gomock.Controller) *MockRequester {
	mock := &MockRequester{ctrl: ctrl}
	mock.recorder = &MockRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRequester) EXPECT() *MockRequesterMockRecorder {
	return m.recorder
}

// EntityByID mocks base method
func (m *MockRequester) EntityByID(arg0 flow.Identifier, arg1 flow.IdentityFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EntityByID", arg0, arg1)
}

// EntityByID indicates an expected call of EntityByID
func (mr *MockRequesterMockRecorder) EntityByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityByID", reflect.TypeOf((*MockRequester)(nil).EntityByID), arg0, arg1)
}

// Force mocks base method
func (m *MockRequester) Force() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Force")
}

// Force indicates an expected call of Force
func (mr *MockRequesterMockRecorder) Force() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Force", reflect.TypeOf((*MockRequester)(nil).Force))
}

// Query mocks base method
func (m *MockRequester) Query(arg0 flow.Identifier, arg1 flow.IdentityFilter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Query", arg0, arg1)
}

// Query indicates an expected call of Query
func (mr *MockRequesterMockRecorder) Query(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockRequester)(nil).Query), arg0, arg1)
}