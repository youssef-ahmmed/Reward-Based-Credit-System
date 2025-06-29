package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"Start/internal/types"
	"errors"
	"github.com/google/uuid"
	"time"
)

type creditPackageService struct {
	repo *repository.Repository
}

func NewCreditPackageService(repo *repository.Repository) CreditPackageService {
	return &creditPackageService{repo: repo}
}

func (s *creditPackageService) GetAllCreditPackages(page, limit int, activeFilter *bool) ([]types.CreditCreditPackageResponse, types.PaginationMeta, error) {
	results, total, err := s.repo.GetPaginatedPackages(page, limit, activeFilter)
	if err != nil {
		return nil, types.PaginationMeta{}, err
	}

	totalPages := (int(total) + limit - 1) / limit
	var res []types.CreditCreditPackageResponse
	for _, p := range results {
		res = append(res, types.CreditCreditPackageResponse{
			ID:           p.ID,
			Name:         p.Name,
			PriceEGP:     p.PriceEGP,
			Credits:      p.Credits,
			RewardPoints: p.RewardPoints,
			IsActive:     p.IsActive,
			CreatedAt:    p.CreatedAt.Format(time.RFC3339),
		})
	}

	return res, types.PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}, nil
}

func (s *creditPackageService) GetCreditPackageByID(id string) (*types.CreditCreditPackageResponse, error) {
	pkg, err := s.repo.GetCreditPackageByID(id)
	if err != nil {
		return nil, err
	}
	return &types.CreditCreditPackageResponse{
		ID:           pkg.ID,
		Name:         pkg.Name,
		PriceEGP:     pkg.PriceEGP,
		Credits:      pkg.Credits,
		RewardPoints: pkg.RewardPoints,
		IsActive:     pkg.IsActive,
		CreatedAt:    pkg.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *creditPackageService) CreateCreditPackage(input types.CreateCreditPackageRequest) (*types.CreditCreditPackageResponse, error) {
	pkg := &store.CreditPackage{
		ID:           uuid.NewString(),
		Name:         input.Name,
		PriceEGP:     input.PriceEGP,
		Credits:      input.Credits,
		RewardPoints: input.RewardPoints,
		IsActive:     input.IsActive,
		CreatedAt:    time.Now(),
	}

	if err := s.repo.CreateCreditPackage(pkg); err != nil {
		return nil, err
	}

	return &types.CreditCreditPackageResponse{
		ID:           pkg.ID,
		Name:         pkg.Name,
		PriceEGP:     pkg.PriceEGP,
		Credits:      pkg.Credits,
		RewardPoints: pkg.RewardPoints,
		IsActive:     pkg.IsActive,
		CreatedAt:    pkg.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *creditPackageService) UpdateCreditPackages(id string, input types.UpdateCreditPackageRequest) (*types.CreditCreditPackageResponse, error) {
	pkg, err := s.repo.GetCreditPackageByID(id)
	if err != nil {
		return nil, errors.New("package not found")
	}

	pkg.Name = input.Name
	pkg.PriceEGP = input.PriceEGP
	pkg.Credits = input.Credits
	pkg.RewardPoints = input.RewardPoints
	pkg.IsActive = input.IsActive

	if err := s.repo.UpdateCreditPackage(pkg); err != nil {
		return nil, err
	}

	return &types.CreditCreditPackageResponse{
		ID:           pkg.ID,
		Name:         pkg.Name,
		PriceEGP:     pkg.PriceEGP,
		Credits:      pkg.Credits,
		RewardPoints: pkg.RewardPoints,
		IsActive:     pkg.IsActive,
		CreatedAt:    pkg.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *creditPackageService) DeleteCreditPackage(id string) error {
	return s.repo.DeleteCreditPackage(id)
}
