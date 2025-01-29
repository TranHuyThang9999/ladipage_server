package entities

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ID        int64          `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}
