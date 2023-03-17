// Code generated with goa v3.11.2, DO NOT EDIT.
//
// chatter protocol buffer definition
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.1
// source: goadesign_goagen_chatter.proto

package chatterpb

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
	Chatter_Login_FullMethodName     = "/chatter.Chatter/Login"
	Chatter_Echoer_FullMethodName    = "/chatter.Chatter/Echoer"
	Chatter_Listener_FullMethodName  = "/chatter.Chatter/Listener"
	Chatter_Summary_FullMethodName   = "/chatter.Chatter/Summary"
	Chatter_Subscribe_FullMethodName = "/chatter.Chatter/Subscribe"
	Chatter_History_FullMethodName   = "/chatter.Chatter/History"
)

// ChatterClient is the client API for Chatter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatterClient interface {
	// Creates a valid JWT token for auth to chat.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Echoes the message sent by the client.
	Echoer(ctx context.Context, opts ...grpc.CallOption) (Chatter_EchoerClient, error)
	// Listens to the messages sent by the client.
	Listener(ctx context.Context, opts ...grpc.CallOption) (Chatter_ListenerClient, error)
	// Summarizes the chat messages sent by the client.
	Summary(ctx context.Context, opts ...grpc.CallOption) (Chatter_SummaryClient, error)
	// Subscribe to events sent when new chat messages are added.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Chatter_SubscribeClient, error)
	// Returns the chat messages sent to the server.
	History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (Chatter_HistoryClient, error)
}

type chatterClient struct {
	cc grpc.ClientConnInterface
}

func NewChatterClient(cc grpc.ClientConnInterface) ChatterClient {
	return &chatterClient{cc}
}

func (c *chatterClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, Chatter_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatterClient) Echoer(ctx context.Context, opts ...grpc.CallOption) (Chatter_EchoerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chatter_ServiceDesc.Streams[0], Chatter_Echoer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterEchoerClient{stream}
	return x, nil
}

type Chatter_EchoerClient interface {
	Send(*EchoerStreamingRequest) error
	Recv() (*EchoerResponse, error)
	grpc.ClientStream
}

type chatterEchoerClient struct {
	grpc.ClientStream
}

func (x *chatterEchoerClient) Send(m *EchoerStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatterEchoerClient) Recv() (*EchoerResponse, error) {
	m := new(EchoerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) Listener(ctx context.Context, opts ...grpc.CallOption) (Chatter_ListenerClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chatter_ServiceDesc.Streams[1], Chatter_Listener_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterListenerClient{stream}
	return x, nil
}

type Chatter_ListenerClient interface {
	Send(*ListenerStreamingRequest) error
	CloseAndRecv() (*ListenerResponse, error)
	grpc.ClientStream
}

type chatterListenerClient struct {
	grpc.ClientStream
}

func (x *chatterListenerClient) Send(m *ListenerStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatterListenerClient) CloseAndRecv() (*ListenerResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ListenerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) Summary(ctx context.Context, opts ...grpc.CallOption) (Chatter_SummaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chatter_ServiceDesc.Streams[2], Chatter_Summary_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterSummaryClient{stream}
	return x, nil
}

type Chatter_SummaryClient interface {
	Send(*SummaryStreamingRequest) error
	CloseAndRecv() (*ChatSummaryCollection, error)
	grpc.ClientStream
}

type chatterSummaryClient struct {
	grpc.ClientStream
}

func (x *chatterSummaryClient) Send(m *SummaryStreamingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatterSummaryClient) CloseAndRecv() (*ChatSummaryCollection, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ChatSummaryCollection)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Chatter_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chatter_ServiceDesc.Streams[3], Chatter_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chatter_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type chatterSubscribeClient struct {
	grpc.ClientStream
}

func (x *chatterSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatterClient) History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (Chatter_HistoryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chatter_ServiceDesc.Streams[4], Chatter_History_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatterHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chatter_HistoryClient interface {
	Recv() (*HistoryResponse, error)
	grpc.ClientStream
}

type chatterHistoryClient struct {
	grpc.ClientStream
}

func (x *chatterHistoryClient) Recv() (*HistoryResponse, error) {
	m := new(HistoryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatterServer is the server API for Chatter service.
// All implementations must embed UnimplementedChatterServer
// for forward compatibility
type ChatterServer interface {
	// Creates a valid JWT token for auth to chat.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Echoes the message sent by the client.
	Echoer(Chatter_EchoerServer) error
	// Listens to the messages sent by the client.
	Listener(Chatter_ListenerServer) error
	// Summarizes the chat messages sent by the client.
	Summary(Chatter_SummaryServer) error
	// Subscribe to events sent when new chat messages are added.
	Subscribe(*SubscribeRequest, Chatter_SubscribeServer) error
	// Returns the chat messages sent to the server.
	History(*HistoryRequest, Chatter_HistoryServer) error
	mustEmbedUnimplementedChatterServer()
}

// UnimplementedChatterServer must be embedded to have forward compatible implementations.
type UnimplementedChatterServer struct {
}

func (UnimplementedChatterServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedChatterServer) Echoer(Chatter_EchoerServer) error {
	return status.Errorf(codes.Unimplemented, "method Echoer not implemented")
}
func (UnimplementedChatterServer) Listener(Chatter_ListenerServer) error {
	return status.Errorf(codes.Unimplemented, "method Listener not implemented")
}
func (UnimplementedChatterServer) Summary(Chatter_SummaryServer) error {
	return status.Errorf(codes.Unimplemented, "method Summary not implemented")
}
func (UnimplementedChatterServer) Subscribe(*SubscribeRequest, Chatter_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedChatterServer) History(*HistoryRequest, Chatter_HistoryServer) error {
	return status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedChatterServer) mustEmbedUnimplementedChatterServer() {}

// UnsafeChatterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatterServer will
// result in compilation errors.
type UnsafeChatterServer interface {
	mustEmbedUnimplementedChatterServer()
}

func RegisterChatterServer(s grpc.ServiceRegistrar, srv ChatterServer) {
	s.RegisterService(&Chatter_ServiceDesc, srv)
}

func _Chatter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chatter_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatterServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chatter_Echoer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatterServer).Echoer(&chatterEchoerServer{stream})
}

type Chatter_EchoerServer interface {
	Send(*EchoerResponse) error
	Recv() (*EchoerStreamingRequest, error)
	grpc.ServerStream
}

type chatterEchoerServer struct {
	grpc.ServerStream
}

func (x *chatterEchoerServer) Send(m *EchoerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatterEchoerServer) Recv() (*EchoerStreamingRequest, error) {
	m := new(EchoerStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chatter_Listener_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatterServer).Listener(&chatterListenerServer{stream})
}

type Chatter_ListenerServer interface {
	SendAndClose(*ListenerResponse) error
	Recv() (*ListenerStreamingRequest, error)
	grpc.ServerStream
}

type chatterListenerServer struct {
	grpc.ServerStream
}

func (x *chatterListenerServer) SendAndClose(m *ListenerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatterListenerServer) Recv() (*ListenerStreamingRequest, error) {
	m := new(ListenerStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chatter_Summary_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatterServer).Summary(&chatterSummaryServer{stream})
}

type Chatter_SummaryServer interface {
	SendAndClose(*ChatSummaryCollection) error
	Recv() (*SummaryStreamingRequest, error)
	grpc.ServerStream
}

type chatterSummaryServer struct {
	grpc.ServerStream
}

func (x *chatterSummaryServer) SendAndClose(m *ChatSummaryCollection) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatterSummaryServer) Recv() (*SummaryStreamingRequest, error) {
	m := new(SummaryStreamingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chatter_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatterServer).Subscribe(m, &chatterSubscribeServer{stream})
}

type Chatter_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type chatterSubscribeServer struct {
	grpc.ServerStream
}

func (x *chatterSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Chatter_History_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatterServer).History(m, &chatterHistoryServer{stream})
}

type Chatter_HistoryServer interface {
	Send(*HistoryResponse) error
	grpc.ServerStream
}

type chatterHistoryServer struct {
	grpc.ServerStream
}

func (x *chatterHistoryServer) Send(m *HistoryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Chatter_ServiceDesc is the grpc.ServiceDesc for Chatter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chatter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatter.Chatter",
	HandlerType: (*ChatterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Chatter_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echoer",
			Handler:       _Chatter_Echoer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Listener",
			Handler:       _Chatter_Listener_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Summary",
			Handler:       _Chatter_Summary_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _Chatter_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "History",
			Handler:       _Chatter_History_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "goadesign_goagen_chatter.proto",
}
