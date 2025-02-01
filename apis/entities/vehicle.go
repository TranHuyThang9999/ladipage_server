package entities

type CreateVehicleRequest struct {
	VehicleCategoryID int64  `json:"vehicle_category_id" binding:"required"` // Khóa ngoại, bắt buộc
	ModelName         string `json:"model_name,omitempty"`                   // Xforce, Xpander...
	Variant           string `json:"variant,omitempty"`                      // Ultimate, Premium, Exceed...
	VersionYear       int    `json:"version_year"`                           // 2024
	BasePrice         string `json:"base_price,omitempty"`                   // 705000000
	PromotionalPrice  string `json:"promotional_price,omitempty"`            // Giá khuyến mãi nếu có

	// Thông số kỹ thuật
	Color        string `json:"color,omitempty"`        // Màu xe
	Transmission string `json:"transmission,omitempty"` // Số tự động/số sàn
	Engine       string `json:"engine,omitempty"`       // Động cơ
	FuelType     string `json:"fuel_type,omitempty"`    // Loại nhiên liệu
	Seating      int    `json:"seating,omitempty"`      // Số chỗ ngồi

	// Trạng thái
	Status   string    `json:"status"`             // available, discontinued... (Bắt buộc)
	Featured bool      `json:"featured,omitempty"` // Xe nổi bật
	Note     string    `json:"note,omitempty"`     // Ghi chú
	Urls     []*string `json:"urls,omitempty"`
}
type GetCreateVehicles struct {
	ID               int64  `json:"id" binding:"required"`
	VehicleCategory  string `json:"vehicle_category" `
	ModelName        string `json:"model_name,omitempty"`        // Xforce, Xpander...
	Variant          string `json:"variant,omitempty"`           // Ultimate, Premium, Exceed...
	VersionYear      int    `json:"version_year"`                // 2024
	BasePrice        string `json:"base_price,omitempty"`        // 705000000
	PromotionalPrice string `json:"promotional_price,omitempty"` // Giá khuyến mãi nếu có

	// Thông số kỹ thuật
	Color        string `json:"color,omitempty"`        // Màu xe
	Transmission string `json:"transmission,omitempty"` // Số tự động/số sàn
	Engine       string `json:"engine,omitempty"`       // Động cơ
	FuelType     string `json:"fuel_type,omitempty"`    // Loại nhiên liệu
	Seating      int    `json:"seating,omitempty"`      // Số chỗ ngồi

	// Trạng thái
	Status   string `json:"status"`             // available, discontinued... (Bắt buộc)
	Featured bool   `json:"featured,omitempty"` // Xe nổi bật
	Note     string `json:"note,omitempty"`     // Ghi chú
}

type UpdateCreateVehicles struct {
	ID                int64  `json:"id" binding:"required"`
	VehicleCategoryID int64  `json:"vehicle_category_id" `        // Khóa ngoại, bắt buộc
	ModelName         string `json:"model_name,omitempty"`        // Xforce, Xpander...
	Variant           string `json:"variant,omitempty"`           // Ultimate, Premium, Exceed...
	VersionYear       int    `json:"version_year"`                // 2024
	BasePrice         string `json:"base_price,omitempty"`        // 705000000
	PromotionalPrice  string `json:"promotional_price,omitempty"` // Giá khuyến mãi nếu có

	// Thông số kỹ thuật
	Color        string `json:"color,omitempty"`        // Màu xe
	Transmission string `json:"transmission,omitempty"` // Số tự động/số sàn
	Engine       string `json:"engine,omitempty"`       // Động cơ
	FuelType     string `json:"fuel_type,omitempty"`    // Loại nhiên liệu
	Seating      int    `json:"seating,omitempty"`      // Số chỗ ngồi

	// Trạng thái
	Status   string `json:"status"`             // available, discontinued... (Bắt buộc)
	Featured bool   `json:"featured,omitempty"` // Xe nổi bật
	Note     string `json:"note,omitempty"`     // Ghi chú
}
