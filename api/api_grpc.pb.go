// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoChatClient is the client API for GoChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoChatClient interface {
	Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
}

type goChatClient struct {
	cc grpc.ClientConnInterface
}

func NewGoChatClient(cc grpc.ClientConnInterface) GoChatClient {
	return &goChatClient{cc}
}

func (c *goChatClient) Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, "/api.GoChat/Chat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoChatServer is the server API for GoChat service.
// All implementations must embed UnimplementedGoChatServer
// for forward compatibility
type GoChatServer interface {
	Chat(context.Context, *ChatRequest) (*ChatResponse, error)
	mustEmbedUnimplementedGoChatServer()
}

// UnimplementedGoChatServer must be embedded to have forward compatible implementations.
type UnimplementedGoChatServer struct {
}

func (UnimplementedGoChatServer) Chat(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedGoChatServer) mustEmbedUnimplementedGoChatServer() {}

// UnsafeGoChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoChatServer will
// result in compilation errors.
type UnsafeGoChatServer interface {
	mustEmbedUnimplementedGoChatServer()
}

func RegisterGoChatServer(s grpc.ServiceRegistrar, srv GoChatServer) {
	s.RegisterService(&GoChat_ServiceDesc, srv)
}

func _GoChat_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoChatServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GoChat/Chat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoChatServer).Chat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoChat_ServiceDesc is the grpc.ServiceDesc for GoChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.GoChat",
	HandlerType: (*GoChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Chat",
			Handler:    _GoChat_Chat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
