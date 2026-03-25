package favorites

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	favoritesConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", favoritesConfigurator.addFavoriteHandler)
	router.Get("/{id}", favoritesConfigurator.favoriteByIdHandler)
	router.Get("/user/{id}", favoritesConfigurator.favoritesByUserHandler)
	router.Delete("/{id}", favoritesConfigurator.deleteFavoriteHandler)
	return router
}
