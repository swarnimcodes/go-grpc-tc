package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	pb "github.com/swarnimcodes/go-grpc-tc/github.com/swarnimcodes/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	c      pb.UserServiceClient
	ctx    context.Context
	conn   *grpc.ClientConn
	cancel context.CancelFunc
)

// abstract out the initialising
// this prevents re-initialisation for each test
func init() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	var err error
	conn, err = grpc.NewClient("dns:///localhost:50051", opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()

	c = pb.NewUserServiceClient(conn)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
}

func TestGetUserById_ValidID(t *testing.T) {
	req := &pb.UserIdRequest{Id: 1}
	resp, err := c.GetUserById(ctx, req)
	if err != nil {
		t.Fatalf("GetUserById failed: %v", err)
	}
	assert.Equal(t, int32(1), resp.User.Id)
}

func TestUsersByIds_ValidIds(t *testing.T) {
	req := &pb.UserIdsRequest{Ids: []int32{1, 2}}
	resp, err := c.GetUsersByIds(ctx, req)
	if err != nil {
		t.Fatalf("GetUsersByIds failed: %v", err)
	}
	assert.Equal(t, 2, len(resp.Users))
	assert.Equal(t, int32(1), resp.Users[0].Id)
	assert.Equal(t, int32(2), resp.Users[1].Id)
}

func TestGetUsersByIds_PartialValidIDs(t *testing.T) {
	req := &pb.UserIdsRequest{Ids: []int32{1, 10}}
	resp, err := c.GetUsersByIds(ctx, req)
	if err != nil {
		t.Fatalf("GetUsersByIds failed: %v", err)
	}
	assert.Equal(t, 1, len(resp.Users))
	assert.Equal(t, int32(1), resp.Users[0].Id)
}

func TestGetUsersByIds_InvalidIDs(t *testing.T) {
	req := &pb.UserIdsRequest{Ids: []int32{999, 1000}}
	_, err := c.GetUsersByIds(ctx, req)
	assert.Error(t, err)
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("Expected grpc error status, got %v", err)
	}
	assert.Equal(t, codes.NotFound, st.Code())
}

func TestGetUsersByIds_SingleValidID(t *testing.T) {
	req := &pb.UserIdsRequest{Ids: []int32{2}}
	resp, err := c.GetUsersByIds(ctx, req)
	if err != nil {
		t.Fatalf("GetUsersByIds failed: %v", err)
	}
	assert.Equal(t, 1, len(resp.Users))
	assert.Equal(t, int32(2), resp.Users[0].Id)
}

func TestSearchUsers_ByFirstName(t *testing.T) {
	req := &pb.SearchRequest{Fname: "John", SearchByPhoneNumber: true, SearchByMarriageStatus: true, Phone: 8237828372, Married: true}
	resp, err := c.SearchUsers(ctx, req)
	if err != nil {
		t.Fatalf("SearchUsers failed: %v", err)
	}
	assert.Equal(t, 1, len(resp.Users))
}

func TestSearchUsers_ByCity(t *testing.T) {
	req := &pb.SearchRequest{City: "CA", SearchByMarriageStatus: false, SearchByPhoneNumber: false}
	resp, err := c.SearchUsers(ctx, req)
	if err != nil {
		t.Fatalf("SearchUsers failed: %v", err)
	}
	for _, user := range resp.Users {
		assert.Equal(t, "CA", user.City)
	}
}

func TestSearchUsers_ByMarriageStatus(t *testing.T) {
	req := &pb.SearchRequest{Married: true, SearchByMarriageStatus: true, SearchByPhoneNumber: false}
	resp, err := c.SearchUsers(ctx, req)
	if err != nil {
		t.Fatalf("SearchUsers failed: %v", err)
	}
	for _, user := range resp.Users {
		assert.Equal(t, true, user.Married)
	}
}

func TestSearchUsers_NoMatch(t *testing.T) {
	req := &pb.SearchRequest{Fname: "Nonexistent", SearchByPhoneNumber: false, SearchByMarriageStatus: false}
	_, err := c.SearchUsers(ctx, req)
	assert.Error(t, err)
	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("Expected grpc error status, got %v", err)
	}
	assert.Equal(t, codes.NotFound, st.Code())
}

func TestMain(m *testing.M) {
	code := m.Run()
	cancel()
	conn.Close()
	os.Exit(code)
}
