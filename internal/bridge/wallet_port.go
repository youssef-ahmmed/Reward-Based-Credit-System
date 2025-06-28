package bridge

import (
	"Start/internal/user"
	"gorm.io/gorm"
)

type WalletPort interface {
	AddToWallet(userID string, credits int, points int) error
	GetWallet(userID string) (*user.Wallet, error)
	DeductPointsTx(tx *gorm.DB, userID string, points int) error
}
