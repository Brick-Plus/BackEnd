package orders

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type OrdersConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *OrdersConfigurator {
	return &OrdersConfigurator{configuration}
}

func orderToModel(orders []*dbmodel.Order) []models.Order {
	ordersEdited := []models.Order{}
	for _, order := range orders {
		productIDs := []int{}
		for _, product := range order.Products {
			productIDs = append(productIDs, int(product.ID))
		}
		orderEdited := models.Order{
			Reference:  order.Reference,
			IdUser:     order.IdUser,
			ProductIDs: productIDs,
			ExpDate:    order.ExpDate,
			OrderDate:  order.OrderDate,
			PriceTTC:   order.PriceTTC,
			Devise:     order.Devise,
		}
		ordersEdited = append(ordersEdited, orderEdited)
	}
	return ordersEdited
}

func (config *OrdersConfigurator) orderByReferenceHandler(w http.ResponseWriter, r *http.Request) {
	ref := chi.URLParam(r, "ref")
	order, err := config.OrdersRepository.FindOrderByReference(ref)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Order not found"})
		return
	}
	render.JSON(w, r, orderToModel([]*dbmodel.Order{order}))
}

func (config *OrdersConfigurator) ordersByUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	orders, err := config.OrdersRepository.FindOrdersByUserId(userId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to load orders for this user"})
		return
	}
	render.JSON(w, r, orderToModel(orders))
}

func (config *OrdersConfigurator) addOrderHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Order{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	// Existence checks
	user, err := config.UsersRepository.FindUserById(strconv.Itoa(req.IdUser))
	if err != nil || len(user) == 0 {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	// Fetch products
	products := []dbmodel.Product{}
	for _, pid := range req.ProductIDs {
		p, err := config.ProductsRepository.FindProductById(strconv.Itoa(pid))
		if err != nil || len(p) == 0 {
			render.JSON(w, r, map[string]string{"error": "Product not found: " + strconv.Itoa(pid)})
			return
		}
		products = append(products, *p[0])
	}

	newOrder := &dbmodel.Order{
		Reference:  req.Reference,
		IdUser:     req.IdUser,
		Products:   products,
		ExpDate:    req.ExpDate,
		OrderDate:  req.OrderDate,
		PriceTTC:   req.PriceTTC,
		Devise:     req.Devise,
	}
	_, err = config.OrdersRepository.Create(newOrder)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create order"})
		return
	}
	render.JSON(w, r, map[string]string{"success": "Order successfully created"})
}

func (config *OrdersConfigurator) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	ref := chi.URLParam(r, "ref")
	order, err := config.OrdersRepository.FindOrderByReference(ref)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Order not found"})
		return
	}
	config.OrdersRepository.Delete(order)
	render.JSON(w, r, map[string]string{"success": "Order successfully deleted"})
}
