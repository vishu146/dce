// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import awsiface "github.com/Optum/dce/pkg/awsiface"

import mock "github.com/stretchr/testify/mock"
import time "time"

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CalculateTotalSpend provides a mock function with given fields: startDate, endDate
func (_m *Service) CalculateTotalSpend(startDate time.Time, endDate time.Time) (float64, error) {
	ret := _m.Called(startDate, endDate)

	var r0 float64
	if rf, ok := ret.Get(0).(func(time.Time, time.Time) float64); ok {
		r0 = rf(startDate, endDate)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time, time.Time) error); ok {
		r1 = rf(startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCostExplorer provides a mock function with given fields: costExplorer
func (_m *Service) SetCostExplorer(costExplorer awsiface.CostExplorerAPI) {
	_m.Called(costExplorer)
}