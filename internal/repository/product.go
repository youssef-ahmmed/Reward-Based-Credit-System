package repository

import (
	"Start/internal/store"
	"Start/internal/types"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (r *Repository) GetAllProducts(filters types.ProductFilters, page, limit int, sortBy, sortOrder string) ([]store.Product, int64, error) {
	var products []store.Product
	var count int64

	query := r.db.Model(&store.Product{})

	if filters.CategoryID != "" {
		query = query.Where("category_id = ?", filters.CategoryID)
	}
	if filters.IsOffer != nil {
		query = query.Where("is_offer = ?", *filters.IsOffer)
	}
	if filters.MinPoints > 0 {
		query = query.Where("redemption_points >= ?", filters.MinPoints)
	}
	if filters.MaxPoints > 0 {
		query = query.Where("redemption_points <= ?", filters.MaxPoints)
	}

	query.Count(&count)

	if sortBy != "" {
		orderClause := fmt.Sprintf("%s %s", sortBy, sortOrder)
		query = query.Order(orderClause)
	}

	err := query.Limit(limit).Offset((page - 1) * limit).Find(&products).Error
	return products, count, err
}

func (r *Repository) SearchProducts(queryStr string, filters types.ProductFilters, page, limit int) ([]store.Product, int64, error) {
	var products []store.Product
	var count int64

	query := r.db.Model(&store.Product{}).
		Where("name ILIKE ? OR description ILIKE ?", "%"+queryStr+"%", "%"+queryStr+"%")

	if filters.CategoryID != "" {
		query = query.Where("category_id = ?", filters.CategoryID)
	}
	if filters.IsOffer != nil {
		query = query.Where("is_offer = ?", *filters.IsOffer)
	}
	if filters.MinPoints > 0 {
		query = query.Where("redemption_points >= ?", filters.MinPoints)
	}
	if filters.MaxPoints > 0 {
		query = query.Where("redemption_points <= ?", filters.MaxPoints)
	}

	query.Count(&count)
	err := query.Limit(limit).Offset((page - 1) * limit).Find(&products).Error
	return products, count, err
}

func (r *Repository) CreateProduct(p *store.Product) error {
	return r.db.Create(p).Error
}

func (r *Repository) UpdateProduct(p *store.Product) error {
	return r.db.Save(p).Error
}

func (r *Repository) DeleteProduct(id string) error {
	return r.db.Delete(&store.Product{}, "id = ?", id).Error
}

func (r *Repository) GetProductByID(id string) (*store.Product, error) {
	var p store.Product
	err := r.db.First(&p, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &p, err
}

func (r *Repository) FetchProductsByPreferences(includeCats, excludeCats []string, minPoints, maxPoints, limit int) ([]store.Product, error) {
	query := r.db.Preload("Category").
		Where("redemption_points >= ? AND redemption_points <= ?", minPoints, maxPoints).
		Limit(limit)

	if len(includeCats) > 0 {
		query = query.Where("category_id IN ?", includeCats)
	}
	if len(excludeCats) > 0 {
		query = query.Where("category_id NOT IN ?", excludeCats)
	}

	var products []store.Product
	err := query.Find(&products).Error
	return products, err
}
