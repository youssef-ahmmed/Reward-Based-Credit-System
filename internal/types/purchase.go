package types

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
	Credits           int                `json:"credits"`
	CreatedAt         string             `json:"createdAt"`
	CreditPackageInfo *SimplePackageInfo `json:"creditPackage,omitempty"`
}

type SimplePackageInfo struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
