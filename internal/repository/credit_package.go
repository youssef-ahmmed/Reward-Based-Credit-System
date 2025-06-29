package repository

import (
	"Start/internal/store"
)

func (r *Repository) GetPaginatedPackages(page, limit int, activeFilter *bool) ([]store.CreditPackage, int64, error) {
	var packages []store.CreditPackage
	var total int64

	query := r.db.Model(&store.CreditPackage{})

	if activeFilter != nil {
		query = query.Where("is_active = ?", *activeFilter)
	}

	err := query.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&packages).Error
	return packages, total, err
}

func (r *Repository) GetCreditPackageByID(id string) (*store.CreditPackage, error) {
	var pkg store.CreditPackage
	err := r.db.First(&pkg, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *Repository) CreateCreditPackage(pkg *store.CreditPackage) error {
	return r.db.Create(pkg).Error
}

func (r *Repository) UpdateCreditPackage(pkg *store.CreditPackage) error {
	return r.db.Save(pkg).Error
}

func (r *Repository) DeleteCreditPackage(id string) error {
	return r.db.Delete(&store.CreditPackage{}, "id = ?", id).Error
}
