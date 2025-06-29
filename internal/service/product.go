package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"Start/internal/types"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type productService struct {
	repo *repository.Repository
}

func NewProductService(repo *repository.Repository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts(filters types.ProductFilters, page, limit int, sortBy, sortOrder string) ([]store.Product, types.PaginationMeta, error) {
	validSort := map[string]bool{"name": true, "redemption_points": true, "stock_quantity": true}
	if !validSort[sortBy] {
		sortBy = "created_at"
	}
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	prods, total, err := s.repo.GetAllProducts(filters, page, limit, sortBy, sortOrder)
	if err != nil {
		return nil, types.PaginationMeta{}, err
	}

	totalPages := (int(total) + limit - 1) / limit
	return prods, types.PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}, nil
}

func (s *productService) SearchProducts(query string, filters types.ProductFilters, page, limit int) ([]store.Product, types.PaginationMeta, error) {
	prods, total, err := s.repo.SearchProducts(query, filters, page, limit)
	if err != nil {
		return nil, types.PaginationMeta{}, err
	}
	totalPages := (int(total) + limit - 1) / limit
	return prods, types.PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}, nil
}

func (s *productService) CreateProduct(input *types.CreateProductRequest) (*types.ProductResponse, error) {
	cat, err := s.repo.GetCategoryByID(input.CategoryID)
	if err != nil || cat == nil {
		return nil, errors.New("invalid category")
	}

	tagsJSON, err := json.Marshal(input.Tags)
	if err != nil {
		return nil, err
	}

	p := &store.Product{
		ID:               uuid.NewString(),
		Name:             input.Name,
		Description:      input.Description,
		CategoryID:       input.CategoryID,
		RedemptionPoints: input.RedemptionPoints,
		StockQuantity:    input.StockQuantity,
		IsOffer:          input.IsOffer,
		CreatedAt:        time.Now(),
		Tags:             tagsJSON,
	}

	if input.ImageURL != nil {
		p.ImageURL = *input.ImageURL
	}

	if err := s.repo.CreateProduct(p); err != nil {
		return nil, err
	}

	return ToProductResponse(p, cat), nil
}

func (s *productService) UpdateProduct(id string, input *types.UpdateProductRequest) (*types.ProductResponse, error) {
	existing, err := s.repo.GetProductByID(id)
	if err != nil || existing == nil {
		return nil, errors.New("product not found")
	}

	if input.Name != nil {
		existing.Name = *input.Name
	}
	if input.Description != nil {
		existing.Description = *input.Description
	}
	if input.CategoryID != nil {
		cat, _ := s.repo.GetCategoryByID(*input.CategoryID)
		if cat == nil {
			return nil, errors.New("invalid category ID")
		}
		existing.CategoryID = *input.CategoryID
	}
	if input.RedemptionPoints != nil {
		existing.RedemptionPoints = *input.RedemptionPoints
	}
	if input.StockQuantity != nil {
		existing.StockQuantity = *input.StockQuantity
	}
	if input.IsOffer != nil {
		existing.IsOffer = *input.IsOffer
	}
	if input.ImageURL != nil {
		existing.ImageURL = *input.ImageURL
	}
	if input.Tags != nil {
		tagsJSON, err := json.Marshal(input.Tags)
		if err != nil {
			return nil, err
		}
		existing.Tags = tagsJSON
	}

	if err := s.repo.UpdateProduct(existing); err != nil {
		return nil, err
	}

	category, _ := s.repo.GetCategoryByID(existing.CategoryID)

	return ToProductResponse(existing, category), nil
}

func (s *productService) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}
