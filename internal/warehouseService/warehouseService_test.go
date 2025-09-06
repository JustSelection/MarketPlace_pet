package warehouseService

import (
	"MarketPlace_Pet/internal/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	tests := []struct {
		name      string
		input     models.Product
		mockSetup func(m *MockWarehouseRepository, input models.Product)
		wantErr   bool
	}{
		{
			name:  "успешное создание",
			input: models.Product{Name: "Test Product", Price: 100.0, Quantity: 10},
			mockSetup: func(m *MockWarehouseRepository, input models.Product) {
				m.On("Create", mock.MatchedBy(func(p models.Product) bool {
					return p.Name == input.Name && p.Price == input.Price && p.Quantity == input.Quantity
				})).Return(input, nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка создания",
			input: models.Product{Name: "Bad Product", Price: 0, Quantity: 0},
			mockSetup: func(m *MockWarehouseRepository, input models.Product) {
				m.On("Create", mock.AnythingOfType("models.Product")).Return(models.Product{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockWarehouseRepository)
			tt.mockSetup(mockRepo, tt.input)

			// Убедитесь, что mockRepo передан в сервис
			service := NewWarehouseService(mockRepo)

			// Для цены используем указатель на float32, передавая значение вместо nil
			price := float32(tt.input.Price)
			_, err := service.CreateProduct(tt.input.Name, "", &price, tt.input.Quantity)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockWarehouseRepository)
		want      []models.Product
		wantErr   bool
	}{
		{
			name: "успешное получение всех продуктов",
			mockSetup: func(m *MockWarehouseRepository) {
				m.On("GetAll").Return([]models.Product{
					{ID: "1", Name: "Product 1", Price: 100.0, Quantity: 10},
					{ID: "2", Name: "Product 2", Price: 200.0, Quantity: 5},
				}, nil)
			},
			want: []models.Product{
				{ID: "1", Name: "Product 1", Price: 100.0, Quantity: 10},
				{ID: "2", Name: "Product 2", Price: 200.0, Quantity: 5},
			},
			wantErr: false,
		},
		{
			name: "ошибка репозитория",
			mockSetup: func(m *MockWarehouseRepository) {
				m.On("GetAll").Return(nil, errors.New("db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockWarehouseRepository)
			tt.mockSetup(mockRepo)

			service := NewWarehouseService(mockRepo)
			result, err := service.GetAllProducts()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(m *MockWarehouseRepository, id string)
		want      models.Product
		wantErr   bool
	}{
		{
			name: "успешное получение",
			id:   "1",
			mockSetup: func(m *MockWarehouseRepository, id string) {
				m.On("GetByID", id).Return(models.Product{
					ID: id, Name: "Test Product", Price: 100.0, Quantity: 10}, nil)
			},
			want:    models.Product{ID: "1", Name: "Test Product", Price: 100.0, Quantity: 10},
			wantErr: false,
		},
		{
			name: "ошибка получения",
			id:   "99",
			mockSetup: func(m *MockWarehouseRepository, id string) {
				m.On("GetByID", id).Return(models.Product{}, errors.New("not found"))
			},
			want:    models.Product{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockWarehouseRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewWarehouseService(mockRepo)
			result, err := service.GetProductByID(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
