// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import cluster "github.com/ovrclk/akash/provider/cluster"
import mock "github.com/stretchr/testify/mock"
import types "github.com/ovrclk/akash/types"

// Cluster is an autogenerated mock type for the Cluster type
type Cluster struct {
	mock.Mock
}

// Reserve provides a mock function with given fields: _a0, _a1
func (_m *Cluster) Reserve(_a0 types.OrderID, _a1 *types.DeploymentGroup) (cluster.Reservation, error) {
	ret := _m.Called(_a0, _a1)

	var r0 cluster.Reservation
	if rf, ok := ret.Get(0).(func(types.OrderID, *types.DeploymentGroup) cluster.Reservation); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cluster.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.OrderID, *types.DeploymentGroup) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unreserve provides a mock function with given fields: _a0, _a1
func (_m *Cluster) Unreserve(_a0 types.OrderID, _a1 types.ResourceList) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.OrderID, types.ResourceList) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
