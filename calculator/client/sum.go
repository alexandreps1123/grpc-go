package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 10,
		B: 7,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sum: %v\n", res.Result)
}
