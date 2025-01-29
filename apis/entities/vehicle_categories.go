package entities

type CreateVehicleCategoriesRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatorID int64  `json:"-"`
}

type UpdateVehicleCategoriesRequest struct {
	ID   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
