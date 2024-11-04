package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"yaws/internal/server/api"
	"yaws/internal/store/postgresql/models"
	"yaws/mocks/yaws/int/store"
	"yaws/mocks/yaws/int/transactional"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	mockStorage       *store.MockStore
	mockTransactional *transactional.MockTransactional
	logger            zerolog.Logger
}

func TestServerTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}

func (ts *ServerTestSuite) SetupSuite() {
	ts.mockStorage = new(store.MockStore)
	ts.mockTransactional = new(transactional.MockTransactional)
	ts.logger = zerolog.New(os.Stderr)
}

func (ts *ServerTestSuite) AfterTest(_, _ string) {
	ts.mockStorage.AssertExpectations(ts.T())
	ts.mockTransactional.AssertExpectations(ts.T())
}

func (ts *ServerTestSuite) TearDownSuite() {
	// ts.realsomething.Close()
}

func (ts *ServerTestSuite) TestGetCustomers() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("GetCustomers", int32(25), int32(0)).Return(nil, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.GetCustomers(c, api.GetCustomersParams{})) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `null`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestAddCustomers() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(ts.T(), server.AddCustomers(c)) {
		assert.Equal(ts.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(ts.T(), `{"error":"No customers to add"}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestDeleteCustomerById() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("DeleteCustomerById", int32(1)).Return(models.Customer{}, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.DeleteCustomerById(c, int32(1))) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"address":"","email":"","id":0,"name":"","phone":""}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestGetCustomerById() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("GetCustomerById", int32(1)).Return(models.Customer{}, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.GetCustomerById(c, int32(1))) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"address":"","email":"","id":0,"name":"","phone":""}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestUpdateCustomerById() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("UpdateCustomerById", models.Customer{}, int32(1)).Return(models.Customer{}, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.UpdateCustomerById(c, int32(1))) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"address":"","email":"","id":0,"name":"","phone":""}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestGetOrders() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("GetOrders", int32(25), int32(0), "", "").Return([]models.Order{}, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.GetOrders(c, api.GetOrdersParams{})) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"Order":null,"discount":""}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestCreateOrder() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(ts.T(), server.CreateOrder(c)) {
		assert.Equal(ts.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(ts.T(), `{"error":"Customer ID is required"}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestGetOrderById() {
	var (
		testOrderId = uuid.New()
	)
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("GetOrderById", testOrderId).Return(models.Order{
		ID: testOrderId,
	}, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.GetOrderById(c, testOrderId)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"customer_id":0,"id":"`+testOrderId.String()+`","payment_status":null,"products":null,"status":null,"total_price":""}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestUpdateOrderStatus() {
	var (
		testOrderId = uuid.New()
	)
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("UpdateOrderStatus", models.OrderStatus{
		OrderID: testOrderId,
		Status:  "<nil>",
	}).Return(models.Order{
		ID:         testOrderId,
		CustomerID: 1,
	}, nil).Once()
	ts.mockStorage.On("GetCustomerById", int32(1)).Return(models.Customer{}, nil).Once()
	ts.mockTransactional.On("Send", mock.Anything, mock.Anything,
		"Order Status update", "Your order has been updated").Return(nil).Once()
	// Assertions
	if assert.NoError(ts.T(), server.UpdateOrderStatus(c, testOrderId)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"customer_id":1,"id":"`+testOrderId.String()+`","payment_status":null,"products":null,"status":null,"total_price":""}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestPaymentWebhook() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(ts.T(), server.PaymentWebhook(c)) {
		assert.Equal(ts.T(), http.StatusUnauthorized, rec.Code)
		assert.Equal(ts.T(), `{"error":"Unauthorized"}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestGetProducts() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("GetProducts", int32(25), int32(0), int32(0)).Return(nil, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.GetProducts(c, api.GetProductsParams{})) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `null`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestAddProducts() {
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(ts.T(), server.AddProducts(c)) {
		assert.Equal(ts.T(), http.StatusBadRequest, rec.Code)
		assert.Equal(ts.T(), `{"error":"Bad Request"}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestDeleteProductById() {
	var (
		testProductId = uuid.New()
	)
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("DeleteProductById", testProductId).Return(models.Product{
		ID:          testProductId,
		Description: "",
	}, nil).Once()
	// Assertions
	if assert.NoError(ts.T(), server.DeleteProductById(c, testProductId)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"description":"","id":"`+testProductId.String()+`","name":"","price":"","quantity":0}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestGetProductById() {
	var (
		testProductId = uuid.New()
	)
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("GetProductById", testProductId).Return(models.Product{
		ID:          testProductId,
		Description: "",
	}, nil).Once()

	// Assertions
	if assert.NoError(ts.T(), server.GetProductById(c, testProductId)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"description":"","id":"`+testProductId.String()+`","name":"","price":"","quantity":0}`+"\n", rec.Body.String())
	}
}

func (ts *ServerTestSuite) TestUpdateProductById() {
	var (
		testProductId = uuid.New()
	)
	server := NewWebStoreServer(ts.logger, ts.mockStorage, ts.mockTransactional)
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"id":"`+testProductId.String()+`"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Expectations
	ts.mockStorage.On("UpdateProductById", models.Product{
		ID: testProductId,
	}, testProductId).Return(models.Product{
		ID:          testProductId,
		Description: "",
	}, nil).Once()
	// Assertions
	if assert.NoError(ts.T(), server.UpdateProductById(c, testProductId)) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `{"description":"","id":"`+testProductId.String()+`","name":"","price":"","quantity":0}`+"\n", rec.Body.String())
	}
}
