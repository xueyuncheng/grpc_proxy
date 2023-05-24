package main

import (
	"context"
	"grpc_test/pb/testpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/mwitkow/grpc-proxy/proxy"
)

func main() {
	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		outCtx := metadata.NewOutgoingContext(ctx, md.Copy())

		cc, err := grpc.DialContext(ctx, "localhost:8081", grpc.WithInsecure(), grpc.WithCodec(proxy.Codec()))
		if err != nil {
			return outCtx, nil, err
		}

		return outCtx, cc, nil
	}

	s := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)

	testpb.RegisterTestServiceServer(s, &Server{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s.Serve(lis)
}

type Server struct {
	testpb.UnimplementedTestServiceServer
}

func (s *Server) TestMethod(ctx context.Context, req *testpb.TestRequest) (*testpb.TestResponse, error) {
	return &testpb.TestResponse{Message: req.Name}, nil
}
