package notices

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	noticesConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Get("/{id}", noticesConfigurator.noticeByIdHandler)
	router.Get("/product/{id}", noticesConfigurator.noticesByProductHandler)
	router.Get("/user/{id}", noticesConfigurator.noticesByUserHandler)

	router.Group(func(r chi.Router) {
		r.Use(security.UserMiddleware)

		r.Post("/", noticesConfigurator.addNoticeHandler)
		r.Delete("/{id}", noticesConfigurator.deleteNoticeHandler)
		r.Put("/{id}", noticesConfigurator.editNoticeHandler)
	})

	return router
}
