package models

type Product struct {
	ID          string  `json:"product_id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
