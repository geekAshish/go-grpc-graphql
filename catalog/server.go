package catalog

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	serv := grpc.NewServer()
	reflection.Register(serv)
	pb.RegisterAccountServiceServer(serv, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           s,
	})
	return serv.Serve(lis)
}

func (s *grpcServer) PostProduct(ctx context.Context, r *pb.PostProductRequest) (*pb.PostProductResponse, error) {
	p, err := s.service.PostProduct(ctx, r.Name, r.Description, r.Price)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.PostProductResponse{
		Product: &pb.Product{
			Id: p.ID,
			Description: p.Description,
			Name: p.Name,
			Price: p.Price,
		},
	}, nil
}

func (s *grpcServer) GetProduct(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	p, err := s.service.GetProduct(ctx, r.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetProductResponse{
		Product: &pb.Product{
			Id: p.ID,
			Description: p.Description,
			Name: p.Name,
			Price: p.Price,
		},
	}, nil
}

func (s *grpcServer) GetProducts(ctx context.Context, r *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var res []Product;
	var err error

	if r.Query != "" {
		res, err = s.service.SearchProducts(ctx, r.Query, r.Skip, r.Take)
	} else if len(r.Ids) != 0 {
		res, err = s.service.GetProductsByIDs(ctx, r.Ids)
	} else {
		res, err = s.service.GetProducts(ctx, r.Skip, r.Take)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	products := []*pb.Product{}

	for _, p := range res {
		products = append(products, &pb.Product{
			ID: p.ID,
			Name: p.Name,
			Description: p.Description,
			Price: p.Price,
		})
	}

	return &pb.GetProductsResponse{
		Products: products,
	}, nil
}




