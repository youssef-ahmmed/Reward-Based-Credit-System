package store

import "time"

type Purchase struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	UserID          string    `json:"user_id"`
	CreditPackageID string    `json:"credit_package_id"`
	Status          string    `json:"status"`
	Credits         int       `json:"credits"`
	CreatedAt       time.Time `json:"created_at"`

	CreditPackage CreditPackage `gorm:"foreignKey:CreditPackageID"`
}
