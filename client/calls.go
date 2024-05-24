package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/thakurabhiv/go-grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Unable to get response: %v", err)
	}

	log.Printf("Message from server: %s", res.Message)
}

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Unable to open stream: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		log.Printf("Sending name: %s", name)
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Error sending name: %v", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	messages, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}

	for _, message := range messages.Messages {
		log.Printf("Received message: %s", message)
	}

	log.Printf("Client streaming ended")
}

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Error sending names: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}

		log.Printf("Message received: %v", res.Message)
	}
}

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error opening bidirectional stream: %v", err)
	}

	// adding waitgroup
	// to wait for both Send and Recv operation to finish
	var wg sync.WaitGroup

	// for receiving messages
	// we will open new goroutine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}

			wg.Done()
			log.Printf("Recevied message: %v", res.Message)
		}
	}()

	//send names one by one via stream
	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}

		err = stream.Send(req)
		if err != nil {
			log.Fatalf("Error sending name: %v", name)
		}

		wg.Add(1)
		time.Sleep(500 * time.Millisecond)
	}

	wg.Wait()
	log.Println("Streaming closed")
}
