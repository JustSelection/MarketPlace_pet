package warehouseService

import (
	"MarketPlace_Pet/internal/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type WarehouseRepository interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product) (models.Product, error)
	GetByID(productID string) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(productID string) error
}

type warehouseRepository struct {
	db *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository { return &warehouseRepository{db: db} }

func (r *warehouseRepository) GetAll() ([]models.Product, error) {
	products := make([]models.Product, 0)

	result := r.db.Where("deleted_at IS NULL").Find(&products)

	if result.Error != nil {
		return nil, fmt.Errorf("repo: could not get all products: %w", result.Error)
	}

	return products, nil
}

func (r *warehouseRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *warehouseRepository) GetByID(productID string) (models.Product, error) {
	var product models.Product

	result := r.db.Where("product_id = ? AND deleted_at IS NULL", productID).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Product{}, fmt.Errorf("repo: product not found: %w", result.Error)
		}
		return models.Product{}, fmt.Errorf("repo: could not get product: %w", result.Error)
	}
	return product, nil
}

func (r *warehouseRepository) Update(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return models.Product{}, fmt.Errorf("repo: could not update product: %w", err)
	}
	return product, nil
}

func (r *warehouseRepository) Delete(productID string) error {
	result := r.db.Where("product_id = ? AND deleted_at IS NULL", productID).Delete(&models.Product{})
	if result.Error != nil {
		return fmt.Errorf("repo: could not delete product: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("repo: product with id %s not found: %w", productID, result.Error)
	}
	return nil
}
