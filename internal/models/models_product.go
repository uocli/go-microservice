package models

type Product struct {
	ProductID string  `gorm:"primaryKey" json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	VendorID  string  `gorm:"foreignKey" json:"VendorId"`
}
