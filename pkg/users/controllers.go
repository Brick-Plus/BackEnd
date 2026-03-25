package users

import (
	"BrickPlus/config"
	"BrickPlus/database/dbmodel"
	"BrickPlus/pkg/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserConfigurator struct {
	*config.Config
}

func New(configuration *config.Config) *UserConfigurator {
	return &UserConfigurator{configuration}
}

func userToModel(users []*dbmodel.User) []models.User {
	userToModel := &models.User{}
	usersEdited := []models.User{}
	for _, user := range users {
		userToModel.Name = user.Name
		userToModel.Surname = user.Surname
		userToModel.Country = user.Country
		userToModel.Email = user.Email
		userToModel.PhoneNumber = user.PhoneNumber
		userToModel.IsSubscribed = user.IsSubscribed
		userToModel.DeliveryAddress = user.DeliveryAddress
		userToModel.BillingAddress = user.BillingAddress
		usersEdited = append(usersEdited, *userToModel)
	}
	return usersEdited
}

func (config *UserConfigurator) userByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	user, err := config.UsersRepository.FindUserById(userId)
	if err != nil {
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"Error": "Failed to load the wanted user"})
		return
	}
	userEdited := userToModel(user)
	render.JSON(w, r, userEdited)
}

func (config *UserConfigurator) addUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, 403)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	addUser := &dbmodel.User{Name: req.Name, Surname: req.Surname, Country: req.Country, IsAdmin: false, Email: req.Email, PhoneNumber: req.PhoneNumber, Password: req.Password, IsSubscribed: req.IsSubscribed, DeliveryAddress: req.DeliveryAddress, BillingAddress: req.BillingAddress}
	config.UsersRepository.Create(addUser)
	render.JSON(w, r, map[string]string{"success": "New user successfully added"})
}

func (config *UserConfigurator) deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	user, err := config.UsersRepository.FindUserById(userId)
	if err != nil {
		render.Status(r, 404)
		render.JSON(w, r, map[string]string{"Error": "Failed to find the wanted user"})
		return
	}
	config.UsersRepository.Delete(user[0])
	render.JSON(w, r, map[string]string{"success": "User successfully deleted"})
}

func (config *UserConfigurator) editUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.User{}
	userId := chi.URLParam(r, "id")
	if err := render.Bind(r, req); err != nil {
		render.Status(r, 403)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	updatedUser := &dbmodel.User{Name: req.Name, Surname: req.Surname, Country: req.Country, IsAdmin: req.IsAdmin, Email: req.Email, PhoneNumber: req.PhoneNumber, Password: req.Password, IsSubscribed: req.IsSubscribed, DeliveryAddress: req.DeliveryAddress, BillingAddress: req.BillingAddress}
	config.UsersRepository.Update(updatedUser, userId)
	render.JSON(w, r, map[string]string{"success": "User successfully updated"})
}
