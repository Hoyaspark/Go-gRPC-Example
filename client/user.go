package client

import (
	userpb "github.com/hoyaspark/go-grpc-example/proto/user"
	"google.golang.org/grpc"
	"sync"
)

var (
	once sync.Once
	cli  userpb.UserClient
)

func GetUserClient(host string) userpb.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(host, grpc.WithBlock(), grpc.WithInsecure())

		cli = userpb.NewUserClient(conn)
	})

	return cli
}
