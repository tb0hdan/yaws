package store

import (
	"yaws/internal/store/postgresql"

	"github.com/labstack/echo/v4"
)

type Store interface {
	Connect() error
	GetOrders(ctx echo.Context) error
	CreateOrder(ctx echo.Context) error
	GetOrderById(ctx echo.Context, id string) error
	UpdateOrderStatus(ctx echo.Context, id string) error
	PaymentWebhook(ctx echo.Context) error
	GetProducts(ctx echo.Context) error
	AddProducts(ctx echo.Context) error
	DeleteProductById(ctx echo.Context, id string) error
	GetProductById(ctx echo.Context, id string) error
	UpdateProductById(ctx echo.Context, id string) error
}

const (
	PostgreSQL = "postgresql"
)

func New(storeType, connection string) Store {
	switch storeType {
	case PostgreSQL:
		return postgresql.Store{Connection: connection}
	default:
		return nil
	}
}