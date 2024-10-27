package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID        int32 `gorm:"primaryKey"`
	Name      string
	Email     string
	Phone     string
	Address   string
	Orders    []Order `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name        string
	Price       string
	Description string
	Quantity    int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Order struct {
	gorm.Model
	ID            uuid.UUID `gorm:"primaryKey;type:uuid"`
	CustomerID    int32
	PaymentStatus string
	Products      []Product `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Status        string
	TotalPrice    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Webhook struct {
	// This is a placeholder for the webhook
	OrderID       uuid.UUID
	PaymentStatus string
}
