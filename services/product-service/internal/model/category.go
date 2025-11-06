package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name          string     `gorm:"not null" json:"name"`
	Description   string     `gorm:"not null" json:"description"`
	ParentID      *uuid.UUID `gorm:"type:uuid" json:"parent_id"` // Nullable → for top-level categories
	Parent        *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Subcategories []Category `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subcategories,omitempty"`
	CreatedBy     uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}
