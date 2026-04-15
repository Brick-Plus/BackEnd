package main

import (
	"BrickPlus/config"
	"BrickPlus/pkg/notices"
	"BrickPlus/pkg/categories"
	"BrickPlus/pkg/products"
	"BrickPlus/pkg/themes"
	"BrickPlus/pkg/users"
	"BrickPlus/pkg/favorites"
	"BrickPlus/pkg/orders"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter();
	router.Mount("/api/v1/users", users.Routes(configuration))
	router.Mount("/api/v1/products", products.Routes(configuration))
	router.Mount("/api/v1/themes", themes.Routes(configuration))
	router.Mount("/api/v1/categories", categories.Routes(configuration))
	router.Mount("/api/v1/notices", notices.Routes(configuration))
	router.Mount("/api/v1/favorites", favorites.Routes(configuration))
	router.Mount("/api/v1/orders", orders.Routes(configuration))
	router.Handle("/config.openapi.yaml", http.FileServer(http.Dir("./")))
	fs := http.FileServer(http.Dir("./swagger-ui"))
	router.Handle("/*", fs)

	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Println("Serving on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server error:", err)
	}
}

