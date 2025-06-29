package repository

import (
	"Start/internal/store"
	"gorm.io/gorm"
	"time"
)

func (r *Repository) SumPointsEarned() (int, error) {
	var total int
	err := r.db.Model(&store.Wallet{}).Select("SUM(points_balance)").Scan(&total).Error
	return total, err
}

func (r *Repository) CreateWallet(wallet *store.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *Repository) GetWalletByUserID(userID string) (*store.Wallet, error) {
	var wallet store.Wallet
	if err := r.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *Repository) AddToWallet(userID string, credits int, points int) error {
	return r.db.Model(&store.Wallet{}).
		Where("user_id = ?", userID).
		Updates(map[string]interface{}{
			"credits_balance": gorm.Expr("credits_balance + ?", credits),
			"points_balance":  gorm.Expr("points_balance + ?", points),
			"updated_at":      time.Now(),
		}).Error
}

func (r *Repository) UpdateWallet(wallet *store.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *Repository) UpdateWalletCredits(userID string, amount int) error {
	return r.db.Model(&store.Wallet{}).Where("user_id = ?", userID).
		UpdateColumn("credits_balance", gorm.Expr("credits_balance + ?", amount)).Error
}

func (r *Repository) UpdateWalletPoints(userID string, amount int) error {
	return r.db.Model(&store.Wallet{}).Where("user_id = ?", userID).
		Update("points_balance", gorm.Expr("points_balance + ?", amount)).Error
}

func (r *Repository) DeductPointsTx(tx *gorm.DB, userID string, points int) error {
	return tx.Model(&store.Wallet{}).Where("user_id = ? AND points_balance >= ?", userID, points).
		UpdateColumn("points_balance", gorm.Expr("points_balance - ?", points)).Error
}
