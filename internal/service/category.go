package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"Start/internal/types"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryService struct {
	repo *repository.Repository
}

func NewCategoryService(repo *repository.Repository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(c *types.CreateCategoryRequest) (*types.CategoryResponse, error) {
	newCategory := &store.Category{
		ID:          uuid.NewString(),
		Name:        c.Name,
		Description: c.Description,
	}

	if c.ParentCategoryID != "" {
		parentCat, err := s.repo.GetCategoryByID(c.ParentCategoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("parent category does not exist")
			}
			return nil, err
		}
		newCategory.ParentCategoryID = &parentCat.ID
	}

	if err := s.repo.CreateCategory(newCategory); err != nil {
		return nil, err
	}

	return &types.CategoryResponse{
		ID:               newCategory.ID,
		Name:             newCategory.Name,
		Description:      newCategory.Description,
		ParentCategoryID: newCategory.ParentCategoryID,
	}, nil
}

func (s *categoryService) GetAllCategories(parentID *string) ([]store.Category, error) {
	return s.repo.GetAllCategories(parentID)
}

func (s *categoryService) GetCategoryDetails(categoryID string, page, limit int) (*types.CategoryDetailsResponse, error) {
	cat, err := s.repo.GetCategoryByID(categoryID)
	if err != nil || cat == nil {
		return nil, errors.New("category not found")
	}

	products, total, err := s.repo.GetProductsByCategoryID(categoryID, page, limit)
	if err != nil {
		return nil, err
	}

	var productResponses []*types.ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, ToProductResponse(p, cat))
	}

	totalPages := (total + limit - 1) / limit

	return &types.CategoryDetailsResponse{
		Category: cat,
		Products: productResponses,
		Pagination: &types.PaginationMeta{
			CurrentPage:  page,
			ItemsPerPage: limit,
			TotalItems:   total,
			TotalPages:   totalPages,
		},
	}, nil
}

func (s *categoryService) UpdateCategory(id string, input *types.UpdateCategoryRequest) (*types.CategoryResponse, error) {
	existing, err := s.repo.GetCategoryByID(id)
	if err != nil || existing == nil {
		return nil, errors.New("category not found")
	}

	if input.Name != nil {
		existing.Name = *input.Name
	}
	if input.Description != nil {
		existing.Description = *input.Description
	}
	if input.ParentCategoryID != nil {
		if *input.ParentCategoryID != "" {
			parent, _ := s.repo.GetCategoryByID(*input.ParentCategoryID)
			if parent == nil {
				return nil, errors.New("parent category not found")
			}
			existing.ParentCategoryID = input.ParentCategoryID
		} else {
			existing.ParentCategoryID = nil
		}
	}

	if err := s.repo.UpdateCategory(existing); err != nil {
		return nil, err
	}

	return &types.CategoryResponse{
		ID:               existing.ID,
		Name:             existing.Name,
		Description:      existing.Description,
		ParentCategoryID: existing.ParentCategoryID,
	}, nil
}

func (s *categoryService) DeleteCategory(id string) error {
	return s.repo.DeleteCategory(id)
}
