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
	)
	log.Println("Database migrated succesfully")
}