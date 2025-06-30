package migration

import (
	"Start/internal/store"
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running auto-migrations...")

	err := db.AutoMigrate(
		&store.User{},
		&store.Wallet{},
		&store.Product{},
		&store.Category{},
		&store.CreditPackage{},
		&store.Purchase{},
		&store.Redemption{},
	)
	if err != nil {
		log.Printf("Migration failed: %v", err)
		return err
	}

	log.Println("Auto-migration completed successfully.")
	return nil
}
