package service

import (
	"Start/internal/repository"
	"Start/internal/store"
	"Start/internal/types"
	"strconv"
	"time"
)

type aiService struct {
	repo *repository.Repository
}

func NewAIService(repo *repository.Repository) AIService {
	return &aiService{repo}
}

func (s *aiService) RecommendProducts(req types.RecommendationRequest) (*types.RecommendationResponse, error) {
	products, err := s.repo.FetchProductsByPreferences(
		req.UserPreferences.Categories,
		req.UserPreferences.ExcludeCategories,
		req.UserPreferences.PriceRange.MinPoints,
		req.UserPreferences.PriceRange.MaxPoints,
		req.Limit,
	)
	if err != nil {
		return nil, err
	}

	var recommendations []types.RecommendedProduct
	for _, p := range products {
		var confidence = 0.85
		if parsedID, err := strconv.Atoi(p.ID); err == nil {
			confidence = 0.8 + 0.1*float64(parsedID%2)
		}

		recommendations = append(recommendations, types.RecommendedProduct{
			Product: struct {
				ID               string         `json:"id"`
				Name             string         `json:"name"`
				Description      string         `json:"description"`
				Category         store.Category `json:"category"`
				RedemptionPoints int            `json:"redemption_points"`
				StockQuantity    int            `json:"stock_quantity"`
				IsOffer          bool           `json:"is_offer"`
			}{
				ID:               p.ID,
				Name:             p.Name,
				Description:      p.Description,
				Category:         store.Category{ID: p.Category.ID, Name: p.Category.Name},
				RedemptionPoints: p.RedemptionPoints,
				StockQuantity:    p.StockQuantity,
				IsOffer:          p.IsOffer,
			},
			ConfidenceScore: confidence,
			Reason:          "Matched category and price range",
		})
	}

	return &types.RecommendationResponse{
		Recommendations: recommendations,
		RecommendationMeta: struct {
			AlgorithmVersion string `json:"algorithm_version"`
			GeneratedAt      string `json:"generated_at"`
			UserSegment      string `json:"user_segment"`
		}{
			AlgorithmVersion: "v1.0-mock",
			GeneratedAt:      time.Now().Format(time.RFC3339),
			UserSegment:      "general",
		},
	}, nil
}
