package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/swarnimcodes/go-grpc-tc/github.com/swarnimcodes/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient("dns:///localhost:50051", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// get user by their id
	user, err := c.GetUserById(ctx, &pb.UserIdRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v\n", err)
	}
	fmt.Printf("User: %v\n", user)

	// get user list by ids
	users, err := c.GetUsersByIds(ctx, &pb.UserIdsRequest{Ids: []int32{1, 2}})
	if err != nil {
		log.Fatalf("could not get user list: %v", err)
	}
	fmt.Printf("Users: %v\n", users)

	// search users based on multiple criterias
	// adding the fields SearchByPhoneNumber and SearchByMarriageStatus is mandatory
	// if either field is true, the corresponding field must be provided for search

	searchRes, err := c.SearchUsers(ctx, &pb.SearchRequest{Fname: "John", SearchByPhoneNumber: true, Phone: 8237828372, SearchByMarriageStatus: false})
	if err != nil {
		log.Fatalf("could not fetch users: %v", err)
	}
	fmt.Printf("Search Result: %v\n", searchRes)

}
