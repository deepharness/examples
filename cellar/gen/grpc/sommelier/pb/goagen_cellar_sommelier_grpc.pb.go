// Code generated with goa v3.16.1, DO NOT EDIT.
//
// sommelier protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/cellar/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: goagen_cellar_sommelier.proto

package sommelierpb

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
	Sommelier_Pick_FullMethodName = "/sommelier.Sommelier/Pick"
)

// SommelierClient is the client API for Sommelier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SommelierClient interface {
	// Pick implements pick.
	Pick(ctx context.Context, in *PickRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error)
}

type sommelierClient struct {
	cc grpc.ClientConnInterface
}

func NewSommelierClient(cc grpc.ClientConnInterface) SommelierClient {
	return &sommelierClient{cc}
}

func (c *sommelierClient) Pick(ctx context.Context, in *PickRequest, opts ...grpc.CallOption) (*StoredBottleCollection, error) {
	out := new(StoredBottleCollection)
	err := c.cc.Invoke(ctx, Sommelier_Pick_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SommelierServer is the server API for Sommelier service.
// All implementations must embed UnimplementedSommelierServer
// for forward compatibility
type SommelierServer interface {
	// Pick implements pick.
	Pick(context.Context, *PickRequest) (*StoredBottleCollection, error)
	mustEmbedUnimplementedSommelierServer()
}

// UnimplementedSommelierServer must be embedded to have forward compatible implementations.
type UnimplementedSommelierServer struct {
}

func (UnimplementedSommelierServer) Pick(context.Context, *PickRequest) (*StoredBottleCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pick not implemented")
}
func (UnimplementedSommelierServer) mustEmbedUnimplementedSommelierServer() {}

// UnsafeSommelierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SommelierServer will
// result in compilation errors.
type UnsafeSommelierServer interface {
	mustEmbedUnimplementedSommelierServer()
}

func RegisterSommelierServer(s grpc.ServiceRegistrar, srv SommelierServer) {
	s.RegisterService(&Sommelier_ServiceDesc, srv)
}

func _Sommelier_Pick_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PickRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SommelierServer).Pick(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sommelier_Pick_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SommelierServer).Pick(ctx, req.(*PickRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sommelier_ServiceDesc is the grpc.ServiceDesc for Sommelier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sommelier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sommelier.Sommelier",
	HandlerType: (*SommelierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pick",
			Handler:    _Sommelier_Pick_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goagen_cellar_sommelier.proto",
}
