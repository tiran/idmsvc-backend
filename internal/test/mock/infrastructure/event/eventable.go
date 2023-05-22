// Code generated by mockery v2.16.0. DO NOT EDIT.

package event

import (
	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
	mock "github.com/stretchr/testify/mock"
)

// Eventable is an autogenerated mock type for the Eventable type
type Eventable struct {
	mock.Mock
}

// OnMessage provides a mock function with given fields: msg
func (_m *Eventable) OnMessage(msg *kafka.Message) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*kafka.Message) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEventable interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventable creates a new instance of Eventable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventable(t mockConstructorTestingTNewEventable) *Eventable {
	mock := &Eventable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
