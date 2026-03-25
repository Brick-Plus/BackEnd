package models

import "net/http"

type Notice struct {
	IdUser    int    `json:"id_user"`
	IdProduct int    `json:"id_product"`
	Content   string `json:"content"`
}

func (n *Notice) Bind(r *http.Request) error {
	return nil
}