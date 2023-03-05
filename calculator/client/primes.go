package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doFactorStream was invoked")
	req := &pb.PrimeRequest{
		Number: 500000,
	}

	stream, err := c.Primes(context.Background(), req)
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

		log.Printf("Factor Stream: %v\n", msg.Result)
	}
}
