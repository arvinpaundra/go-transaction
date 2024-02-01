// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	model "clean-arch/internal/model"
)

// OrderDetailRepository is an autogenerated mock type for the OrderDetailRepository type
type OrderDetailRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields: ctx
func (_m *OrderDetailRepository) FindAll(ctx *gin.Context) ([]model.OrderDetail, error) {
	ret := _m.Called(ctx)

	var r0 []model.OrderDetail
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context) ([]model.OrderDetail, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context) []model.OrderDetail); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.OrderDetail)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: ctx, orderDetails
func (_m *OrderDetailRepository) Insert(ctx *gin.Context, orderDetails []model.OrderDetail) error {
	ret := _m.Called(ctx, orderDetails)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gin.Context, []model.OrderDetail) error); ok {
		r0 = rf(ctx, orderDetails)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderDetailRepository creates a new instance of OrderDetailRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderDetailRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderDetailRepository {
	mock := &OrderDetailRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
