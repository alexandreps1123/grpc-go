package main

import (
	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
)

func (s *Server) PrimeStream(in *pb.PrimeRequest, stream pb.PrimeService_PrimeStreamServer) error {
	number := in.GetNumber()
	var divisor uint32 = 2
	for {
		if number == 1 {
			break
		}

		if number%divisor == 0 {
			stream.Send(&pb.calculatorResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
