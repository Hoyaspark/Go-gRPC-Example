package data

import "go-grpc-example/proto/post"

type PostData struct {
	UserId string
	Posts  []*post.PostMessage
}

var UserPosts = []*PostData{
	{
		UserId: "1",
		Posts: []*post.PostMessage{
			{
				PostId: "1",
				Author: "",
				Title:  "gRPC 구축하기 (1)",
				Body:   "gRPC를 구축하려면 이렇게 하면 된다",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
			{
				PostId: "2",
				Author: "",
				Title:  "gRPC 구축하기 (2)",
				Body:   "gRPC를 구축은 이렇다",
				Tags:   []string{"gRPC", "Golang", "server", "coding", "protobuf"},
			},
		},
	},
}
