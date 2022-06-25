// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	fasthttp "github.com/valyala/fasthttp"

	mock "github.com/stretchr/testify/mock"
)

// AuthMwInterface is an autogenerated mock type for the AuthMwInterface type
type AuthMwInterface struct {
	mock.Mock
}

// VerifyAuth provides a mock function with given fields: ctx
func (_m *AuthMwInterface) VerifyAuth(ctx *fasthttp.RequestCtx) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fasthttp.RequestCtx) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuthMwInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthMwInterface creates a new instance of AuthMwInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthMwInterface(t mockConstructorTestingTNewAuthMwInterface) *AuthMwInterface {
	mock := &AuthMwInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}