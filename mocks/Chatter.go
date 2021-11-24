// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	api "github.com/gautamrege/gochat/api"
	mock "github.com/stretchr/testify/mock"
)

// Chatter is an autogenerated mock type for the Chatter type
type Chatter struct {
	mock.Mock
}

// Input provides a mock function with given fields:
func (_m *Chatter) Input() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Moderate provides a mock function with given fields: _a0
func (_m *Chatter) Moderate(_a0 api.ChatRequest) {
	_m.Called(_a0)
}

// Render provides a mock function with given fields: _a0
func (_m *Chatter) Render(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}