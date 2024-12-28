package model

import "gorm.io/gorm"

// Cart DB Model
type Cart struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"not null"`
	CartItems []CartItem `gorm:"many2many:cart_items_map;"`
}

func (cart *Cart) GetAPIResponseObject() *CartResponse {
	var cartItemResponses []CartItemResponse = []CartItemResponse{}
	for _, cartItem := range cart.CartItems {
		cartItemResponses = append(cartItemResponses, *cartItem.GetAPIResponseObject())
	}
	return &CartResponse{ID: cart.ID, CartItems: cartItemResponses, UserID: cart.UserID}
}
