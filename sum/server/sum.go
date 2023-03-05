package main

import (
	"context"

	pb "github.com/alexandreps1123/grpc-go/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	c := in.GetA() + in.GetB()
	return &pb.SumResponse{Result: c}, nil
}

func (s *Server) FactorStream(in *pb.FactorRequest, stream pb.SumService_FactorStreamServer) error {
	number := in.GetNumber()
	var divisor uint32 = 2
	for {
		if number == 1 {
			break
		}

		if number%divisor == 0 {
			stream.Send(&pb.SumResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
