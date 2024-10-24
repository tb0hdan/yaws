package server

import (
	"fmt"

	"yaws/internal/server/api"
	"yaws/internal/store/postgresql/models"
)

func FromAPIProductToModelsProduct(product api.Product) models.Product {
	return models.Product{ // nolint:exhaustruct
		Name:     product.Name,
		Price:    product.Price,
		Quantity: *product.Quantity,
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
		Name:     product.Name,
		Price:    product.Price,
		Quantity: &product.Quantity,
	}
}

func FromModelsProductListToAPIProductList(products []models.Product) api.ProductList {
	var productList api.ProductList
	for _, product := range products {
		productList = append(productList, FromModelsProductToAPIProduct(product))
	}
	return productList
}

func FromAPIOrderToModelsOrder(order api.Order) models.Order {
	return models.Order{ // nolint:exhaustruct
		CustomerID:    order.CustomerId,
		PaymentStatus: fmt.Sprint(order.PaymentStatus),
		Status:        fmt.Sprint(order.Status),
		TotalPrice:    order.TotalPrice,
	}
}

func FromModelsOrderToAPIOrder(order models.Order) api.Order {
	return api.Order{ // nolint:exhaustruct
		CustomerId:    order.CustomerID,
		PaymentStatus: order.PaymentStatus,
		Status:        order.Status,
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
		Id:   customer.ID,
		Name: customer.Name,
	}
}

func FromAPICustomerToModelsCustomer(customer api.Customer) models.Customer {
	return models.Customer{ // nolint:exhaustruct
		Name: customer.Name,
	}
}

func FromAPICustomerListToModelsCustomerList(customers api.CustomerList) []models.Customer {
	customerList := make([]models.Customer, 0, len(customers))
	for _, customer := range customers {
		customerList = append(customerList, FromAPICustomerToModelsCustomer(customer))
	}
	return customerList
}
