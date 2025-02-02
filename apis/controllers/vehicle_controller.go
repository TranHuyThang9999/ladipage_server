package controllers

import (
	"github.com/gin-gonic/gin"
	"ladipage_server/apis/entities"
	"ladipage_server/apis/resources"
	"ladipage_server/core/services"
)

type VehicleController struct {
	vehicle *services.VehicleService
	base    *baseController
	reso    *resources.Resource
}

func NewVehicleController(
	vehicle *services.VehicleService,
	base *baseController,
	reso *resources.Resource,
) *VehicleController {
	return &VehicleController{
		vehicle: vehicle,
		base:    base,
		reso:    reso,
	}
}

func (u *VehicleController) AddVehicle(ctx *gin.Context) {
	var req entities.CreateVehicleRequest
	if !u.base.Bind(ctx, &req) {
		return
	}

	err := u.vehicle.Add(ctx, &req)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.CreatedSuccess(ctx)
}

func (u *VehicleController) DeleteVehicle(ctx *gin.Context) {
	panic("implement me")
}

func (u *VehicleController) GetVehicles(ctx *gin.Context) {
	resp, err := u.vehicle.ListVehicle(ctx)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}
	u.reso.Response(ctx, resp)
}

func (u *VehicleController) GetListFileVehicleById(ctx *gin.Context) {
	objectID, ok := u.base.GetParamTypeNumber(ctx, "objectID")
	if !ok {
		return
	}
	listFile, err := u.vehicle.ListFileByObjectID(ctx, objectID)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.Response(ctx, listFile)
}

func (u *VehicleController) UpdateVehicleById(ctx *gin.Context) {
	var req entities.UpdateCreateVehicles
	if !u.base.Bind(ctx, &req) {
		return
	}

	err := u.vehicle.UpdateVehicleByID(ctx, &req)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.UpdatedSuccess(ctx)
}

func (u *VehicleController) DeleteVehicleById(ctx *gin.Context) {
	id, ok := u.base.GetParamTypeNumber(ctx, "id")
	if !ok {
		return
	}
	err := u.vehicle.DeleteVehicleByID(ctx, id)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.DeletedSuccess(ctx)
}

func (u *VehicleController) AddListFileByObjectID(ctx *gin.Context) {
	var req entities.CreateFilesRequest
	if !u.base.Bind(ctx, &req) {
		return
	}
	userID, ok := u.base.GetUserID(ctx)
	if !ok {
		return
	}
	req.CreatorID = userID
	err := u.vehicle.AddListFileByObjectID(ctx, &req)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.CreatedSuccess(ctx)
}

func (u *VehicleController) DeleteListFileByID(ctx *gin.Context) {
	var req entities.DeleteFilesRequest
	if !u.base.Bind(ctx, &req) {
		return
	}
	err := u.vehicle.DeleteListFileByID(ctx, &req)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.DeletedSuccess(ctx)
}

func (u *VehicleController) GetVehiclesByVehicleCategoryID(ctx *gin.Context) {
	vehicleCategoryID, ok := u.base.GetParamTypeNumber(ctx, "vehicleCategoryID")
	if !ok {
		return
	}
	resp, err := u.vehicle.GetVehiclesByVehicleCategoryID(ctx, vehicleCategoryID)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}
	u.reso.Response(ctx, resp)
}
