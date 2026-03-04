package main

import (
	"BrickPlus/config"
	"BrickPlus/pkg/products"
	"BrickPlus/pkg/themes"
	"BrickPlus/pkg/users"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter();
	router.Mount("/api/users", users.Routes(configuration))
	router.Mount("/api/products", products.Routes(configuration))
	router.Mount("/api/themes", themes.Routes(configuration))
	router.Mount("/api/categories", themes.Routes(configuration))
	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	router := Routes(configuration)

	log.Println("Serving on :8080")
	http.ListenAndServe(":8080", router)
}