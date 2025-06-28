package user

import (
	"gorm.io/gorm"
)

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
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByID(id string) (*User, error) {
	var user User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdatePassword(userID string, hashedPassword string) error {
	return r.db.Model(&User{}).Where("id = ?", userID).Update("password_hash", hashedPassword).Error
}

func (r *Repository) IsUsernameTaken(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *Repository) CreateWallet(wallet *Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *Repository) GetWalletByUserID(userID string) (*Wallet, error) {
	var wallet Wallet
	if err := r.db.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *Repository) UpdateWallet(wallet *Wallet) error {
	return r.db.Save(wallet).Error
}
