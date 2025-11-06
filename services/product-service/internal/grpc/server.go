package grpc

import (
	"context"
	"product-service/internal/repository"
	pb "product-service/proto"
)

type ProductGRPCServer struct {
	pb.UnimplementedProductServiceServer
	repo repository.ProductRepository
}

func NewProductGRPCServer(repo repository.ProductRepository) *ProductGRPCServer {
	return &ProductGRPCServer{repo: repo}
}

func (s *ProductGRPCServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := s.repo.GetByID(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{
		Id:         product.ID.String(),
		Name:       product.Name,
		Price:      product.Price,
		Stock:      int32(product.Quantity),
		CategoryId: product.CategoryID.String(),
		CreatedBy:  product.CreatedBy.String(),
	}, nil
}
