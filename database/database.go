package database

import (
	"log"
	"BrickPlus/database/dbmodel"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&dbmodel.User{},
		&dbmodel.Product{},
		&dbmodel.Categorie{},
		&dbmodel.Theme{},
		&dbmodel.Notice{},
		&dbmodel.Favorite{},
		&dbmodel.Order{},
	); err != nil {
		log.Panicln("Migration failed:", err)
	}
	log.Println("Database migrated succesfully")
}