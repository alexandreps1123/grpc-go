package main

import (
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

/*
 *	Request stream
 *	Response stream
 */
func (s *Server) StreamHelloWorldStream(stream pb.HelloWorldService_StreamHelloWorldStreamServer) error {
	log.Println("StreamHelloWorldStream was invoked")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		res := "Hello" + req.GetName() + "!"
		err = stream.Send(&pb.HelloWorldResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
	}
}
