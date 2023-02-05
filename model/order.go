package model

type Order struct {
	OrderID string  `json:"order_id"`
	Product Product `json:"product"`
	Address Address `json:"address"`
	Date    Date    `json:"date"`
}
