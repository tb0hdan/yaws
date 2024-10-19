package postgresql

import "github.com/labstack/echo/v4"

type Store struct {
	Connection string
}

func (s Store) GetOrders(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) CreateOrder(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetOrderById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) UpdateOrderStatus(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) PaymentWebhook(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) AddProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) DeleteProductById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) GetProductById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (s Store) UpdateProductById(ctx echo.Context, id string) error {
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
