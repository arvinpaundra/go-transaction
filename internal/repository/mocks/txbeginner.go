// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	repository "clean-arch/internal/repository"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// TxBeginner is an autogenerated mock type for the TxBeginner type
type TxBeginner struct {
	mock.Mock
}

// Begin provides a mock function with given fields: opts
func (_m *TxBeginner) Begin(opts ...*sql.TxOptions) (repository.Transaction, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 repository.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(...*sql.TxOptions) (repository.Transaction, error)); ok {
		return rf(opts...)
	}
	if rf, ok := ret.Get(0).(func(...*sql.TxOptions) repository.Transaction); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(...*sql.TxOptions) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTxBeginner creates a new instance of TxBeginner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxBeginner(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxBeginner {
	mock := &TxBeginner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}