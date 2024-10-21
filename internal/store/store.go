package store

import (
	"yaws/internal/server/api"
	"yaws/internal/store/postgresql"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Store interface {
	Connect() error
	GetOrders(ctx echo.Context) (api.OrderList, error)
	CreateOrder(ctx echo.Context) (api.Order, error)
	GetOrderById(ctx echo.Context, id uuid.UUID) (*api.Order, error)
	UpdateOrderStatus(ctx echo.Context, id uuid.UUID) error
	PaymentWebhook(ctx echo.Context) error
	GetProducts(ctx echo.Context, params api.GetProductsParams) (api.ProductList, error)
	AddProducts(ctx echo.Context) (api.ProductList, error)
	DeleteProductById(ctx echo.Context, id uuid.UUID) (api.Product, error)
	GetProductById(ctx echo.Context, id uuid.UUID) (api.Product, error)
	UpdateProductById(ctx echo.Context, id uuid.UUID) error
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
