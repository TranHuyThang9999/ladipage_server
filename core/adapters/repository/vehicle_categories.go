package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"ladipage_server/core/adapters"
	"ladipage_server/core/domain"
)

type vehicleCategory struct {
	db *adapters.Pgsql
}

func NewVehicleCategoryRepository(db *adapters.Pgsql) domain.RepositoryVehicleCategory {
	return &vehicleCategory{
		db: db,
	}
}

func (v *vehicleCategory) AddVehicleCategory(ctx context.Context, vehicleCategory *domain.VehicleCategory) error {
	result := v.db.DB().Create(vehicleCategory)
	return result.Error
}

func (v *vehicleCategory) ListVehicleCategories(ctx context.Context) ([]*domain.VehicleCategory, error) {
	var vehicleCategories []*domain.VehicleCategory
	result := v.db.DB().Find(&vehicleCategories)
	return vehicleCategories, result.Error
}

func (v *vehicleCategory) DeleteVehicleCategoryByID(ctx context.Context, id int64) error {
	result := v.db.DB().WithContext(ctx).Delete(&domain.VehicleCategory{}, id)
	return result.Error
}

func (v *vehicleCategory) UpdateVehicleCategoryByID(ctx context.Context, vehicleCategory *domain.VehicleCategory) error {
	result := v.db.DB().WithContext(ctx).Save(vehicleCategory)
	return result.Error
}
func (v *vehicleCategory) GetVehicleCategoryByName(ctx context.Context, name string) (*domain.VehicleCategory, error) {
	var vehicleCategory domain.VehicleCategory
	err := v.db.DB().WithContext(ctx).Where("name = ?", name).First(&vehicleCategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &vehicleCategory, err
}
