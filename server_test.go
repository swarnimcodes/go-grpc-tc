package main

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	pb "github.com/swarnimcodes/go-grpc-tc/github.com/swarnimcodes/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetUserById_ValidId(t *testing.T) {
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

	req := &pb.UserIdRequest{Id: 1}
	resp, err := c.GetUserById(ctx, req)
	if err != nil {
		t.Fatalf("Failed to get user by ID: %v", err)
	}
	assert.Equal(t, int32(1), resp.User.Id, "Expected user ID does not match")
}
