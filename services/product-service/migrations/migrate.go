package migrations

import (
	"log"
	"product-service/internal/model"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) {
	err := db.AutoMigrate(&model.Category{}, &model.Product{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Migration completed: Category & Product tables ready")
}
