// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: authentication/v1/authentication.proto

package authentication

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Authentication_PublishToken_FullMethodName = "/auth.v1.Authentication/PublishToken"
	Authentication_RefreshToken_FullMethodName = "/auth.v1.Authentication/RefreshToken"
	Authentication_RevokeToken_FullMethodName  = "/auth.v1.Authentication/RevokeToken"
)

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	PublishToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PublishTokenResponse, error)
	RefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PublishTokenResponse, error)
	RevokeToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) PublishToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PublishTokenResponse, error) {
	out := new(PublishTokenResponse)
	err := c.cc.Invoke(ctx, Authentication_PublishToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RefreshToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PublishTokenResponse, error) {
	out := new(PublishTokenResponse)
	err := c.cc.Invoke(ctx, Authentication_RefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RevokeToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Authentication_RevokeToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	PublishToken(context.Context, *emptypb.Empty) (*PublishTokenResponse, error)
	RefreshToken(context.Context, *emptypb.Empty) (*PublishTokenResponse, error)
	RevokeToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) PublishToken(context.Context, *emptypb.Empty) (*PublishTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishToken not implemented")
}
func (UnimplementedAuthenticationServer) RefreshToken(context.Context, *emptypb.Empty) (*PublishTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) RevokeToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_PublishToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).PublishToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_PublishToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).PublishToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_RefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RefreshToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RevokeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RevokeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authentication_RevokeToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RevokeToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishToken",
			Handler:    _Authentication_PublishToken_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _Authentication_RefreshToken_Handler,
		},
		{
			MethodName: "RevokeToken",
			Handler:    _Authentication_RevokeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication/v1/authentication.proto",
}
