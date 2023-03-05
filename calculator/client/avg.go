package main

import (
	"context"
	"log"
	"time"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func doAVG(c pb.CalculatorServiceClient) {

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.AVG(context.Background())
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println(res.GetResult())
}
