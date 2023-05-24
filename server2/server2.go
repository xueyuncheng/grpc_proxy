package main

import (
	"context"
	"grpc_test/pb/testpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()

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

func (s *Server) TestMethod(ctx context.Context, req *testpb.TestRequest) (*testpb.TestResponse, error) {
	return &testpb.TestResponse{Message: req.Name}, nil
}
