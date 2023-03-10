package main

import (
	"context"
	"io"
	"log"

	pb "github.com/alexandreps1123/grpc-go/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBlog(c pb.BlogServiceClient) {
	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		log.Println(res)
	}
}
