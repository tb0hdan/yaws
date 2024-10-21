package server

import (
	"net/http"

	"yaws/internal/server/api"
	"yaws/pkg/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type WebStoreServer struct {
	// This is a placeholder for the server
	logger zerolog.Logger
	store  types.Store
}

func (w *WebStoreServer) GetOrders(ctx echo.Context, params api.GetOrdersParams) error {
	orders, err := w.store.GetOrders(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (w *WebStoreServer) CreateOrder(ctx echo.Context) error {
	order, err := w.store.CreateOrder(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, order)
}

func (w *WebStoreServer) GetOrderById(ctx echo.Context, id uuid.UUID) error {
	order, err := w.store.GetOrderById(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}

	if order == nil {
		return ctx.JSON(http.StatusNotFound, "Order not found")
	}

	return ctx.JSON(http.StatusOK, order)
}

func (w *WebStoreServer) UpdateOrderStatus(ctx echo.Context, id uuid.UUID) error {
	err := w.store.UpdateOrderStatus(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, "Order updated")
}

func (w *WebStoreServer) PaymentWebhook(ctx echo.Context) error {
	err := w.store.PaymentWebhook(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, "Payment webhook received")
}

func (w *WebStoreServer) GetProducts(ctx echo.Context, params api.GetProductsParams) error {
	products, err := w.store.GetProducts(ctx, params)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, products)
}

func (w *WebStoreServer) AddProducts(ctx echo.Context) error {
	products, err := w.store.AddProducts(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, products)
}

func (w *WebStoreServer) DeleteProductById(ctx echo.Context, id uuid.UUID) error {
	product, err := w.store.DeleteProductById(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, product)
}

func (w *WebStoreServer) GetProductById(ctx echo.Context, id uuid.UUID) error {
	product, err := w.store.GetProductById(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, product)
}

func (w *WebStoreServer) UpdateProductById(ctx echo.Context, id uuid.UUID) error {
	err := w.store.UpdateProductById(ctx, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
	}
	return ctx.JSON(http.StatusOK, "Product updated")
}

func NewWebStoreServer(logger zerolog.Logger, store types.Store) WebStoreServer {
	return WebStoreServer{logger: logger, store: store}
}
