// Code generated by mockery v2.33.0. DO NOT EDIT.

package am

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMessageHandler is an autogenerated mock type for the MessageHandler type
type MockMessageHandler[I IncomingMessage] struct {
	mock.Mock
}

// HandleMessage provides a mock function with given fields: ctx, msg
func (_m *MockMessageHandler[I]) HandleMessage(ctx context.Context, msg I) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, I) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockMessageHandler creates a new instance of MockMessageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMessageHandler[I IncomingMessage](t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMessageHandler[I] {
	mock := &MockMessageHandler[I]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
