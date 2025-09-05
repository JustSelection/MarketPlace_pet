package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        string         `json:"order_id" gorm:"column:order_id;primaryKey"`
	UserID    string         `json:"user_id" gorm:"not null;index"`
	Items     []OrderItem    `json:"cart_items"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type OrderItem struct {
	OrderID   string         `json:"order_id" gorm:"primaryKey"`
	ProductID string         `json:"product_id" gorm:"not null;index"`
	UserID    string         `json:"user_id" gorm:"not null;index"`
	Quantity  int            `json:"quantity"`
	Price     float32        `json:"price"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
