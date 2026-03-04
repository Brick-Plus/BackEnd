package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
)

var homeTpl = template.Must(template.New("home.html").Funcs(template.FuncMap{
	"euro": func(cents int) string {
		return fmt.Sprintf("%.2f", float64(cents)/100)
	},
	"formatPrice": func(cents int64, currency string) string {
		return fmt.Sprintf("%.2f %s", float64(cents)/100, strings.ToUpper(currency))
	},
	"upper": func(s string) string {
		return strings.ToUpper(s)
	},
}).ParseFiles("templates/home.html"))
var cartTpl = template.Must(template.New("cart.html").Funcs(template.FuncMap{
	"euro": func(cents int) string {
		return fmt.Sprintf("%.2f", float64(cents)/100)
	},
}).ParseFiles("templates/cart.html"))

var successTpl = template.Must(template.ParseFiles("templates/success.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	notif := ""
	if c, err := r.Cookie("notif"); err == nil {
		notif = c.Value
		http.SetCookie(w, &http.Cookie{Name: "notif", Value: "", Path: "/", MaxAge: -1})
	}
	products, err := getStripeProducts()
	if err != nil {
		fmt.Println("Erreur Stripe:", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
		return
	}
	// Récupère le thème sélectionné dans l'URL
	selectedTheme := r.URL.Query().Get("theme")
	var filtered []StripeProduct
	if selectedTheme != "" {
		for _, p := range products {
			if strings.EqualFold(p.Theme, selectedTheme) {
				filtered = append(filtered, p)
			}
		}
	} else {
		filtered = products
	}
	// Récupère toutes les valeurs de thème pour la navbar
	themeSet := map[string]struct{}{}
	for _, p := range products {
		if p.Theme != "" {
			themeSet[p.Theme] = struct{}{}
		}
	}
	var themes []string
	for t := range themeSet {
		themes = append(themes, t)
	}
	data := struct {
		Products      []StripeProduct
		Notif         string
		Themes        []string
		SelectedTheme string
	}{filtered, notif, themes, selectedTheme}
	if err := homeTpl.ExecuteTemplate(w, "home.html", data); err != nil {
		fmt.Println("Erreur template home.html:", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}

// Handler panier (cart)
func cartHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cart")
	cart := make(map[string]int)
	if err == nil {
		cart = parseCart(cookie.Value)
	}
	type CartItem struct {
		Name  string
		Qty   int
		Total int
	}
	var items []CartItem
	total := 0
	for id, qty := range cart {
		for _, a := range articles {
			if fmt.Sprintf("%d", a.ID) == id {
				items = append(items, CartItem{Name: a.Name, Qty: qty, Total: a.Price * qty})
				total += a.Price * qty
			}
		}
	}
	data := struct {
		Items []CartItem
		Total int
	}{items, total}
	if err := cartTpl.ExecuteTemplate(w, "cart.html", data); err != nil {
		fmt.Println("Erreur template cart.html:", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}

type StripeProduct struct {
	ID       string
	Name     string
	Price    int64 // en centimes
	Currency string
	Stock    int64 // à récupérer via metadata ou autre
	Active   bool
	Theme    string
	ImageURL string
}

// Récupère les produits Stripe
func getStripeProducts() ([]StripeProduct, error) {
	stripe.Key = "sk_test_51SsJ75E2CJz60xIxFoq0o8X7vQPU1sYVytuY4qmTTu9kDb9TcadkvePvznX6aOaniJvDOOe5DJeHQQog5MVaCOJN00uB2rjXBF"
	var products []StripeProduct
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("active", "", "true")
	i := product.List(params)
	for i.Next() {
		p := i.Product()
		// Récupère le prix associé au produit
		priceParams := &stripe.PriceListParams{Product: stripe.String(p.ID)}
		priceParams.Filters.AddFilter("active", "", "true")
		pr := price.List(priceParams)
		var unitPrice int64
		var currency string
		if pr.Next() {
			unitPrice = pr.Price().UnitAmount
			currency = string(pr.Price().Currency)
		}
		stock := int64(0)
		if val, ok := p.Metadata["stock"]; ok {
			s, _ := strconv.ParseInt(val, 10, 64)
			stock = s
		}
		theme := ""
		if val, ok := p.Metadata["Thème"]; ok {
			theme = val
		}
		imageURL := ""
		if len(p.Images) > 0 {
			imageURL = p.Images[0]
		}
		products = append(products, StripeProduct{
			ID:       p.ID,
			Name:     p.Name,
			Price:    unitPrice,
			Currency: currency,
			Stock:    stock,
			Active:   p.Active,
			Theme:    theme,
			ImageURL: imageURL,
		})
	}
	return products, nil
}

func buyHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cart")
	cart := make(map[string]int)
	if err == nil {
		cart = parseCart(cookie.Value)
	}
	var lineItems []*stripe.CheckoutSessionLineItemParams
	for id, qty := range cart {
		for _, a := range articles {
			if fmt.Sprintf("%d", a.ID) == id && qty > 0 {
				lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						Currency: stripe.String("eur"),
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String(a.Name),
						},
						UnitAmount: stripe.Int64(int64(a.Price)),
					},
					Quantity: stripe.Int64(int64(qty)),
				})
			}
		}
	}
	if len(lineItems) == 0 {
		http.Redirect(w, r, "/cart", http.StatusSeeOther)
		return
	}
	stripe.Key = "sk_test_51SsJ75E2CJz60xIxFoq0o8X7vQPU1sYVytuY4qmTTu9kDb9TcadkvePvznX6aOaniJvDOOe5DJeHQQog5MVaCOJN00uB2rjXBF"
	domain := "http://localhost:8081"
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
		Mode:               stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:         stripe.String(domain + "/success"),
		CancelURL:          stripe.String(domain + "/cart"),
	}
	sess, err := session.New(params)
	if err != nil {
		http.Error(w, "Erreur Stripe", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, sess.URL, http.StatusSeeOther)
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := successTpl.Execute(w, nil); err != nil {
		fmt.Println("Erreur template success.html:", err)
		http.Error(w, "Erreur interne", http.StatusInternalServerError)
	}
}
