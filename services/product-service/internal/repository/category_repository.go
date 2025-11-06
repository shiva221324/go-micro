package repository

import (
	"context"
	"product-service/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *model.Category) error
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetByNameAndParent(ctx context.Context, name string, parentID *uuid.UUID) (*model.Category, error)
	// UpdateCategory(ctx context.Context, id string, category *model.Category) error
	// DeleteCategory(ctx context.Context, id string) error
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) CreateCategory(ctx context.Context, category *model.Category) error {
	return r.db.WithContext(ctx).Create(category).Error
}

func (r *categoryRepo) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).First(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepo) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepo) GetByNameAndParent(ctx context.Context, name string, parentID *uuid.UUID) (*model.Category, error) {
	var category model.Category
	err := r.db.WithContext(ctx).Where("name = ? AND parent_id = ?", name, parentID).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}
