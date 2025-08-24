package graphql

import (
	"github.com/geekAshish/go-grpc-graphql-micro/account"
	"github.com/geekAshish/go-grpc-graphql-micro/catalog"
	"github.com/geekAshish/go-grpc-graphql-micro/graphql"
	"github.com/geekAshish/go-grpc-graphql-micro/order"
)

type Server struct {
	accountClient *account.Client
	catalogClient *catalog.Client
	orderClient   *order.Client
}

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		return nil, err
	}

	catalogClient, err := catalog.NewClient(catalogUrl)
	if err != nil {
		// catlog is depend on account client
		accountClient.Close()
		return nil, err
	}

	orderClient, err := order.NewClient(orderUrl)
	if err != nil {
		// order is depend on account, catalog client
		accountClient.Close()
		catalogClient.Close()
		return nil, err
	}

	return &Server{
		accountClient,
		catalogClient,
		orderClient,
	}, nil
}

func (s *Server) Mutation() MutationResolver {
	return &mutationResolver{
		server: s,
	}
}

func (s *Server) Query() QueryResolver {
	return &queryResolver{
		server: s,
	}
}

func (s *Server) Account() AccountResolver {
	return &accountResolver{
		server: s,
	}
}

func (s *Server) ExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema((Config{
		Resolvers: s,
	}),
	)
}
