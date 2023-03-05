package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

/*
 *	Request stream
 *	Response unary
 */
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

		log.Printf("Processing: %v\n", req.GetName())

		res += fmt.Sprintf("Hello %s!\n", req.GetName())
	}
}
