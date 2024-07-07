package main

import (
	"context"
	"log"
	"time"

	"github.com/swarnimcodes/go-grpc-tc/github.com/swarnimcodes/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up the client options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// Create a new client connection
	conn, err := grpc.NewClient("dns:///localhost:50051", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := userpb.NewUserServiceClient(conn)

	userCtx, userCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer userCancel()

	user := &userpb.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	// Create User
	createRes, err := c.CreateUser(userCtx, &userpb.CreateUserRequest{User: user})
	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}
	log.Printf("User created: %v", createRes.User)

	// Get User
	getRes, err := c.GetUser(userCtx, &userpb.GetUserRequest{Id: createRes.User.Id})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User retrieved: %v", getRes.User)

	// Update User
	user.Id = createRes.User.Id
	user.Name = "Jane Doe"
	updateRes, err := c.UpdateUser(userCtx, &userpb.UpdateUserRequest{User: user})
	if err != nil {
		log.Fatalf("could not update user: %v", err)
	}
	log.Printf("User updated: %v", updateRes.User)

	// List Users
	listRes, err := c.ListUsers(userCtx, &userpb.ListUsersRequest{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	log.Printf("Users: %v", listRes.Users)

	// Delete User
	deleteRes, err := c.DeleteUser(userCtx, &userpb.DeleteUserRequest{Id: user.Id})
	if err != nil {
		log.Fatalf("could not delete user: %v", err)
	}
	log.Printf("User deleted: %v", deleteRes.Result)
}
