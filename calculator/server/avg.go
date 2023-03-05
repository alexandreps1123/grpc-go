package main

import (
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func (s *Server) AVG(stream pb.CalculatorService_AVGServer) error {
	var avg float32 = 0
	var sum float32 = 0
	var divisor float32 = 1

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: avg,
			})
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		sum += float32(req.GetNumber())
		avg = sum / divisor
		divisor++
	}
}
