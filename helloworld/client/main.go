package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

var addr string = "localhost:50051"

func main() {
	/*
	* Failed to connect:
	*	grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials())
	* 	explicitly or set credentials
	 */
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewHelloWorldServiceClient(conn)

	doHelloWorld(c)
}
