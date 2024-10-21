package server

import (
	"fmt"

	"yaws/internal/server/api"
	"yaws/internal/store/postgresql/models"
)

func FromAPIProductToModelsProduct(product api.Product) models.Product {
	return models.Product{
		Name:     product.Name,
		Price:    product.Price,
		Quantity: *product.Quantity,
	}
}

func FromAPIProductListToModelsProductList(products api.ProductList) []models.Product {
	var productList []models.Product
	for _, product := range products {
		productList = append(productList, FromAPIProductToModelsProduct(product))
	}
	return productList
}

func FromModelsProductToAPIProduct(product models.Product) api.Product {
	return api.Product{
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
	return models.Order{
		CustomerId:    order.CustomerId,
		PaymentStatus: fmt.Sprint(order.PaymentStatus),
		Status:        fmt.Sprint(order.Status),
		TotalPrice:    order.TotalPrice,
	}
}

func FromAPIOrderListToModelsOrderList(orders api.OrderList) []models.Order {
	var orderList []models.Order
	for _, order := range orders.Order {
		orderList = append(orderList, FromAPIOrderToModelsOrder(order))
	}
	return orderList
}

func FromModelsOrderToAPIOrder(order models.Order) api.Order {
	return api.Order{
		CustomerId: order.CustomerId,
		// PaymentStatus: order.PaymentStatus,
		// Status:        api.Status(order.Status),
		TotalPrice: order.TotalPrice,
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
		OrderId:       webhook.OrderId,
		PaymentStatus: webhook.PaymentStatus,
	}
}
