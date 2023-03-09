package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doHelloWorldWithDeadline(c pb.HelloWorldServiceClient, timeout time.Duration) {
	log.Println("doHelloWorldWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.HelloWorldRequest{
		Name: "Alexandre",
	}

	res, err := c.HelloWorldWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error %v\n", err)
		}
	}

	log.Printf("Hello World With Deadline: %s\n", res.Result)
}
