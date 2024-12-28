package repository

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	DB *gorm.DB
}

func (repo *GormOrderRepository) GetAllOrders() ([]model.Order, error) {
	var orders []model.Order
	err := repo.DB.Preload("Cart").Preload("Cart.CartItems").Preload("Cart.CartItems.Product").Preload("Cart.User").Find(&orders).Error
	return orders, err
}

func (repo *GormOrderRepository) GetOrderByID(id uint) (*model.Order, error) {
	var order model.Order
	err := repo.DB.Preload("Cart").Preload("Cart.CartItems").Preload("Cart.CartItems.Product").Preload("Cart.User").First(&order, id).Error
	return &order, err
}

func (repo *GormOrderRepository) CreateOrder(order *model.Order) error {
	return repo.DB.Create(order).Error
}

func (repo *GormOrderRepository) UpdateOrder(order *model.Order) error {
	return repo.DB.Save(order).Error
}

func (repo *GormOrderRepository) DeleteOrder(id uint) error {
	return repo.DB.Delete(&model.Order{}, id).Error
}
