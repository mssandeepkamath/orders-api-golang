package model

type Address struct {
	Flat    string `json:"flat_number"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state,omitempty"`
	Pincode string `json:"pincode"`
}
