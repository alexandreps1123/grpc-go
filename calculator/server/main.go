package main

import (
	"log"
	"net"

	pb "github.com/alexandreps1123/grpc-go/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50001"

type Server struct {
	pb.calculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listining on %s\n", addr)

	s := grpc.NewServer()

	pb.RegistercalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
