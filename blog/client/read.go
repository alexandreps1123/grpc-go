package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/blog/proto"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
