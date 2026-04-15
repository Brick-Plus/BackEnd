package products

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ProductsConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *ProductsConfigurator {
	return &ProductsConfigurator{configuration}
}

func productToModel(products []*dbmodel.Product) []models.Product {
	productToModel := &models.Product{}
	productsEdited := []models.Product{}
	for _, product := range products {
		productToModel.NameProduct = product.NameProduct
		productToModel.Description = product.Description
		productToModel.LEGOReference = product.LEGOReference
		productToModel.NbPieces = product.NbPieces
		productToModel.NbFigurines = product.NbFigurines
		productToModel.Price = product.Price
		productToModel.State = product.State
		productToModel.IdTheme = product.IdTheme
		productToModel.IdCategorie = product.IdCategorie
		productToModel.SubTheme = product.SubTheme
		productToModel.Stock = product.Stock
		productToModel.Weight = product.Weight
		productToModel.MainPicture = product.MainPicture
		productsEdited = append(productsEdited, *productToModel)
	}
	return productsEdited
}

func (config *ProductsConfigurator) productByIdHandler(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	product, err := config.ProductsRepository.FindProductById(productId)
	if err != nil {
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted product"})
		return
	}
	userEdited := productToModel(product)
	render.JSON(w, r, userEdited)
}

func (config *ProductsConfigurator) addProductHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Product{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, 401)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	addUser := &dbmodel.Product{NameProduct: req.NameProduct, Description: req.Description, LEGOReference: req.LEGOReference, NbPieces: req.NbPieces, NbFigurines: req.NbFigurines, Price: req.Price, State: req.State, IdTheme: req.IdTheme, IdCategorie: req.IdCategorie, SubTheme: req.SubTheme, Stock: req.Stock, Weight: req.Weight, MainPicture: req.MainPicture}
	config.ProductsRepository.Create(addUser)
	render.JSON(w, r, map[string]string{"success": "New product successfully added"})
}

func (config *ProductsConfigurator) deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "id")
	product, err := config.ProductsRepository.FindProductById(productId)
	if err != nil {
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted user"})
		return
	}
	config.ProductsRepository.Delete(product[0])
	render.JSON(w, r, map[string]string{"success": "Product successfully deleted"})
}

func (config *ProductsConfigurator) editProductHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Product{}
	productId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.Status(r, 401)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedProduct := &dbmodel.Product{NameProduct: req.NameProduct, Description: req.Description, LEGOReference: req.LEGOReference, NbPieces: req.NbPieces, NbFigurines: req.NbFigurines, Price: req.Price, State: req.State, IdTheme: req.IdTheme, IdCategorie: req.IdCategorie, SubTheme: req.SubTheme, Stock: req.Stock, Weight: req.Weight, MainPicture: req.MainPicture}
	config.ProductsRepository.Update(updatedProduct, productId)
	render.JSON(w, r, map[string]string{"success": "Product successfully updated"})
}
