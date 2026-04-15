package dbmodel

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	NameProduct string `json:"name_product"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	IdTheme int
	Theme Theme `gorm:"foreignKey:IdTheme"` 
	IdCategorie int
	Categorie Categorie `gorm:"foreignKey:IdCategorie"`
	SubTheme string `json:"sub_theme"`
	NbPieces int `json:"nb_pieces"`
	NbFigurines int `json:"nb_figurines"`
	LEGOReference string `json:"LEGO_reference"`
	Weight float32 `json:"weight"`
	MainPicture string `json:"main_picture"`
	State string `json:"state"`
	Stock int `json:"stock"`
}

type ProductsRepository interface{
	Create(newProduct *Product)(*Product, error)
	FindProductById(productId string)([]*Product, error)
	Delete(productToDelete *Product) error
	Update(productToUpdate *Product, productId string) error
}

type productsRepository struct{
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) ProductsRepository {
	return &productsRepository{db: db}
}

func (r *productsRepository) Create(product *Product) (*Product, error){
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productsRepository) FindProductById(productId string)([]*Product, error){
	var entry []*Product
	if err := r.db.Where("id = ?", productId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *productsRepository) Delete(productToDelete *Product) error {
	if err := r.db.Delete(productToDelete).Error; err != nil{
		return err
	}
	return nil
}

func (r *productsRepository) Update(productToUpdate *Product, productId string) error {
	if err := r.db.Where("id = ?", productId).Updates(productToUpdate).Error; err != nil {
		return err
	}
	return nil
}