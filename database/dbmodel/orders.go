package dbmodel

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Reference string    `json:"reference" gorm:"type:varchar(255);uniqueIndex"`
	IdUser    int       `json:"id_user"`
	User      User      `gorm:"foreignKey:IdUser"`
	Products  []Product `gorm:"many2many:order_products;"`
	ExpDate   time.Time `json:"exp_date"`
	OrderDate time.Time `json:"order_date"`
	PriceTTC  float64   `json:"price_ttc"`
	Devise    string    `json:"devise"`
}

type OrdersRepository interface {
	Create(newOrder *Order) (*Order, error)
	FindOrdersByUserId(userId string) ([]*Order, error)
	FindOrderByReference(reference string) (*Order, error)
	Delete(orderToDelete *Order) error
	Update(orderToUpdate *Order, reference string) error
}

type ordersRepository struct {
	db *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) OrdersRepository {
	return &ordersRepository{db: db}
}

func (r *ordersRepository) Create(order *Order) (*Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *ordersRepository) FindOrdersByUserId(userId string) ([]*Order, error) {
	var entries []*Order
	if err := r.db.Preload("Products").Where("id_user = ?", userId).Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *ordersRepository) FindOrderByReference(reference string) (*Order, error) {
	var entry Order
	if err := r.db.Preload("Products").Where("reference = ?", reference).First(&entry).Error; err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *ordersRepository) Delete(orderToDelete *Order) error {
	if err := r.db.Delete(orderToDelete).Error; err != nil {
		return err
	}
	return nil
}

func (r *ordersRepository) Update(orderToUpdate *Order, reference string) error {
	if err := r.db.Where("reference = ?", reference).Updates(orderToUpdate).Error; err != nil {
		return err
	}
	return nil
}
