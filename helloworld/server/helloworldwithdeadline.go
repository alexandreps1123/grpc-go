package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) HelloWorldWithDeadline(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	log.Printf("HelloWorldWithDeadline was invoked!")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.HelloWorldResponse{
		Result: "Hello " + in.GetName(),
	}, nil
}
