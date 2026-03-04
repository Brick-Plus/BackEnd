package models

import "net/http"

type Categorie struct {
	Categorie string `json:"categorie"`
}

func (c *Categorie) Bind(r *http.Request)error{
	return nil
}