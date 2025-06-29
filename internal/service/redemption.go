package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"Start/internal/types"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type redemptionService struct {
	repo *repository.Repository
}

func NewRedemptionService(repo *repository.Repository) RedemptionService {
	return &redemptionService{repo: repo}
}

func (s *redemptionService) CreateRedemption(userID string, input types.CreateRedemptionRequest) (*types.RedemptionResponse, error) {
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
	wallet, err := s.repo.GetWalletByUserID(userID)
	if err != nil || wallet == nil {
		return nil, errors.New("user wallet not found")
	}
	if wallet.PointsBalance < pointsRequired {
		return nil, errors.New("insufficient points")
	}

	redemptionID := uuid.NewString()
	now := time.Now()

	if err := s.repo.WithTx(func(tx *gorm.DB) error {
		if err := s.repo.DeductPointsTx(tx, userID, pointsRequired); err != nil {
			return err
		}
		if err := s.repo.DecrementStockTx(tx, product.ID, input.Quantity); err != nil {
			return err
		}
		r := &store.Redemption{
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

	return &types.RedemptionResponse{
		ID: redemptionID,
		Product: types.RedemptionProduct{
			ID:           product.ID,
			Name:         product.Name,
			RewardPoints: product.RedemptionPoints,
		},
		Quantity:   input.Quantity,
		PointsUsed: pointsRequired,
		CreatedAt:  now.Format(time.RFC3339),
	}, nil
}

func (s *redemptionService) GetUserRedemptions(userID string, page, limit int) ([]*types.RedemptionResponse, int64, error) {
	records, total, err := s.repo.ListRedemptionsByUser(userID, page, limit)
	if err != nil {
		return nil, 0, err
	}

	var responses []*types.RedemptionResponse
	for _, r := range records {
		responses = append(responses, &types.RedemptionResponse{
			ID: r.ID,
			Product: types.RedemptionProduct{
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

func (s *redemptionService) GetRedemptionByID(userID, id string) (*types.RedemptionResponse, error) {
	r, err := s.repo.GetRedemptionByID(id)
	if err != nil || r == nil {
		return nil, errors.New("not found")
	}
	if r.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return &types.RedemptionResponse{
		ID: r.ID,
		Product: types.RedemptionProduct{
			ID:           r.Product.ID,
			Name:         r.Product.Name,
			RewardPoints: r.Product.RedemptionPoints,
		},
		Quantity:   r.Quantity,
		PointsUsed: r.Quantity * r.Product.RedemptionPoints,
		CreatedAt:  r.CreatedAt.Format(time.RFC3339),
	}, nil
}
