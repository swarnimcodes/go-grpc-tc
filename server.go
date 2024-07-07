package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/swarnimcodes/go-grpc-tc/github.com/swarnimcodes/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	userpb.UnimplementedUserServiceServer
	users map[int32]*userpb.User
	mu    sync.Mutex
	idSeq int32
}

func (s *server) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := req.GetUser()
	s.idSeq++
	user.Id = s.idSeq
	s.users[user.Id] = user

	return &userpb.CreateUserResponse{User: user}, nil
}

func (s *server) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, exists := s.users[req.GetId()]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "User not found!")
	}
	return &userpb.GetUserResponse{User: user}, nil
}

func (s *server) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user := req.GetUser()
	_, exists := s.users[user.Id]
	if !exists {
		return nil, status.Error(codes.NotFound, "User not found!")
	}
	s.users[user.Id] = user

	return &userpb.UpdateUserResponse{User: user}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.users[req.GetId()]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "User not found!")
	}
	delete(s.users, req.GetId())
	return &userpb.DeleteUserResponse{Result: "User deleted"}, nil
}

func (s *server) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var users []*userpb.User
	for _, user := range s.users {
		users = append(users, user)
	}

	return &userpb.ListUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{users: make(map[int32]*userpb.User)})

	log.Printf("server listening at: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
