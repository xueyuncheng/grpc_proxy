package main

import (
	"context"
	"grpc_test/pb/testpb"
	"log"
	"net"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer(grpc.CustomCodec(proxy.Codec()))

	testpb.RegisterTestServiceServer(s, &Server{})

	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	s.Serve(lis)
}

type Server struct {
	testpb.UnimplementedTestServiceServer
}

func (s *Server) TestMethod2(ctx context.Context, req *testpb.TestRequest) (*testpb.TestResponse, error) {
	return &testpb.TestResponse{}, nil
}
