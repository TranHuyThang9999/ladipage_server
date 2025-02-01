package domain

import (
	"context"
	"ladipage_server/apis/entities"

	"gorm.io/gorm"
)

// Loại xe
type VehicleCategory struct {
	entities.Model
	CreatorID int64  `json:"creator_id"`
	Name      string `json:"name"  gorm:"index:idx_name,unique"`
}

type RepositoryVehicleCategory interface {
	AddVehicleCategory(ctx context.Context, db *gorm.DB, vehicleCategory *VehicleCategory) error
	ListVehicleCategories(ctx context.Context) ([]*VehicleCategory, error)
	DeleteVehicleCategoryByID(ctx context.Context, id int64) error
	UpdateVehicleCategoryByID(ctx context.Context, vehicleCategory *VehicleCategory) error
	GetVehicleCategoryByName(ctx context.Context, name string) (*VehicleCategory, error)
	ExistsByName(ctx context.Context, id int64, newName string) (int64, error)
	GetVehicleCategoryByID(ctx context.Context, id int64) (*VehicleCategory, error)
	GetVehicleCategoriesByIDs(ctx context.Context, ids []int64) ([]*VehicleCategory, error)
}
