package models

import (
	"net/http"
	"time"
)

type Order struct {
	Reference string    `json:"reference"`
	IdUser    int       `json:"id_user"`
	ProductIDs []int    `json:"product_ids"` // List of product IDs in the order
	ExpDate   time.Time `json:"exp_date"`
	OrderDate time.Time `json:"order_date"`
	PriceTTC  float64   `json:"price_ttc"`
	Devise    string    `json:"devise"`
}

func (o *Order) Bind(r *http.Request) error {
	return nil
}
