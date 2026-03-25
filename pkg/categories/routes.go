package categories

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(config *config.Config) chi.Router{
	CategorieConfigurator := New(config)
	router := chi.NewRouter()
	router.Get("/{id}", CategorieConfigurator.categorieByIdHandler)

	router.Group(func(r chi.Router) {
		r.Use(security.AdminMiddleware)

		r.Post("/", CategorieConfigurator.addCategorieHandler)
		r.Delete("/{id}", CategorieConfigurator.deleteCategorieHandler)
		r.Put("/{id}", CategorieConfigurator.updateCategorieHandler)
	})
	
	return router
}