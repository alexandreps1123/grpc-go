package main

import (
	"context"
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
