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
