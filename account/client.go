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

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	r, err := c.service.GetAccount(
		ctx,
		&pb.GetAccountRequest{Id: id},
	)

	if err != nil {
		return nil, err
	}

	return &Account{
		ID: r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (c *Client) GetAccounts(ctx context.Context, take uint64, skip uint64) ([]Account, error) {
	res, err := c.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		},
	)
	
	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	for _, p := range res.Accounts {
		accounts = append(
			accounts,
			Account{
				ID:   p.Id,
				Name: p.Name,
			},
		)
	}

	return accounts, nil
}
