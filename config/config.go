package config

import (
	"BrickPlus/database"
	"BrickPlus/database/dbmodel"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	UsersRepository      dbmodel.UsersRepository
	ProductsRepository   dbmodel.ProductsRepository
	ThemesRepository     dbmodel.ThemesRepository
	CategoriesRepository dbmodel.CategoriesRepository
	NoticesRepository    dbmodel.NoticesRepository
	FavoritesRepository   dbmodel.FavoritesRepository
	OrdersRepository      dbmodel.OrdersRepository
}

func New() (*Config, error) {
	godotenv.Load()
	config := Config{}
	dsn := os.Getenv("DB_URL")
	databaseSession, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	database.Migrate(databaseSession)
	config.UsersRepository = dbmodel.NewUsersRepository(databaseSession)
	config.ProductsRepository = dbmodel.NewProductsRepository(databaseSession)
	config.ThemesRepository = dbmodel.NewThemesRepository(databaseSession)
	config.CategoriesRepository = dbmodel.NewCategoriesRepository(databaseSession)
	config.NoticesRepository = dbmodel.NewNoticesRepository(databaseSession)
	config.FavoritesRepository = dbmodel.NewFavoritesRepository(databaseSession)
	config.OrdersRepository = dbmodel.NewOrdersRepository(databaseSession)
	return &config, nil
}
