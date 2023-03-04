package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/sum/proto"
)

func doHelloWorld(c pb.HelloWorldServiceClient) {
	log.Println("doHelloWorld was invoked")
	res, err := c.HelloWorld(context.Background(), &pb.HelloWorldRequest{
		Name: "Alexandre",
	})
	if err != nil {
		log.Fatalf("Could no greet: %v\n", err)
	}

	log.Printf("Hello World: %s\n", res.Result)
}
