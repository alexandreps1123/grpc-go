package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {

	reqs := []*pb.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	waitc := make(chan int)

	go func() {
		for _, req := range reqs {
			log.Printf("Send req: %v\n", req)
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

			log.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}
