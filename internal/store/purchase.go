package store

import "time"

type Purchase struct {
	ID              string `gorm:"primaryKey"`
	UserID          string
	CreditPackageID string
	Status          string
	CreatedAt       time.Time

	CreditPackage CreditPackage `gorm:"foreignKey:CreditPackageID"`
}
