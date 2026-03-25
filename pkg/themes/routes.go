package themes

import (
	"BrickPlus/config"
	"BrickPlus/security"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	ThemesConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Get("/{id}", ThemesConfigurator.themeByIdHandler)

	router.Group(func(r chi.Router) {
		r.Use(security.AdminMiddleware)

		r.Post("/", ThemesConfigurator.addThemeHandler)
		r.Delete("/{id}", ThemesConfigurator.deleteThemehandler)
		r.Put("/{id}", ThemesConfigurator.editThemeHandler)
	})

	return router
}
