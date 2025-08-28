package models

type Product struct {
	ID          string  `json:"product_id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
}
