package main

import (
	"context"
	"grpc_test/pb/commpb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/mwitkow/grpc-proxy/proxy"
)

func main() {
	cc, err := grpc.DialContext(context.Background(), "localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		outCtx := metadata.NewOutgoingContext(ctx, md.Copy())

		return outCtx, cc, nil
	}

	s := grpc.NewServer(
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)

	// testpb.RegisterTestServiceServer(s, &Server{})
	commpb.RegisterCommServiceServer(s, &Server{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s.Serve(lis)
}

type Server struct {
	commpb.UnimplementedCommServiceServer
}
