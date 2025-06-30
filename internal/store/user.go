package store

import "time"

type User struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username" gorm:"uniqueIndex"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`   // "user" or "admin"
	Status       string    `json:"status"` // "suspended" or "active", "banned"
	CreatedAt    time.Time `json:"created_at"`

	Wallet Wallet `gorm:"foreignKey:UserID"`
}
