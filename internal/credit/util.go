package credit

import "time"

func ToPurchaseResponse(p *Purchase, pkg *CreditPackage) *PurchaseResponse {
	return &PurchaseResponse{
		ID:              p.ID,
		UserID:          p.UserID,
		CreditPackageID: p.CreditPackageID,
		Status:          p.Status,
		CreatedAt:       p.CreatedAt.Format(time.RFC3339),
		CreditPackageInfo: &SimplePackageInfo{
			ID:    pkg.ID,
			Name:  pkg.Name,
			Price: pkg.PriceEGP,
		},
	}
}
