package entities

type CreateVehicleCategoriesRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatorID int64  `json:"-"`
}
