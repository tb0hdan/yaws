package server

import (
	"fmt"

	"yaws/internal/server/api"
	"yaws/internal/store/postgresql/models"
)

func FromAPIProductToModelsProduct(product api.Product) models.Product {
	var (
		description string
	)

	if product.Description != nil {
		description = *product.Description
	}

	return models.Product{ // nolint:exhaustruct
		ID:          product.Id,
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: description,
	}
}

func FromAPIProductListToModelsProductList(products api.ProductList) []models.Product {
	productList := make([]models.Product, 0, len(products))

	for _, product := range products {
		productList = append(productList, FromAPIProductToModelsProduct(product))
	}
	return productList
}

func FromModelsProductToAPIProduct(product models.Product) api.Product {
	return api.Product{ // nolint:exhaustruct
		Id:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		Description: &product.Description,
	}
}

func FromModelsProductListToAPIProductList(products []models.Product) api.ProductList {
	var productList api.ProductList
	for _, product := range products {
		productList = append(productList, FromModelsProductToAPIProduct(product))
	}
	return productList
}

func FromAPILineItemToModelsProduct(lineItem api.LineItem) models.Product {
	return models.Product{ // nolint:exhaustruct
		ID:       lineItem.ProductId,
		Quantity: lineItem.Quantity,
	}
}

func FromAPILineItemListToModelsProductList(lineItems []api.LineItem) []models.Product {
	productList := make([]models.Product, 0, len(lineItems))

	for _, lineItem := range lineItems {
		productList = append(productList, FromAPILineItemToModelsProduct(lineItem))
	}
	return productList
}

func FromAPIOrderToModelsOrder(order api.Order) models.Order {
	var (
		paymentStatus = "unpaid"
		orderStatus   = "pending"
	)

	if order.PaymentStatus != nil {
		paymentStatus = fmt.Sprint(*order.PaymentStatus)
	}

	if order.Status != nil {
		orderStatus = fmt.Sprint(*order.Status)
	}

	return models.Order{ // nolint:exhaustruct
		CustomerID:    order.CustomerId,
		PaymentStatus: paymentStatus,
		Status:        orderStatus,
		TotalPrice:    order.TotalPrice,
		Products:      FromAPILineItemListToModelsProductList(order.Products),
	}
}

func FromModelsOrderToAPIOrder(order models.Order) api.Order {
	var (
		paymentStatus,
		orderStatus interface{}
	)

	if order.PaymentStatus != "" {
		paymentStatus = order.PaymentStatus
	}

	if order.Status != "" {
		orderStatus = order.Status
	}

	return api.Order{ // nolint:exhaustruct
		CustomerId:    order.CustomerID,
		PaymentStatus: &paymentStatus,
		Status:        &orderStatus,
		TotalPrice:    order.TotalPrice,
	}
}

func FromModelsOrderListToAPIOrderList(orders []models.Order) api.OrderList {
	var orderList api.OrderList
	for _, order := range orders {
		orderList.Order = append(orderList.Order, FromModelsOrderToAPIOrder(order))
	}
	return orderList
}

func FromAPIWebhookToModelsWebhook(webhook api.Webhook) models.Webhook {
	return models.Webhook{
		OrderID:       webhook.OrderId,
		PaymentStatus: webhook.PaymentStatus,
	}
}

func FromModelsCustomerListToAPICustomerList(customers []models.Customer) api.CustomerList {
	var customerList api.CustomerList
	for _, customer := range customers {
		customerList = append(customerList, FromModelsCustomerToAPICustomer(customer))
	}
	return customerList
}

func FromModelsCustomerToAPICustomer(customer models.Customer) api.Customer {
	return api.Customer{ // nolint:exhaustruct
		Id:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
	}
}

func FromAPICustomerToModelsCustomer(customer api.Customer) models.Customer {
	return models.Customer{ // nolint:exhaustruct
		Name:    customer.Name,
		Email:   customer.Email,
		Phone:   customer.Phone,
		Address: customer.Address,
	}
}

func FromAPICustomerListToModelsCustomerList(customers api.CustomerList) []models.Customer {
	customerList := make([]models.Customer, 0, len(customers))
	for _, customer := range customers {
		customerList = append(customerList, FromAPICustomerToModelsCustomer(customer))
	}
	return customerList
}
