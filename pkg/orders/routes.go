package orders

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	ordersConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", ordersConfigurator.addOrderHandler)
	router.Get("/{ref}", ordersConfigurator.orderByReferenceHandler)
	router.Get("/user/{id}", ordersConfigurator.ordersByUserHandler)
	router.Delete("/{ref}", ordersConfigurator.deleteOrderHandler)
	return router
}
