package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/thakurabhiv/go-grpc-demo/proto"
)

type helloServer struct {
	pb.GreetServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Println("Called SayHello")
	return &pb.HelloResponse{
		Message: "Hello buddy",
	}, nil
}

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("Sending messages")
			return stream.SendAndClose(&pb.MessageList{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}

		log.Printf("Received name: %v", req.Name)
		messages = append(messages, fmt.Sprintf("Hello, %s", req.Name))
	}
}

func (s *helloServer) SayHelloServerStreaming(names *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	for _, name := range names.Names {
		res := &pb.HelloResponse{
			Message: fmt.Sprintf("Hello, %s", name),
		}

		err := stream.Send(res)
		log.Printf("Sending message for %v", name)
		if err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}

	return nil
}

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		res := &pb.HelloResponse{
			Message: fmt.Sprintf("Hello, %s", req.Name),
		}

		log.Printf("Sending response for name: %s", req.Name)
		if err = stream.Send(res); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
}
