// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package liiklus

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// LiiklusServiceClient is the client API for LiiklusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiiklusServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LiiklusService_SubscribeClient, error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (LiiklusService_ReceiveClient, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error)
	GetEndOffsets(ctx context.Context, in *GetEndOffsetsRequest, opts ...grpc.CallOption) (*GetEndOffsetsReply, error)
}

type liiklusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiiklusServiceClient(cc grpc.ClientConnInterface) LiiklusServiceClient {
	return &liiklusServiceClient{cc}
}

var liiklusServicePublishStreamDesc = &grpc.StreamDesc{
	StreamName: "Publish",
}

func (c *liiklusServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error) {
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var liiklusServiceSubscribeStreamDesc = &grpc.StreamDesc{
	StreamName:    "Subscribe",
	ServerStreams: true,
}

func (c *liiklusServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (LiiklusService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, liiklusServiceSubscribeStreamDesc, "/com.github.bsideup.liiklus.LiiklusService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &liiklusServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LiiklusService_SubscribeClient interface {
	Recv() (*SubscribeReply, error)
	grpc.ClientStream
}

type liiklusServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *liiklusServiceSubscribeClient) Recv() (*SubscribeReply, error) {
	m := new(SubscribeReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var liiklusServiceReceiveStreamDesc = &grpc.StreamDesc{
	StreamName:    "Receive",
	ServerStreams: true,
}

func (c *liiklusServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (LiiklusService_ReceiveClient, error) {
	stream, err := c.cc.NewStream(ctx, liiklusServiceReceiveStreamDesc, "/com.github.bsideup.liiklus.LiiklusService/Receive", opts...)
	if err != nil {
		return nil, err
	}
	x := &liiklusServiceReceiveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LiiklusService_ReceiveClient interface {
	Recv() (*ReceiveReply, error)
	grpc.ClientStream
}

type liiklusServiceReceiveClient struct {
	grpc.ClientStream
}

func (x *liiklusServiceReceiveClient) Recv() (*ReceiveReply, error) {
	m := new(ReceiveReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var liiklusServiceAckStreamDesc = &grpc.StreamDesc{
	StreamName: "Ack",
}

func (c *liiklusServiceClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/Ack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var liiklusServiceGetOffsetsStreamDesc = &grpc.StreamDesc{
	StreamName: "GetOffsets",
}

func (c *liiklusServiceClient) GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error) {
	out := new(GetOffsetsReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/GetOffsets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var liiklusServiceGetEndOffsetsStreamDesc = &grpc.StreamDesc{
	StreamName: "GetEndOffsets",
}

func (c *liiklusServiceClient) GetEndOffsets(ctx context.Context, in *GetEndOffsetsRequest, opts ...grpc.CallOption) (*GetEndOffsetsReply, error) {
	out := new(GetEndOffsetsReply)
	err := c.cc.Invoke(ctx, "/com.github.bsideup.liiklus.LiiklusService/GetEndOffsets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiiklusServiceService is the service API for LiiklusService service.
// Fields should be assigned to their respective handler implementations only before
// RegisterLiiklusServiceService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type LiiklusServiceService struct {
	Publish       func(context.Context, *PublishRequest) (*PublishReply, error)
	Subscribe     func(*SubscribeRequest, LiiklusService_SubscribeServer) error
	Receive       func(*ReceiveRequest, LiiklusService_ReceiveServer) error
	Ack           func(context.Context, *AckRequest) (*emptypb.Empty, error)
	GetOffsets    func(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error)
	GetEndOffsets func(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error)
}

func (s *LiiklusServiceService) publish(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *LiiklusServiceService) subscribe(_ interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.Subscribe(m, &liiklusServiceSubscribeServer{stream})
}
func (s *LiiklusServiceService) receive(_ interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return s.Receive(m, &liiklusServiceReceiveServer{stream})
}
func (s *LiiklusServiceService) ack(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *LiiklusServiceService) getOffsets(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/GetOffsets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetOffsets(ctx, req.(*GetOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func (s *LiiklusServiceService) getEndOffsets(_ interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return s.GetEndOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     s,
		FullMethod: "/com.github.bsideup.liiklus.LiiklusService/GetEndOffsets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return s.GetEndOffsets(ctx, req.(*GetEndOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

type LiiklusService_SubscribeServer interface {
	Send(*SubscribeReply) error
	grpc.ServerStream
}

type liiklusServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *liiklusServiceSubscribeServer) Send(m *SubscribeReply) error {
	return x.ServerStream.SendMsg(m)
}

type LiiklusService_ReceiveServer interface {
	Send(*ReceiveReply) error
	grpc.ServerStream
}

type liiklusServiceReceiveServer struct {
	grpc.ServerStream
}

func (x *liiklusServiceReceiveServer) Send(m *ReceiveReply) error {
	return x.ServerStream.SendMsg(m)
}

// RegisterLiiklusServiceService registers a service implementation with a gRPC server.
func RegisterLiiklusServiceService(s grpc.ServiceRegistrar, srv *LiiklusServiceService) {
	srvCopy := *srv
	if srvCopy.Publish == nil {
		srvCopy.Publish = func(context.Context, *PublishRequest) (*PublishReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
		}
	}
	if srvCopy.Subscribe == nil {
		srvCopy.Subscribe = func(*SubscribeRequest, LiiklusService_SubscribeServer) error {
			return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
		}
	}
	if srvCopy.Receive == nil {
		srvCopy.Receive = func(*ReceiveRequest, LiiklusService_ReceiveServer) error {
			return status.Errorf(codes.Unimplemented, "method Receive not implemented")
		}
	}
	if srvCopy.Ack == nil {
		srvCopy.Ack = func(context.Context, *AckRequest) (*emptypb.Empty, error) {
			return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
		}
	}
	if srvCopy.GetOffsets == nil {
		srvCopy.GetOffsets = func(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetOffsets not implemented")
		}
	}
	if srvCopy.GetEndOffsets == nil {
		srvCopy.GetEndOffsets = func(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method GetEndOffsets not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "com.github.bsideup.liiklus.LiiklusService",
		Methods: []grpc.MethodDesc{
			{
				MethodName: "Publish",
				Handler:    srvCopy.publish,
			},
			{
				MethodName: "Ack",
				Handler:    srvCopy.ack,
			},
			{
				MethodName: "GetOffsets",
				Handler:    srvCopy.getOffsets,
			},
			{
				MethodName: "GetEndOffsets",
				Handler:    srvCopy.getEndOffsets,
			},
		},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Subscribe",
				Handler:       srvCopy.subscribe,
				ServerStreams: true,
			},
			{
				StreamName:    "Receive",
				Handler:       srvCopy.receive,
				ServerStreams: true,
			},
		},
		Metadata: "LiiklusService.proto",
	}

	s.RegisterService(&sd, nil)
}

// NewLiiklusServiceService creates a new LiiklusServiceService containing the
// implemented methods of the LiiklusService service in s.  Any unimplemented
// methods will result in the gRPC server returning an UNIMPLEMENTED status to the client.
// This includes situations where the method handler is misspelled or has the wrong
// signature.  For this reason, this function should be used with great care and
// is not recommended to be used by most users.
func NewLiiklusServiceService(s interface{}) *LiiklusServiceService {
	ns := &LiiklusServiceService{}
	if h, ok := s.(interface {
		Publish(context.Context, *PublishRequest) (*PublishReply, error)
	}); ok {
		ns.Publish = h.Publish
	}
	if h, ok := s.(interface {
		Subscribe(*SubscribeRequest, LiiklusService_SubscribeServer) error
	}); ok {
		ns.Subscribe = h.Subscribe
	}
	if h, ok := s.(interface {
		Receive(*ReceiveRequest, LiiklusService_ReceiveServer) error
	}); ok {
		ns.Receive = h.Receive
	}
	if h, ok := s.(interface {
		Ack(context.Context, *AckRequest) (*emptypb.Empty, error)
	}); ok {
		ns.Ack = h.Ack
	}
	if h, ok := s.(interface {
		GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error)
	}); ok {
		ns.GetOffsets = h.GetOffsets
	}
	if h, ok := s.(interface {
		GetEndOffsets(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error)
	}); ok {
		ns.GetEndOffsets = h.GetEndOffsets
	}
	return ns
}

// UnstableLiiklusServiceService is the service API for LiiklusService service.
// New methods may be added to this interface if they are added to the service
// definition, which is not a backward-compatible change.  For this reason,
// use of this type is not recommended.
type UnstableLiiklusServiceService interface {
	Publish(context.Context, *PublishRequest) (*PublishReply, error)
	Subscribe(*SubscribeRequest, LiiklusService_SubscribeServer) error
	Receive(*ReceiveRequest, LiiklusService_ReceiveServer) error
	Ack(context.Context, *AckRequest) (*emptypb.Empty, error)
	GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error)
	GetEndOffsets(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error)
}
