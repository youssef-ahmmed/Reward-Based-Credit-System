package repository

import "Start/internal/store"

func (r *Repository) GetAllCategories(parentID *string) ([]store.Category, error) {
	var categories []store.Category
	query := r.db.Model(&store.Category{})
	if parentID != nil {
		query = query.Where("parent_category_id = ?", *parentID)
	}
	err := query.Find(&categories).Error
	return categories, err
}

func (r *Repository) GetProductsByCategoryID(categoryID string, page, limit int) ([]*store.Product, int, error) {
	var products []*store.Product
	var total int64

	r.db.Model(&store.Product{}).Where("category_id = ?", categoryID).Count(&total)

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

func (r *Repository) CreateCategory(c *store.Category) error {
	return r.db.Create(c).Error
}

func (r *Repository) UpdateCategory(c *store.Category) error {
	return r.db.Save(c).Error
}

func (r *Repository) DeleteCategory(id string) error {
	return r.db.Delete(&store.Category{}, "id = ?", id).Error
}

func (r *Repository) GetCategoryByID(id string) (*store.Category, error) {
	var category store.Category
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
