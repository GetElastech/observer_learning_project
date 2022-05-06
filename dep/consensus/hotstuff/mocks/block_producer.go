// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"

	model "github.com/onflow/flow-go/consensus/hotstuff/model"
)

// BlockProducer is an autogenerated mock type for the BlockProducer type
type BlockProducer struct {
	mock.Mock
}

// MakeBlockProposal provides a mock function with given fields: qc, view
func (_m *BlockProducer) MakeBlockProposal(qc *flow.QuorumCertificate, view uint64) (*model.Proposal, error) {
	ret := _m.Called(qc, view)

	var r0 *model.Proposal
	if rf, ok := ret.Get(0).(func(*flow.QuorumCertificate, uint64) *model.Proposal); ok {
		r0 = rf(qc, view)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Proposal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*flow.QuorumCertificate, uint64) error); ok {
		r1 = rf(qc, view)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}