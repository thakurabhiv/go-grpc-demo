package main

import (
	"log"
	"net"

	pb "github.com/thakurabhiv/go-grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Unable to get port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("Starting grpc server at: %v", listener.Addr())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Unable to serve grpc: %v", err)
	}
}
