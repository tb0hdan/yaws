package server

import (
	"net/http"

	"yaws/internal/server/api"
	"yaws/internal/store"
	"yaws/internal/transactional"
	"yaws/pkg/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

const (
	// DefaultLimit is the default limit for pagination.
	DefaultLimit = int32(25)
	// DefaultOffset is the default offset for pagination.
	DefaultOffset = int32(0)
	// DefaultMinQuantity is the default minimum quantity for products.
	DefaultMinQuantity = int32(0)
)

type WebStoreServer struct {
	// This is a placeholder for the server
	logger zerolog.Logger
	store  store.Store
	sender *transactional.Sender
}

func (w *WebStoreServer) GetCustomers(ctx echo.Context, params api.GetCustomersParams) error {
	limit := DefaultLimit
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := DefaultOffset
	if params.Offset != nil {
		offset = *params.Offset
	}
	customers, err := w.store.GetCustomers(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerListToAPICustomerList(customers))
}

func (w *WebStoreServer) AddCustomers(ctx echo.Context) error {
	var (
		req api.CustomerList
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}
	if len(req) == 0 {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError("No customers to add"))
	}

	customers, err := w.store.AddCustomers(FromAPICustomerListToModelsCustomerList(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerListToAPICustomerList(customers))
}

func (w *WebStoreServer) DeleteCustomerById(ctx echo.Context, id int32) error {
	customer, err := w.store.DeleteCustomerById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerToAPICustomer(customer))
}

func (w *WebStoreServer) GetCustomerById(ctx echo.Context, id int32) error {
	customer, err := w.store.GetCustomerById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerToAPICustomer(customer))
}

func (w *WebStoreServer) UpdateCustomerById(ctx echo.Context, id int32) error {
	var (
		req api.Customer
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}

	customer, err := w.store.UpdateCustomerById(FromAPICustomerToModelsCustomer(req), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsCustomerToAPICustomer(customer))
}

func (w *WebStoreServer) GetOrders(ctx echo.Context, params api.GetOrdersParams) error {
	limit := DefaultLimit
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := DefaultOffset
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
		return ctx.JSON(http.StatusInternalServerError, utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsOrderListToAPIOrderList(orders))
}

func (w *WebStoreServer) CreateOrder(ctx echo.Context) error {
	var (
		req     api.Order
		lineMap = make(map[uuid.UUID]api.LineItem)
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}
	if req.CustomerId == 0 {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError("Customer ID is required"))
	}
	// Check for line item duplicates
	for _, line := range req.Products {
		if _, ok := lineMap[line.Id]; ok {
			return ctx.JSON(http.StatusBadRequest, utils.ToAPIError("Duplicate line items"))
		}
		lineMap[line.Id] = line
	}
	// Create order
	order, err := w.store.CreateOrder(FromAPIOrderToModelsOrder(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	// Send email
	customer, err := w.store.GetCustomerById(req.CustomerId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	err = w.sender.SendSimple(customer,
		"Order Confirmation", "Your order has been received")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}

	return ctx.JSON(http.StatusOK, FromModelsOrderToAPIOrder(order))
}

func (w *WebStoreServer) GetOrderById(ctx echo.Context, id uuid.UUID) error {
	order, err := w.store.GetOrderById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}

	return ctx.JSON(http.StatusOK, FromModelsOrderToAPIOrder(order))
}

func (w *WebStoreServer) UpdateOrderStatus(ctx echo.Context, id uuid.UUID) error {
	var (
		req api.OrderStatus
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}

	order, err := w.store.UpdateOrderStatus(FromAPIOrderStatusToModelsOrderStatus(req, id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	customer, err := w.store.GetCustomerById(order.CustomerID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	err = w.sender.SendSimple(customer,
		"Order Status update", "Your order has been updated")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsOrderToAPIOrder(order))
}

func (w *WebStoreServer) PaymentWebhook(ctx echo.Context) error {
	var (
		req api.Webhook
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}
	// Primitive check for customer ID
	if ctx.Request().Header.Get("X-Customer-Id") != "123" {
		return ctx.JSON(http.StatusUnauthorized, utils.ToAPIError("Unauthorized"))
	}
	//
	err := w.store.PaymentWebhook(FromAPIWebhookToModelsWebhook(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	order, err := w.store.GetOrderById(req.OrderId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	customer, err := w.store.GetCustomerById(order.CustomerID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	err = w.sender.SendSimple(customer,
		"Payment Status update", "Your payment was received")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}

	return ctx.JSON(http.StatusOK, "Payment webhook received")
}

func (w *WebStoreServer) GetProducts(ctx echo.Context, params api.GetProductsParams) error {
	limit := DefaultLimit
	if params.Limit != nil {
		limit = *params.Limit
	}
	offset := DefaultOffset
	if params.Offset != nil {
		offset = *params.Offset
	}
	minQuantity := DefaultMinQuantity
	if params.MinQuantity != nil {
		minQuantity = *params.MinQuantity
	}

	products, err := w.store.GetProducts(limit, offset, minQuantity)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductListToAPIProductList(products))
}

func (w *WebStoreServer) AddProducts(ctx echo.Context) error {
	var (
		req        api.ProductList
		productMap = make(map[uuid.UUID]api.Product)
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}
	if len(req) == 0 {
		return ctx.JSON(http.StatusBadRequest,
			utils.ToAPIError("Bad Request"))
	}
	// Validate product list
	for _, product := range req {
		if product.Quantity <= 0 {
			return ctx.JSON(http.StatusBadRequest,
				utils.ToAPIError("Quantity must be greater than 0"))
		}
		price, err := decimal.NewFromString(product.Price)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				utils.ToAPIError("Invalid price"))
		}
		if price.LessThanOrEqual(decimal.NewFromInt(0)) {
			return ctx.JSON(http.StatusBadRequest, "Price must be greater than 0")
		}
		if _, ok := productMap[product.Id]; ok {
			return ctx.JSON(http.StatusBadRequest,
				utils.ToAPIError("Duplicate products"))
		}
		productMap[product.Id] = product
	}
	//
	products, err := w.store.AddProducts(FromAPIProductListToModelsProductList(req))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductListToAPIProductList(products))
}

func (w *WebStoreServer) DeleteProductById(ctx echo.Context, id uuid.UUID) error {
	product, err := w.store.DeleteProductById(id)
	if err != nil {
		if err.Error() == "record not found" {
			return ctx.JSON(http.StatusNotFound,
				utils.ToAPIError(errors.Wrap(err, "Not Found")))
		}
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductToAPIProduct(product))
}

func (w *WebStoreServer) GetProductById(ctx echo.Context, id uuid.UUID) error {
	product, err := w.store.GetProductById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductToAPIProduct(product))
}

func (w *WebStoreServer) UpdateProductById(ctx echo.Context, id uuid.UUID) error {
	var (
		req api.Product
	)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest,
			utils.ToAPIError(errors.Wrap(err, "Bad Request")))
	}

	product, err := w.store.UpdateProductById(FromAPIProductToModelsProduct(req), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError,
			utils.ToAPIError(errors.Wrap(err, "Internal Server Error")))
	}
	return ctx.JSON(http.StatusOK, FromModelsProductToAPIProduct(product))
}

func NewWebStoreServer(logger zerolog.Logger, store store.Store, sender *transactional.Sender) WebStoreServer {
	return WebStoreServer{logger: logger, store: store, sender: sender}
}
