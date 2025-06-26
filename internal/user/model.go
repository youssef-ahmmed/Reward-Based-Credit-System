package user

import "time"

type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`   // "user" or "admin"
	Status       string    `json:"status"` // "suspended" or "active", "banned"
	CreatedAt    time.Time `json:"created_at"`
}

type Wallet struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	PointsBalance  int       `json:"points_balance"`
	CreditsBalance int       `json:"credits_balance"`
	UpdatedAt      time.Time `json:"updated_at"`
}
