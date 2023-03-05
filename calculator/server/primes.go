package main

import (
	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func (s *Server) PrimeStream(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	var number = in.GetNumber()
	var divisor uint64 = 2

	for {
		if number == 1 {
			break
		}

		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
