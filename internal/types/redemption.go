package types

type CreateRedemptionRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type RedemptionResponse struct {
	ID         string            `json:"id"`
	Product    RedemptionProduct `json:"product"`
	Quantity   int               `json:"quantity"`
	PointsUsed int               `json:"points_used"`
	CreatedAt  string            `json:"created_at"`
}

type RedemptionProduct struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	RewardPoints int    `json:"reward_points"`
}
