package dbmodel

import "gorm.io/gorm"

type Categorie struct {
	gorm.Model
	Categorie string `json:"categorie"`
}

type CategoriesRepository interface {
	Create(categorie *Categorie)(*Categorie, error)
	FindCategorieById(categorieId string) ([]*Categorie, error)
	CategorieToDelete(categorie *Categorie) error
	CategorieToUpdate(categorie *Categorie, categorieId string) error
}

type categoriesRepository struct{
	db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) CategoriesRepository {
	return &categoriesRepository{db: db}
}

func (r *categoriesRepository) Create(categorie *Categorie) (*Categorie, error){
	if err := r.db.Create(categorie).Error; err != nil{
		return nil, err
	}
	return categorie, nil
}

func (r *categoriesRepository) FindCategorieById(categorieId string) ([]*Categorie, error){
	var entry []*Categorie
	if err := r.db.Where("id = ?", categorieId).Find(&entry).Error; err != nil{
		return nil, err
	}
	return entry, nil
}

func (r *categoriesRepository) CategorieToDelete(categorie *Categorie) error {
	if err := r.db.Delete(categorie).Error; err != nil{
		return err
	}
	return nil
}

func (r *categoriesRepository) CategorieToUpdate(categorie *Categorie, categorieId string) error {
	if err := r.db.Where("id = ?", categorieId).Updates(categorie).Error; err != nil {
		return err
	}
	return nil
}