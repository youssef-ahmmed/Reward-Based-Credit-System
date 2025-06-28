package credit

type CreateCreditPackageRequest struct {
	Name         string  `json:"name" binding:"required"`
	PriceEGP     float64 `json:"priceEgp" binding:"required"`
	Credits      int     `json:"credits" binding:"required"`
	RewardPoints int     `json:"rewardPoints" binding:"required"`
	IsActive     bool    `json:"isActive" binding:"required"`
}

type UpdateCreditPackageRequest CreateCreditPackageRequest

type CreditPackageResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	PriceEGP     float64 `json:"priceEgp"`
	Credits      int     `json:"credits"`
	RewardPoints int     `json:"rewardPoints"`
	IsActive     bool    `json:"isActive"`
	CreatedAt    string  `json:"createdAt"`
}

type PaginationMeta struct {
	CurrentPage  int `json:"currentPage"`
	TotalPages   int `json:"totalPages"`
	TotalItems   int `json:"totalItems"`
	ItemsPerPage int `json:"itemsPerPage"`
}

type PaginatedResponse struct {
	Packages   []CreditPackageResponse `json:"packages"`
	Pagination PaginationMeta          `json:"pagination"`
}
