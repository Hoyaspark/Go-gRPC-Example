package main

import (
	"context"
	"github.com/hoyaspark/go-grpc-example/data"
	userpb "github.com/hoyaspark/go-grpc-example/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const portNumber = 9000

type userServer struct {
	userpb.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.GetUserId()

	var userMessage *userpb.UserMessage

	for _, u := range data.UserData {
		if u.UserId != userId {
			continue
		}

		userMessage = u

		break
	}

	return &userpb.GetUserResponse{UserMessage: userMessage}, nil
}

func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UserData))

	for i, u := range data.UserData {
		userMessages[i] = u
	}

	return &userpb.ListUsersResponse{UserMessages: userMessages}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(portNumber))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpc := grpc.NewServer()
	userpb.RegisterUserServer(grpc, &userServer{})

	log.Printf("start gRPC server on %d port", portNumber)

	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
