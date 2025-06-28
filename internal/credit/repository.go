package credit

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetPaginatedPackages(page, limit int, activeFilter *bool) ([]CreditPackage, int64, error) {
	var packages []CreditPackage
	var total int64

	query := r.db.Model(&CreditPackage{})

	if activeFilter != nil {
		query = query.Where("is_active = ?", *activeFilter)
	}

	err := query.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&packages).Error
	return packages, total, err
}

func (r *Repository) GetByID(id string) (*CreditPackage, error) {
	var pkg CreditPackage
	err := r.db.First(&pkg, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *Repository) Create(pkg *CreditPackage) error {
	return r.db.Create(pkg).Error
}

func (r *Repository) Update(pkg *CreditPackage) error {
	return r.db.Save(pkg).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&CreditPackage{}, "id = ?", id).Error
}

func (r *Repository) CreatePurchase(p *Purchase) error {
	return r.db.Create(p).Error
}

func (r *Repository) GetUserPurchases(userID, status string, page, limit int) ([]Purchase, int64, error) {
	var purchases []Purchase
	var count int64

	q := r.db.Where("user_id = ?", userID)
	if status != "" {
		q = q.Where("status = ?", status)
	}

	if err := q.Model(&Purchase{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	err := q.Preload("CreditPackage").
		Offset((page - 1) * limit).
		Limit(limit).
		Order("created_at DESC").
		Find(&purchases).Error

	return purchases, count, err
}

func (r *Repository) GetPurchaseByID(id string) (*Purchase, error) {
	var p Purchase
	err := r.db.Preload("CreditPackage").Where("id = ?", id).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}
