package models

import (
	"errors"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Country string `json:"country"`
	Email string `json:"email"`
	PhoneNumber string  `json:"phone_number"`
	Password string `json:"password"`
	IsAdmin bool `json:"is_admin"`
	IsSubscribed bool `json:"is_subscribed"`
	DeliveryAddress string `json:"delivery_address"`
	BillingAddress string `json:"billing_address"`
}

func (u *User) Bind(r *http.Request) error {
	if r.Method == "POST"{
		if u.IsAdmin {
		return errors.New("User cannot be an admin !")
	}
		return nil
	}
	return nil
}