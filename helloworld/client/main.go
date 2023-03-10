package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/alexandreps1123/grpc-go/helloworld/proto"
)

var addr string = "localhost:50051"

func main() {
	/*
	* Failed to connect:
	*	grpc: no transport security set (use grpc.WithTransportCredentials(insecure.NewCredentials())
	* 	explicitly or set credentials
	 */
	// ssl
	opts := []grpc.DialOption{}
	tls := true // change that to false if needed

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust cretificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewHelloWorldServiceClient(conn)

	// doHelloWorld(c)
	// doHelloWorldStream(c)
	// doStreamHelloWorld(c)
	// doStreamHelloWorldStream(c)
	doHelloWorldWithDeadline(c, 5*time.Second)
}
