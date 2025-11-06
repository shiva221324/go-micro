package repository

import (
	"context"
	"gorm.io/gorm"
	"product-service/internal/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	GetProductById(ctx context.Context, id string) (*model.Product, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	UpdateProduct(ctx context.Context, id string, product *model.Product) error
	DeleteProduct(ctx context.Context, id string) error
	GetProductsByCategoryId(ctx context.Context, categoryId string) ([]model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) CreateProduct(ctx context.Context, product *model.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepo) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).First(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, id string, product *model.Product) error {
	return r.db.WithContext(ctx).Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepo) DeleteProduct(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Product{}, "id = ?", id).Error
}

func (r *productRepo) GetProductsByCategoryId(ctx context.Context, categoryId string) ([]model.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
