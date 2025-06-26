package credit

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetActivePackages() ([]Package, error) {
	var packages []Package
	err := r.db.Where("is_active = ?", true).Find(&packages).Error
	return packages, err
}

func (r *Repository) CreatePurchase(p *Purchase) error {
	return r.db.Create(p).Error
}
