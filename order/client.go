package order

import (
	"context"
	"log"
	"time"

	pb "github.com/geekAshish/go-grpc-graphql-micro/order/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.OrderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := pb.NewOrderServiceClient(conn)

	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostOrder(ctx context.Context, accountId string, products []OrderedProduct) (*Order, error) {
	protoProduct := []*pb.PostOrderRequest_OrderProduct{}

	for _, p := range products {
		protoProduct = append(protoProduct, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  p.Quantity,
		})
	}

	r, err := c.service.PostOrder(
		ctx,
		&pb.PostOrderRequest{
			AccountId: accountId,
			Products:  protoProduct,
		},
	)

	if err != nil {
		return nil, err
	}

	newOrder := r.Order
	newOrderCreatedAt := time.Time{}

	newOrderCreatedAt.UnmarshalBinary(newOrder.CreatedAt)

	return &Order{
		ID:         newOrder.Id,
		CreatedAt:  newOrderCreatedAt,
		TotalPrice: newOrder.TotalPrice,
		AccountID:  newOrder.AccountId,
		Products:   products,
	}, nil
}

func (c *Client) GetOrdersForAccount(ctx context.Context, accountId string) ([]Order, error) {
	r, err := c.service.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{
		AccountId: accountId,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	orders := []Order{}
	for _, orderProto := range r.Orders {
		newOrder := Order{
			ID:         orderProto.Id,
			TotalPrice: orderProto.TotalPrice,
			AccountID:  orderProto.AccountId,
		}
		newOrder.CreatedAt = time.Time{}
		newOrder.CreatedAt.UnmarshalBinary(orderProto.CreatedAt)
		products := []OrderedProduct{}

		for _, p := range orderProto.Products {
			products = append(products, OrderedProduct{
				ID:          p.Id,
				Quantity:    p.Quantity,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			})
		}
		newOrder.Products = products
		orders = append(orders, newOrder)
	}
	return orders, nil
}
