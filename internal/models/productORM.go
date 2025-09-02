package models

import "gorm.io/gorm"

type Product struct {
	ID          string         `json:"product_id" gorm:"column:product_id;primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Price       float32        `json:"price"`
	Quantity    int            `json:"quantity"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type UpdateProduct struct {
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	Price       *float32       `json:"price,omitempty"`
	Quantity    *int           `json:"quantity,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
