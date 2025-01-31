package services

import (
	"context"
	"ladipage_server/apis/entities"
	"ladipage_server/common/logger"
	"ladipage_server/common/utils"
	"ladipage_server/core/constant"
	customerrors "ladipage_server/core/custom_errors"
	"ladipage_server/core/domain"
	"strings"
)

type VehicleCategoriesService struct {
	vehicle domain.RepositoryVehicleCategory
	file    domain.RepositoryFileDescriptors
	logger  *logger.Logger
}

func NewVehicleCategoriesService(vehicle domain.RepositoryVehicleCategory,
	file domain.RepositoryFileDescriptors,
	logger *logger.Logger) *VehicleCategoriesService {
	return &VehicleCategoriesService{
		vehicle: vehicle,
		logger:  logger,
		file:    file,
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

func (svc *VehicleCategoriesService) FindAll(ctx context.Context) ([]*entities.ListVehicleCategories, *customerrors.CustomError) {
	var list = make([]*entities.ListVehicleCategories, 0)

	vehicles, err := svc.vehicle.ListVehicleCategories(ctx)
	if err != nil {
		svc.logger.Error("ListVehicleCategories Failed", err)
		return nil, customerrors.ErrDB
	}
	for _, v := range vehicles {
		list = append(list, &entities.ListVehicleCategories{
			ID:        v.ID,
			Name:      v.Name,
			CreatedAt: v.CreatedAt,
		})
	}
	return list, nil
}

func (svc *VehicleCategoriesService) UpdateVehicleCategoryByID(ctx context.Context, req *entities.UpdateVehicleCategoriesRequest) *customerrors.CustomError {
	nameVehicle := strings.TrimSpace(req.Name)
	checkVehicle, err := svc.vehicle.GetVehicleCategoryByID(ctx, req.ID)
	if err != nil {
		svc.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	if checkVehicle == nil {
		svc.logger.Warn("Vehicle category not found")
		return customerrors.ErrNotFound
	}
	vehicleExists, err := svc.vehicle.ExistsByName(ctx, req.ID, nameVehicle)
	if err != nil {
		svc.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	if vehicleExists != 0 {
		svc.logger.Warn("Vehicle category exists with the same name")
		return customerrors.ErrCategoryExists
	}
	err = svc.vehicle.UpdateVehicleCategoryByID(ctx, &domain.VehicleCategory{
		Model: entities.Model{
			ID: req.ID,
		},
		Name: nameVehicle,
	})
	if err != nil {
		svc.logger.Error("UpdateVehicleCategory Failed", err)
		return customerrors.ErrDB
	}
	return nil
}

func (svc *VehicleCategoriesService) DeleteVehicleCategoryByID(ctx context.Context, id int64) *customerrors.CustomError {
	checkVehicle, err := svc.vehicle.GetVehicleCategoryByID(ctx, id)
	if err != nil {
		svc.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	if checkVehicle == nil {
		svc.logger.Warn("Vehicle category not found")
		return customerrors.ErrNotFound
	}
	err = svc.vehicle.DeleteVehicleCategoryByID(ctx, id)
	if err != nil {
		svc.logger.Error("DeleteVehicleCategory Failed", err)
		return customerrors.ErrDB
	}
	return nil
}
func (u *VehicleCategoriesService) AddListFileByObjectID(ctx context.Context, req *entities.CreateFilesRequest) *customerrors.CustomError {
	var listFileAdd []*domain.FileDescriptors
	checkVehicle, err := u.vehicle.GetVehicleCategoryByID(ctx, req.ObjectID)
	if err != nil {
		u.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	if checkVehicle == nil {
		u.logger.Warn("Vehicle category not found")
		return customerrors.ErrNotFound
	}
	for _, url := range req.Url {
		listFileAdd = append(listFileAdd, &domain.FileDescriptors{
			CreatorID:  req.CreatorID,
			ObjectID:   req.ObjectID,
			Url:        *url,
			TypeObject: constant.TypeObjectVehicleCategories,
		})
	}

	err = u.file.AddListFileWith(ctx, listFileAdd)
	if err != nil {
		u.logger.Error("error add list file", err)
		return customerrors.ErrDB
	}

	return nil
}

func (u *VehicleCategoriesService) DeleteListFileByID(ctx context.Context, req *entities.DeleteFilesRequest) *customerrors.CustomError {
	checkVehicle, err := u.vehicle.GetVehicleCategoryByID(ctx, req.ObjectID)
	if err != nil {
		u.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	if checkVehicle == nil {
		u.logger.Warn("Vehicle category not found")
		return customerrors.ErrNotFound
	}

	err = u.file.DeleteListFileByObjectID(ctx, req.IDs)
	if err != nil {
		u.logger.Error("error delete list file", err)
		return customerrors.ErrDB
	}

	return nil
}

func (u *VehicleCategoriesService) ListFileByObjectID(ctx context.Context, objectID int64) ([]*entities.ListFileByObjectID,
	*customerrors.CustomError) {
	var listFileByObjectID = make([]*entities.ListFileByObjectID, 0)
	checkVehicle, err := u.vehicle.GetVehicleCategoryByID(ctx, objectID)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}
	if checkVehicle == nil {
		u.logger.Warn("Vehicle category not found")
		return nil, customerrors.ErrNotFound
	}
	listfile, err := u.file.ListByObjectID(ctx, objectID)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}
	for _, v := range listfile {
		listFileByObjectID = append(listFileByObjectID, &entities.ListFileByObjectID{
			ID:       v.ID,
			ObjectID: v.ObjectID,
			Url:      v.Url,
		})
	}
	return listFileByObjectID, nil
}
