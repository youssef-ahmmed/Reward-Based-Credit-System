package store

import "time"

type Redemption struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Status    string    `json:"status"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`

	Product Product `gorm:"foreignKey:ProductID" json:"product"`
}
