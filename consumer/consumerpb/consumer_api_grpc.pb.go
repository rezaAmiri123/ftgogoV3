// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: consumerpb/consumer_api.proto

package consumerpb

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

const (
	ConsumerService_RegisterConsumer_FullMethodName        = "/consumerpb.ConsumerService/RegisterConsumer"
	ConsumerService_GetConsumer_FullMethodName             = "/consumerpb.ConsumerService/GetConsumer"
	ConsumerService_UpdateConsumer_FullMethodName          = "/consumerpb.ConsumerService/UpdateConsumer"
	ConsumerService_GetAddress_FullMethodName              = "/consumerpb.ConsumerService/GetAddress"
	ConsumerService_UpdateAddress_FullMethodName           = "/consumerpb.ConsumerService/UpdateAddress"
	ConsumerService_RemoveAddress_FullMethodName           = "/consumerpb.ConsumerService/RemoveAddress"
	ConsumerService_ValidateOrderByConsumer_FullMethodName = "/consumerpb.ConsumerService/ValidateOrderByConsumer"
)

// ConsumerServiceClient is the client API for ConsumerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerServiceClient interface {
	RegisterConsumer(ctx context.Context, in *RegisterConsumerRequest, opts ...grpc.CallOption) (*RegisterConsumerResponse, error)
	GetConsumer(ctx context.Context, in *GetConsumerRequest, opts ...grpc.CallOption) (*GetConsumerResponse, error)
	UpdateConsumer(ctx context.Context, in *UpdateConsumerRequest, opts ...grpc.CallOption) (*UpdateConsumerResponse, error)
	GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error)
	UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error)
	RemoveAddress(ctx context.Context, in *RemoveAddressRequest, opts ...grpc.CallOption) (*RemoveAddressResponse, error)
	ValidateOrderByConsumer(ctx context.Context, in *ValidateOrderByConsumerRequest, opts ...grpc.CallOption) (*ValidateOrderByConsumerResponse, error)
}

type consumerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerServiceClient(cc grpc.ClientConnInterface) ConsumerServiceClient {
	return &consumerServiceClient{cc}
}

func (c *consumerServiceClient) RegisterConsumer(ctx context.Context, in *RegisterConsumerRequest, opts ...grpc.CallOption) (*RegisterConsumerResponse, error) {
	out := new(RegisterConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_RegisterConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetConsumer(ctx context.Context, in *GetConsumerRequest, opts ...grpc.CallOption) (*GetConsumerResponse, error) {
	out := new(GetConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) UpdateConsumer(ctx context.Context, in *UpdateConsumerRequest, opts ...grpc.CallOption) (*UpdateConsumerResponse, error) {
	out := new(UpdateConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_UpdateConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) GetAddress(ctx context.Context, in *GetAddressRequest, opts ...grpc.CallOption) (*GetAddressResponse, error) {
	out := new(GetAddressResponse)
	err := c.cc.Invoke(ctx, ConsumerService_GetAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) UpdateAddress(ctx context.Context, in *UpdateAddressRequest, opts ...grpc.CallOption) (*UpdateAddressResponse, error) {
	out := new(UpdateAddressResponse)
	err := c.cc.Invoke(ctx, ConsumerService_UpdateAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) RemoveAddress(ctx context.Context, in *RemoveAddressRequest, opts ...grpc.CallOption) (*RemoveAddressResponse, error) {
	out := new(RemoveAddressResponse)
	err := c.cc.Invoke(ctx, ConsumerService_RemoveAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerServiceClient) ValidateOrderByConsumer(ctx context.Context, in *ValidateOrderByConsumerRequest, opts ...grpc.CallOption) (*ValidateOrderByConsumerResponse, error) {
	out := new(ValidateOrderByConsumerResponse)
	err := c.cc.Invoke(ctx, ConsumerService_ValidateOrderByConsumer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerServiceServer is the server API for ConsumerService service.
// All implementations must embed UnimplementedConsumerServiceServer
// for forward compatibility
type ConsumerServiceServer interface {
	RegisterConsumer(context.Context, *RegisterConsumerRequest) (*RegisterConsumerResponse, error)
	GetConsumer(context.Context, *GetConsumerRequest) (*GetConsumerResponse, error)
	UpdateConsumer(context.Context, *UpdateConsumerRequest) (*UpdateConsumerResponse, error)
	GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error)
	UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error)
	RemoveAddress(context.Context, *RemoveAddressRequest) (*RemoveAddressResponse, error)
	ValidateOrderByConsumer(context.Context, *ValidateOrderByConsumerRequest) (*ValidateOrderByConsumerResponse, error)
	mustEmbedUnimplementedConsumerServiceServer()
}

// UnimplementedConsumerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerServiceServer struct {
}

func (UnimplementedConsumerServiceServer) RegisterConsumer(context.Context, *RegisterConsumerRequest) (*RegisterConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) GetConsumer(context.Context, *GetConsumerRequest) (*GetConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) UpdateConsumer(context.Context, *UpdateConsumerRequest) (*UpdateConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) GetAddress(context.Context, *GetAddressRequest) (*GetAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddress not implemented")
}
func (UnimplementedConsumerServiceServer) UpdateAddress(context.Context, *UpdateAddressRequest) (*UpdateAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAddress not implemented")
}
func (UnimplementedConsumerServiceServer) RemoveAddress(context.Context, *RemoveAddressRequest) (*RemoveAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAddress not implemented")
}
func (UnimplementedConsumerServiceServer) ValidateOrderByConsumer(context.Context, *ValidateOrderByConsumerRequest) (*ValidateOrderByConsumerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateOrderByConsumer not implemented")
}
func (UnimplementedConsumerServiceServer) mustEmbedUnimplementedConsumerServiceServer() {}

// UnsafeConsumerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerServiceServer will
// result in compilation errors.
type UnsafeConsumerServiceServer interface {
	mustEmbedUnimplementedConsumerServiceServer()
}

func RegisterConsumerServiceServer(s grpc.ServiceRegistrar, srv ConsumerServiceServer) {
	s.RegisterService(&ConsumerService_ServiceDesc, srv)
}

func _ConsumerService_RegisterConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).RegisterConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_RegisterConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).RegisterConsumer(ctx, req.(*RegisterConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetConsumer(ctx, req.(*GetConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_UpdateConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).UpdateConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_UpdateConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).UpdateConsumer(ctx, req.(*UpdateConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_GetAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).GetAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_GetAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).GetAddress(ctx, req.(*GetAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_UpdateAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).UpdateAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_UpdateAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).UpdateAddress(ctx, req.(*UpdateAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_RemoveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).RemoveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_RemoveAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).RemoveAddress(ctx, req.(*RemoveAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerService_ValidateOrderByConsumer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateOrderByConsumerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerServiceServer).ValidateOrderByConsumer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConsumerService_ValidateOrderByConsumer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerServiceServer).ValidateOrderByConsumer(ctx, req.(*ValidateOrderByConsumerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsumerService_ServiceDesc is the grpc.ServiceDesc for ConsumerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsumerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consumerpb.ConsumerService",
	HandlerType: (*ConsumerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterConsumer",
			Handler:    _ConsumerService_RegisterConsumer_Handler,
		},
		{
			MethodName: "GetConsumer",
			Handler:    _ConsumerService_GetConsumer_Handler,
		},
		{
			MethodName: "UpdateConsumer",
			Handler:    _ConsumerService_UpdateConsumer_Handler,
		},
		{
			MethodName: "GetAddress",
			Handler:    _ConsumerService_GetAddress_Handler,
		},
		{
			MethodName: "UpdateAddress",
			Handler:    _ConsumerService_UpdateAddress_Handler,
		},
		{
			MethodName: "RemoveAddress",
			Handler:    _ConsumerService_RemoveAddress_Handler,
		},
		{
			MethodName: "ValidateOrderByConsumer",
			Handler:    _ConsumerService_ValidateOrderByConsumer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consumerpb/consumer_api.proto",
}
