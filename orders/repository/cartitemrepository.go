package repository

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"gorm.io/gorm"
)

type GormCartItemRepository struct {
	DB *gorm.DB
}

func (repo *GormCartItemRepository) GetAllCartItems() ([]model.CartItem, error) {
	var cartItems []model.CartItem
	err := repo.DB.Find(&cartItems).Error
	return cartItems, err
}

func (repo *GormCartItemRepository) GetCartItemByID(id uint) (*model.CartItem, error) {
	var cartItems model.CartItem
	err := repo.DB.First(&cartItems, id).Error
	return &cartItems, err
}

func (repo *GormCartItemRepository) CreateCartItem(cartItem *model.CartItem) error {
	return repo.DB.Create(cartItem).Error
}

func (repo *GormCartItemRepository) CreateCartItems(cartItems *[]model.CartItem) error {
	return repo.DB.Create(cartItems).Error
}

func (repo *GormCartItemRepository) UpdateCartItem(cartItem *model.CartItem) error {
	return repo.DB.Save(cartItem).Error
}

func (repo *GormCartItemRepository) DeleteCartItem(id uint) error {
	return repo.DB.Delete(&model.CartItem{}, id).Error
}
