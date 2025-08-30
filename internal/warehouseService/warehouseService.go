package warehouseService

import (
	"MarketPlace_Pet/internal/models"
	"fmt"

	"github.com/google/uuid"
)

type WarehouseService interface {
	GetAllProducts() ([]models.Product, error)
	CreateProduct(name, description string, price float64, quantity int) (models.Product, error)
	GetProductByID(productID string) (models.Product, error)
	UpdateProductByID(productID string, name, description *string, price *float64, quantity *int) (models.Product, error)
	DeleteProduct(productID string) error
}

type warehouseService struct {
	warehouseRepo WarehouseRepository
}

func NewWarehouseService(warehouseRepo WarehouseRepository) WarehouseService {
	return &warehouseService{warehouseRepo: warehouseRepo}
}

func (s *warehouseService) GetAllProducts() ([]models.Product, error) {
	return s.warehouseRepo.GetAll()
}

func (s *warehouseService) CreateProduct(name, description string, price float64, quantity int) (models.Product, error) {
	product := models.Product{}
	product.ID = uuid.New().String()
	product.Name = name
	product.Description = description
	product.Price = price
	product.Quantity = quantity

	return s.warehouseRepo.Create(product)
}

func (s *warehouseService) GetProductByID(productID string) (models.Product, error) {
	return s.warehouseRepo.GetByID(productID)
}

func (s *warehouseService) UpdateProductByID(productID string, name, description *string, price *float64, quantity *int) (models.Product, error) {
	product, err := s.warehouseRepo.GetByID(productID)
	if err != nil {
		return models.Product{}, fmt.Errorf("service: product not found: %w", err)
	}
	if name != nil {
		product.Name = *name
	}
	if description != nil {
		product.Description = *description
	}
	if price != nil {
		product.Price = *price
	}
	if quantity != nil {
		product.Quantity = *quantity
	}
	return s.warehouseRepo.Update(product)
}

func (s *warehouseService) DeleteProduct(productID string) error {
	return s.warehouseRepo.Delete(productID)
}
