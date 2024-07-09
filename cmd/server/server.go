package main

import (
	"fmt"
	"log"
	"net"

	"context"
	"sync"

	pb "github.com/swarnimcodes/go-grpc-tc/github.com/swarnimcodes/userpb"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ID      int32
	Fname   string
	City    string
	Phone   int64
	Height  float32
	Married bool
}

var users = []User{
	{ID: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: false},
	{ID: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
	{ID: 3, Fname: "John", City: "CA", Phone: 8237828372, Height: 5.9, Married: true},
	{ID: 4, Fname: "Emily", City: "CA", Phone: 8981273892, Height: 5.1, Married: true},
}

type server struct {
	pb.UnimplementedUserServiceServer
	users map[int32]*pb.User
	mu    sync.Mutex
}

func (s *server) GetUserById(ctx context.Context, req *pb.UserIdRequest) (*pb.UserResponse, error) {
	log.Printf("received `GetUserById` request for ID `%d`\n", req.Id)
	s.mu.Lock()
	defer s.mu.Unlock()
	var foundUser *pb.User

	for _, user := range s.users {
		if req.Id == user.Id {
			foundUser = user
			log.Printf("found user with the requested ID `%d`", req.Id)
			break
		}
	}

	if foundUser == nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	} else {
		return &pb.UserResponse{User: foundUser}, nil
	}
}

func (s *server) GetUsersByIds(ctx context.Context, req *pb.UserIdsRequest) (*pb.UsersResponse, error) {
	log.Printf("received `GetUsersByIds` request for IDs `%v`\n", req.GetIds())
	s.mu.Lock()
	defer s.mu.Unlock()
	var result []*pb.User
	var numMatches int32
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				log.Printf("matched user with ID `%d`", id)
				numMatches++
				result = append(result, &pb.User{
					Id:      user.Id,
					Fname:   user.Fname,
					City:    user.City,
					Phone:   user.Phone,
					Height:  user.Height,
					Married: user.Married,
				})
			}
		}
	}
	if numMatches == 0 {
		return nil, status.Errorf(codes.NotFound, "no users found for requested IDs")
	} else {
		return &pb.UsersResponse{Users: result}, nil
	}
}

// the filtering fields `SearchByPhoneNumber` and `SearchByMarriageStatus` are mandatory
// this allows us to know if filtering should be done on their basis or not
func (s *server) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	log.Println("received `SearchUsers` request")
	var result []*pb.User
	var numMatches int32 = 0
	for _, user := range s.users {
		if (req.Fname == "" || req.Fname == user.Fname) &&
			(req.City == "" || req.City == user.City) &&
			((req.SearchByMarriageStatus && req.Married == user.Married) || (!req.SearchByMarriageStatus)) &&
			((req.SearchByPhoneNumber && req.Phone == user.Phone) || (!req.SearchByPhoneNumber)) {
			numMatches++
			log.Printf("matched user with ID `%d`", user.Id)
			result = append(result, &pb.User{
				Id:      user.Id,
				Fname:   user.Fname,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			})
		}
	}
	if numMatches == 0 {
		return nil, status.Error(codes.NotFound, "no users matched the search criteria")
	} else {
		return &pb.UsersResponse{Users: result}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	srv := &server{
		users: make(map[int32]*pb.User),
	}
	// initialize the server with some users
	for _, user := range users {
		srv.users[user.ID] = &pb.User{
			Id:      user.ID,
			Fname:   user.Fname,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		}
	}

	pb.RegisterUserServiceServer(s, srv)

	fmt.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
