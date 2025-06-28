package product

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllProducts(filters ProductFilters, page, limit int, sortBy, sortOrder string) ([]Product, int64, error) {
	var products []Product
	var count int64

	query := r.db.Model(&Product{})

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

func (r *Repository) SearchProducts(queryStr string, filters ProductFilters, page, limit int) ([]Product, int64, error) {
	var products []Product
	var count int64

	query := r.db.Model(&Product{}).
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

func (r *Repository) CreateProduct(p *Product) error {
	return r.db.Create(p).Error
}

func (r *Repository) UpdateProduct(p *Product) error {
	return r.db.Save(p).Error
}

func (r *Repository) DeleteProduct(id string) error {
	return r.db.Delete(&Product{}, "id = ?", id).Error
}

func (r *Repository) GetProductByID(id string) (*Product, error) {
	var p Product
	err := r.db.First(&p, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &p, err
}

func (r *Repository) GetAllCategories(parentID *string) ([]Category, error) {
	var categories []Category
	query := r.db.Model(&Category{})
	if parentID != nil {
		query = query.Where("parent_category_id = ?", *parentID)
	}
	err := query.Find(&categories).Error
	return categories, err
}

func (r *Repository) GetProductsByCategoryID(categoryID string, page, limit int) ([]*Product, int, error) {
	var products []*Product
	var total int64

	r.db.Model(&Product{}).Where("category_id = ?", categoryID).Count(&total)

	offset := (page - 1) * limit
	err := r.db.Where("category_id = ?", categoryID).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&products).Error

	if err != nil {
		return nil, 0, err
	}

	return products, int(total), nil
}

func (r *Repository) CreateCategory(c *Category) error {
	return r.db.Create(c).Error
}

func (r *Repository) UpdateCategory(c *Category) error {
	return r.db.Save(c).Error
}

func (r *Repository) DeleteCategory(id string) error {
	return r.db.Delete(&Category{}, "id = ?", id).Error
}

func (r *Repository) GetCategoryByID(id string) (*Category, error) {
	var category Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *Repository) WithTx(fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *Repository) DecrementStockTx(tx *gorm.DB, productID string, quantity int) error {
	return tx.Model(&Product{}).Where("id = ? AND stock_quantity >= ?", productID, quantity).
		UpdateColumn("stock_quantity", gorm.Expr("stock_quantity - ?", quantity)).Error
}

func (r *Repository) ListRedemptionsByUser(userID string, page, limit int) ([]*Redemption, int64, error) {
	var redemptions []*Redemption
	var total int64

	offset := (page - 1) * limit

	err := r.db.Model(&Redemption{}).Where("user_id = ?", userID).Count(&total).Error
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

func (r *Repository) GetRedemptionByID(id string) (*Redemption, error) {
	var rdm Redemption
	err := r.db.Preload("Product").First(&rdm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &rdm, nil
}
