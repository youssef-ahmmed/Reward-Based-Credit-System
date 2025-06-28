package credit

import "time"

type CreditPackage struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	PriceEGP     float64   `json:"price_egp"`
	Credits      int       `json:"credits"`
	RewardPoints int       `json:"reward_points"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}

type Purchase struct {
	ID              string `gorm:"primaryKey"`
	UserID          string
	CreditPackageID string
	Status          string
	CreatedAt       time.Time

	CreditPackage CreditPackage `gorm:"foreignKey:CreditPackageID"`
}
