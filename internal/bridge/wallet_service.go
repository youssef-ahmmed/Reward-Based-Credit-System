package bridge

import "Start/internal/user"

type WalletServiceBridge struct {
	UserService user.Service
}

func NewWalletServiceBridge(userSvc user.Service) *WalletServiceBridge {
	return &WalletServiceBridge{UserService: userSvc}
}

func (w *WalletServiceBridge) AddToWallet(userID string, credits, points int) error {
	return w.UserService.AddToWallet(userID, credits, points)
}
