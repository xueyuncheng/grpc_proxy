package main

import (
	"context"
	"grpc_test/pb/testpb"
	"log"

	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithInsecure(), grpc.WithCodec(proxy.Codec()))
	if err != nil {
		log.Fatalln(err)
	}

	cli := testpb.NewTestServiceClient(conn)
	resp, err := cli.TestMethod(ctx, &testpb.TestRequest{Name: "test"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.Message)
}
