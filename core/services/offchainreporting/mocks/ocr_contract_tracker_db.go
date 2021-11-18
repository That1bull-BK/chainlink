// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	"github.com/smartcontractkit/chainlink/core/services/pg"
	offchainaggregator "github.com/smartcontractkit/libocr/gethwrappers/offchainaggregator"
	mock "github.com/stretchr/testify/mock"
)

// OCRContractTrackerDB is an autogenerated mock type for the OCRContractTrackerDB type
type OCRContractTrackerDB struct {
	mock.Mock
}

// LoadLatestRoundRequested provides a mock function with given fields:
func (_m *OCRContractTrackerDB) LoadLatestRoundRequested() (offchainaggregator.OffchainAggregatorRoundRequested, error) {
	ret := _m.Called()

	var r0 offchainaggregator.OffchainAggregatorRoundRequested
	if rf, ok := ret.Get(0).(func() offchainaggregator.OffchainAggregatorRoundRequested); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(offchainaggregator.OffchainAggregatorRoundRequested)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveLatestRoundRequested provides a mock function with given fields: tx, rr
func (_m *OCRContractTrackerDB) SaveLatestRoundRequested(tx pg.Queryer, rr offchainaggregator.OffchainAggregatorRoundRequested) error {
	ret := _m.Called(tx, rr)

	var r0 error
	if rf, ok := ret.Get(0).(func(pg.Queryer, offchainaggregator.OffchainAggregatorRoundRequested) error); ok {
		r0 = rf(tx, rr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
