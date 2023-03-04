package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/sum/proto"
)

func (s *Server) HelloWorld(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	log.Printf("HelloWorld functionwas invoked with %v\n", in)

	return &pb.HelloWorldResponse{Result: "Hello" + in.Hello}, nil
}
