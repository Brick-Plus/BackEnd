package dbmodel

import "gorm.io/gorm"

type Theme struct {
	gorm.Model
	Theme string `json:"theme"`
	Image string `json:"image"`
}

type ThemesRepository interface {
	Create(theme *Theme)(*Theme, error)
	FindThemeById(themeId string)([]*Theme, error)
	ThemeToDelete(theme *Theme) error
	ThemeToUpdate(theme *Theme, themeId string) error
}

type themesRepository struct {
	db *gorm.DB
}

func NewThemesRepository(db *gorm.DB) ThemesRepository {
	return &themesRepository{db: db}
}

func (r *themesRepository) Create(theme *Theme)(*Theme, error){
	if err := r.db.Create(theme).Error; err != nil{
		return nil, err
	}
	return theme, nil
}

func (r *themesRepository) FindThemeById(themeId string)([]*Theme, error){
	var entry []*Theme
	if err := r.db.Where("id = ?", themeId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *themesRepository) ThemeToDelete(theme *Theme) error {
	if err := r.db.Delete(theme).Error; err != nil {
		return err
	}
	return nil
}

func (r *themesRepository) ThemeToUpdate(theme *Theme, themeId string) error {
	if err := r.db.Where("id = ?", themeId).Updates(theme).Error; err != nil {
		return err
	}
	return nil
}