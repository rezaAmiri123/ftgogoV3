// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: kitchenpb/kitchen_api.proto

package kitchenpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	KitchenService_CreateTicket_FullMethodName        = "/kitchenpb.KitchenService/CreateTicket"
	KitchenService_GetTicket_FullMethodName           = "/kitchenpb.KitchenService/GetTicket"
	KitchenService_GetRestaurant_FullMethodName       = "/kitchenpb.KitchenService/GetRestaurant"
	KitchenService_AcceptTicket_FullMethodName        = "/kitchenpb.KitchenService/AcceptTicket"
	KitchenService_ConfirmCreateTicket_FullMethodName = "/kitchenpb.KitchenService/ConfirmCreateTicket"
)

// KitchenServiceClient is the client API for KitchenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitchenServiceClient interface {
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error)
	GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*GetRestaurantResponse, error)
	AcceptTicket(ctx context.Context, in *AcceptTicketRequest, opts ...grpc.CallOption) (*AcceptTicketResponse, error)
	ConfirmCreateTicket(ctx context.Context, in *ConfirmCreateTicketRequest, opts ...grpc.CallOption) (*ConfirmCreateTicketResponse, error)
}

type kitchenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKitchenServiceClient(cc grpc.ClientConnInterface) KitchenServiceClient {
	return &kitchenServiceClient{cc}
}

func (c *kitchenServiceClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, KitchenService_CreateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTicketResponse)
	err := c.cc.Invoke(ctx, KitchenService_GetTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenServiceClient) GetRestaurant(ctx context.Context, in *GetRestaurantRequest, opts ...grpc.CallOption) (*GetRestaurantResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRestaurantResponse)
	err := c.cc.Invoke(ctx, KitchenService_GetRestaurant_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenServiceClient) AcceptTicket(ctx context.Context, in *AcceptTicketRequest, opts ...grpc.CallOption) (*AcceptTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AcceptTicketResponse)
	err := c.cc.Invoke(ctx, KitchenService_AcceptTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kitchenServiceClient) ConfirmCreateTicket(ctx context.Context, in *ConfirmCreateTicketRequest, opts ...grpc.CallOption) (*ConfirmCreateTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConfirmCreateTicketResponse)
	err := c.cc.Invoke(ctx, KitchenService_ConfirmCreateTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KitchenServiceServer is the server API for KitchenService service.
// All implementations must embed UnimplementedKitchenServiceServer
// for forward compatibility.
type KitchenServiceServer interface {
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error)
	GetRestaurant(context.Context, *GetRestaurantRequest) (*GetRestaurantResponse, error)
	AcceptTicket(context.Context, *AcceptTicketRequest) (*AcceptTicketResponse, error)
	ConfirmCreateTicket(context.Context, *ConfirmCreateTicketRequest) (*ConfirmCreateTicketResponse, error)
	mustEmbedUnimplementedKitchenServiceServer()
}

// UnimplementedKitchenServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedKitchenServiceServer struct{}

func (UnimplementedKitchenServiceServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedKitchenServiceServer) GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedKitchenServiceServer) GetRestaurant(context.Context, *GetRestaurantRequest) (*GetRestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}
func (UnimplementedKitchenServiceServer) AcceptTicket(context.Context, *AcceptTicketRequest) (*AcceptTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptTicket not implemented")
}
func (UnimplementedKitchenServiceServer) ConfirmCreateTicket(context.Context, *ConfirmCreateTicketRequest) (*ConfirmCreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmCreateTicket not implemented")
}
func (UnimplementedKitchenServiceServer) mustEmbedUnimplementedKitchenServiceServer() {}
func (UnimplementedKitchenServiceServer) testEmbeddedByValue()                        {}

// UnsafeKitchenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitchenServiceServer will
// result in compilation errors.
type UnsafeKitchenServiceServer interface {
	mustEmbedUnimplementedKitchenServiceServer()
}

func RegisterKitchenServiceServer(s grpc.ServiceRegistrar, srv KitchenServiceServer) {
	// If the following call pancis, it indicates UnimplementedKitchenServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&KitchenService_ServiceDesc, srv)
}

func _KitchenService_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KitchenService_CreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitchenService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KitchenService_GetTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitchenService_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KitchenService_GetRestaurant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).GetRestaurant(ctx, req.(*GetRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitchenService_AcceptTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).AcceptTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KitchenService_AcceptTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).AcceptTicket(ctx, req.(*AcceptTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KitchenService_ConfirmCreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmCreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServiceServer).ConfirmCreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KitchenService_ConfirmCreateTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServiceServer).ConfirmCreateTicket(ctx, req.(*ConfirmCreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KitchenService_ServiceDesc is the grpc.ServiceDesc for KitchenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KitchenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitchenpb.KitchenService",
	HandlerType: (*KitchenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTicket",
			Handler:    _KitchenService_CreateTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _KitchenService_GetTicket_Handler,
		},
		{
			MethodName: "GetRestaurant",
			Handler:    _KitchenService_GetRestaurant_Handler,
		},
		{
			MethodName: "AcceptTicket",
			Handler:    _KitchenService_AcceptTicket_Handler,
		},
		{
			MethodName: "ConfirmCreateTicket",
			Handler:    _KitchenService_ConfirmCreateTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kitchenpb/kitchen_api.proto",
}
