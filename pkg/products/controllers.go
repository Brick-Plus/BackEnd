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
	productsEdited := []models.Product{}
	for _, product := range products {
		pModel := models.Product{
			NameProduct:   product.NameProduct,
			Description:   product.Description,
			LEGOReference: product.LEGOReference,
			NbPieces:      product.NbPieces,
			NbFigurines:   product.NbFigurines,
			Price:         product.Price,
			State:         product.State,
			IdTheme:       product.IdTheme,
			IdCategorie:   product.IdCategorie,
			SubTheme:      product.SubTheme,
			Stock:         product.Stock,
			Weight:        product.Weight,
			MainPicture:   product.MainPicture,
		}
		productsEdited = append(productsEdited, pModel)
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
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	newProduct := &dbmodel.Product{
		NameProduct:   req.NameProduct,
		Description:   req.Description,
		LEGOReference: req.LEGOReference,
		NbPieces:      req.NbPieces,
		NbFigurines:   req.NbFigurines,
		Price:         req.Price,
		State:         req.State,
		IdTheme:       req.IdTheme,
		IdCategorie:   req.IdCategorie,
		SubTheme:      req.SubTheme,
		Stock:         req.Stock,
		Weight:        req.Weight,
		MainPicture:   req.MainPicture,
	}
	config.ProductsRepository.Create(newProduct)
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
		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedProduct := &dbmodel.Product{
		NameProduct:   req.NameProduct,
		Description:   req.Description,
		LEGOReference: req.LEGOReference,
		NbPieces:      req.NbPieces,
		NbFigurines:   req.NbFigurines,
		Price:         req.Price,
		State:         req.State,
		IdTheme:       req.IdTheme,
		IdCategorie:   req.IdCategorie,
		SubTheme:      req.SubTheme,
		Stock:         req.Stock,
		Weight:        req.Weight,
		MainPicture:   req.MainPicture,
	}
	config.ProductsRepository.Update(updatedProduct, productId)
	render.JSON(w, r, map[string]string{"success": "Product successfully updated"})
}
