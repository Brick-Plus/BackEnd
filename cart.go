package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func removeFromCartHandler(w http.ResponseWriter, r *http.Request) {
   if r.Method != http.MethodPost {
	   http.Redirect(w, r, "/cart", http.StatusSeeOther)
	   return
   }
   id := r.FormValue("id")
   cookie, err := r.Cookie("cart")
   cart := make(map[string]int)
   if err == nil {
	   cart = parseCart(cookie.Value)
   }
   delete(cart, id)
   setCartCookie(w, cart)
   http.Redirect(w, r, "/cart", http.StatusSeeOther)
}

func addToCartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	id := r.FormValue("id")
	qty := r.FormValue("qty")
	cookie, err := r.Cookie("cart")
	cart := make(map[string]int)
	if err == nil {
		cart = parseCart(cookie.Value)
	}
	cart[id] += atoi(qty)
	setCartCookie(w, cart)
	http.SetCookie(w, &http.Cookie{
		Name:   "notif",
		Value:  id + ":" + qty,
		Path:   "/",
		MaxAge: 5, // 5 secondes
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func parseCart(val string) map[string]int {
	cart := make(map[string]int)
	pairs := strings.Split(val, ",")
	for _, pair := range pairs {
		if pair == "" {
			continue
		}
		kv := strings.Split(pair, ":")
		if len(kv) == 2 {
			cart[kv[0]] = atoi(kv[1])
		}
	}
	return cart
}

func setCartCookie(w http.ResponseWriter, cart map[string]int) {
	var pairs []string
	for id, qty := range cart {
		if qty > 0 {
			pairs = append(pairs, id+":"+fmt.Sprintf("%d", qty))
		}
	}
	val := strings.Join(pairs, ",")
	http.SetCookie(w, &http.Cookie{
		Name:   "cart",
		Value:  val,
		Path:   "/",
		MaxAge: 3600 * 24,
	})
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	if n < 1 {
		n = 1
	}
	return n
}
