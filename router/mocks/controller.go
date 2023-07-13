// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// GetSample provides a mock function with given fields: _a0
func (_m *Controller) GetSample(_a0 *gin.Context) {
	_m.Called(_a0)
}

// Ping provides a mock function with given fields: _a0
func (_m *Controller) Ping(_a0 *gin.Context) {
	_m.Called(_a0)
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
