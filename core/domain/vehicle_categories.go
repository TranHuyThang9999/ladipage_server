package domain

import (
	"context"
	"ladipage_server/apis/entities"
)

// Loại xe
type VehicleCategory struct {
	entities.Model
	CreatorID int64  `json:"creator_id"`
	Name      string `json:"name"  gorm:"index:idx_name,unique"`
}

type RepositoryVehicleCategory interface {
	AddVehicleCategory(ctx context.Context, vehicleCategory *VehicleCategory) error
	ListVehicleCategories(ctx context.Context) ([]*VehicleCategory, error)
	DeleteVehicleCategoryByID(ctx context.Context, id int64) error
	UpdateVehicleCategoryByID(ctx context.Context, vehicleCategory *VehicleCategory) error
	GetVehicleCategoryByName(ctx context.Context, name string) (*VehicleCategory, error)
}
