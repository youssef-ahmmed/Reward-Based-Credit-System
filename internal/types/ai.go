package types

import "Start/internal/store"

type RecommendationRequest struct {
	UserPreferences struct {
		Categories        []string `json:"categories"`         // <== UUIDs as strings
		ExcludeCategories []string `json:"exclude_categories"` // <== UUIDs as strings
		PriceRange        struct {
			MinPoints int `json:"min_points"`
			MaxPoints int `json:"max_points"`
		} `json:"price_range"`
	} `json:"user_preferences"`
	Limit   int    `json:"limit"`
	Context string `json:"context"`
}

type RecommendedProduct struct {
	Product struct {
		ID               string         `json:"id"`
		Name             string         `json:"name"`
		Description      string         `json:"description"`
		Category         store.Category `json:"category"`
		RedemptionPoints int            `json:"redemption_points"`
		StockQuantity    int            `json:"stock_quantity"`
		IsOffer          bool           `json:"is_offer"`
	} `json:"product"`
	ConfidenceScore float64 `json:"confidence_score"`
	Reason          string  `json:"reason"`
}

type RecommendationResponse struct {
	Recommendations    []RecommendedProduct `json:"recommendations"`
	RecommendationMeta struct {
		AlgorithmVersion string `json:"algorithm_version"`
		GeneratedAt      string `json:"generated_at"`
		UserSegment      string `json:"user_segment"`
	} `json:"recommendation_meta"`
}
