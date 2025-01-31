package domain

import (
	"context"
	"ladipage_server/apis/entities"
)

type Vehicle struct {
	entities.Model
	VehicleCategoryID int64           `gorm:"not null;index"` // Khóa ngoại
	VehicleCategory   VehicleCategory `gorm:"foreignKey:VehicleCategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ModelName         string          `gorm:"type:varchar(100);not null"`  // Xforce, Xpander...
	Variant           string          `gorm:"type:varchar(50)"`            // Ultimate, Premium, Exceed...
	VersionYear       int             `gorm:"type:int"`                    // 2024
	BasePrice         string          `json:"base_price,omitempty"`        // 705000000
	PromotionalPrice  string          `json:"promotional_price,omitempty"` // Giá khuyến mãi nếu có

	// Thông số kỹ thuật
	Color        string `gorm:"type:varchar(30)"` // Màu xe
	Transmission string `gorm:"type:varchar(30)"` // Số tự động/số sàn
	Engine       string `gorm:"type:varchar(50)"` // Động cơ
	FuelType     string `gorm:"type:varchar(30)"` // Loại nhiên liệu
	Seating      int    `gorm:"type:int"`         // Số chỗ ngồi

	// Trạng thái
	Status   string `gorm:"type:varchar(20);not null"` // available, discontinued...
	Featured bool   `gorm:"default:false"`             // Xe nổi bật
	Note     string `gorm:"type:text"`
}

func (v *Vehicle) TableName() string {
	return "vehicles"
}

type RepositoryVehicle interface {
	AddVehicle(ctx context.Context, vehicle *Vehicle) error
	ListVehicles(ctx context.Context) ([]*Vehicle, error)
	DeleteVehicleByID(ctx context.Context, id int64) error
	UpdateVehicleByID(ctx context.Context, vehicle *Vehicle) error
	GetVehicleByID(ctx context.Context, id int64) (*Vehicle, error)
	GetVehicleByModelName(ctx context.Context, modelName string) (*Vehicle, error)
	CheckVehicleExists(ctx context.Context, id int64) (int64, error)
}
