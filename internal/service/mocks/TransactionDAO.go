// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/blazee5/finance-tracker/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// TransactionDAO is an autogenerated mock type for the TransactionDAO type
type TransactionDAO struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, transaction
func (_m *TransactionDAO) Create(ctx context.Context, transaction models.Transaction) (interface{}, error) {
	ret := _m.Called(ctx, transaction)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Transaction) (interface{}, error)); ok {
		return rf(ctx, transaction)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Transaction) interface{}); ok {
		r0 = rf(ctx, transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Transaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TransactionDAO) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAnalyze provides a mock function with given fields: ctx, id
func (_m *TransactionDAO) GetAnalyze(ctx context.Context, id string) ([]models.Analyze, error) {
	ret := _m.Called(ctx, id)

	var r0 []models.Analyze
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Analyze, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Analyze); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Analyze)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransaction provides a mock function with given fields: ctx, id
func (_m *TransactionDAO) GetTransaction(ctx context.Context, id string) (models.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (models.Transaction, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Transaction)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactions provides a mock function with given fields: ctx, userID
func (_m *TransactionDAO) GetTransactions(ctx context.Context, userID string) ([]models.Transaction, error) {
	ret := _m.Called(ctx, userID)

	var r0 []models.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]models.Transaction, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Transaction); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, transaction
func (_m *TransactionDAO) Update(ctx context.Context, transaction models.Transaction) (int64, error) {
	ret := _m.Called(ctx, transaction)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Transaction) (int64, error)); ok {
		return rf(ctx, transaction)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.Transaction) int64); ok {
		r0 = rf(ctx, transaction)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.Transaction) error); ok {
		r1 = rf(ctx, transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransactionDAO creates a new instance of TransactionDAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionDAO {
	mock := &TransactionDAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
