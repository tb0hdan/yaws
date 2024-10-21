package postgresql

import (
	"yaws/internal/server/api"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Store struct {
	Connection string
}

func (s Store) PaymentWebhook(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetOrders(ctx echo.Context) (api.OrderList, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) CreateOrder(ctx echo.Context) (api.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetOrderById(ctx echo.Context, id uuid.UUID) (*api.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) UpdateOrderStatus(ctx echo.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetProducts(ctx echo.Context, params api.GetProductsParams) (api.ProductList, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) AddProducts(ctx echo.Context) (api.ProductList, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) DeleteProductById(ctx echo.Context, id uuid.UUID) (api.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetProductById(ctx echo.Context, id uuid.UUID) (api.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s Store) UpdateProductById(ctx echo.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
func (s Store) Connect() error {
	//TODO implement me
	panic("implement me")
}

func New() Store {
	return Store{}
}
