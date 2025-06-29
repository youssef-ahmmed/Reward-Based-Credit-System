package repository

import (
	"Start/internal/store"
	"errors"
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

func (r *Repository) FetchAllRedemptions(page, limit int, status, dateFrom, dateTo string) ([]*store.Redemption, int, error) {
	var redemptions []*store.Redemption
	var count int64

	query := r.db.Preload("Product").Model(&store.Redemption{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if dateFrom != "" {
		query = query.Where("created_at >= ?", dateFrom)
	}
	if dateTo != "" {
		query = query.Where("created_at <= ?", dateTo)
	}
	query.Count(&count)

	err := query.Order("created_at desc").Offset((page - 1) * limit).Limit(limit).Find(&redemptions).Error
	return redemptions, int(count), err
}

func (r *Repository) UpdateRedemptionStatus(id, status string) error {
	return r.db.Model(&store.Redemption{}).Where("id = ?", id).
		Updates(map[string]interface{}{"status": status}).Error
}

func (r *Repository) FindRedemptionByID(id string) (*store.Redemption, error) {
	var redemption store.Redemption
	err := r.db.First(&redemption, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &redemption, err
}
