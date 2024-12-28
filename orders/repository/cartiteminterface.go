package repository

import "github.com/sms2sakthivel/order-manager/orders/model"

type CartItemRepository interface {
	GetAllCartItemss() ([]model.CartItem, error)
	GetCartItemByID(id uint) (*model.CartItem, error)
	CreateCartItem(order *model.CartItem) error
	CreateCartItems(order []*model.CartItem) error
	UpdateCartItem(order *model.CartItem) error
	DeleteCartItem(id uint) error
}
