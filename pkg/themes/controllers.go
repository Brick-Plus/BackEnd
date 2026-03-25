package themes

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ThemeConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *ThemeConfigurator{
	return &ThemeConfigurator{configuration}
}

func ThemeToModel(themes []*dbmodel.Theme) []models.Theme{
	themeToModel := &models.Theme{}
	themeEdited := []models.Theme{}
	for _, theme := range themes{
		themeToModel.Image = theme.Image
		themeToModel.Theme = theme.Image
		themeEdited = append(themeEdited, *themeToModel)
	}
	return themeEdited
}

func (config *ThemeConfigurator) themeByIdHandler(w http.ResponseWriter, r *http.Request){
	themeId := chi.URLParam(r, "id")
	theme, err := config.ThemesRepository.FindThemeById(themeId)
	if err != nil{
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted theme"})
		return
	}
	themeEdited := ThemeToModel(theme)
	render.JSON(w, r, themeEdited)
}

func (config *ThemeConfigurator) addThemeHandler(w http.ResponseWriter, r *http.Request){
	req := &models.Theme{}
	if err := render.Bind(r, req); err != nil{
		render.Status(r, 403)
		render.JSON(w, r, map[string]string{"Error" : err.Error()})
		return
	}
	addTheme := &dbmodel.Theme{Theme: req.Theme, Image: req.Image}
	config.ThemesRepository.Create(addTheme)
	render.JSON(w, r, map[string]string{"success" : "New theme successfully added !"})
}

func (config *ThemeConfigurator) deleteThemehandler(w http.ResponseWriter, r *http.Request){
	themeId := chi.URLParam(r, "id")
	theme, err := config.ThemesRepository.FindThemeById(themeId)
	if err != nil{
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted theme"})
	}
	config.ThemesRepository.ThemeToDelete(theme[0])
	render.JSON(w, r, map[string]string{"success": "The theme was deleted successfully !"})
}

func (config *ThemeConfigurator) editThemeHandler(w http.ResponseWriter, r *http.Request){
	req := &models.Theme{}
	themeId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil{
		render.Status(r, 403)
		render.JSON(w, r, map[string]string{"Error": err.Error()})
		return
	}
	updatedTheme := &dbmodel.Theme{Theme: req.Theme, Image: req.Image}
	config.ThemesRepository.ThemeToUpdate(updatedTheme, themeId)
	render.JSON(w, r, map[string]string{"Success": "Theme successfully updated !"})
}