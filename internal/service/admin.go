package service

import (
	"Start/internal/repository"
	"Start/internal/types"
	"errors"
)

type adminService struct {
	repo *repository.Repository
}

func NewAdminService(repo *repository.Repository) AdminService {
	return &adminService{repo: repo}
}

func (s *adminService) GetAdminDashboardStats() (*types.DashboardStatsResponse, error) {
	totalUsers, err := s.repo.CountTotalUsers()
	if err != nil {
		return nil, err
	}

	totalOrders, err := s.repo.CountTotalPurchases()
	if err != nil {
		return nil, err
	}

	creditsIssued, err := s.repo.SumCreditsIssued()
	if err != nil {
		return nil, err
	}

	pointsEarned, err := s.repo.SumPointsEarned()
	if err != nil {
		return nil, err
	}

	return &types.DashboardStatsResponse{
		TotalUsers:    totalUsers,
		TotalOrders:   totalOrders,
		CreditsIssued: HumanizeNumber(creditsIssued),
		PointsEarned:  HumanizeNumber(pointsEarned),
	}, nil
}

func (s *adminService) GetAllUsers(page, limit int, search, sortBy, sortOrder string) ([]*types.UserDTO, int, error) {
	users, total, err := s.repo.FetchAllUsers(page, limit, search, sortBy, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	var result []*types.UserDTO
	for _, u := range users {
		result = append(result, ToUserResponse(u))
	}

	return result, total, nil
}

func (s *adminService) GetAllPurchases(page, limit int, status, dateFrom, dateTo string) ([]*types.PurchaseResponse, int, error) {
	purchases, total, err := s.repo.FetchAllPurchases(page, limit, status, dateFrom, dateTo)
	if err != nil {
		return nil, 0, err
	}

	var result []*types.PurchaseResponse
	for _, p := range purchases {
		result = append(result, ToPurchaseResponse(p, &p.CreditPackage))
	}

	return result, total, nil
}

func (s *adminService) GetAllRedemptions(page, limit int, status, dateFrom, dateTo string) ([]*types.RedemptionResponse, int, error) {
	redemptions, total, err := s.repo.FetchAllRedemptions(page, limit, status, dateFrom, dateTo)
	if err != nil {
		return nil, 0, err
	}

	var result []*types.RedemptionResponse
	for _, r := range redemptions {
		result = append(result, ToRedemptionResponse(r))
	}

	return result, total, nil
}

func (s *adminService) UpdateRedemptionStatus(id, status, notes string) error {
	if status != "pending" && status != "delivered" && status != "cancelled" {
		return errors.New("invalid status")
	}

	r, err := s.repo.FindRedemptionByID(id)
	if err != nil {
		return err
	}
	if r == nil {
		return errors.New("not found")
	}

	return s.repo.UpdateRedemptionStatus(id, status, notes)
}

func (s *adminService) ManageUserCredits(userID, action string, amount int) error {
	if action != "add" && action != "subtract" {
		return errors.New("invalid action")
	}

	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	if action == "add" {
		return s.repo.UpdateWalletCredits(userID, amount)
	}
	return s.repo.UpdateWalletCredits(userID, -amount)
}

func (s *adminService) UpdateUserStatus(userID, status string) error {
	if status != "active" && status != "suspended" && status != "banned" {
		return errors.New("invalid status")
	}

	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	return s.repo.UpdateUserStatus(userID, status)
}
