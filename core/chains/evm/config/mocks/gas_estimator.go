// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	common "github.com/ethereum/go-ethereum/common"
	assets "github.com/smartcontractkit/chainlink/v2/core/assets"

	config "github.com/smartcontractkit/chainlink/v2/core/chains/evm/config"

	mock "github.com/stretchr/testify/mock"
)

// GasEstimator is an autogenerated mock type for the GasEstimator type
type GasEstimator struct {
	mock.Mock
}

// BlockHistory provides a mock function with given fields:
func (_m *GasEstimator) BlockHistory() config.BlockHistory {
	ret := _m.Called()

	var r0 config.BlockHistory
	if rf, ok := ret.Get(0).(func() config.BlockHistory); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.BlockHistory)
		}
	}

	return r0
}

// BumpMin provides a mock function with given fields:
func (_m *GasEstimator) BumpMin() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// BumpPercent provides a mock function with given fields:
func (_m *GasEstimator) BumpPercent() uint16 {
	ret := _m.Called()

	var r0 uint16
	if rf, ok := ret.Get(0).(func() uint16); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint16)
	}

	return r0
}

// BumpThreshold provides a mock function with given fields:
func (_m *GasEstimator) BumpThreshold() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// BumpTxDepth provides a mock function with given fields:
func (_m *GasEstimator) BumpTxDepth() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// EIP1559DynamicFees provides a mock function with given fields:
func (_m *GasEstimator) EIP1559DynamicFees() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FeeCapDefault provides a mock function with given fields:
func (_m *GasEstimator) FeeCapDefault() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// LimitDefault provides a mock function with given fields:
func (_m *GasEstimator) LimitDefault() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// LimitJobType provides a mock function with given fields:
func (_m *GasEstimator) LimitJobType() config.LimitJobType {
	ret := _m.Called()

	var r0 config.LimitJobType
	if rf, ok := ret.Get(0).(func() config.LimitJobType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(config.LimitJobType)
		}
	}

	return r0
}

// LimitMax provides a mock function with given fields:
func (_m *GasEstimator) LimitMax() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// LimitMultiplier provides a mock function with given fields:
func (_m *GasEstimator) LimitMultiplier() float32 {
	ret := _m.Called()

	var r0 float32
	if rf, ok := ret.Get(0).(func() float32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float32)
	}

	return r0
}

// LimitTransfer provides a mock function with given fields:
func (_m *GasEstimator) LimitTransfer() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// Mode provides a mock function with given fields:
func (_m *GasEstimator) Mode() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PriceDefault provides a mock function with given fields:
func (_m *GasEstimator) PriceDefault() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// PriceMax provides a mock function with given fields:
func (_m *GasEstimator) PriceMax() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// PriceMaxKey provides a mock function with given fields: _a0
func (_m *GasEstimator) PriceMaxKey(_a0 common.Address) *assets.Wei {
	ret := _m.Called(_a0)

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func(common.Address) *assets.Wei); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// PriceMin provides a mock function with given fields:
func (_m *GasEstimator) PriceMin() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// TipCapDefault provides a mock function with given fields:
func (_m *GasEstimator) TipCapDefault() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

// TipCapMin provides a mock function with given fields:
func (_m *GasEstimator) TipCapMin() *assets.Wei {
	ret := _m.Called()

	var r0 *assets.Wei
	if rf, ok := ret.Get(0).(func() *assets.Wei); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*assets.Wei)
		}
	}

	return r0
}

type mockConstructorTestingTNewGasEstimator interface {
	mock.TestingT
	Cleanup(func())
}

// NewGasEstimator creates a new instance of GasEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGasEstimator(t mockConstructorTestingTNewGasEstimator) *GasEstimator {
	mock := &GasEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
