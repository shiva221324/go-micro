package migrations

import (
	"auth-service/internal/model"
	"log"

	"gorm.io/gorm"
)

// Run performs database migrations
func Run(db *gorm.DB) {
	log.Println("🚀 Running migrations...")

	// AutoMigrate will create or update tables based on model structs
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migrations completed successfully")
}
