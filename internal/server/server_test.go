package server

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"yaws/internal/server/api"
	"yaws/mocks/yaws/int/store"
	"yaws/mocks/yaws/int/transactional"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
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
	ts.mockStorage.On("GetCustomers", int32(25), int32(0)).Return(nil, nil)

	// Assertions
	if assert.NoError(ts.T(), server.GetCustomers(c, api.GetCustomersParams{})) {
		assert.Equal(ts.T(), http.StatusOK, rec.Code)
		assert.Equal(ts.T(), `null`+"\n", rec.Body.String())
	}
}
