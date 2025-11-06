package service

import (
	"context"
	"product-service/internal/model"
	"product-service/internal/repository"
)

type ProductService interface {
	CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error)
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	UpdateProduct(ctx context.Context, id string, product *model.Product) (*model.Product, error)
	DeleteProduct(ctx context.Context, id string) error
	GetProductsByCategoryId(ctx context.Context, categoryId string) ([]model.Product, error)
}

type productService struct {
	productRepo     repository.ProductRepository
	categoryService CategoryService
}

func NewProductService(productRepo repository.ProductRepository, categoryService CategoryService) ProductService {
	return &productService{productRepo, categoryService}
}

func (s *productService) CreateProduct(ctx context.Context, product *model.Product) (*model.Product, error) {
	err := s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	product, err := s.productRepo.GetProductById(ctx, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	return s.productRepo.GetAllProducts(ctx)
}

func (s *productService) UpdateProduct(ctx context.Context, id string, product *model.Product) (*model.Product, error) {
	err := s.productRepo.UpdateProduct(ctx, id, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepo.DeleteProduct(ctx, id)
}

func (s *productService) GetProductsByCategoryId(ctx context.Context, categoryId string) ([]model.Product, error) {
	return s.productRepo.GetProductsByCategoryId(ctx, categoryId)
}
