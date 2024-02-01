// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	repository "clean-arch/internal/repository"

	mock "github.com/stretchr/testify/mock"
)

// Transaction is an autogenerated mock type for the Transaction type
type Transaction struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *Transaction) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderDetailRepository provides a mock function with given fields:
func (_m *Transaction) OrderDetailRepository() repository.OrderDetailRepository {
	ret := _m.Called()

	var r0 repository.OrderDetailRepository
	if rf, ok := ret.Get(0).(func() repository.OrderDetailRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.OrderDetailRepository)
		}
	}

	return r0
}

// OrderRepository provides a mock function with given fields:
func (_m *Transaction) OrderRepository() repository.OrderRepository {
	ret := _m.Called()

	var r0 repository.OrderRepository
	if rf, ok := ret.Get(0).(func() repository.OrderRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.OrderRepository)
		}
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *Transaction) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository provides a mock function with given fields:
func (_m *Transaction) UserRepository() repository.UserRepository {
	ret := _m.Called()

	var r0 repository.UserRepository
	if rf, ok := ret.Get(0).(func() repository.UserRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.UserRepository)
		}
	}

	return r0
}

// NewTransaction creates a new instance of Transaction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransaction(t interface {
	mock.TestingT
	Cleanup(func())
}) *Transaction {
	mock := &Transaction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}