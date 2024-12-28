package model

import (
	products "github.com/sms2sakthivel/product-manager/products/model"
	users "github.com/sms2sakthivel/user-manager/users/model"
	"gorm.io/gorm"
)

// Order DB Model
type Order struct {
	ID     uint `gorm:"primaryKey"`
	CartID uint `gorm:"not null"`

	// Foreign key constraints
	Cart Cart `gorm:"forignKey:CartID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

func (order *Order) GetAPIResponseObject() *OrderResponse {
	return &OrderResponse{ID: order.ID, CartID: order.CartID, Cart: &order.Cart}
}

// Cart DB Model
type Cart struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `gorm:"not null"`
	CartItems []CartItem `gorm:"many2many:cart_items_map;"`

	// Foreign key constraints
	User users.User `gorm:"forignKey:UserID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

func (cart *Cart) GetAPIResponseObject() *CartResponse {
	var cartItemResponses []CartItemResponse = []CartItemResponse{}
	for _, cartItem := range cart.CartItems {
		cartItemResponses = append(cartItemResponses, *cartItem.GetAPIResponseObject())
	}
	return &CartResponse{ID: cart.ID, CartItems: cartItemResponses, User: cart.User}
}

type CartItem struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
	Discount  int  `gorm:"not null"`

	// Foreign key constraints
	Product products.Product `gorm:"forignKey:ProductID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

func (cartItem *CartItem) GetAPIResponseObject() *CartItemResponse {
	return &CartItemResponse{ID: cartItem.ID, ProductID: cartItem.ProductID, Discount: cartItem.Discount, Product: &cartItem.Product}
}

type CartItemsMap struct {
	CartID     uint `gorm:"primaryKey"` // Foreign key to Cart
	CartItemID uint `gorm:"primaryKey"` // Foreign key to CartItem

	// Foreign key constraints
	Cart     Cart     `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
	CartItem CartItem `gorm:"foreignKey:CartItemID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}
