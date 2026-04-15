package dbmodel

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Surname string `json:"surname"`
	Country string `json:"country"`
	ProfilPicture string `json:"pp_url"`
	IsAdmin bool `json:"is_admin"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
	IsSubscribed bool `json:"is_subscribed"`
	DeliveryAddress string `json:"delivery_address"`
	BillingAddress string `json:"billing_address"`
}

type UsersRepository interface{
	Create(newUser *User)(*User, error)
	FindUserById(userId string)([]*User, error)
	FindUserByEmail(email string)(*User, error)
	Delete(userToDelete *User) error
	Update(userToUpdate *User, userId string) error
}

type usersRepository struct{
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &usersRepository{db: db}
}

func (r *usersRepository) Create(user *User) (*User, error){
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *usersRepository) FindUserById(userId string)([]*User, error){
	var entry []*User
	if err := r.db.Where("id = ?", userId).Find(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *usersRepository) FindUserByEmail(email string) (*User, error) {
	var entry User
	if err := r.db.Where("email = ?", email).First(&entry).Error; err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *usersRepository) Delete(userToDelete *User) error {
	if userToDelete.IsAdmin {
		return errors.New("cannot delete an admin user")
	}
	if err := r.db.Delete(userToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *usersRepository) Update(userToUpdate *User, userId string) error {
	if err := r.db.Where("id = ?", userId).Updates(userToUpdate).Error; err != nil {
		return err
	}
	return nil
}