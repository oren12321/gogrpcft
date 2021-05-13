// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// TransferClient is the client API for Transfer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferClient interface {
	Receive(ctx context.Context, in *Info, opts ...grpc.CallOption) (Transfer_ReceiveClient, error)
	Send(ctx context.Context, opts ...grpc.CallOption) (Transfer_SendClient, error)
}

type transferClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferClient(cc grpc.ClientConnInterface) TransferClient {
	return &transferClient{cc}
}

func (c *transferClient) Receive(ctx context.Context, in *Info, opts ...grpc.CallOption) (Transfer_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transfer_ServiceDesc.Streams[0], "/Io.Transfer/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Transfer_ReceiveClient interface {
	Recv() (*Packet, error)
	grpc.ClientStream
}

type transferReceiveClient struct {
	grpc.ClientStream
}

func (x *transferReceiveClient) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *transferClient) Send(ctx context.Context, opts ...grpc.CallOption) (Transfer_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transfer_ServiceDesc.Streams[1], "/Io.Transfer/Send", opts...)
	if err != nil {
		return nil, err
	}
	x := &transferSendClient{stream}
	return x, nil
}

type Transfer_SendClient interface {
	Send(*Packet) error
	CloseAndRecv() (*Status, error)
	grpc.ClientStream
}

type transferSendClient struct {
	grpc.ClientStream
}

func (x *transferSendClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transferSendClient) CloseAndRecv() (*Status, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Status)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransferServer is the server API for Transfer service.
// All implementations must embed UnimplementedTransferServer
// for forward compatibility
type TransferServer interface {
	Receive(*Info, Transfer_ReceiveServer) error
	Send(Transfer_SendServer) error
	mustEmbedUnimplementedTransferServer()
}

// UnimplementedTransferServer must be embedded to have forward compatible implementations.
type UnimplementedTransferServer struct {
}

func (UnimplementedTransferServer) Receive(*Info, Transfer_ReceiveServer) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedTransferServer) Send(Transfer_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedTransferServer) mustEmbedUnimplementedTransferServer() {}

// UnsafeTransferServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServer will
// result in compilation errors.
type UnsafeTransferServer interface {
	mustEmbedUnimplementedTransferServer()
}

func RegisterTransferServer(s grpc.ServiceRegistrar, srv TransferServer) {
	s.RegisterService(&Transfer_ServiceDesc, srv)
}

func _Transfer_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Info)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TransferServer).Receive(m, &transferReceiveServer{stream})
}

type Transfer_ReceiveServer interface {
	Send(*Packet) error
	grpc.ServerStream
}

type transferReceiveServer struct {
	grpc.ServerStream
}

func (x *transferReceiveServer) Send(m *Packet) error {
	return x.ServerStream.SendMsg(m)
}

func _Transfer_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransferServer).Send(&transferSendServer{stream})
}

type Transfer_SendServer interface {
	SendAndClose(*Status) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type transferSendServer struct {
	grpc.ServerStream
}

func (x *transferSendServer) SendAndClose(m *Status) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transferSendServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Transfer_ServiceDesc is the grpc.ServiceDesc for Transfer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transfer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Io.Transfer",
	HandlerType: (*TransferServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receive",
			Handler:       _Transfer_Receive_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Send",
			Handler:       _Transfer_Send_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "io.proto",
}
