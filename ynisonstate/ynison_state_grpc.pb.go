// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: ynisonstate/ynison_state.proto

package ynisonstate

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

// YnisonStateServiceClient is the client API for YnisonStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YnisonStateServiceClient interface {
	PutYnisonState(ctx context.Context, opts ...grpc.CallOption) (YnisonStateService_PutYnisonStateClient, error)
}

type ynisonStateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYnisonStateServiceClient(cc grpc.ClientConnInterface) YnisonStateServiceClient {
	return &ynisonStateServiceClient{cc}
}

func (c *ynisonStateServiceClient) PutYnisonState(ctx context.Context, opts ...grpc.CallOption) (YnisonStateService_PutYnisonStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &YnisonStateService_ServiceDesc.Streams[0], "/ynison_state.YnisonStateService/PutYnisonState", opts...)
	if err != nil {
		return nil, err
	}
	x := &ynisonStateServicePutYnisonStateClient{stream}
	return x, nil
}

type YnisonStateService_PutYnisonStateClient interface {
	Send(*PutYnisonStateRequest) error
	Recv() (*PutYnisonStateResponse, error)
	grpc.ClientStream
}

type ynisonStateServicePutYnisonStateClient struct {
	grpc.ClientStream
}

func (x *ynisonStateServicePutYnisonStateClient) Send(m *PutYnisonStateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ynisonStateServicePutYnisonStateClient) Recv() (*PutYnisonStateResponse, error) {
	m := new(PutYnisonStateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YnisonStateServiceServer is the server API for YnisonStateService service.
// All implementations must embed UnimplementedYnisonStateServiceServer
// for forward compatibility
type YnisonStateServiceServer interface {
	PutYnisonState(YnisonStateService_PutYnisonStateServer) error
	mustEmbedUnimplementedYnisonStateServiceServer()
}

// UnimplementedYnisonStateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYnisonStateServiceServer struct {
}

func (UnimplementedYnisonStateServiceServer) PutYnisonState(YnisonStateService_PutYnisonStateServer) error {
	return status.Errorf(codes.Unimplemented, "method PutYnisonState not implemented")
}
func (UnimplementedYnisonStateServiceServer) mustEmbedUnimplementedYnisonStateServiceServer() {}

// UnsafeYnisonStateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YnisonStateServiceServer will
// result in compilation errors.
type UnsafeYnisonStateServiceServer interface {
	mustEmbedUnimplementedYnisonStateServiceServer()
}

func RegisterYnisonStateServiceServer(s grpc.ServiceRegistrar, srv YnisonStateServiceServer) {
	s.RegisterService(&YnisonStateService_ServiceDesc, srv)
}

func _YnisonStateService_PutYnisonState_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(YnisonStateServiceServer).PutYnisonState(&ynisonStateServicePutYnisonStateServer{stream})
}

type YnisonStateService_PutYnisonStateServer interface {
	Send(*PutYnisonStateResponse) error
	Recv() (*PutYnisonStateRequest, error)
	grpc.ServerStream
}

type ynisonStateServicePutYnisonStateServer struct {
	grpc.ServerStream
}

func (x *ynisonStateServicePutYnisonStateServer) Send(m *PutYnisonStateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ynisonStateServicePutYnisonStateServer) Recv() (*PutYnisonStateRequest, error) {
	m := new(PutYnisonStateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// YnisonStateService_ServiceDesc is the grpc.ServiceDesc for YnisonStateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YnisonStateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ynison_state.YnisonStateService",
	HandlerType: (*YnisonStateServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutYnisonState",
			Handler:       _YnisonStateService_PutYnisonState_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ynisonstate/ynison_state.proto",
}