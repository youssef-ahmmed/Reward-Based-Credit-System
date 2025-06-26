package product

import "time"

type Category struct {
	ID               string  `json:"id"`
	ParentCategoryID *string `json:"parent_category_id,omitempty"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
}

type Product struct {
	ID               string    `json:"id"`
	CategoryID       string    `json:"category_id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	RedemptionPoints int       `json:"redemption_points"`
	StockQuantity    int       `json:"stock_quantity"`
	IsOffer          bool      `json:"is_offer"`
	CreatedAt        time.Time `json:"created_at"`
	ImageURL         string    `json:"image_url"`
	Tags             []string  `json:"tags"`
}

type Redemption struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
