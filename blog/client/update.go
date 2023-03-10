package main

import (
	"context"
	"log"

	pb "github.com/alexandreps1123/grpc-go/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Certamente não é Alexandre",
		Title:    "Não Tamo aí",
		Content:  "Não é um conteúdo duvidoso para um blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println("Blog was updated!")
}
