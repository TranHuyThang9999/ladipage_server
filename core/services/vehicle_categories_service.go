package services

import (
	"context"
	"ladipage_server/apis/entities"
	"ladipage_server/common/logger"
	"ladipage_server/common/utils"
	customerrors "ladipage_server/core/custom_errors"
	"ladipage_server/core/domain"
	"strings"
)

type VehicleCategoriesService struct {
	vehicle domain.RepositoryVehicleCategory
	logger  *logger.Logger
}

func NewVehicleCategoriesService(vehicle domain.RepositoryVehicleCategory,
	logger *logger.Logger) *VehicleCategoriesService {
	return &VehicleCategoriesService{
		vehicle: vehicle,
		logger:  logger,
	}
}
func (svc *VehicleCategoriesService) Add(ctx context.Context, req *entities.CreateVehicleCategoriesRequest) *customerrors.CustomError {
	nameVehicle := strings.TrimSpace(req.Name)
	vehicle, err := svc.vehicle.GetVehicleCategoryByName(ctx, nameVehicle)
	if err != nil {
		svc.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	if vehicle != nil {
		svc.logger.Warn("Vehicle category already exists")
		return customerrors.ErrCategoryExists
	}
	model := &domain.VehicleCategory{
		Model: entities.Model{
			ID: utils.GenUUID(),
		},
		CreatorID: req.CreatorID,
		Name:      nameVehicle,
	}
	err = svc.vehicle.AddVehicleCategory(ctx, model)
	if err != nil {
		svc.logger.Error("AddVehicleCategory Failed", err)
		return customerrors.ErrDB
	}
	return nil
}

func (svc *VehicleCategoriesService) FindAll(ctx context.Context) ([]*domain.VehicleCategory, *customerrors.CustomError) {
	vehicles, err := svc.vehicle.ListVehicleCategories(ctx)
	if err != nil {
		svc.logger.Error("ListVehicleCategories Failed", err)
		return nil, customerrors.ErrDB
	}
	return vehicles, nil
}
