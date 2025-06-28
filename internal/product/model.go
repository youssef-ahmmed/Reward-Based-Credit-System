package product

import (
	"gorm.io/datatypes"
	"time"
)

type Category struct {
	ID               string     `gorm:"primaryKey" json:"id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	ParentCategoryID *string    `json:"parent_category_id"`
	Children         []Category `gorm:"foreignKey:ParentCategoryID" json:"children,omitempty"`
}

type Product struct {
	ID               string         `gorm:"primaryKey" json:"id"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	CategoryID       string         `json:"category_id"`
	Category         Category       `gorm:"foreignKey:CategoryID" json:"category"`
	RedemptionPoints int            `json:"redemption_points"`
	StockQuantity    int            `json:"stock_quantity"`
	IsOffer          bool           `json:"is_offer"`
	CreatedAt        time.Time      `json:"created_at"`
	ImageURL         string         `json:"image_url"`
	Tags             datatypes.JSON `json:"tags"`
}

type Redemption struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
