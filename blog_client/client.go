package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AdityaKshettri/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client!")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}
	defer conn.Close()

	c := blogpb.NewBlogServiceClient(conn)

	createBlog(c)
}

func createBlog(c blogpb.BlogServiceClient) {
	fmt.Println("Creating the Blog!")
	blog := &blogpb.Blog{
		AuthorId: "Aditya",
		Title:    "My First Blog",
		Content:  "Content for my first blog",
	}
	req := &blogpb.CreateBlogRequest{
		Blog: blog,
	}
	res, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	}
	fmt.Println("Blog has been created!", res)
}
