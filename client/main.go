package main

import (
	"fmt"
	"log"

	"github.com/thakurabhiv/go-grpc-demo/proto"
	pb "github.com/thakurabhiv/go-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9000"
)

func main() {
	addr := fmt.Sprintf("localhost%v", port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while conneting: %v", err)
	}

	client := proto.NewGreetServiceClient(conn)

	names := &pb.NameList{
		Names: []string{"Abhishek", "Aditya", "Sunny", "Pankaj", "John", "Jane"},
	}

	// callSayHello(client)
	// callSayHelloClientStreaming(client, names)
	// callSayHelloServerStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)
}
