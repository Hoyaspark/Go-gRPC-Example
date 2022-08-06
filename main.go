package main

import (
	"context"
	"github.com/hoyaspark/go-grpc-example/data"
	postpb "github.com/hoyaspark/go-grpc-example/proto/post"
	userpb "github.com/hoyaspark/go-grpc-example/proto/user"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const portNumber = 9001

type postServer struct {
	postpb.PostServer

	userCli userpb.UserClient
}

func (s *postServer) ListPostsByUserId(ctx context.Context, req *postpb.ListPostsByUserIdRequest) (*postpb.ListPostsByUserIdResponse, error) {
	userId := req.GetUserId()

	res, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userId}, grpc.EmptyCallOption{})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var postMessages []*postpb.PostMessage

	for _, v := range data.UserPosts {
		if v.UserId != userId {
			continue
		}

		for _, a := range v.Posts {
			a.Author = res.GetUserMessage().GetName()
		}

		postMessages = v.Posts
		break
	}

	return &postpb.ListPostsByUserIdResponse{PostMessages: postMessages}, nil

}

func (s *postServer) ListPosts(ctx context.Context, req *postpb.ListPostsRequest) (*postpb.ListPostsResponse, error) {

	var postMessages []*postpb.PostMessage

	for _, up := range data.UserPosts {
		userId := up.UserId

		res, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: userId}, grpc.EmptyCallOption{})

		if err != nil {
			log.Println(err)
			return nil, err
		}

		for _, u := range up.Posts {
			u.Author = res.UserMessage.Name
		}

		postMessages = append(postMessages, up.Posts...)
	}

	return &postpb.ListPostsResponse{PostMessages: postMessages}, nil
}

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
