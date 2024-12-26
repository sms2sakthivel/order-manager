package repository

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
)

type OrderRepository interface {
	GetAllOrders() ([]model.Order, error)
	GetOrderByID(id uint) (*model.Order, error)
	CreateOrder(order *model.Order) error
	UpdateOrder(order *model.Order) error
	DeleteOrder(id uint) error
}
