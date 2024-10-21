package server

import (
	"net/http"

	"yaws/internal/server/api"
	"yaws/internal/transactional"
	"yaws/pkg/types"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type WebStoreServer struct {
	// This is a placeholder for the server
	logger zerolog.Logger
	store  types.Store
	sender transactional.Transactional
}

func (w *WebStoreServer) GetCustomers(ctx echo.Context, params api.GetCustomersParams) error {
	limit := int32(25)
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := int32(0)
	if params.Offset != nil {
		offset = *params.Offset
	}
	customers, err := w.store.GetCustomers(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerListToAPICustomerList(customers))
}

func (w *WebStoreServer) AddCustomers(ctx echo.Context) error {
	var (
		req api.CustomerList
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}
	if len(req) == 0 {
		return ctx.JSON(http.StatusBadRequest, "No customers to add")
	}

	customers, err := w.store.AddCustomers(FromAPICustomerListToModelsCustomerList(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerListToAPICustomerList(customers))
}

func (w *WebStoreServer) DeleteCustomerById(ctx echo.Context, id int32) error {
	customer, err := w.store.DeleteCustomerById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerToAPICustomer(customer))
}

func (w *WebStoreServer) GetCustomerById(ctx echo.Context, id int32) error {
	customer, err := w.store.GetCustomerById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerToAPICustomer(customer))
}

func (w *WebStoreServer) UpdateCustomerById(ctx echo.Context, id int32) error {
	var (
		req api.Customer
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}

	customer, err := w.store.UpdateCustomerById(FromAPICustomerToModelsCustomer(req), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerToAPICustomer(customer))
}

func (w *WebStoreServer) GetOrders(ctx echo.Context, params api.GetOrdersParams) error {
	limit := int32(25)
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := int32(0)
	if params.Offset != nil {
		offset = *params.Offset
	}
	status := ""
	if params.Status != nil {
		status = *params.Status
	}
	paymentStatus := ""
	if params.PaymentStatus != nil {
		paymentStatus = *params.PaymentStatus
	}
	orders, err := w.store.GetOrders(limit, offset, status, paymentStatus)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsOrderListToAPIOrderList(orders))
}

func (w *WebStoreServer) CreateOrder(ctx echo.Context) error {
	var (
		req api.Order
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}
	if req.CustomerId == 0 {
		return ctx.JSON(http.StatusBadRequest, "Customer ID is required")
	}

	order, err := w.store.CreateOrder(FromAPIOrderToModelsOrder(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	// Send email
	customer, err := w.store.GetCustomerById(req.CustomerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	err = w.sender.Send(
		types.Contact{Name: "YAWS", Email: "yaws@example.com"},
		types.Contact{Name: customer.Name, Email: customer.Email},
		"Order Confirmation", "Your order has been received")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}

	return ctx.JSON(http.StatusOK, FromModelsOrderToAPIOrder(order))
}

func (w *WebStoreServer) GetOrderById(ctx echo.Context, id uuid.UUID) error {
	order, err := w.store.GetOrderById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}

	return ctx.JSON(http.StatusOK, FromModelsOrderToAPIOrder(order))
}

func (w *WebStoreServer) UpdateOrderStatus(ctx echo.Context, id uuid.UUID) error {
	var (
		req api.Order
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}

	order, err := w.store.UpdateOrderStatus(FromAPIOrderToModelsOrder(req), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	customer, err := w.store.GetCustomerById(req.CustomerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	err = w.sender.Send(
		types.Contact{Name: "YAWS", Email: "yaws@example.com"},
		types.Contact{Name: customer.Name, Email: customer.Email},
		"Order Status update", "Your order has been updated")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsOrderToAPIOrder(order))
}

func (w *WebStoreServer) PaymentWebhook(ctx echo.Context) error {
	var (
		req api.Webhook
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}
	// Primitive check for customer ID
	if ctx.Request().Header.Get("X-Customer-ID") != "123" {
		return ctx.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	//
	err := w.store.PaymentWebhook(FromAPIWebhookToModelsWebhook(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	order, err := w.store.GetOrderById(req.OrderId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	customer, err := w.store.GetCustomerById(order.CustomerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	err = w.sender.Send(
		types.Contact{Name: "YAWS", Email: "yaws@example.com"},
		types.Contact{Name: customer.Name, Email: customer.Email},
		"Payment Status update", "Your payment was received")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}

	return ctx.JSON(http.StatusOK, "Payment webhook received")
}

func (w *WebStoreServer) GetProducts(ctx echo.Context, params api.GetProductsParams) error {
	limit := int32(25)
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := int32(0)
	if params.Offset != nil {
		offset = *params.Offset
	}
	minQuantity := int32(0)
	if params.MinQuantity != nil {
		minQuantity = *params.MinQuantity
	}

	products, err := w.store.GetProducts(limit, offset, minQuantity)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductListToAPIProductList(products))
}

func (w *WebStoreServer) AddProducts(ctx echo.Context) error {
	var (
		req api.ProductList
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}
	if len(req) == 0 {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	products, err := w.store.AddProducts(FromAPIProductListToModelsProductList(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductListToAPIProductList(products))
}

func (w *WebStoreServer) DeleteProductById(ctx echo.Context, id uuid.UUID) error {
	product, err := w.store.DeleteProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductToAPIProduct(product))
}

func (w *WebStoreServer) GetProductById(ctx echo.Context, id uuid.UUID) error {
	product, err := w.store.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductToAPIProduct(product))
}

func (w *WebStoreServer) UpdateProductById(ctx echo.Context, id uuid.UUID) error {
	var (
		req api.Product
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errors.Wrap(err, "Bad Request"))
	}

	product, err := w.store.UpdateProductById(FromAPIProductToModelsProduct(req), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errors.Wrap(err, "Internal Server Error"))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductToAPIProduct(product))
}

func NewWebStoreServer(logger zerolog.Logger, store types.Store, sender transactional.Transactional) WebStoreServer {
	return WebStoreServer{logger: logger, store: store, sender: sender}
}
