package server

import (
	"yaws/pkg/types"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type WebStoreServer struct {
	// This is a placeholder for the server
	logger zerolog.Logger
	store  types.Store
}

func (w WebStoreServer) GetOrders(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) CreateOrder(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) GetOrderById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) UpdateOrderStatus(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) PaymentWebhook(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) GetProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) AddProducts(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) DeleteProductById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) GetProductById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (w WebStoreServer) UpdateProductById(ctx echo.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewWebStoreServer(logger zerolog.Logger, store types.Store) WebStoreServer {
	return WebStoreServer{logger: logger, store: store}
}
