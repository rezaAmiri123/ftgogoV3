// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	application "github.com/rezaAmiri123/ftgogoV3/consumer/internal/application"

	domain "github.com/rezaAmiri123/ftgogoV3/consumer/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// GetConsumer provides a mock function with given fields: ctx, get
func (_m *App) GetConsumer(ctx context.Context, get application.GetConsumer) (*domain.Consumer, error) {
	ret := _m.Called(ctx, get)

	var r0 *domain.Consumer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, application.GetConsumer) (*domain.Consumer, error)); ok {
		return rf(ctx, get)
	}
	if rf, ok := ret.Get(0).(func(context.Context, application.GetConsumer) *domain.Consumer); ok {
		r0 = rf(ctx, get)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Consumer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, application.GetConsumer) error); ok {
		r1 = rf(ctx, get)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConsumerAddress provides a mock function with given fields: ctx, get
func (_m *App) GetConsumerAddress(ctx context.Context, get application.GetConsumerAddress) (domain.Address, error) {
	ret := _m.Called(ctx, get)

	var r0 domain.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, application.GetConsumerAddress) (domain.Address, error)); ok {
		return rf(ctx, get)
	}
	if rf, ok := ret.Get(0).(func(context.Context, application.GetConsumerAddress) domain.Address); ok {
		r0 = rf(ctx, get)
	} else {
		r0 = ret.Get(0).(domain.Address)
	}

	if rf, ok := ret.Get(1).(func(context.Context, application.GetConsumerAddress) error); ok {
		r1 = rf(ctx, get)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterConsumer provides a mock function with given fields: ctx, register
func (_m *App) RegisterConsumer(ctx context.Context, register application.RegisterConsumer) error {
	ret := _m.Called(ctx, register)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, application.RegisterConsumer) error); ok {
		r0 = rf(ctx, register)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveConsumerAddress provides a mock function with given fields: ctx, remove
func (_m *App) RemoveConsumerAddress(ctx context.Context, remove application.RemoveConsumerAddress) error {
	ret := _m.Called(ctx, remove)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, application.RemoveConsumerAddress) error); ok {
		r0 = rf(ctx, remove)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateConsumerAddress provides a mock function with given fields: ctx, update
func (_m *App) UpdateConsumerAddress(ctx context.Context, update application.UpdateConsumerAddress) error {
	ret := _m.Called(ctx, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, application.UpdateConsumerAddress) error); ok {
		r0 = rf(ctx, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateOrderByConsumer provides a mock function with given fields: ctx, validate
func (_m *App) ValidateOrderByConsumer(ctx context.Context, validate application.ValidateOrderByConsumer) error {
	ret := _m.Called(ctx, validate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, application.ValidateOrderByConsumer) error); ok {
		r0 = rf(ctx, validate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewApp creates a new instance of App. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewApp(t interface {
	mock.TestingT
	Cleanup(func())
}) *App {
	mock := &App{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}