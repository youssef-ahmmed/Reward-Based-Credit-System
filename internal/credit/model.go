package credit

import "time"

type Package struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	PriceEGP     float64 `json:"price_egp"`
	Credits      int     `json:"credits"`
	RewardPoints int     `json:"reward_points"`
	IsActive     bool    `json:"is_active"`
}

type Purchase struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	CreditPackageID string    `json:"credit_package_id"`
	CreatedAt       time.Time `json:"created_at"`
}
