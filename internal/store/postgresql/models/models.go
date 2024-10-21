package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Id      int32 `gorm:"primaryKey"`
	Name    string
	Email   string
	Phone   string
	Address string
	Orders  []Order
}

type Product struct {
	gorm.Model
	Id       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name     string
	Price    string
	Quantity int32
}

type Order struct {
	gorm.Model
	Id            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CustomerId    int32
	PaymentStatus string
	Products      []Product
	Status        string
	TotalPrice    string
	CreatedAt     string
	UpdatedAt     string
}

type Webhook struct {
	// This is a placeholder for the webhook
	OrderId       uuid.UUID
	PaymentStatus string
}
