// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: message/v1/message.proto

package message

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

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*Message, error)
	DeleteMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.message.v1.MessageService/CreateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.message.v1.MessageService/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/snack.message.v1.MessageService/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) DeleteMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.message.v1.MessageService/DeleteMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	CreateMessage(context.Context, *CreateMessageRequest) (*emptypb.Empty, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*emptypb.Empty, error)
	GetMessage(context.Context, *MessageID) (*Message, error)
	DeleteMessage(context.Context, *MessageID) (*emptypb.Empty, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedMessageServiceServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedMessageServiceServer) GetMessage(context.Context, *MessageID) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedMessageServiceServer) DeleteMessage(context.Context, *MessageID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMessage not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.MessageService/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.MessageService/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.MessageService/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetMessage(ctx, req.(*MessageID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_DeleteMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).DeleteMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.MessageService/DeleteMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).DeleteMessage(ctx, req.(*MessageID))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snack.message.v1.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _MessageService_CreateMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _MessageService_UpdateMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _MessageService_GetMessage_Handler,
		},
		{
			MethodName: "DeleteMessage",
			Handler:    _MessageService_DeleteMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/v1/message.proto",
}

// PinMessageServiceClient is the client API for PinMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PinMessageServiceClient interface {
	PinMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPinMessage(ctx context.Context, in *PinnedMessage, opts ...grpc.CallOption) (*PinnedMessage, error)
	UnPinMessage(ctx context.Context, in *PinnedMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pinMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPinMessageServiceClient(cc grpc.ClientConnInterface) PinMessageServiceClient {
	return &pinMessageServiceClient{cc}
}

func (c *pinMessageServiceClient) PinMessage(ctx context.Context, in *MessageID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.message.v1.PinMessageService/PinMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinMessageServiceClient) GetPinMessage(ctx context.Context, in *PinnedMessage, opts ...grpc.CallOption) (*PinnedMessage, error) {
	out := new(PinnedMessage)
	err := c.cc.Invoke(ctx, "/snack.message.v1.PinMessageService/GetPinMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pinMessageServiceClient) UnPinMessage(ctx context.Context, in *PinnedMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/snack.message.v1.PinMessageService/UnPinMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PinMessageServiceServer is the server API for PinMessageService service.
// All implementations must embed UnimplementedPinMessageServiceServer
// for forward compatibility
type PinMessageServiceServer interface {
	PinMessage(context.Context, *MessageID) (*emptypb.Empty, error)
	GetPinMessage(context.Context, *PinnedMessage) (*PinnedMessage, error)
	UnPinMessage(context.Context, *PinnedMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedPinMessageServiceServer()
}

// UnimplementedPinMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPinMessageServiceServer struct {
}

func (UnimplementedPinMessageServiceServer) PinMessage(context.Context, *MessageID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinMessage not implemented")
}
func (UnimplementedPinMessageServiceServer) GetPinMessage(context.Context, *PinnedMessage) (*PinnedMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPinMessage not implemented")
}
func (UnimplementedPinMessageServiceServer) UnPinMessage(context.Context, *PinnedMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnPinMessage not implemented")
}
func (UnimplementedPinMessageServiceServer) mustEmbedUnimplementedPinMessageServiceServer() {}

// UnsafePinMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PinMessageServiceServer will
// result in compilation errors.
type UnsafePinMessageServiceServer interface {
	mustEmbedUnimplementedPinMessageServiceServer()
}

func RegisterPinMessageServiceServer(s grpc.ServiceRegistrar, srv PinMessageServiceServer) {
	s.RegisterService(&PinMessageService_ServiceDesc, srv)
}

func _PinMessageService_PinMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinMessageServiceServer).PinMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.PinMessageService/PinMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinMessageServiceServer).PinMessage(ctx, req.(*MessageID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PinMessageService_GetPinMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinnedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinMessageServiceServer).GetPinMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.PinMessageService/GetPinMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinMessageServiceServer).GetPinMessage(ctx, req.(*PinnedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _PinMessageService_UnPinMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinnedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PinMessageServiceServer).UnPinMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/snack.message.v1.PinMessageService/UnPinMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PinMessageServiceServer).UnPinMessage(ctx, req.(*PinnedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// PinMessageService_ServiceDesc is the grpc.ServiceDesc for PinMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PinMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snack.message.v1.PinMessageService",
	HandlerType: (*PinMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PinMessage",
			Handler:    _PinMessageService_PinMessage_Handler,
		},
		{
			MethodName: "GetPinMessage",
			Handler:    _PinMessageService_GetPinMessage_Handler,
		},
		{
			MethodName: "UnPinMessage",
			Handler:    _PinMessageService_UnPinMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message/v1/message.proto",
}
