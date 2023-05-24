// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: testpb.proto

package testpb

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
	TestService_TestMethod_FullMethodName  = "/testpb.TestService/TestMethod"
	TestService_TestMethod2_FullMethodName = "/testpb.TestService/TestMethod2"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	TestMethod(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	TestMethod2(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestMethod(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, TestService_TestMethod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestMethod2(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, TestService_TestMethod2_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	TestMethod(context.Context, *TestRequest) (*TestResponse, error)
	TestMethod2(context.Context, *TestRequest) (*TestResponse, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) TestMethod(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod not implemented")
}
func (UnimplementedTestServiceServer) TestMethod2(context.Context, *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestMethod2 not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_TestMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_TestMethod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestMethod(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestMethod2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestMethod2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_TestMethod2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestMethod2(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testpb.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestMethod",
			Handler:    _TestService_TestMethod_Handler,
		},
		{
			MethodName: "TestMethod2",
			Handler:    _TestService_TestMethod2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testpb.proto",
}
