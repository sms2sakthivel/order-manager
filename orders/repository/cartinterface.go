package repository

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"gorm.io/gorm"
)

type CartRepository interface {
	GetAllCarts() ([]model.Cart, error)
	GetCartByID(id uint) (*model.Cart, error)
	CreateCart(order *model.Cart) error
	UpdateCart(order *model.Cart) error
	DeleteCart(id uint) error
	Init(db *gorm.DB)
}
