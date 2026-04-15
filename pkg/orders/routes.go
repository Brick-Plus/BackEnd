package orders

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	ordersConfigurator := New(configuration)
	router := chi.NewRouter()
	
	router.With(security.AdminMiddleware).Get("/{ref}", ordersConfigurator.orderByReferenceHandler)

	router.Group(func(r chi.Router) {
		r.Use(security.UserMiddleware)

		r.Post("/", ordersConfigurator.addOrderHandler)
		r.Get("/user/{id}", ordersConfigurator.ordersByUserHandler)
		r.Delete("/{ref}", ordersConfigurator.deleteOrderHandler)
	})
	return router
}
