package models

import "net/http"

type Product struct {
	NameProduct   string `json:"name_product"`
	Description   string `json:"description"`
	Price         float64 `json:"price"`
	IdTheme       int
	IdCategorie   int
	SubTheme      string  `json:"sub_theme"`
	NbPieces      int     `json:"nb_pieces"`
	NbFigurines   int     `json:"nb_figurines"`
	LEGOReference string  `json:"LEGO_reference"`
	Weight        float32 `json:"weight"`
	MainPicture   string  `json:"main_picture"`
	State         string  `json:"state"`
	Stock         int     `json:"stock"`
}

func (p *Product) Bind(r *http.Request) error {
	return nil
}
