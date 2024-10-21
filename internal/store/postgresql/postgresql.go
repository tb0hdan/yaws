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

func (s *Store) GetOrders(limit, offset int32, status, paymentStatus string) ([]models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) CreateOrder(order models.Order) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetOrderById(id uuid.UUID) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) UpdateOrderStatus(order models.Order, id uuid.UUID) (models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) PaymentWebhook(webhook models.Webhook) error {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetProducts(limit, offset, minQuantity int32) ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) AddProducts(products []models.Product) ([]models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) DeleteProductById(id uuid.UUID) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) GetProductById(id uuid.UUID) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) UpdateProductById(product models.Product, id uuid.UUID) (models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Store) Connect() error {
	db, err := gorm.Open(postgres.Open(s.Connection), &gorm.Config{})
	if err != nil {
		return err
	}
	s.db = db
	// Migrate the schema
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Product{})
	return nil
}

func New() Store {
	return Store{}
}
