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

	// Create Blog
	fmt.Println("Creating the Blog!")
	blog := &blogpb.Blog{
		AuthorId: "Aditya",
		Title:    "My First Blog",
		Content:  "Content for my first blog",
	}
	createBlogReq := &blogpb.CreateBlogRequest{
		Blog: blog,
	}
	createBlogRes, createBlogErr := c.CreateBlog(context.Background(), createBlogReq)
	if createBlogErr != nil {
		log.Fatalf("Unexpected Error: %v\n", createBlogErr)
	}
	fmt.Println("Blog has been created!", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	// Read Blog
	fmt.Println("Reading the Blog!")
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		Id: "12345",
	})
	if readBlogErr != nil {
		fmt.Printf("Error while reading: %v\n", readBlogErr)
	}
	readBlogRes, readBlogErr = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		Id: blogID,
	})
	if readBlogErr != nil {
		fmt.Printf("Error while reading: %v\n", readBlogErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)
}
