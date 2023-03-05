package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

func (s *Server) HelloWorld(ctx context.Context, in *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	log.Printf("HelloWorld function was invoked with %v\n", in)

	return &pb.HelloWorldResponse{Result: "Hello " + in.Name}, nil
}

func (s *Server) HelloWorldStream(in *pb.HelloWorldRequest, stream pb.HelloWorldService_HelloWorldStreamServer) error {
	log.Printf("HelloWorldStream function was invoked with %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.Name, i)

		time.Sleep(2 * time.Second)
		stream.Send(&pb.HelloWorldResponse{
			Result: res,
		})
	}
	return nil
}

func (s *Server) StreamHelloWorld(stream pb.HelloWorldService_StreamHelloWorldServer) error {
	log.Println("StreamHelloWorld function was invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloWorldResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.GetName())
	}

	return nil
}
