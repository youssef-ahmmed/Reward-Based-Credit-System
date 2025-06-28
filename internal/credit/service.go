package credit

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllCreditPackages(page, limit int, activeFilter *bool) ([]CreditPackageResponse, PaginationMeta, error) {
	results, total, err := s.repo.GetPaginatedPackages(page, limit, activeFilter)
	if err != nil {
		return nil, PaginationMeta{}, err
	}

	totalPages := (int(total) + limit - 1) / limit
	var res []CreditPackageResponse
	for _, p := range results {
		res = append(res, CreditPackageResponse{
			ID:           p.ID,
			Name:         p.Name,
			PriceEGP:     p.PriceEGP,
			Credits:      p.Credits,
			RewardPoints: p.RewardPoints,
			IsActive:     p.IsActive,
			CreatedAt:    p.CreatedAt.Format(time.RFC3339),
		})
	}

	return res, PaginationMeta{
		CurrentPage:  page,
		TotalPages:   totalPages,
		TotalItems:   int(total),
		ItemsPerPage: limit,
	}, nil
}

func (s *Service) GetCreditPackageByID(id string) (*CreditPackageResponse, error) {
	pkg, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &CreditPackageResponse{
		ID:           pkg.ID,
		Name:         pkg.Name,
		PriceEGP:     pkg.PriceEGP,
		Credits:      pkg.Credits,
		RewardPoints: pkg.RewardPoints,
		IsActive:     pkg.IsActive,
		CreatedAt:    pkg.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) CreateCreditPackage(input CreateCreditPackageRequest) (*CreditPackageResponse, error) {
	pkg := &CreditPackage{
		ID:           uuid.NewString(),
		Name:         input.Name,
		PriceEGP:     input.PriceEGP,
		Credits:      input.Credits,
		RewardPoints: input.RewardPoints,
		IsActive:     input.IsActive,
		CreatedAt:    time.Now(),
	}

	if err := s.repo.Create(pkg); err != nil {
		return nil, err
	}

	return &CreditPackageResponse{
		ID:           pkg.ID,
		Name:         pkg.Name,
		PriceEGP:     pkg.PriceEGP,
		Credits:      pkg.Credits,
		RewardPoints: pkg.RewardPoints,
		IsActive:     pkg.IsActive,
		CreatedAt:    pkg.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) UpdateCreditPackages(id string, input UpdateCreditPackageRequest) (*CreditPackageResponse, error) {
	pkg, err := s.repo.GetByID(id)
	if err != nil {
		return nil, errors.New("package not found")
	}

	pkg.Name = input.Name
	pkg.PriceEGP = input.PriceEGP
	pkg.Credits = input.Credits
	pkg.RewardPoints = input.RewardPoints
	pkg.IsActive = input.IsActive

	if err := s.repo.Update(pkg); err != nil {
		return nil, err
	}

	return &CreditPackageResponse{
		ID:           pkg.ID,
		Name:         pkg.Name,
		PriceEGP:     pkg.PriceEGP,
		Credits:      pkg.Credits,
		RewardPoints: pkg.RewardPoints,
		IsActive:     pkg.IsActive,
		CreatedAt:    pkg.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) DeleteCreditPackage(id string) error {
	return s.repo.Delete(id)
}
