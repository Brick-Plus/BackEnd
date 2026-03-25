package dbmodel

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	IdUser    int     `json:"id_user"`
	User      User    `gorm:"foreignKey:IdUser"`
	IdProduct int     `json:"id_product"`
	Product   Product `gorm:"foreignKey:IdProduct"`
}

type FavoritesRepository interface {
	Create(newFavorite *Favorite) (*Favorite, error)
	FindFavoritesByUserId(userId string) ([]*Favorite, error)
	FindFavoriteById(favoriteId string) ([]*Favorite, error)
	Delete(favoriteToDelete *Favorite) error
	// FindByUserAndProduct checks if a favorite already exists
	FindByUserAndProduct(userId int, productId int) ([]*Favorite, error)
}

type favoritesRepository struct {
	db *gorm.DB
}

func NewFavoritesRepository(db *gorm.DB) FavoritesRepository {
	return &favoritesRepository{db: db}
}

func (r *favoritesRepository) Create(favorite *Favorite) (*Favorite, error) {
	if err := r.db.Create(favorite).Error; err != nil {
		return nil, err
	}
	return favorite, nil
}

func (r *favoritesRepository) FindFavoritesByUserId(userId string) ([]*Favorite, error) {
	var entries []*Favorite
	if err := r.db.Where("id_user = ?", userId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *favoritesRepository) FindFavoriteById(favoriteId string) ([]*Favorite, error) {
	var entries []*Favorite
	if err := r.db.Where("id = ?", favoriteId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *favoritesRepository) Delete(favoriteToDelete *Favorite) error {
	if err := r.db.Delete(favoriteToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *favoritesRepository) FindByUserAndProduct(userId int, productId int) ([]*Favorite, error) {
	var entries []*Favorite
	if err := r.db.Where("id_user = ? AND id_product = ?", userId, productId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}
