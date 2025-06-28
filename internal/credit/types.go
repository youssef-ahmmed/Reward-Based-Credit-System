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

type CreatePurchaseRequest struct {
	CreditPackageID string                 `json:"creditPackageId" binding:"required"`
	PaymentMethod   string                 `json:"paymentMethod" binding:"required"`
	PaymentDetails  map[string]interface{} `json:"paymentDetails"`
}

type PurchaseResponse struct {
	ID                string             `json:"id"`
	UserID            string             `json:"userId"`
	CreditPackageID   string             `json:"creditPackageId"`
	Status            string             `json:"status"`
	CreatedAt         string             `json:"createdAt"`
	CreditPackageInfo *SimplePackageInfo `json:"creditPackage,omitempty"`
}

type SimplePackageInfo struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
