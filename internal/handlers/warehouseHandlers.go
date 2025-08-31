package handlers

import (
	"MarketPlace_Pet/internal/warehouseService"
	"MarketPlace_Pet/internal/web/warehouse"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type WarehouseHandler struct {
	service warehouseService.WarehouseService
}

func NewWarehouseHandler(service warehouseService.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{service: service}
}

func (h *WarehouseHandler) GetWarehouse(_ context.Context, _ warehouse.GetWarehouseRequestObject) (warehouse.GetWarehouseResponseObject, error) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		return nil, fmt.Errorf("handler: could not get all products: %w", err)
	}

	responseProds := make([]warehouse.WarehouseProduct, len(products))
	for i, product := range products {
		responseProds[i] = warehouse.WarehouseProduct{
			Name:      product.Name,
			Price:     product.Price,
			Quantity:  product.Quantity,
			ProductID: product.ID,
		}
	}

	return warehouse.GetWarehouse200JSONResponse(responseProds), nil
}

func (h *WarehouseHandler) PostWarehouse(_ context.Context, req warehouse.PostWarehouseRequestObject) (warehouse.PostWarehouseResponseObject, error) {
	if req.Body.Quantity < 1 || req.Body.Price < 0.1 || req.Body.Description == "" || req.Body.Name == "" {
		return warehouse.PostWarehouse400Response{}, nil
	}

	product, err := h.service.CreateProduct(
		req.Body.Name,
		req.Body.Description,
		&req.Body.Price,
		req.Body.Quantity,
	)
	if err != nil {
		return nil, fmt.Errorf("handler: could not create product: %w", err)
	}

	return warehouse.PostWarehouse201JSONResponse(warehouse.WarehouseProduct{
		Name:        product.Name,
		Description: &product.Description,
		ProductID:   product.ID,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}), nil
}

func (h *WarehouseHandler) DeleteWarehouseProductId(_ context.Context, req warehouse.DeleteWarehouseProductIdRequestObject) (warehouse.DeleteWarehouseProductIdResponseObject, error) {
	err := h.service.DeleteProduct(req.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return warehouse.DeleteWarehouseProductId404Response{}, nil
		}
		return nil, fmt.Errorf("handler: could not delete product: %w", err)
	}

	return warehouse.DeleteWarehouseProductId204Response{}, nil
}

func (h *WarehouseHandler) GetWarehouseProductId(_ context.Context, req warehouse.GetWarehouseProductIdRequestObject) (warehouse.GetWarehouseProductIdResponseObject, error) {
	product, err := h.service.GetProductByID(req.ProductId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return warehouse.GetWarehouseProductId404Response{}, nil
		}
		return nil, fmt.Errorf("handler: could not get product: %w", err)
	}
	return warehouse.GetWarehouseProductId200JSONResponse(warehouse.WarehouseProduct{
		Name:        product.Name,
		Price:       product.Price,
		Quantity:    product.Quantity,
		ProductID:   product.ID,
		Description: &product.Description,
	}), nil
}

func (h *WarehouseHandler) PatchWarehouseProductId(_ context.Context, req warehouse.PatchWarehouseProductIdRequestObject) (warehouse.PatchWarehouseProductIdResponseObject, error) {
	if (req.Body.Quantity != nil && *req.Body.Quantity < 1) ||
		(req.Body.Price != nil && *req.Body.Price < 0.1) ||
		(req.Body.Description != nil && *req.Body.Description == "") ||
		(req.Body.Name != nil && *req.Body.Name == "") {
		return warehouse.PatchWarehouseProductId400Response{}, nil
	}
	if req.Body.Quantity == nil && req.Body.Name == nil && req.Body.Description == nil && req.Body.Price == nil {
		return warehouse.PatchWarehouseProductId400Response{}, nil
	}

	product, err := h.service.UpdateProductByID(
		req.ProductId,
		req.Body.Name,
		req.Body.Description,
		req.Body.Price,
		req.Body.Quantity,
	)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return warehouse.PatchWarehouseProductId404Response{}, nil
		}
		return nil, fmt.Errorf("handler: could not update product: %w", err)
	}
	return warehouse.PatchWarehouseProductId200JSONResponse(warehouse.WarehouseProduct{
		Name:        product.Name,
		Description: &product.Description,
		ProductID:   product.ID,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}), nil
}
