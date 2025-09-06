package warehouseService

import (
	"MarketPlace_Pet/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockWarehouseRepository struct {
	mock.Mock
}

func (m *MockWarehouseRepository) GetAll() ([]models.Product, error) {
	args := m.Called()
	if res := args.Get(0); res != nil {
		return res.([]models.Product), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockWarehouseRepository) Create(product models.Product) (models.Product, error) {
	args := m.Called(product)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockWarehouseRepository) GetByID(productID string) (models.Product, error) {
	args := m.Called(productID)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockWarehouseRepository) Update(product models.Product) (models.Product, error) {
	args := m.Called(product)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockWarehouseRepository) Delete(productID string) error {
	args := m.Called(productID)
	return args.Error(0)
}
