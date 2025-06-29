package repository

import (
	"Start/internal/store"
)

func (r *Repository) CreatePurchase(p *store.Purchase) error {
	return r.db.Create(p).Error
}

func (r *Repository) GetUserPurchases(userID, status string, page, limit int) ([]store.Purchase, int64, error) {
	var purchases []store.Purchase
	var count int64

	q := r.db.Where("user_id = ?", userID)
	if status != "" {
		q = q.Where("status = ?", status)
	}

	if err := q.Model(&store.Purchase{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := q.Preload("CreditPackage").
		Offset((page - 1) * limit).
		Limit(limit).
		Order("created_at DESC").
		Find(&purchases).Error

	return purchases, count, err
}

func (r *Repository) GetPurchaseByID(id string) (*store.Purchase, error) {
	var p store.Purchase
	err := r.db.Preload("CreditPackage").Where("id = ?", id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *Repository) CountTotalPurchases() (int, error) {
	var count int64
	err := r.db.Model(&store.Purchase{}).Count(&count).Error
	return int(count), err
}

func (r *Repository) SumCreditsIssued() (int, error) {
	var total int
	err := r.db.Model(&store.Purchase{}).Select("SUM(credits)").Scan(&total).Error
	return total, err
}

func (r *Repository) FetchAllPurchases(page, limit int, status, dateFrom, dateTo string) ([]*store.Purchase, int, error) {
	var purchases []*store.Purchase
	var count int64

	query := r.db.Model(&store.Purchase{}).
		Preload("CreditPackage")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if dateFrom != "" {
		query = query.Where("created_at >= ?", dateFrom)
	}
	if dateTo != "" {
		query = query.Where("created_at <= ?", dateTo)
	}
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.
		Offset((page - 1) * limit).
		Limit(limit).
		Order("created_at DESC").
		Find(&purchases).Error

	return purchases, int(count), err
}
