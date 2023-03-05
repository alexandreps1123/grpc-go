package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

/*
 *	Request unary
 *	Response stream
 */
func doHelloWorldStream(c pb.HelloWorldServiceClient) {
	log.Println("doHelloWorldStream was invoked")
	req := &pb.HelloWorldRequest{
		Name: "Alexandre",
	}

	stream, err := c.HelloWorldStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		log.Printf("Hello World Stream: %s\n", msg.Result)
	}

}
