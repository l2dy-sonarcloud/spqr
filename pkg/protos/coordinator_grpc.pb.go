// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: protos/coordinator.proto

package proto

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
	TopologyService_OpenRouter_FullMethodName        = "/spqr.TopologyService/OpenRouter"
	TopologyService_GetRouterStatus_FullMethodName   = "/spqr.TopologyService/GetRouterStatus"
	TopologyService_CloseRouter_FullMethodName       = "/spqr.TopologyService/CloseRouter"
	TopologyService_UpdateCoordinator_FullMethodName = "/spqr.TopologyService/UpdateCoordinator"
	TopologyService_GetCoordinator_FullMethodName    = "/spqr.TopologyService/GetCoordinator"
)

// TopologyServiceClient is the client API for TopologyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TopologyServiceClient interface {
	OpenRouter(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRouterStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRouterStatusReply, error)
	CloseRouter(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCoordinator(ctx context.Context, in *UpdateCoordinatorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetCoordinator(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCoordinatorResponse, error)
}

type topologyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopologyServiceClient(cc grpc.ClientConnInterface) TopologyServiceClient {
	return &topologyServiceClient{cc}
}

func (c *topologyServiceClient) OpenRouter(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TopologyService_OpenRouter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyServiceClient) GetRouterStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRouterStatusReply, error) {
	out := new(GetRouterStatusReply)
	err := c.cc.Invoke(ctx, TopologyService_GetRouterStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyServiceClient) CloseRouter(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TopologyService_CloseRouter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyServiceClient) UpdateCoordinator(ctx context.Context, in *UpdateCoordinatorRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TopologyService_UpdateCoordinator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topologyServiceClient) GetCoordinator(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetCoordinatorResponse, error) {
	out := new(GetCoordinatorResponse)
	err := c.cc.Invoke(ctx, TopologyService_GetCoordinator_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopologyServiceServer is the server API for TopologyService service.
// All implementations must embed UnimplementedTopologyServiceServer
// for forward compatibility
type TopologyServiceServer interface {
	OpenRouter(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetRouterStatus(context.Context, *emptypb.Empty) (*GetRouterStatusReply, error)
	CloseRouter(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	UpdateCoordinator(context.Context, *UpdateCoordinatorRequest) (*emptypb.Empty, error)
	GetCoordinator(context.Context, *emptypb.Empty) (*GetCoordinatorResponse, error)
	mustEmbedUnimplementedTopologyServiceServer()
}

// UnimplementedTopologyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTopologyServiceServer struct {
}

func (UnimplementedTopologyServiceServer) OpenRouter(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OpenRouter not implemented")
}
func (UnimplementedTopologyServiceServer) GetRouterStatus(context.Context, *emptypb.Empty) (*GetRouterStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouterStatus not implemented")
}
func (UnimplementedTopologyServiceServer) CloseRouter(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseRouter not implemented")
}
func (UnimplementedTopologyServiceServer) UpdateCoordinator(context.Context, *UpdateCoordinatorRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoordinator not implemented")
}
func (UnimplementedTopologyServiceServer) GetCoordinator(context.Context, *emptypb.Empty) (*GetCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoordinator not implemented")
}
func (UnimplementedTopologyServiceServer) mustEmbedUnimplementedTopologyServiceServer() {}

// UnsafeTopologyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TopologyServiceServer will
// result in compilation errors.
type UnsafeTopologyServiceServer interface {
	mustEmbedUnimplementedTopologyServiceServer()
}

func RegisterTopologyServiceServer(s grpc.ServiceRegistrar, srv TopologyServiceServer) {
	s.RegisterService(&TopologyService_ServiceDesc, srv)
}

func _TopologyService_OpenRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).OpenRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_OpenRouter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).OpenRouter(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyService_GetRouterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).GetRouterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_GetRouterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).GetRouterStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyService_CloseRouter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).CloseRouter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_CloseRouter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).CloseRouter(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyService_UpdateCoordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCoordinatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).UpdateCoordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_UpdateCoordinator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).UpdateCoordinator(ctx, req.(*UpdateCoordinatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopologyService_GetCoordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopologyServiceServer).GetCoordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TopologyService_GetCoordinator_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopologyServiceServer).GetCoordinator(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TopologyService_ServiceDesc is the grpc.ServiceDesc for TopologyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TopologyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.TopologyService",
	HandlerType: (*TopologyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OpenRouter",
			Handler:    _TopologyService_OpenRouter_Handler,
		},
		{
			MethodName: "GetRouterStatus",
			Handler:    _TopologyService_GetRouterStatus_Handler,
		},
		{
			MethodName: "CloseRouter",
			Handler:    _TopologyService_CloseRouter_Handler,
		},
		{
			MethodName: "UpdateCoordinator",
			Handler:    _TopologyService_UpdateCoordinator_Handler,
		},
		{
			MethodName: "GetCoordinator",
			Handler:    _TopologyService_GetCoordinator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/coordinator.proto",
}
