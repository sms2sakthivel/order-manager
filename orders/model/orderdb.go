package model

import "gorm.io/gorm"

// Order DB Model
type Order struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	CartID uint `gorm:"not null"`

	// Foreign key constraints
	Cart Cart `gorm:"forignKey:CartID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

func (order *Order) GetAPIResponseObject() *OrderResponse {
	return &OrderResponse{ID: order.ID, CartID: order.CartID, Cart: *order.Cart.GetAPIResponseObject()}
}
