package controllers

import (
	"ladipage_server/apis/entities"
	"ladipage_server/apis/resources"
	"ladipage_server/core/services"

	"github.com/gin-gonic/gin"
)

type VehicleCategoriesController struct {
	vehicle *services.VehicleCategoriesService
	base    *baseController
	reso    *resources.Resource
}

func NewVehicleCategoriesController(
	vehicle *services.VehicleCategoriesService,
	base *baseController,
	reso *resources.Resource,
) *VehicleCategoriesController {
	return &VehicleCategoriesController{
		vehicle: vehicle,
		base:    base,
		reso:    reso,
	}
}

func (u *VehicleCategoriesController) AddVehicle(ctx *gin.Context) {
	var req entities.CreateVehicleCategoriesRequest
	if !u.base.Bind(ctx, &req) {
		return
	}
	userID, ok := u.base.GetUserID(ctx)
	if !ok {
		return
	}
	req.CreatorID = userID
	err := u.vehicle.Add(ctx, &req)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}

	u.reso.CreatedSuccess(ctx)
}
func (u *VehicleCategoriesController) ListVehicle(ctx *gin.Context) {
	vehicles, err := u.vehicle.FindAll(ctx)
	if err != nil {
		u.reso.Error(ctx, err)
		return
	}
	u.reso.Response(ctx, vehicles)
}
