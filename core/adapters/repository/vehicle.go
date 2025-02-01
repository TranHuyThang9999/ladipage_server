package repository

import (
	"context"
	"ladipage_server/core/adapters"
	"ladipage_server/core/domain"

	"gorm.io/gorm"
)

type vehicleRepository struct {
	db *adapters.Pgsql
}

func NewVehicleRepository(db *adapters.Pgsql) domain.RepositoryVehicle {
	return &vehicleRepository{
		db: db,
	}
}

// AddVehicle implements domain.RepositoryVehicle.
func (v *vehicleRepository) AddVehicle(ctx context.Context, db *gorm.DB, vehicle *domain.Vehicle) error {
	result := db.WithContext(ctx).Create(vehicle)
	return result.Error
}

// CheckVehicleExists implements domain.RepositoryVehicle.
func (v *vehicleRepository) CheckVehicleExists(ctx context.Context, id int64) (int64, error) {
	panic("unimplemented")
}

// DeleteVehicleByID implements domain.RepositoryVehicle.
func (v *vehicleRepository) DeleteVehicleByID(ctx context.Context, id int64) error {
	result := v.db.DB().WithContext(ctx).Delete(&domain.Vehicle{}, id)
	return result.Error
}

// GetVehicleByID implements domain.RepositoryVehicle.
func (v *vehicleRepository) GetVehicleByID(ctx context.Context, id int64) (*domain.Vehicle, error) {
	panic("unimplemented")
}

// GetVehicleByModelName implements domain.RepositoryVehicle.
func (v *vehicleRepository) GetVehicleByModelName(ctx context.Context, modelName string) (*domain.Vehicle, error) {
	panic("unimplemented")
}

// ListVehicles implements domain.RepositoryVehicle.
func (v *vehicleRepository) ListVehicles(ctx context.Context) ([]*domain.Vehicle, error) {
	var vehicles = make([]*domain.Vehicle, 0)
	result := v.db.DB().WithContext(ctx).Find(&vehicles)
	return vehicles, result.Error
}

// UpdateVehicleByID implements domain.RepositoryVehicle.
func (v *vehicleRepository) UpdateVehicleByID(ctx context.Context, vehicle *domain.Vehicle) error {
	result := v.db.DB().WithContext(ctx).Save(vehicle)
	return result.Error
}

func (v *vehicleRepository) CheckDuplicateVehicle(ctx context.Context, vehicleCategoryID int64, modelName string) (int64, error) {
	var count int64
	result := v.db.DB().WithContext(ctx).
		Model(&domain.Vehicle{}).
		Where("vehicle_category_id = ? AND model_name = ?", vehicleCategoryID, modelName).
		Count(&count)
	return count, result.Error
}
