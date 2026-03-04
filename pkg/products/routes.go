package products

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	ProductsConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", ProductsConfigurator.addProductHandler)
	router.Get("/{id}", ProductsConfigurator.productByIdHandler)
	router.Delete("/{id}", ProductsConfigurator.deleteProductHandler)
	router.Put("/{id}", ProductsConfigurator.editProductHandler)
	return router
}