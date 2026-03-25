package models

import "net/http"

type Favorite struct {
	IdUser    int `json:"id_user"`
	IdProduct int `json:"id_product"`
}

func (f *Favorite) Bind(r *http.Request) error {
	return nil
}
