package notices

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	noticesConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", noticesConfigurator.addNoticeHandler)
	router.Get("/{id}", noticesConfigurator.noticeByIdHandler)
	router.Get("/product/{id}", noticesConfigurator.noticesByProductHandler)
	router.Get("/user/{id}", noticesConfigurator.noticesByUserHandler)
	router.Delete("/{id}", noticesConfigurator.deleteNoticeHandler)
	router.Put("/{id}", noticesConfigurator.editNoticeHandler)
	return router
}