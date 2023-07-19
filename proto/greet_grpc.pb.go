// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: proto/greet.proto

package proto

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error)
	SayHelloServerStream(ctx context.Context, in *NameList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamClient, error)
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamClient, error)
	SayHelloBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidirectionalStreamClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) SayHello(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/greet_service.GreetService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) SayHelloServerStream(ctx context.Context, in *NameList, opts ...grpc.CallOption) (GreetService_SayHelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet_service.GreetService/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_SayHelloServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], "/greet_service.GreetService/SayHelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloClientStreamClient{stream}
	return x, nil
}

type GreetService_SayHelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*MessageList, error)
	grpc.ClientStream
}

type greetServiceSayHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloClientStreamClient) CloseAndRecv() (*MessageList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) SayHelloBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (GreetService_SayHelloBidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], "/greet_service.GreetService/SayHelloBidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceSayHelloBidirectionalStreamClient{stream}
	return x, nil
}

type GreetService_SayHelloBidirectionalStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type greetServiceSayHelloBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *greetServiceSayHelloBidirectionalStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceSayHelloBidirectionalStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	SayHello(context.Context, *NoParam) (*HelloResponse, error)
	SayHelloServerStream(*NameList, GreetService_SayHelloServerStreamServer) error
	SayHelloClientStream(GreetService_SayHelloClientStreamServer) error
	SayHelloBidirectionalStream(GreetService_SayHelloBidirectionalStreamServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) SayHello(context.Context, *NoParam) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloServerStream(*NameList, GreetService_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloClientStream(GreetService_SayHelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (UnimplementedGreetServiceServer) SayHelloBidirectionalStream(GreetService_SayHelloBidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectionalStream not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet_service.GreetService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).SayHello(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NameList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).SayHelloServerStream(m, &greetServiceSayHelloServerStreamServer{stream})
}

type GreetService_SayHelloServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type greetServiceSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloClientStream(&greetServiceSayHelloClientStreamServer{stream})
}

type GreetService_SayHelloClientStreamServer interface {
	SendAndClose(*MessageList) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloClientStreamServer) SendAndClose(m *MessageList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_SayHelloBidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).SayHelloBidirectionalStream(&greetServiceSayHelloBidirectionalStreamServer{stream})
}

type GreetService_SayHelloBidirectionalStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greetServiceSayHelloBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *greetServiceSayHelloBidirectionalStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceSayHelloBidirectionalStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GreetService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _GreetService_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _GreetService_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectionalStream",
			Handler:       _GreetService_SayHelloBidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}
