package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

func (s *Server) HelloWorld(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	log.Printf("HelloWorld function was invoked with %v\n", in)

	return &pb.HelloWorldResponse{Result: "Hello " + in.Name}, nil
}
