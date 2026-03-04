package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/buy", buyHandler)
	http.HandleFunc("/success", successHandler)
	http.HandleFunc("/add-to-cart", addToCartHandler)
	http.HandleFunc("/cart", cartHandler)
	http.HandleFunc("/remove-from-cart", removeFromCartHandler)
	// Sert les fichiers statiques (CSS)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	log.Println("Serveur lancé sur :8081 ...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
