// Code generated by mockery v2.16.0. DO NOT EDIT.

package presenter

import (
	model "github.com/podengo-project/idmsvc-backend/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	public "github.com/podengo-project/idmsvc-backend/internal/api/public"
)

// HostPresenter is an autogenerated mock type for the HostPresenter type
type HostPresenter struct {
	mock.Mock
}

// HostConf provides a mock function with given fields: domain
func (_m *HostPresenter) HostConf(domain *model.Domain) (*public.HostConfResponseSchema, error) {
	ret := _m.Called(domain)

	var r0 *public.HostConfResponseSchema
	if rf, ok := ret.Get(0).(func(*model.Domain) *public.HostConfResponseSchema); ok {
		r0 = rf(domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*public.HostConfResponseSchema)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHostPresenter interface {
	mock.TestingT
	Cleanup(func())
}

// NewHostPresenter creates a new instance of HostPresenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHostPresenter(t mockConstructorTestingTNewHostPresenter) *HostPresenter {
	mock := &HostPresenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}