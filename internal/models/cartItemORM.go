package models

type UserCartItem struct {
	ProductID   string  `json:"product_id" gorm:"primaryKey"`
	UserID      string  `json:"user_id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}
