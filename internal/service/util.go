package service

import (
	"Start/internal/store"
	"Start/internal/types"
	"encoding/json"
	"fmt"
	"time"
)

func ToPurchaseResponse(p *store.Purchase, pkg *store.CreditPackage) *types.PurchaseResponse {
	return &types.PurchaseResponse{
		ID:              p.ID,
		UserID:          p.UserID,
		CreditPackageID: p.CreditPackageID,
		Status:          p.Status,
		CreatedAt:       p.CreatedAt.Format(time.RFC3339),
		CreditPackageInfo: &types.SimplePackageInfo{
			ID:    pkg.ID,
			Name:  pkg.Name,
			Price: pkg.PriceEGP,
		},
	}
}

func ToProductResponse(p *store.Product, c *store.Category) *types.ProductResponse {
	var image *string
	if p.ImageURL != "" {
		image = &p.ImageURL
	}

	var tags []string
	if err := json.Unmarshal(p.Tags, &tags); err != nil {
		tags = []string{}
	}

	return &types.ProductResponse{
		ID:               p.ID,
		Name:             p.Name,
		Description:      p.Description,
		Category:         &types.CategorySummary{ID: c.ID, Name: c.Name},
		RedemptionPoints: p.RedemptionPoints,
		StockQuantity:    p.StockQuantity,
		IsOffer:          p.IsOffer,
		ImageURL:         image,
		Tags:             tags,
		CreatedAt:        p.CreatedAt.Format(time.RFC3339),
	}
}

func HumanizeNumber(n int) string {
	switch {
	case n >= 1_000_000:
		return fmt.Sprintf("%.1fM", float64(n)/1_000_000)
	case n >= 1_000:
		return fmt.Sprintf("%.1fK", float64(n)/1_000)
	default:
		return fmt.Sprintf("%d", n)
	}
}

func ToUserResponse(u *store.User) *types.UserDTO {
	return &types.UserDTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Role:      u.Role,
	}
}

func ToRedemptionResponse(r *store.Redemption) *types.RedemptionResponse {
	return &types.RedemptionResponse{
		ID: r.ID,
		Product: types.RedemptionProduct{
			ID:   r.Product.ID,
			Name: r.Product.Name,
		},
		Quantity:   r.Quantity,
		PointsUsed: r.Quantity * r.Product.RedemptionPoints,
		CreatedAt:  r.CreatedAt.Format(time.RFC3339),
	}
}
