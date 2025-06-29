package store

import "time"

type Wallet struct {
	ID             string `json:"id" gorm:"primaryKey"`
	UserID         string `json:"user_id" gorm:"uniqueIndex"`
	User           *User
	PointsBalance  int       `json:"points_balance"`
	CreditsBalance int       `json:"credits_balance"`
	UpdatedAt      time.Time `json:"updated_at"`
}
