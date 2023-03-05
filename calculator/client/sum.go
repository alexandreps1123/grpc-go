package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func docalculator(c pb.calculatorServiceClient) {
	res, err := c.calculator(context.Background(), &pb.calculatorRequest{
		A: 10,
		B: 7,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.Result)
}

func doFactorStream(c pb.calculatorServiceClient) {
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
