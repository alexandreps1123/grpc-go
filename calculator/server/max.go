package main

import (
	"io"
	"log"
	"math"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	var max uint32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		if float64(max) < math.Max(float64(max), float64(req.Number)) {
			max = req.GetNumber()

			err := stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error: %v\n", err)
			}
		}

	}
}
