package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	blog := &pb.Blog{
		AuthorId: "Alexandre",
		Title:    "Tamo aí",
		Content:  "Conteúdo duvidoso para um blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Printf("Blog have been created: %s\n", res.Id)
	return res.Id
}
