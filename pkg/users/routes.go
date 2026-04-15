package users

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	UserConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", UserConfigurator.addUserHandler)
	router.With(security.UserMiddleware).Get("/{id}", UserConfigurator.userByIdHandler)
	router.With(security.UserMiddleware).Delete("/{id}", UserConfigurator.deleteUserHandler)
	router.With(security.UserMiddleware).Put("/{id}", UserConfigurator.editUserHandler)
	return router
}