package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModelID struct {
	ID        uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time       `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
