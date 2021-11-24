// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	api "github.com/gautamrege/gochat/api"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// GoChatClient is an autogenerated mock type for the GoChatClient type
type GoChatClient struct {
	mock.Mock
}

// Chat provides a mock function with given fields: ctx, in, opts
func (_m *GoChatClient) Chat(ctx context.Context, in *api.ChatRequest, opts ...grpc.CallOption) (*api.ChatResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *api.ChatResponse
	if rf, ok := ret.Get(0).(func(context.Context, *api.ChatRequest, ...grpc.CallOption) *api.ChatResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.ChatResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *api.ChatRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
