package themes

import (
	"BrickPlus/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router{
	ThemesConfigurator := New(configuration)
	router := chi.NewRouter()
	router.Post("/", ThemesConfigurator.addThemeHandler)
	router.Get("/{id}", ThemesConfigurator.themeByIdHandler)
	router.Delete("/{id}", ThemesConfigurator.deleteThemehandler)
	router.Put("/{id}", ThemesConfigurator.editThemeHandler)
	return router
}