package store

import (
	"gorm.io/datatypes"
	"time"
)

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
