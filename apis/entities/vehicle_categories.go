package entities

import "time"

type CreateVehicleCategoriesRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatorID int64  `json:"-"`
}

type UpdateVehicleCategoriesRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type ListVehicleCategories struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
