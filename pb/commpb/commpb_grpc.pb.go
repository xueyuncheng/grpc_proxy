// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: commpb.proto

package commpb

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
	CommService_SendMsg_FullMethodName = "/commpb.CommService/SendMsg"
)

// CommServiceClient is the client API for CommService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommServiceClient interface {
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
}

type commServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommServiceClient(cc grpc.ClientConnInterface) CommServiceClient {
	return &commServiceClient{cc}
}

func (c *commServiceClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, CommService_SendMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommServiceServer is the server API for CommService service.
// All implementations must embed UnimplementedCommServiceServer
// for forward compatibility
type CommServiceServer interface {
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
	mustEmbedUnimplementedCommServiceServer()
}

// UnimplementedCommServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommServiceServer struct {
}

func (UnimplementedCommServiceServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedCommServiceServer) mustEmbedUnimplementedCommServiceServer() {}

// UnsafeCommServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommServiceServer will
// result in compilation errors.
type UnsafeCommServiceServer interface {
	mustEmbedUnimplementedCommServiceServer()
}

func RegisterCommServiceServer(s grpc.ServiceRegistrar, srv CommServiceServer) {
	s.RegisterService(&CommService_ServiceDesc, srv)
}

func _CommService_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommServiceServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommService_SendMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommServiceServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CommService_ServiceDesc is the grpc.ServiceDesc for CommService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "commpb.CommService",
	HandlerType: (*CommServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsg",
			Handler:    _CommService_SendMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commpb.proto",
}
