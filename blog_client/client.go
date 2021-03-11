package main

import (
	"context"
	"fmt"
	"io"
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

	// Update Blog
	fmt.Println("Updating the Blog!")
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Aditya! updated",
		Title:    "My First Blog! updated",
		Content:  "Content for my first blog! updated",
	}
	updateReq := &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), updateReq)
	if updateErr != nil {
		fmt.Printf("Error happened while updating Blog: %v\n", updateErr)
	}
	fmt.Printf("Blog updated: %v\n", updateRes)

	// Delete Blog
	deleteReq := &blogpb.DeleteBlogRequest{
		Id: blogID,
	}
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), deleteReq)
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting Blog: %v\n", deleteErr)
	}
	fmt.Printf("Blog deleted: %v\n", deleteRes)

	// List All Blogs
	listReq := &blogpb.ListBlogRequest{}
	stream, err := c.ListBlog(context.Background(), listReq)
	if err != nil {
		log.Fatalf("Error while calling ListBlog RPC : %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened : %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
