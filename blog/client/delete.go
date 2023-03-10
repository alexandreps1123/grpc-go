package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}
