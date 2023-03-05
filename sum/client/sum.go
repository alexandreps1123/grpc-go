package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/sum/proto"
)

func doSum(c pb.SumServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 10,
		B: 7,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result)
}

func doFactorStream(c pb.SumServiceClient) {
	log.Println("doFactorStream was invoked")
	req := &pb.FactorRequest{
		Number: 500000,
	}

	stream, err := c.FactorStream(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		log.Printf("doFactorStream: %v\n", msg.Result)
	}
}
