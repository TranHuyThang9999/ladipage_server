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
