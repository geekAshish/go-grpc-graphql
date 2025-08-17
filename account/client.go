package account

import (
	"context"

	"github.com/geekAshish/go-grpc-graphql-micro/account/pb/github.com/geekAshish/go-grpc-graphql-micro/account/pb"
	"google.golang.org/grpc"
)


type Client struct {
	conn *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewAccountServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close();
}

func (c *Client) PostAccount(ctx context.Context, name string) (*Account, error) {
	r, err := c.service.PostAccount(
		ctx,
		&pb.PostAccountRequest{},
	)

	if err != nil {
		return nil, err
	}

	return &Account{
		ID: r.Account.Id,
		Name: r.Account.Name,
	}, nil
}
