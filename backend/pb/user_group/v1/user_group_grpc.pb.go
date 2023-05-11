// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user_group/v1/user_group.proto

package user_group

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

// UserGroupServiceClient is the client API for UserGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserGroupServiceClient interface {
	CreateUserGroup(ctx context.Context, in *CreateUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserGroup(ctx context.Context, in *UserGroupID, opts ...grpc.CallOption) (*UserGroup, error)
	UpdateUserGroup(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteUserGroup(ctx context.Context, in *UserGroupID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserGroupServiceClient(cc grpc.ClientConnInterface) UserGroupServiceClient {
	return &userGroupServiceClient{cc}
}

func (c *userGroupServiceClient) CreateUserGroup(ctx context.Context, in *CreateUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.UserGroupService/CreateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupServiceClient) GetUserGroup(ctx context.Context, in *UserGroupID, opts ...grpc.CallOption) (*UserGroup, error) {
	out := new(UserGroup)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.UserGroupService/GetUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupServiceClient) UpdateUserGroup(ctx context.Context, in *UpdateUserGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.UserGroupService/UpdateUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userGroupServiceClient) DeleteUserGroup(ctx context.Context, in *UserGroupID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.UserGroupService/DeleteUserGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserGroupServiceServer is the server API for UserGroupService service.
// All implementations must embed UnimplementedUserGroupServiceServer
// for forward compatibility
type UserGroupServiceServer interface {
	CreateUserGroup(context.Context, *CreateUserGroupRequest) (*emptypb.Empty, error)
	GetUserGroup(context.Context, *UserGroupID) (*UserGroup, error)
	UpdateUserGroup(context.Context, *UpdateUserGroupRequest) (*emptypb.Empty, error)
	DeleteUserGroup(context.Context, *UserGroupID) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserGroupServiceServer()
}

// UnimplementedUserGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserGroupServiceServer struct {
}

func (UnimplementedUserGroupServiceServer) CreateUserGroup(context.Context, *CreateUserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserGroup not implemented")
}
func (UnimplementedUserGroupServiceServer) GetUserGroup(context.Context, *UserGroupID) (*UserGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGroup not implemented")
}
func (UnimplementedUserGroupServiceServer) UpdateUserGroup(context.Context, *UpdateUserGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserGroup not implemented")
}
func (UnimplementedUserGroupServiceServer) DeleteUserGroup(context.Context, *UserGroupID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGroup not implemented")
}
func (UnimplementedUserGroupServiceServer) mustEmbedUnimplementedUserGroupServiceServer() {}

// UnsafeUserGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserGroupServiceServer will
// result in compilation errors.
type UnsafeUserGroupServiceServer interface {
	mustEmbedUnimplementedUserGroupServiceServer()
}

func RegisterUserGroupServiceServer(s grpc.ServiceRegistrar, srv UserGroupServiceServer) {
	s.RegisterService(&UserGroupService_ServiceDesc, srv)
}

func _UserGroupService_CreateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServiceServer).CreateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.UserGroupService/CreateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServiceServer).CreateUserGroup(ctx, req.(*CreateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroupService_GetUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServiceServer).GetUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.UserGroupService/GetUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServiceServer).GetUserGroup(ctx, req.(*UserGroupID))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroupService_UpdateUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServiceServer).UpdateUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.UserGroupService/UpdateUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServiceServer).UpdateUserGroup(ctx, req.(*UpdateUserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserGroupService_DeleteUserGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserGroupServiceServer).DeleteUserGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.UserGroupService/DeleteUserGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserGroupServiceServer).DeleteUserGroup(ctx, req.(*UserGroupID))
	}
	return interceptor(ctx, in, info, handler)
}

// UserGroupService_ServiceDesc is the grpc.ServiceDesc for UserGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snack.user_group.v1.UserGroupService",
	HandlerType: (*UserGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserGroup",
			Handler:    _UserGroupService_CreateUserGroup_Handler,
		},
		{
			MethodName: "GetUserGroup",
			Handler:    _UserGroupService_GetUserGroup_Handler,
		},
		{
			MethodName: "UpdateUserGroup",
			Handler:    _UserGroupService_UpdateUserGroup_Handler,
		},
		{
			MethodName: "DeleteUserGroup",
			Handler:    _UserGroupService_DeleteUserGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_group/v1/user_group.proto",
}

// MemberServiceClient is the client API for MemberService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MemberServiceClient interface {
	AddUserToMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveUserFromMember(ctx context.Context, in *RemoveUserFromMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ChangeMemberGrant(ctx context.Context, in *Member, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type memberServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMemberServiceClient(cc grpc.ClientConnInterface) MemberServiceClient {
	return &memberServiceClient{cc}
}

func (c *memberServiceClient) AddUserToMember(ctx context.Context, in *Member, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.MemberService/AddUserToMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) RemoveUserFromMember(ctx context.Context, in *RemoveUserFromMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.MemberService/RemoveUserFromMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *memberServiceClient) ChangeMemberGrant(ctx context.Context, in *Member, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.user_group.v1.MemberService/ChangeMemberGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MemberServiceServer is the server API for MemberService service.
// All implementations must embed UnimplementedMemberServiceServer
// for forward compatibility
type MemberServiceServer interface {
	AddUserToMember(context.Context, *Member) (*emptypb.Empty, error)
	RemoveUserFromMember(context.Context, *RemoveUserFromMemberRequest) (*emptypb.Empty, error)
	ChangeMemberGrant(context.Context, *Member) (*emptypb.Empty, error)
	mustEmbedUnimplementedMemberServiceServer()
}

// UnimplementedMemberServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMemberServiceServer struct {
}

func (UnimplementedMemberServiceServer) AddUserToMember(context.Context, *Member) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToMember not implemented")
}
func (UnimplementedMemberServiceServer) RemoveUserFromMember(context.Context, *RemoveUserFromMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromMember not implemented")
}
func (UnimplementedMemberServiceServer) ChangeMemberGrant(context.Context, *Member) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeMemberGrant not implemented")
}
func (UnimplementedMemberServiceServer) mustEmbedUnimplementedMemberServiceServer() {}

// UnsafeMemberServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MemberServiceServer will
// result in compilation errors.
type UnsafeMemberServiceServer interface {
	mustEmbedUnimplementedMemberServiceServer()
}

func RegisterMemberServiceServer(s grpc.ServiceRegistrar, srv MemberServiceServer) {
	s.RegisterService(&MemberService_ServiceDesc, srv)
}

func _MemberService_AddUserToMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).AddUserToMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.MemberService/AddUserToMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).AddUserToMember(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_RemoveUserFromMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).RemoveUserFromMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.MemberService/RemoveUserFromMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).RemoveUserFromMember(ctx, req.(*RemoveUserFromMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MemberService_ChangeMemberGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Member)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MemberServiceServer).ChangeMemberGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.user_group.v1.MemberService/ChangeMemberGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MemberServiceServer).ChangeMemberGrant(ctx, req.(*Member))
	}
	return interceptor(ctx, in, info, handler)
}

// MemberService_ServiceDesc is the grpc.ServiceDesc for MemberService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MemberService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snack.user_group.v1.MemberService",
	HandlerType: (*MemberServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserToMember",
			Handler:    _MemberService_AddUserToMember_Handler,
		},
		{
			MethodName: "RemoveUserFromMember",
			Handler:    _MemberService_RemoveUserFromMember_Handler,
		},
		{
			MethodName: "ChangeMemberGrant",
			Handler:    _MemberService_ChangeMemberGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_group/v1/user_group.proto",
}