package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

/*
 *	Request stream
 *	Response stream
 */
func doStreamHelloWorldStream(c pb.HelloWorldServiceClient) {
	log.Println("doStreamHelloWorldStream was invoked")

	stream, err := c.StreamHelloWorldStream(context.Background())
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	reqs := []*pb.HelloWorldRequest{
		{Name: "Alexandre"},
		{Name: "Ranna"},
		{Name: "Raul"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)

			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error: %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
