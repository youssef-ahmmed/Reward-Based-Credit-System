package repository

import (
	"Start/internal/store"
	"errors"
	"gorm.io/gorm"
)

func (r *Repository) CreateUser(user *store.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) IsEmailOrUsernameTaken(email, username string) (bool, error) {
	var count int64
	err := r.db.Model(&store.User{}).
		Where("email = ? OR username = ?", email, username).
		Count(&count).Error
	return count > 0, err
}

func (r *Repository) FindByEmail(email string) (*store.User, error) {
	var user store.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByID(id string) (*store.User, error) {
	var user store.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdatePassword(userID string, hashedPassword string) error {
	return r.db.Model(&store.User{}).Where("id = ?", userID).Update("password_hash", hashedPassword).Error
}

func (r *Repository) IsUsernameTaken(username string) (bool, error) {
	var count int64
	if err := r.db.Model(&store.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) UpdateUser(user *store.User) error {
	return r.db.Save(user).Error
}

func (r *Repository) CountTotalUsers() (int, error) {
	var count int64
	err := r.db.Model(&store.User{}).Count(&count).Error
	return int(count), err
}

func (r *Repository) FetchAllUsers(page, limit int, search, sortBy, sortOrder string) ([]*store.User, int, error) {
	var users []*store.User
	var count int64

	query := r.db.Model(&store.User{})
	if search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query.Count(&count)

	err := query.Order(sortBy + " " + sortOrder).Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	return users, int(count), err
}

func (r *Repository) UpdateUserStatus(userID, status string) error {
	return r.db.Model(&store.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{"status": status}).Error
}

func (r *Repository) FindUserByID(id string) (*store.User, error) {
	var user store.User
	err := r.db.First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}
