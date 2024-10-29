package store

import (
	"yaws/internal/store/postgresql"
	"yaws/internal/store/postgresql/models"

	"github.com/google/uuid"
)

// nolint: interfacebloat
type Store interface {
	Connect() error
	GetCustomers(limit, offset int32) ([]models.Customer, error)
	AddCustomers(customers []models.Customer) ([]models.Customer, error)
	DeleteCustomerById(id int32) (models.Customer, error)
	GetCustomerById(id int32) (models.Customer, error)
	UpdateCustomerById(customer models.Customer, id int32) (models.Customer, error)
	GetOrders(limit, offset int32, status, paymentStatus string) ([]models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	GetOrderById(id uuid.UUID) (models.Order, error)
	UpdateOrderStatus(orderStatus models.OrderStatus) (models.Order, error)
	PaymentWebhook(webhook models.Webhook) error
	GetProducts(limit, offset, minQuantity int32) ([]models.Product, error)
	AddProducts(products []models.Product) ([]models.Product, error)
	DeleteProductById(id uuid.UUID) (models.Product, error)
	GetProductById(id uuid.UUID) (models.Product, error)
	UpdateProductById(product models.Product, id uuid.UUID) (models.Product, error)
}

const (
	PostgreSQL = "postgresql"
)

func New(storeType, connection string) Store {
	switch storeType {
	case PostgreSQL:
		return postgresql.New(connection)
	default:
		return nil
	}
}
