package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) IsEmailOrUsernameTaken(email, username string) (bool, error) {
	var count int64
	err := r.db.Model(&User{}).
		Where("email = ? OR username = ?", email, username).
		Count(&count).Error
	return count > 0, err
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) GetWallet(userID string) (*Wallet, error) {
	var wallet Wallet
	err := r.db.Where("user_id = ?", userID).First(&wallet).Error
	return &wallet, err
}

func (r *Repository) UpdateWallet(wallet *Wallet) error {
	return r.db.Save(wallet).Error
}
