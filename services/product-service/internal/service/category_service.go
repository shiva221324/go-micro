package service

import (
	"context"
	"product-service/internal/model"
	"product-service/internal/repository"

	"github.com/google/uuid"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetAllCategories(ctx context.Context) ([]model.Category, error)
	GetByNameAndParent(ctx context.Context, name string, parentID *uuid.UUID) (*model.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	err := s.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	category, err := s.categoryRepo.GetCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) GetAllCategories(ctx context.Context) ([]model.Category, error) {
	return s.categoryRepo.GetAllCategories(ctx)
}

func (s *categoryService) GetByNameAndParent(ctx context.Context, name string, parentID *uuid.UUID) (*model.Category, error) {
	return s.categoryRepo.GetByNameAndParent(ctx, name, parentID)
}

// func (s *categoryService) UpdateCategory(ctx context.Context, id string, category *model.Category) (*model.Category, error) {
// 	err := s.categoryRepo.UpdateCategory(ctx, id, category)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return category, nil
// }

// func (s *categoryService) DeleteCategory(ctx context.Context, id string) error {
// 	return s.categoryRepo.DeleteCategory(ctx, id)
// }
