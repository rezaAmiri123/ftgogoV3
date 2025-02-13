// Code generated by mockery v2.33.0. DO NOT EDIT.

package domain

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockKitchenRepository is an autogenerated mock type for the KitchenRepository type
type MockKitchenRepository struct {
	mock.Mock
}

// ConfirmCreateTicket provides a mock function with given fields: ctx, ticketID
func (_m *MockKitchenRepository) ConfirmCreateTicket(ctx context.Context, ticketID string) error {
	ret := _m.Called(ctx, ticketID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ticketID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTicket provides a mock function with given fields: ctx, create
func (_m *MockKitchenRepository) CreateTicket(ctx context.Context, create CreateTicket) error {
	ret := _m.Called(ctx, create)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, CreateTicket) error); ok {
		r0 = rf(ctx, create)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockKitchenRepository creates a new instance of MockKitchenRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockKitchenRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockKitchenRepository {
	mock := &MockKitchenRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
