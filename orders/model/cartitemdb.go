package model

import "gorm.io/gorm"

// CartItem DB model
type CartItem struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
	Discount  int  `gorm:"not null"`
}

func (cartItem *CartItem) GetAPIResponseObject() *CartItemResponse {
	return &CartItemResponse{ID: cartItem.ID, ProductID: cartItem.ProductID, Discount: cartItem.Discount, Quantity: cartItem.Quantity}
}

type CartItemsMap struct {
	gorm.Model
	CartID     uint `gorm:"primaryKey"` // Foreign key to Cart
	CartItemID uint `gorm:"primaryKey"` // Foreign key to CartItem

	// Foreign key constraints
	Cart     Cart     `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
	CartItem CartItem `gorm:"foreignKey:CartItemID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}
