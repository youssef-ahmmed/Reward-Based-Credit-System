package product

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetAllProducts() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *Repository) GetOffers() ([]Product, error) {
	var products []Product
	err := r.db.Where("is_offer = ?", true).Find(&products).Error
	return products, err
}

func (r *Repository) CreateRedemption(red *Redemption) error {
	return r.db.Create(red).Error
}
