package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

/*
 *	Request unary
 *	Response stream
 */
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
