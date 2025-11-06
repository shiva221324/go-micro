package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Price       float64    `gorm:"not null"`
	Stock       int        `gorm:"not null"`
	CategoryID  *uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Category    *Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category,omitempty"`
	CreatedBy   uuid.UUID  `gorm:"not null"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return nil
}
