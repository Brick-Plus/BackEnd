package models

import "net/http"

type Theme struct {
	Theme string `json:"theme"`
	Image string `json:"image"`
}

func (t *Theme) Bind(r *http.Request) error {
	return nil
}