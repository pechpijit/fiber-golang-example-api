package models

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
}

type ProductRequest struct {
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
}
