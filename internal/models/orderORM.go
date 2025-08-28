package models

import "time"

type Order struct {
	ID        string     `json:"order_id"`
	CartItems []CartItem `json:"cart_items"`
	CreatedAt time.Time  `json:"created_at"`
}
