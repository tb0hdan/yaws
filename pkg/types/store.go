package types

import (
	"yaws/internal/store/postgresql/models"

	"github.com/google/uuid"
)

type Store interface {
	GetOrders(limit, offset int32, status, paymentStatus string) ([]models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	GetOrderById(id uuid.UUID) (models.Order, error)
	UpdateOrderStatus(order models.Order, id uuid.UUID) (models.Order, error)
	PaymentWebhook(webhook models.Webhook) error
	GetProducts(limit, offset, minQuantity int32) ([]models.Product, error)
	AddProducts([]models.Product) ([]models.Product, error)
	DeleteProductById(id uuid.UUID) (models.Product, error)
	GetProductById(id uuid.UUID) (models.Product, error)
	UpdateProductById(product models.Product, id uuid.UUID) (models.Product, error)
}
