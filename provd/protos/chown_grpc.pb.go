// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: protos/chown.proto

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
	ChownService_ChownSettings_FullMethodName = "/chown.ChownService/ChownSettings"
)

// ChownServiceClient is the client API for ChownService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChownServiceClient interface {
	ChownSettings(ctx context.Context, in *ChownRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chownServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChownServiceClient(cc grpc.ClientConnInterface) ChownServiceClient {
	return &chownServiceClient{cc}
}

func (c *chownServiceClient) ChownSettings(ctx context.Context, in *ChownRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChownService_ChownSettings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChownServiceServer is the server API for ChownService service.
// All implementations must embed UnimplementedChownServiceServer
// for forward compatibility
type ChownServiceServer interface {
	ChownSettings(context.Context, *ChownRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChownServiceServer()
}

// UnimplementedChownServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChownServiceServer struct {
}

func (UnimplementedChownServiceServer) ChownSettings(context.Context, *ChownRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChownSettings not implemented")
}
func (UnimplementedChownServiceServer) mustEmbedUnimplementedChownServiceServer() {}

// UnsafeChownServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChownServiceServer will
// result in compilation errors.
type UnsafeChownServiceServer interface {
	mustEmbedUnimplementedChownServiceServer()
}

func RegisterChownServiceServer(s grpc.ServiceRegistrar, srv ChownServiceServer) {
	s.RegisterService(&ChownService_ServiceDesc, srv)
}

func _ChownService_ChownSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChownServiceServer).ChownSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChownService_ChownSettings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChownServiceServer).ChownSettings(ctx, req.(*ChownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChownService_ServiceDesc is the grpc.ServiceDesc for ChownService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChownService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chown.ChownService",
	HandlerType: (*ChownServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChownSettings",
			Handler:    _ChownService_ChownSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/chown.proto",
}
