package database

import (
	"log"
	"BrickPlus/database/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.Product{},
		&dbmodel.Categorie{},
		&dbmodel.Theme{},
		&dbmodel.Favorite{},
	)
	log.Println("Database migrated succesfully")
}