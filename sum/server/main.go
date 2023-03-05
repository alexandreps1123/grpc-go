package main

import (
	"log"
	"net"

	pb "github.com/alexandreps1123/grpc-go/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50001"

type Server struct {
	pb.SumServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listining on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
