package entities

import "time"

type CreateVehicleCategoriesRequest struct {
	Name      string   `json:"name" binding:"required"`
	Urls      []string `json:"urls,omitempty"`
	CreatorID int64    `json:"-"`
}

type UpdateVehicleCategoriesRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type VehicleCategories struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ListVehicleCategories struct {
	Total             int                  `json:"total,omitempty"`
	VehicleCategories []*VehicleCategories `json:"vehicle_categories,omitempty"`
}
