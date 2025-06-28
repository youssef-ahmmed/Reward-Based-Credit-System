package product

import (
	"Start/internal/bridge"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	repo         *Repository
	walletBridge bridge.WalletPort
}

func NewService(repo *Repository, walletBridge bridge.WalletPort) *Service {
	return &Service{
		repo:         repo,
		walletBridge: walletBridge,
	}
}

func (s *Service) GetAllProducts(filters ProductFilters, page, limit int, sortBy, sortOrder string) ([]Product, PaginationMeta, error) {
	validSort := map[string]bool{"name": true, "redemption_points": true, "stock_quantity": true}
	if !validSort[sortBy] {
		sortBy = "created_at"
	}
	if sortOrder != "asc" {
		sortOrder = "desc"
	}

	prods, total, err := s.repo.GetAllProducts(filters, page, limit, sortBy, sortOrder)
	if err != nil {
		return nil, PaginationMeta{}, err
	}

	totalPages := (int(total) + limit - 1) / limit
	return prods, PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}, nil
}

func (s *Service) SearchProducts(query string, filters ProductFilters, page, limit int) ([]Product, PaginationMeta, error) {
	prods, total, err := s.repo.SearchProducts(query, filters, page, limit)
	if err != nil {
		return nil, PaginationMeta{}, err
	}
	totalPages := (int(total) + limit - 1) / limit
	return prods, PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}, nil
}

func (s *Service) CreateProduct(input *CreateProductRequest) (*ProductResponse, error) {
	cat, err := s.repo.GetCategoryByID(input.CategoryID)
	if err != nil || cat == nil {
		return nil, errors.New("invalid category")
	}

	tagsJSON, err := json.Marshal(input.Tags)
	if err != nil {
		return nil, err
	}

	p := &Product{
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

func (s *Service) UpdateProduct(id string, input *UpdateProductRequest) (*ProductResponse, error) {
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

func (s *Service) DeleteProduct(id string) error {
	return s.repo.DeleteProduct(id)
}

func (s *Service) CreateCategory(c *CreateCategoryRequest) (*CategoryResponse, error) {
	newCategory := &Category{
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

	return &CategoryResponse{
		ID:               newCategory.ID,
		Name:             newCategory.Name,
		Description:      newCategory.Description,
		ParentCategoryID: newCategory.ParentCategoryID,
	}, nil
}

func (s *Service) GetAllCategories(parentID *string) ([]Category, error) {
	return s.repo.GetAllCategories(parentID)
}

func (s *Service) GetCategoryDetails(categoryID string, page, limit int) (*CategoryDetailsResponse, error) {
	cat, err := s.repo.GetCategoryByID(categoryID)
	if err != nil || cat == nil {
		return nil, errors.New("category not found")
	}

	products, total, err := s.repo.GetProductsByCategoryID(categoryID, page, limit)
	if err != nil {
		return nil, err
	}

	var productResponses []*ProductResponse
	for _, p := range products {
		productResponses = append(productResponses, ToProductResponse(p, cat))
	}

	totalPages := (total + limit - 1) / limit

	return &CategoryDetailsResponse{
		Category: cat,
		Products: productResponses,
		Pagination: &PaginationMeta{
			CurrentPage:  page,
			ItemsPerPage: limit,
			TotalItems:   total,
			TotalPages:   totalPages,
		},
	}, nil
}

func (s *Service) UpdateCategory(id string, input *UpdateCategoryRequest) (*CategoryResponse, error) {
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

	return &CategoryResponse{
		ID:               existing.ID,
		Name:             existing.Name,
		Description:      existing.Description,
		ParentCategoryID: existing.ParentCategoryID,
	}, nil
}

func (s *Service) DeleteCategory(id string) error {
	return s.repo.DeleteCategory(id)
}

func (s *Service) CreateRedemption(userID string, input CreateRedemptionRequest) (*RedemptionResponse, error) {
	product, err := s.repo.GetProductByID(input.ProductID)
	if err != nil || product == nil {
		return nil, errors.New("product not found")
	}
	if !product.IsOffer {
		return nil, errors.New("product is not available for redemption")
	}
	if input.Quantity <= 0 {
		return nil, errors.New("invalid quantity")
	}
	if input.Quantity > product.StockQuantity {
		return nil, errors.New("insufficient stock")
	}

	pointsRequired := input.Quantity * product.RedemptionPoints
	wallet, err := s.walletBridge.GetWallet(userID) // via bridge
	if err != nil || wallet == nil {
		return nil, errors.New("user wallet not found")
	}
	if wallet.PointsBalance < pointsRequired {
		return nil, errors.New("insufficient points")
	}

	redemptionID := uuid.NewString()
	now := time.Now()

	if err := s.repo.WithTx(func(tx *gorm.DB) error {
		if err := s.walletBridge.DeductPointsTx(tx, userID, pointsRequired); err != nil {
			return err
		}
		if err := s.repo.DecrementStockTx(tx, product.ID, input.Quantity); err != nil {
			return err
		}
		r := &Redemption{
			ID:        redemptionID,
			UserID:    userID,
			ProductID: product.ID,
			Quantity:  input.Quantity,
			CreatedAt: now,
		}
		return tx.Create(r).Error
	}); err != nil {
		return nil, err
	}

	return &RedemptionResponse{
		ID: redemptionID,
		Product: RedemptionProduct{
			ID:           product.ID,
			Name:         product.Name,
			RewardPoints: product.RedemptionPoints,
		},
		Quantity:   input.Quantity,
		PointsUsed: pointsRequired,
		CreatedAt:  now.Format(time.RFC3339),
	}, nil
}

func (s *Service) GetUserRedemptions(userID string, page, limit int) ([]*RedemptionResponse, int64, error) {
	records, total, err := s.repo.ListRedemptionsByUser(userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []*RedemptionResponse
	for _, r := range records {
		responses = append(responses, &RedemptionResponse{
			ID: r.ID,
			Product: RedemptionProduct{
				ID:           r.Product.ID,
				Name:         r.Product.Name,
				RewardPoints: r.Product.RedemptionPoints,
			},
			Quantity:   r.Quantity,
			PointsUsed: r.Quantity * r.Product.RedemptionPoints,
			CreatedAt:  r.CreatedAt.Format(time.RFC3339),
		})
	}
	return responses, total, nil
}

func (s *Service) GetRedemptionByID(userID, id string) (*RedemptionResponse, error) {
	r, err := s.repo.GetRedemptionByID(id)
	if err != nil || r == nil {
		return nil, errors.New("not found")
	}
	if r.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return &RedemptionResponse{
		ID: r.ID,
		Product: RedemptionProduct{
			ID:           r.Product.ID,
			Name:         r.Product.Name,
			RewardPoints: r.Product.RedemptionPoints,
		},
		Quantity:   r.Quantity,
		PointsUsed: r.Quantity * r.Product.RedemptionPoints,
		CreatedAt:  r.CreatedAt.Format(time.RFC3339),
	}, nil
}
