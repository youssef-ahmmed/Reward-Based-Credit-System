package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"Start/internal/types"
	"errors"
	"github.com/google/uuid"
	"time"
)

type purchaseResponse struct {
	repo *repository.Repository
}

func NewPurchaseService(repo *repository.Repository) PurchaseService {
	return &purchaseResponse{repo: repo}
}

func (s *purchaseResponse) CreatePurchase(userID string, input types.CreatePurchaseRequest) (*types.PurchaseResponse, error) {
	pkg, err := s.repo.GetCreditPackageByID(input.CreditPackageID)
	if err != nil {
		return nil, errors.New("package not found")
	}

	if input.PaymentMethod != "credit_card" {
		return nil, errors.New("payment failed")
	}

	p := &store.Purchase{
		ID:              uuid.NewString(),
		UserID:          userID,
		CreditPackageID: input.CreditPackageID,
		CreatedAt:       time.Now(),
	}

	if err := s.repo.CreatePurchase(p); err != nil {
		return nil, err
	}

	_ = s.repo.AddToWallet(userID, pkg.Credits, pkg.RewardPoints)

	return ToPurchaseResponse(p, pkg), nil
}

func (s *purchaseResponse) GetUserPurchases(userID, status string, page int, limit int) ([]types.PurchaseResponse, types.PaginationMeta, error) {
	purchases, total, err := s.repo.GetUserPurchases(userID, status, page, limit)
	if err != nil {
		return nil, types.PaginationMeta{}, err
	}

	var res []types.PurchaseResponse
	for _, p := range purchases {
		pkg := p.CreditPackage
		res = append(res, *ToPurchaseResponse(&p, &pkg))
	}

	totalPages := (int(total) + limit - 1) / limit
	meta := types.PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}
	return res, meta, nil
}

func (s *purchaseResponse) GetPurchaseByID(userID, purchaseID string) (*types.PurchaseResponse, error) {
	p, err := s.repo.GetPurchaseByID(purchaseID)
	if err != nil {
		return nil, errors.New("not found")
	}
	if p.UserID != userID {
		return nil, errors.New("forbidden")
	}
	return ToPurchaseResponse(p, &p.CreditPackage), nil
}

func (s *purchaseResponse) CountTotalPurchases() (int, error) {
	return s.repo.CountTotalPurchases()
}

func (s *purchaseResponse) SumCreditsIssued() (int, error) {
	return s.repo.SumCreditsIssued()
}
