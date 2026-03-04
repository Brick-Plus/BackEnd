package users

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	UserConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", UserConfigurator.addUserHandler)
	router.Get("/{id}", UserConfigurator.userByIdHandler)
	router.Delete("/{id}", UserConfigurator.deleteUserHandler)
	router.Put("/{id}", UserConfigurator.editUserHandler)
	return router
}