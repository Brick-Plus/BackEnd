package categories

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config) chi.Router{
	CategorieConfigurator := New(config)
	router := chi.NewRouter()
	router.Post("/", CategorieConfigurator.addCategorieHandler)
	router.Get("/{id}", CategorieConfigurator.categorieByIdHandler)
	router.Delete("/{id}", CategorieConfigurator.deleteCategorieHandler)
	router.Put("/{id}", CategorieConfigurator.updateCategorieHandler)
	return router
}