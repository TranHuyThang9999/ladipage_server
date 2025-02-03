package services

import (
	"context"
	"errors"
	"ladipage_server/apis/entities"
	"ladipage_server/common/logger"
	"ladipage_server/common/utils"
	"ladipage_server/core/constant"
	customerrors "ladipage_server/core/custom_errors"
	"ladipage_server/core/domain"
	"strings"

	"gorm.io/gorm"
)

type VehicleService struct {
	vehicle         domain.RepositoryVehicle
	file            domain.RepositoryFileDescriptors
	logger          *logger.Logger
	trans           domain.RepositoryTransactionHelper
	vehicleCategory domain.RepositoryVehicleCategory
}

func NewVehicleService(
	vehicle domain.RepositoryVehicle,
	file domain.RepositoryFileDescriptors,
	trans domain.RepositoryTransactionHelper,
	vehicleCategory domain.RepositoryVehicleCategory,
	logger *logger.Logger) *VehicleService {
	return &VehicleService{
		vehicle:         vehicle,
		logger:          logger,
		file:            file,
		trans:           trans,
		vehicleCategory: vehicleCategory,
	}
}

func (u *VehicleService) Add(ctx context.Context, req *entities.CreateVehicleRequest) *customerrors.CustomError {
	vehicleID := utils.GenUUID()
	carName := strings.TrimSpace(req.ModelName)
	typeVehicleCategory, err := u.vehicleCategory.GetVehicleCategoryByID(ctx, req.VehicleCategoryID)
	if err != nil {
		u.logger.Error("VehicleService - AddVehicle", err)
		return customerrors.ErrDB
	}

	if typeVehicleCategory == nil {
		u.logger.Warn("VehicleService - AddVehicle", errors.New("vehicle category is nil"))
		return customerrors.ErrNotFound
	}

	count, err := u.vehicle.CheckDuplicateVehicle(ctx, req.VehicleCategoryID, carName)
	if err != nil {
		u.logger.Error("VehicleService - AddVehicle", err)
		return customerrors.ErrDB
	}
	if count > 0 {
		u.logger.Warn("VehicleService - AddVehicle", errors.New("vehicle already exists"))
		return customerrors.ErrCategoryExists
	}

	if err := u.trans.Transaction(ctx, func(ctx context.Context, db *gorm.DB) error {
		model := &domain.Vehicle{
			Model: entities.Model{
				ID: vehicleID,
			},
			VehicleCategoryID: req.VehicleCategoryID,
			ModelName:         carName,
			Variant:           req.Variant,
			VersionYear:       req.VersionYear,
			BasePrice:         req.BasePrice,
			PromotionalPrice:  req.PromotionalPrice,
			Color:             req.Color,
			Transmission:      req.Transmission,
			Engine:            req.Engine,
			FuelType:          req.FuelType,
			Seating:           req.Seating,
			Status:            req.Status,
			Featured:          false,
			Note:              req.Note,
		}
		err := u.vehicle.AddVehicle(ctx, db, model)
		if err != nil {
			u.logger.Error("error adding vehicle", err)
			return err
		}

		var listFileDesc = make([]*domain.FileDescriptors, 0)
		if len(req.Urls) > 0 {
			for _, v := range req.Urls {
				listFileDesc = append(listFileDesc, &domain.FileDescriptors{
					Model: &entities.Model{
						ID: utils.GenUUID(),
					},
					ObjectID:   vehicleID,
					Url:        *v,
					TypeObject: constant.TypeObjectVehicle,
				})
			}
			err = u.file.AddListFileWithTransaction(ctx, db, listFileDesc)
			if err != nil {
				u.logger.Error("error adding vehicle", err)
				return err
			}
		}

		return nil
	}); err != nil {
		return customerrors.ErrDB
	}
	return nil
}

func (u *VehicleService) ListFileByObjectID(ctx context.Context, objectID int64) ([]*entities.ListFileByObjectID,
	*customerrors.CustomError) {
	var listFileByObjectID = make([]*entities.ListFileByObjectID, 0)
	//checkVehicle, err := u.vehicle.GetVehicleCategoryByID(ctx, objectID)
	//if err != nil {
	//	u.logger.Error("error database", err)
	//	return nil, customerrors.ErrDB
	//}
	//if checkVehicle == nil {
	//	u.logger.Warn("Vehicle category not found")
	//	return nil, customerrors.ErrNotFound
	//}
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

func (u *VehicleService) ListVehicle(ctx context.Context) ([]*entities.GetCreateVehicles, *customerrors.CustomError) {
	resp, err := u.vehicle.ListVehicles(ctx)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	if len(resp) == 0 {
		return []*entities.GetCreateVehicles{}, nil
	}

	categoryIDs := make(map[int64]bool)
	for _, v := range resp {
		categoryIDs[v.VehicleCategoryID] = true
	}

	categoryIDList := make([]int64, 0, len(categoryIDs))
	for id := range categoryIDs {
		categoryIDList = append(categoryIDList, id)
	}

	vehicleCategories, err := u.vehicleCategory.GetVehicleCategoriesByIDs(ctx, categoryIDList)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	categoryMap := make(map[int64]*domain.VehicleCategory)
	for _, category := range vehicleCategories {
		categoryMap[category.ID] = category
	}

	listVehicle := make([]*entities.GetCreateVehicles, 0, len(resp))
	for _, v := range resp {
		vehicleCategory, exists := categoryMap[v.VehicleCategoryID]
		if exists {
			listVehicle = append(listVehicle, &entities.GetCreateVehicles{
				ID:                v.ID,
				VehicleCategory:   vehicleCategory.Name,
				ModelName:         v.ModelName,
				Variant:           v.Variant,
				VersionYear:       v.VersionYear,
				BasePrice:         v.BasePrice,
				PromotionalPrice:  v.PromotionalPrice,
				Color:             v.Color,
				Transmission:      v.Transmission,
				Engine:            v.Engine,
				FuelType:          v.FuelType,
				Seating:           v.Seating,
				Status:            v.Status,
				Featured:          false,
				Note:              v.Note,
				HCMRollingPrice:   v.HCMRollingPrice,
				HanoiRollingPrice: v.HanoiRollingPrice,
				InstallmentFrom:   v.InstallmentFrom,
			})
		}
	}

	return listVehicle, nil
}

func (u *VehicleService) UpdateVehicleByID(ctx context.Context, req *entities.UpdateCreateVehicles) *customerrors.CustomError {
	err := u.vehicle.UpdateVehicleByID(ctx, &domain.Vehicle{
		Model: entities.Model{
			ID: req.ID,
		},
		VehicleCategoryID: req.VehicleCategoryID,
		ModelName:         req.ModelName,
		Variant:           req.Variant,
		VersionYear:       req.VersionYear,
		BasePrice:         req.BasePrice,
		PromotionalPrice:  req.PromotionalPrice,
		Color:             req.Color,
		Transmission:      req.Transmission,
		Engine:            req.Engine,
		FuelType:          req.FuelType,
		Seating:           req.Seating,
		Status:            req.Status,
		Featured:          true,
		Note:              req.Note,
	})
	if err != nil {
		u.logger.Error("error database", err)
		return customerrors.ErrDB
	}

	return nil
}

func (u *VehicleService) DeleteVehicleByID(ctx context.Context, id int64) *customerrors.CustomError {
	err := u.vehicle.DeleteVehicleByID(ctx, id)
	if err != nil {
		u.logger.Error("error database", err)
		return customerrors.ErrDB
	}
	return nil
}

func (u *VehicleService) AddListFileByObjectID(ctx context.Context, req *entities.CreateFilesRequest) *customerrors.CustomError {
	var listFileAdd []*domain.FileDescriptors
	//checkVehicle, err := u.vehicle.GetVehicleCategoryByID(ctx, req.ObjectID)
	//if err != nil {
	//	u.logger.Error("error database", err)
	//	return customerrors.ErrDB
	//}
	//if checkVehicle == nil {
	//	u.logger.Warn("Vehicle category not found")
	//	return customerrors.ErrNotFound
	//}
	for _, url := range req.Urls {
		listFileAdd = append(listFileAdd, &domain.FileDescriptors{
			Model: &entities.Model{
				ID: utils.GenUUID(),
			},
			CreatorID:  req.CreatorID,
			ObjectID:   req.ObjectID,
			Url:        *url,
			TypeObject: constant.TypeObjectVehicle,
		})
	}

	err := u.file.AddListFileWith(ctx, listFileAdd)
	if err != nil {
		u.logger.Error("error add list file", err)
		return customerrors.ErrDB
	}

	return nil
}

func (u *VehicleService) DeleteListFileByID(ctx context.Context, req *entities.DeleteFilesRequest) *customerrors.CustomError {
	//checkVehicle, err := u.vehicle.GetVehicleCategoryByID(ctx, req.ObjectID)
	//if err != nil {
	//	u.logger.Error("error database", err)
	//	return customerrors.ErrDB
	//}
	//if checkVehicle == nil {
	//	u.logger.Warn("Vehicle category not found")
	//	return customerrors.ErrNotFound
	//}

	err := u.file.DeleteListFileByObjectID(ctx, req.IDs)
	if err != nil {
		u.logger.Error("error delete list file", err)
		return customerrors.ErrDB
	}

	return nil
}

func (u *VehicleService) GetVehiclesByVehicleCategoryID(ctx context.Context, vehicleCategoryID int64) ([]*entities.GetCreateVehicles, *customerrors.CustomError) {
	resp, err := u.vehicle.GetVehiclesByVehicleCategoryID(ctx, vehicleCategoryID)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	if len(resp) == 0 {
		return []*entities.GetCreateVehicles{}, nil
	}
	if len(resp) == 0 {
		return []*entities.GetCreateVehicles{}, nil
	}

	categoryIDs := make(map[int64]bool)
	for _, v := range resp {
		categoryIDs[v.VehicleCategoryID] = true
	}

	categoryIDList := make([]int64, 0, len(categoryIDs))
	for id := range categoryIDs {
		categoryIDList = append(categoryIDList, id)
	}

	vehicleCategories, err := u.vehicleCategory.GetVehicleCategoriesByIDs(ctx, categoryIDList)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	categoryMap := make(map[int64]*domain.VehicleCategory)
	for _, category := range vehicleCategories {
		categoryMap[category.ID] = category
	}

	listVehicle := make([]*entities.GetCreateVehicles, 0, len(resp))
	for _, v := range resp {
		vehicleCategory, exists := categoryMap[v.VehicleCategoryID]
		if exists {
			listVehicle = append(listVehicle, &entities.GetCreateVehicles{
				ID:                v.ID,
				VehicleCategory:   vehicleCategory.Name,
				ModelName:         v.ModelName,
				Variant:           v.Variant,
				VersionYear:       v.VersionYear,
				BasePrice:         v.BasePrice,
				PromotionalPrice:  v.PromotionalPrice,
				Color:             v.Color,
				Transmission:      v.Transmission,
				Engine:            v.Engine,
				FuelType:          v.FuelType,
				Seating:           v.Seating,
				Status:            v.Status,
				Featured:          false,
				Note:              v.Note,
				HCMRollingPrice:   v.HCMRollingPrice,
				HanoiRollingPrice: v.HanoiRollingPrice,
				InstallmentFrom:   v.InstallmentFrom,
			})
		}
	}

	return listVehicle, nil
}

func (u *VehicleService) GetVehiclesByVehicleCategoryIDForPublic(ctx context.Context, vehicleCategoryID int64) ([]*entities.GetCreateVehiclesForPublic, *customerrors.CustomError) {
	// Lấy danh sách vehicles
	vehicles, err := u.vehicle.GetVehiclesByVehicleCategoryID(ctx, vehicleCategoryID)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	if len(vehicles) == 0 {
		return []*entities.GetCreateVehiclesForPublic{}, nil
	}

	// Lấy danh sách vehicle IDs
	vehicleIDs := make([]int64, 0, len(vehicles))
	for _, v := range vehicles {
		vehicleIDs = append(vehicleIDs, v.ID)
	}

	// Lấy tất cả files cho các vehicles một lần
	allFiles, err := u.file.ListByObjectIDs(ctx, vehicleIDs) // Giả sử có method này
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	// Tạo map files theo vehicle ID
	fileMap := make(map[int64][]*string)
	for _, file := range allFiles {
		if _, exists := fileMap[file.ObjectID]; !exists {
			fileMap[file.ObjectID] = make([]*string, 0)
		}
		fileMap[file.ObjectID] = append(fileMap[file.ObjectID], &file.Url)
	}

	// Xử lý categories
	categoryIDs := make(map[int64]bool)
	for _, v := range vehicles {
		categoryIDs[v.VehicleCategoryID] = true
	}

	categoryIDList := make([]int64, 0, len(categoryIDs))
	for id := range categoryIDs {
		categoryIDList = append(categoryIDList, id)
	}

	vehicleCategories, err := u.vehicleCategory.GetVehicleCategoriesByIDs(ctx, categoryIDList)
	if err != nil {
		u.logger.Error("error database", err)
		return nil, customerrors.ErrDB
	}

	categoryMap := make(map[int64]*domain.VehicleCategory)
	for _, category := range vehicleCategories {
		categoryMap[category.ID] = category
	}

	// Tạo kết quả cuối cùng
	listVehicle := make([]*entities.GetCreateVehiclesForPublic, 0, len(vehicles))
	for _, v := range vehicles {
		vehicleCategory, exists := categoryMap[v.VehicleCategoryID]
		if exists {
			listVehicle = append(listVehicle, &entities.GetCreateVehiclesForPublic{
				ID:                v.ID,
				VehicleCategory:   vehicleCategory.Name,
				ModelName:         v.ModelName,
				Variant:           v.Variant,
				VersionYear:       v.VersionYear,
				BasePrice:         v.BasePrice,
				PromotionalPrice:  v.PromotionalPrice,
				Color:             v.Color,
				Transmission:      v.Transmission,
				Engine:            v.Engine,
				FuelType:          v.FuelType,
				Seating:           v.Seating,
				Status:            v.Status,
				Featured:          false,
				Note:              v.Note,
				HCMRollingPrice:   v.HCMRollingPrice,
				HanoiRollingPrice: v.HanoiRollingPrice,
				InstallmentFrom:   v.InstallmentFrom,
				Urls:              fileMap[v.ID], // Lấy URLs từ map
			})
		}
	}

	return listVehicle, nil
}
