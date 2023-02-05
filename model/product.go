package model

type Product struct {
	Id      string  `json:"product_id"`
	Name    string  `json:"product_name"`
	Price   float64 `json:"product_price"`
	Company string  `json:"company_name,omitempty"`
}
