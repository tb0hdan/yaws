// Code generated by mockery v2.46.3. DO NOT EDIT.

package store

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

type MockStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStore) EXPECT() *MockStore_Expecter {
	return &MockStore_Expecter{mock: &_m.Mock}
}

// AddProducts provides a mock function with given fields: ctx
func (_m *MockStore) AddProducts(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for AddProducts")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_AddProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddProducts'
type MockStore_AddProducts_Call struct {
	*mock.Call
}

// AddProducts is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockStore_Expecter) AddProducts(ctx interface{}) *MockStore_AddProducts_Call {
	return &MockStore_AddProducts_Call{Call: _e.mock.On("AddProducts", ctx)}
}

func (_c *MockStore_AddProducts_Call) Run(run func(ctx echo.Context)) *MockStore_AddProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockStore_AddProducts_Call) Return(_a0 error) *MockStore_AddProducts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_AddProducts_Call) RunAndReturn(run func(echo.Context) error) *MockStore_AddProducts_Call {
	_c.Call.Return(run)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *MockStore) Connect() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Connect")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type MockStore_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *MockStore_Expecter) Connect() *MockStore_Connect_Call {
	return &MockStore_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *MockStore_Connect_Call) Run(run func()) *MockStore_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStore_Connect_Call) Return(_a0 error) *MockStore_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_Connect_Call) RunAndReturn(run func() error) *MockStore_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// CreateOrder provides a mock function with given fields: ctx
func (_m *MockStore) CreateOrder(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type MockStore_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockStore_Expecter) CreateOrder(ctx interface{}) *MockStore_CreateOrder_Call {
	return &MockStore_CreateOrder_Call{Call: _e.mock.On("CreateOrder", ctx)}
}

func (_c *MockStore_CreateOrder_Call) Run(run func(ctx echo.Context)) *MockStore_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockStore_CreateOrder_Call) Return(_a0 error) *MockStore_CreateOrder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_CreateOrder_Call) RunAndReturn(run func(echo.Context) error) *MockStore_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteProductById provides a mock function with given fields: ctx, id
func (_m *MockStore) DeleteProductById(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProductById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_DeleteProductById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteProductById'
type MockStore_DeleteProductById_Call struct {
	*mock.Call
}

// DeleteProductById is a helper method to define mock.On call
//   - ctx echo.Context
//   - id string
func (_e *MockStore_Expecter) DeleteProductById(ctx interface{}, id interface{}) *MockStore_DeleteProductById_Call {
	return &MockStore_DeleteProductById_Call{Call: _e.mock.On("DeleteProductById", ctx, id)}
}

func (_c *MockStore_DeleteProductById_Call) Run(run func(ctx echo.Context, id string)) *MockStore_DeleteProductById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStore_DeleteProductById_Call) Return(_a0 error) *MockStore_DeleteProductById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_DeleteProductById_Call) RunAndReturn(run func(echo.Context, string) error) *MockStore_DeleteProductById_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrderById provides a mock function with given fields: ctx, id
func (_m *MockStore) GetOrderById(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_GetOrderById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrderById'
type MockStore_GetOrderById_Call struct {
	*mock.Call
}

// GetOrderById is a helper method to define mock.On call
//   - ctx echo.Context
//   - id string
func (_e *MockStore_Expecter) GetOrderById(ctx interface{}, id interface{}) *MockStore_GetOrderById_Call {
	return &MockStore_GetOrderById_Call{Call: _e.mock.On("GetOrderById", ctx, id)}
}

func (_c *MockStore_GetOrderById_Call) Run(run func(ctx echo.Context, id string)) *MockStore_GetOrderById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStore_GetOrderById_Call) Return(_a0 error) *MockStore_GetOrderById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_GetOrderById_Call) RunAndReturn(run func(echo.Context, string) error) *MockStore_GetOrderById_Call {
	_c.Call.Return(run)
	return _c
}

// GetOrders provides a mock function with given fields: ctx
func (_m *MockStore) GetOrders(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetOrders")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_GetOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOrders'
type MockStore_GetOrders_Call struct {
	*mock.Call
}

// GetOrders is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockStore_Expecter) GetOrders(ctx interface{}) *MockStore_GetOrders_Call {
	return &MockStore_GetOrders_Call{Call: _e.mock.On("GetOrders", ctx)}
}

func (_c *MockStore_GetOrders_Call) Run(run func(ctx echo.Context)) *MockStore_GetOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockStore_GetOrders_Call) Return(_a0 error) *MockStore_GetOrders_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_GetOrders_Call) RunAndReturn(run func(echo.Context) error) *MockStore_GetOrders_Call {
	_c.Call.Return(run)
	return _c
}

// GetProductById provides a mock function with given fields: ctx, id
func (_m *MockStore) GetProductById(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetProductById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_GetProductById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProductById'
type MockStore_GetProductById_Call struct {
	*mock.Call
}

// GetProductById is a helper method to define mock.On call
//   - ctx echo.Context
//   - id string
func (_e *MockStore_Expecter) GetProductById(ctx interface{}, id interface{}) *MockStore_GetProductById_Call {
	return &MockStore_GetProductById_Call{Call: _e.mock.On("GetProductById", ctx, id)}
}

func (_c *MockStore_GetProductById_Call) Run(run func(ctx echo.Context, id string)) *MockStore_GetProductById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStore_GetProductById_Call) Return(_a0 error) *MockStore_GetProductById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_GetProductById_Call) RunAndReturn(run func(echo.Context, string) error) *MockStore_GetProductById_Call {
	_c.Call.Return(run)
	return _c
}

// GetProducts provides a mock function with given fields: ctx
func (_m *MockStore) GetProducts(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetProducts")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_GetProducts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProducts'
type MockStore_GetProducts_Call struct {
	*mock.Call
}

// GetProducts is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockStore_Expecter) GetProducts(ctx interface{}) *MockStore_GetProducts_Call {
	return &MockStore_GetProducts_Call{Call: _e.mock.On("GetProducts", ctx)}
}

func (_c *MockStore_GetProducts_Call) Run(run func(ctx echo.Context)) *MockStore_GetProducts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockStore_GetProducts_Call) Return(_a0 error) *MockStore_GetProducts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_GetProducts_Call) RunAndReturn(run func(echo.Context) error) *MockStore_GetProducts_Call {
	_c.Call.Return(run)
	return _c
}

// PaymentWebhook provides a mock function with given fields: ctx
func (_m *MockStore) PaymentWebhook(ctx echo.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for PaymentWebhook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_PaymentWebhook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PaymentWebhook'
type MockStore_PaymentWebhook_Call struct {
	*mock.Call
}

// PaymentWebhook is a helper method to define mock.On call
//   - ctx echo.Context
func (_e *MockStore_Expecter) PaymentWebhook(ctx interface{}) *MockStore_PaymentWebhook_Call {
	return &MockStore_PaymentWebhook_Call{Call: _e.mock.On("PaymentWebhook", ctx)}
}

func (_c *MockStore_PaymentWebhook_Call) Run(run func(ctx echo.Context)) *MockStore_PaymentWebhook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockStore_PaymentWebhook_Call) Return(_a0 error) *MockStore_PaymentWebhook_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_PaymentWebhook_Call) RunAndReturn(run func(echo.Context) error) *MockStore_PaymentWebhook_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOrderStatus provides a mock function with given fields: ctx, id
func (_m *MockStore) UpdateOrderStatus(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrderStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_UpdateOrderStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOrderStatus'
type MockStore_UpdateOrderStatus_Call struct {
	*mock.Call
}

// UpdateOrderStatus is a helper method to define mock.On call
//   - ctx echo.Context
//   - id string
func (_e *MockStore_Expecter) UpdateOrderStatus(ctx interface{}, id interface{}) *MockStore_UpdateOrderStatus_Call {
	return &MockStore_UpdateOrderStatus_Call{Call: _e.mock.On("UpdateOrderStatus", ctx, id)}
}

func (_c *MockStore_UpdateOrderStatus_Call) Run(run func(ctx echo.Context, id string)) *MockStore_UpdateOrderStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStore_UpdateOrderStatus_Call) Return(_a0 error) *MockStore_UpdateOrderStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_UpdateOrderStatus_Call) RunAndReturn(run func(echo.Context, string) error) *MockStore_UpdateOrderStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateProductById provides a mock function with given fields: ctx, id
func (_m *MockStore) UpdateProductById(ctx echo.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_UpdateProductById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateProductById'
type MockStore_UpdateProductById_Call struct {
	*mock.Call
}

// UpdateProductById is a helper method to define mock.On call
//   - ctx echo.Context
//   - id string
func (_e *MockStore_Expecter) UpdateProductById(ctx interface{}, id interface{}) *MockStore_UpdateProductById_Call {
	return &MockStore_UpdateProductById_Call{Call: _e.mock.On("UpdateProductById", ctx, id)}
}

func (_c *MockStore_UpdateProductById_Call) Run(run func(ctx echo.Context, id string)) *MockStore_UpdateProductById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStore_UpdateProductById_Call) Return(_a0 error) *MockStore_UpdateProductById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStore_UpdateProductById_Call) RunAndReturn(run func(echo.Context, string) error) *MockStore_UpdateProductById_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStore creates a new instance of MockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStore {
	mock := &MockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
