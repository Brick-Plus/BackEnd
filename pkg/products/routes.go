package products

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	ProductsConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Get("/{id}", ProductsConfigurator.productByIdHandler)

	router.Group(func(r chi.Router) {
		r.Use(security.AdminMiddleware)

		r.Post("/", ProductsConfigurator.addProductHandler)
		r.Delete("/{id}", ProductsConfigurator.deleteProductHandler)
		r.Put("/{id}", ProductsConfigurator.editProductHandler)
	})
	return router
}
