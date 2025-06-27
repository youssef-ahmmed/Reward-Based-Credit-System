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
