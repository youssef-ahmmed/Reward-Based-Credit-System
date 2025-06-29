package repository

import (
	"Start/internal/store"
	"gorm.io/gorm"
)

func (r *Repository) WithTx(fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *Repository) DecrementStockTx(tx *gorm.DB, productID string, quantity int) error {
	return tx.Model(&store.Product{}).Where("id = ? AND stock_quantity >= ?", productID, quantity).
		UpdateColumn("stock_quantity", gorm.Expr("stock_quantity - ?", quantity)).Error
}

func (r *Repository) ListRedemptionsByUser(userID string, page, limit int) ([]*store.Redemption, int64, error) {
	var redemptions []*store.Redemption
	var total int64

	offset := (page - 1) * limit

	err := r.db.Model(&store.Redemption{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Product").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&redemptions).Error

	return redemptions, total, err
}

func (r *Repository) GetRedemptionByID(id string) (*store.Redemption, error) {
	var rdm store.Redemption
	err := r.db.Preload("Product").First(&rdm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &rdm, nil
}
