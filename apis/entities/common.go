package entities

import (
	"time"
)

type Model struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at;default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index;column:deleted_at"`
}
