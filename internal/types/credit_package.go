package types

type CreateCreditPackageRequest struct {
	Name         string  `json:"name" binding:"required"`
	PriceEGP     float64 `json:"priceEgp" binding:"required"`
	Credits      int     `json:"credits" binding:"required"`
	RewardPoints int     `json:"rewardPoints" binding:"required"`
	IsActive     bool    `json:"isActive" binding:"required"`
}

type UpdateCreditPackageRequest CreateCreditPackageRequest

type CreditCreditPackageResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	PriceEGP     float64 `json:"priceEgp"`
	Credits      int     `json:"credits"`
	RewardPoints int     `json:"rewardPoints"`
	IsActive     bool    `json:"isActive"`
	CreatedAt    string  `json:"createdAt"`
}

type PaginatedResponse struct {
	Packages   []CreditCreditPackageResponse `json:"packages"`
	Pagination PaginationMeta                `json:"pagination"`
}
