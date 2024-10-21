package postgresql

import (
	"yaws/internal/store/postgresql/models"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	Connection string
	db         *gorm.DB
}

func (s *Store) GetCustomers(limit, offset int32) ([]models.Customer, error) {
	var (
		customers []models.Customer
	)
	return customers, s.db.Limit(int(limit)).Offset(int(offset)).Find(&customers).Error
}

func (s *Store) AddCustomers(customers []models.Customer) ([]models.Customer, error) {
	return customers, s.db.Create(&customers).Error
}

func (s *Store) DeleteCustomerById(id int32) (models.Customer, error) {
	var (
		customer models.Customer
	)
	err := s.db.First(&customer, id)
	if err.Error != nil {
		return customer, err.Error
	}
	return customer, s.db.Delete(&customer).Error
}

func (s *Store) GetCustomerById(id int32) (models.Customer, error) {
	var (
		customer models.Customer
	)
	return customer, s.db.First(&customer, id).Error
}

func (s *Store) UpdateCustomerById(customer models.Customer, id int32) (models.Customer, error) {
	var (
		customerToUpdate models.Customer
	)
	err := s.db.First(&customerToUpdate, id)
	if err.Error != nil {
		return customerToUpdate, err.Error
	}
	customerToUpdate.Name = customer.Name
	customerToUpdate.Email = customer.Email
	customerToUpdate.Phone = customer.Phone
	customerToUpdate.Address = customer.Address
	return customerToUpdate, s.db.Save(&customerToUpdate).Error
}

func (s *Store) GetOrders(limit, offset int32, status, paymentStatus string) ([]models.Order, error) {
	var (
		orders []models.Order
	)
	return orders, s.db.Limit(int(limit)).Offset(int(offset)).Where("status = ? AND payment_status = ?", status, paymentStatus).Find(&orders).Error
}

func (s *Store) CreateOrder(order models.Order) (models.Order, error) {
	return order, s.db.Create(&order).Error
}

func (s *Store) GetOrderById(id uuid.UUID) (models.Order, error) {
	var (
		order models.Order
	)
	return order, s.db.First(&order, id).Error
}

func (s *Store) UpdateOrderStatus(order models.Order, id uuid.UUID) (models.Order, error) {
	var (
		orderToUpdate models.Order
	)
	err := s.db.First(&orderToUpdate, id)
	if err.Error != nil {
		return orderToUpdate, err.Error
	}
	orderToUpdate.Status = order.Status
	return orderToUpdate, s.db.Save(&orderToUpdate).Error
}

func (s *Store) PaymentWebhook(webhook models.Webhook) error {
	var (
		order models.Order
	)
	err := s.db.First(&order, webhook.OrderId)
	if err.Error != nil {
		return err.Error
	}
	order.PaymentStatus = webhook.PaymentStatus
	return s.db.Save(&order).Error
}

func (s *Store) GetProducts(limit, offset, minQuantity int32) ([]models.Product, error) {
	var (
		products []models.Product
	)
	return products, s.db.Limit(int(limit)).Offset(int(offset)).Where("quantity > ?", minQuantity).Find(&products).Error
}

func (s *Store) AddProducts(products []models.Product) ([]models.Product, error) {
	return products, s.db.Create(&products).Error
}

func (s *Store) DeleteProductById(id uuid.UUID) (models.Product, error) {
	var (
		product models.Product
	)
	err := s.db.First(&product, id)
	if err.Error != nil {
		return product, err.Error
	}
	return product, s.db.Delete(&product).Error
}

func (s *Store) GetProductById(id uuid.UUID) (models.Product, error) {
	var (
		product models.Product
	)
	return product, s.db.First(&product, id).Error
}

func (s *Store) UpdateProductById(product models.Product, id uuid.UUID) (models.Product, error) {
	var (
		productToUpdate models.Product
	)
	err := s.db.First(&productToUpdate, id)
	if err.Error != nil {
		return productToUpdate, err.Error
	}
	productToUpdate.Name = product.Name
	productToUpdate.Price = product.Price
	productToUpdate.Quantity = product.Quantity
	return productToUpdate, s.db.Save(&productToUpdate).Error
}

func (s *Store) Connect() error {
	db, err := gorm.Open(postgres.Open(s.Connection), &gorm.Config{})
	if err != nil {
		return err
	}
	s.db = db
	// Migrate the schema
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&models.Customer{})
	if err != nil {
		return err
	}
	return err
}

func New(connection string) *Store {
	return &Store{Connection: connection}
}
