package product

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func parseInt(val string, def int) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

func parseBoolPtr(val string) *bool {
	if strings.ToLower(val) == "true" {
		b := true
		return &b
	} else if strings.ToLower(val) == "false" {
		b := false
		return &b
	}
	return nil
}

func ToProductResponse(p *Product, c *Category) *ProductResponse {
	var image *string
	if p.ImageURL != "" {
		image = &p.ImageURL
	}

	var tags []string
	if err := json.Unmarshal(p.Tags, &tags); err != nil {
		tags = []string{}
	}

	return &ProductResponse{
		ID:               p.ID,
		Name:             p.Name,
		Description:      p.Description,
		Category:         &CategorySummary{ID: c.ID, Name: c.Name},
		RedemptionPoints: p.RedemptionPoints,
		StockQuantity:    p.StockQuantity,
		IsOffer:          p.IsOffer,
		ImageURL:         image,
		Tags:             tags,
		CreatedAt:        p.CreatedAt.Format(time.RFC3339),
	}
}
