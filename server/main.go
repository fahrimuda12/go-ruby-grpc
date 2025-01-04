package main

import (
	"context"
	userpb "go-grpc-simple/gen/go/user/v1"
	"log"
	"net"

	"google.golang.org/grpc"
)

type UserService struct {}

func (u *UserService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	return &userpb.GetUserResponse{
		User: &userpb.User{
			Uuid: req.Uuid,
			Fullname: "Mario",
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9879")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &UserService{})
	grpcServer.Serve(lis)
}