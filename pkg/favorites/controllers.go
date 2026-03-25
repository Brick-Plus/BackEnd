package favorites

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type FavoritesConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *FavoritesConfigurator {
	return &FavoritesConfigurator{configuration}
}

func favoriteToModel(favorites []*dbmodel.Notice) []models.Notice {
	// Not needed as favorite DTO is very simple, 
	// but kept for consistency if ever expanded.
	return nil
}

// favoriteToModel for favorites specifically
func favToModel(favorites []*dbmodel.Favorite) []models.Favorite {
	favsEdited := []models.Favorite{}
	for _, fav := range favorites {
		favEdited := models.Favorite{
			IdUser:    fav.IdUser,
			IdProduct: fav.IdProduct,
		}
		favsEdited = append(favsEdited, favEdited)
	}
	return favsEdited
}

func (config *FavoritesConfigurator) favoriteByIdHandler(w http.ResponseWriter, r *http.Request) {
	favId := chi.URLParam(r, "id")
	favorites, err := config.FavoritesRepository.FindFavoriteById(favId)
	if err != nil || len(favorites) == 0 {
		render.JSON(w, r, map[string]string{"error": "Favorite not found"})
		return
	}
	render.JSON(w, r, favToModel(favorites))
}

func (config *FavoritesConfigurator) favoritesByUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	favorites, err := config.FavoritesRepository.FindFavoritesByUserId(userId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to load favorites for this user"})
		return
	}
	render.JSON(w, r, favToModel(favorites))
}

func (config *FavoritesConfigurator) addFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.Favorite{}
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
	product, err := config.ProductsRepository.FindProductById(strconv.Itoa(req.IdProduct))
	if err != nil || len(product) == 0 {
		render.JSON(w, r, map[string]string{"error": "Product not found"})
		return
	}

	// Double check if already favorited
	existing, err := config.FavoritesRepository.FindByUserAndProduct(req.IdUser, req.IdProduct)
	if err == nil && len(existing) > 0 {
		render.JSON(w, r, map[string]string{"error": "Product already in favorites"})
		return
	}

	newFav := &dbmodel.Favorite{
		IdUser:    req.IdUser,
		IdProduct: req.IdProduct,
	}
	_, err = config.FavoritesRepository.Create(newFav)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to add favorite"})
		return
	}
	render.JSON(w, r, map[string]string{"success": "Favorite successfully added"})
}

func (config *FavoritesConfigurator) deleteFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	favId := chi.URLParam(r, "id")
	favorites, err := config.FavoritesRepository.FindFavoriteById(favId)
	if err != nil || len(favorites) == 0 {
		render.JSON(w, r, map[string]string{"error": "Favorite not found"})
		return
	}
	config.FavoritesRepository.Delete(favorites[0])
	render.JSON(w, r, map[string]string{"success": "Favorite successfully removed"})
}
