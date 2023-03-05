package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

/*
 *	Request stream
 *	Response unary
 */
func doStreamHelloWorld(c pb.HelloWorldServiceClient) {
	log.Println("doStreamHelloWorld was invoked")

	reqs := []*pb.HelloWorldRequest{
		{Name: "Alexandre"},
		{Name: "Ranna"},
		{Name: "Raul"},
	}

	stream, err := c.StreamHelloWorld(context.Background())

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Printf("Stream Hello World: %s\n", res.GetResult())
}
