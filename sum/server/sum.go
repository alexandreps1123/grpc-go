package main

import (
	"context"

	pb "github.com/alexandreps1123/grpc-go/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	c := in.A + in.B
	return &pb.SumResponse{Result: c}, nil
}
