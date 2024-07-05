// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: redirector/ynison-redirector-service.proto

package redirector

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

// YnisonRedirectServiceClient is the client API for YnisonRedirectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YnisonRedirectServiceClient interface {
	GetRedirectToYnison(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error)
}

type ynisonRedirectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYnisonRedirectServiceClient(cc grpc.ClientConnInterface) YnisonRedirectServiceClient {
	return &ynisonRedirectServiceClient{cc}
}

func (c *ynisonRedirectServiceClient) GetRedirectToYnison(ctx context.Context, in *RedirectRequest, opts ...grpc.CallOption) (*RedirectResponse, error) {
	out := new(RedirectResponse)
	err := c.cc.Invoke(ctx, "/ynison_redirect.YnisonRedirectService/GetRedirectToYnison", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YnisonRedirectServiceServer is the server API for YnisonRedirectService service.
// All implementations must embed UnimplementedYnisonRedirectServiceServer
// for forward compatibility
type YnisonRedirectServiceServer interface {
	GetRedirectToYnison(context.Context, *RedirectRequest) (*RedirectResponse, error)
	mustEmbedUnimplementedYnisonRedirectServiceServer()
}

// UnimplementedYnisonRedirectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYnisonRedirectServiceServer struct {
}

func (UnimplementedYnisonRedirectServiceServer) GetRedirectToYnison(context.Context, *RedirectRequest) (*RedirectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRedirectToYnison not implemented")
}
func (UnimplementedYnisonRedirectServiceServer) mustEmbedUnimplementedYnisonRedirectServiceServer() {}

// UnsafeYnisonRedirectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YnisonRedirectServiceServer will
// result in compilation errors.
type UnsafeYnisonRedirectServiceServer interface {
	mustEmbedUnimplementedYnisonRedirectServiceServer()
}

func RegisterYnisonRedirectServiceServer(s grpc.ServiceRegistrar, srv YnisonRedirectServiceServer) {
	s.RegisterService(&YnisonRedirectService_ServiceDesc, srv)
}

func _YnisonRedirectService_GetRedirectToYnison_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedirectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YnisonRedirectServiceServer).GetRedirectToYnison(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ynison_redirect.YnisonRedirectService/GetRedirectToYnison",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YnisonRedirectServiceServer).GetRedirectToYnison(ctx, req.(*RedirectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YnisonRedirectService_ServiceDesc is the grpc.ServiceDesc for YnisonRedirectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YnisonRedirectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ynison_redirect.YnisonRedirectService",
	HandlerType: (*YnisonRedirectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRedirectToYnison",
			Handler:    _YnisonRedirectService_GetRedirectToYnison_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redirector/ynison-redirector-service.proto",
}
