package favorites

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	favoritesConfigurator := New(configuration)
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(security.UserMiddleware)

		r.Post("/", favoritesConfigurator.addFavoriteHandler)
		r.Get("/{id}", favoritesConfigurator.favoriteByIdHandler)
		r.Get("/user/{id}", favoritesConfigurator.favoritesByUserHandler)
		r.Delete("/{id}", favoritesConfigurator.deleteFavoriteHandler)
	})
	
	return router
}
