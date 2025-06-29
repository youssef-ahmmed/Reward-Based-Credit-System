package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type walletService struct {
	repo *repository.Repository
}

func NewWalletService(repo *repository.Repository) WalletService {
	return &walletService{repo: repo}
}

func (s *walletService) GetWallet(userID string) (*store.Wallet, error) {
	wallet, err := s.repo.GetWalletByUserID(userID)
	if err == nil {
		return wallet, nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newWallet := &store.Wallet{
			ID:             uuid.NewString(),
			UserID:         userID,
			CreditsBalance: 0,
			PointsBalance:  0,
			UpdatedAt:      time.Now(),
		}
		if err := s.repo.CreateWallet(newWallet); err != nil {
			return nil, err
		}
		return newWallet, nil
	}

	return nil, err
}

func (s *walletService) DeductPointsTx(tx *gorm.DB, userID string, points int) error {
	return s.repo.DeductPointsTx(tx, userID, points)
}
