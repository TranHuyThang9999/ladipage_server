package repository

import (
	"context"
	"errors"
	"fmt"
	"ladipage_server/core/adapters"
	"ladipage_server/core/domain"

	"gorm.io/gorm"
)

type vehicleCategory struct {
	db *adapters.Pgsql
}

func NewVehicleCategoryRepository(db *adapters.Pgsql) domain.RepositoryVehicleCategory {
	return &vehicleCategory{
		db: db,
	}
}

func (v *vehicleCategory) AddVehicleCategory(ctx context.Context, db *gorm.DB, vehicleCategory *domain.VehicleCategory) error {
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
	result := v.db.DB().WithContext(ctx).Where("id = ?", vehicleCategory.ID).Updates(vehicleCategory)
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

func (v *vehicleCategory) ExistsByName(ctx context.Context, id int64, newName string) (int64, error) {
	var existingID int64

	result := v.db.DB().WithContext(ctx).
		Model(&domain.VehicleCategory{}).
		Select("id").
		Where("LOWER(name) = LOWER(?) AND id != ?", newName, id).
		First(&existingID)

	if result.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	if result.Error != nil {
		return 0, fmt.Errorf("failed to check vehicle category name: %w", result.Error)
	}

	return existingID, nil
}

func (v *vehicleCategory) GetVehicleCategoryByID(ctx context.Context, id int64) (*domain.VehicleCategory, error) {
	var vehicleCategory domain.VehicleCategory
	err := v.db.DB().WithContext(ctx).Where("id = ?", id).First(&vehicleCategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &vehicleCategory, err
}
