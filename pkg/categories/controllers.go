package categories

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CategorieConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *CategorieConfigurator{
	return &CategorieConfigurator{configuration}
}

func CategorieToModel(categories []*dbmodel.Categorie) []models.Categorie{
	categorieToModel := &models.Categorie{}
	categoriesEdited := []models.Categorie{}
	for _, categorie := range categories{
		categorieToModel.Categorie = categorie.Categorie
		categoriesEdited = append(categoriesEdited, *categorieToModel)
	}
	return categoriesEdited
}

func (config *CategorieConfigurator) categorieByIdHandler(w http.ResponseWriter, r *http.Request){
	categorieId := chi.URLParam(r, "id")
	categorie, err := config.CategoriesRepository.FindCategorieById(categorieId)
	if err != nil{
		render.JSON(w, r, map[string]string{"Error": err.Error()})
		return
	}
	categoriesEdited := CategorieToModel(categorie)
	render.JSON(w, r, categoriesEdited)
}

func (config *CategorieConfigurator) addCategorieHandler(w http.ResponseWriter, r *http.Request){
	req := &models.Categorie{}
	if err := render.Bind(r, req); err != nil{
		render.JSON(w, r, map[string]string{"Error": err.Error()})
		return
	}
	addCategorie := &dbmodel.Categorie{Categorie: req.Categorie}
	config.CategoriesRepository.Create(addCategorie)
	render.JSON(w, r, map[string]string{"Success" : "New categorie successfully added !"})
}

func (config *CategorieConfigurator) deleteCategorieHandler(w http.ResponseWriter, r *http.Request){
	categorieId := chi.URLParam(r, "id")
	categorie, err := config.CategoriesRepository.FindCategorieById(categorieId)
	if err != nil{
		render.JSON(w, r, map[string]string{"Error": err.Error()})
		return
	}
	config.CategoriesRepository.CategorieToDelete(categorie[0])
	render.JSON(w, r, map[string]string{"Success": "Categorie successfully deleted !"})
}

func (config *CategorieConfigurator) updateCategorieHandler(w http.ResponseWriter, r *http.Request){
	categorieId := chi.URLParam(r, "id")
	req := &models.Categorie{}
	if err := render.Bind(r, req); err != nil{
		render.JSON(w, r, map[string]string{"Error": err.Error()})
		return
	}
	updatedCategorie := dbmodel.Categorie{Categorie: req.Categorie}
	config.CategoriesRepository.CategorieToUpdate(&updatedCategorie, categorieId)
	render.JSON(w, r, map[string]string{"Success": "Categoriesuccessfully updated !"})
}