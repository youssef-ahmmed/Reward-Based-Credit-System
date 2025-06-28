package bridge

import (
	"Start/internal/user"
	"gorm.io/gorm"
)

type WalletServiceBridge struct {
	UserService user.Service
}

func NewWalletServiceBridge(userSvc user.Service) *WalletServiceBridge {
	return &WalletServiceBridge{UserService: userSvc}
}

func (w *WalletServiceBridge) GetWallet(userID string) (*user.Wallet, error) {
	return w.UserService.GetWallet(userID)
}

func (w *WalletServiceBridge) DeductPointsTx(tx *gorm.DB, userID string, points int) error {
	return w.UserService.DeductPointsTx(tx, userID, points)
}

func (w *WalletServiceBridge) AddToWallet(userID string, credits, points int) error {
	return w.UserService.AddToWallet(userID, credits, points)
}
