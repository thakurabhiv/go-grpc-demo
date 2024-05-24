package main

import (
	"context"

	pb "github.com/thakurabhiv/go-grpc-demo/proto"
)

type helloServer struct {
	pb.GreetServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
