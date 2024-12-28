package model

import (
	products "github.com/sms2sakthivel/product-manager/products/model"
	users "github.com/sms2sakthivel/user-manager/users/model"
)

// Order API Models
type OrderCreateRequest struct {
	CartID uint `json:"cart_id" binding:"required"`
}

type OrderUpdateRequest struct {
	ID     uint `json:"order_id"`
	CartID uint `json:"cart_id"`
}

type OrderResponse struct {
	ID     uint  `json:"order_id"`
	CartID uint  `json:"cart_id"`
	Cart   *Cart `json:"cart,omitempty"`
}

func (ocr *OrderCreateRequest) GetOrderDBObject() *Order {
	return &Order{CartID: ocr.CartID}
}

func (our *OrderUpdateRequest) GetOrderDBObject() *Order {
	return &Order{ID: our.ID, CartID: our.CartID}
}

// Cart API Models
type CartCreateRequest struct {
	UserID    uint              `json:"user_id" binding:"required"`
	CartItems []CartItemRequest `json:"cart_items" binding:"required"`
}

type CartUpdateRequest struct {
	ID        uint              `json:"cart_id" binding:"required"`
	UserID    uint              `json:"user_id" binding:"required"`
	CartItems []CartItemRequest `json:"cart_items" binding:"required"`
}

type CartResponse struct {
	ID        uint               `json:"cart_id"`
	CartItems []CartItemResponse `json:"cart_items" binding:"required"`
	User      users.User         `json:"user" binding:"required"`
}

func (ccr *CartCreateRequest) GetDBObject() *Cart {
	var cartItems []CartItem
	for _, item := range ccr.CartItems {
		cartItems = append(cartItems, CartItem{ProductID: item.ProductID, Quantity: item.Quantity, Discount: item.Discount})
	}
	return &Cart{UserID: ccr.UserID, CartItems: cartItems}
}

func (cur *CartUpdateRequest) GetDBObject() *Cart {
	var cartItems []CartItem
	for _, item := range cur.CartItems {
		cartItems = append(cartItems, CartItem{ProductID: item.ProductID, Quantity: item.Quantity, Discount: item.Discount})
	}
	return &Cart{ID: cur.ID, UserID: cur.UserID, CartItems: cartItems}
}

// CartItem API Models
type CartItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
	Discount  int  `json:"discount" binding:"required"`
}

type CartItemResponse struct {
	ID        uint `json:"cart_item_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
	Discount  int  `json:"discount"`
	Product   *products.Product
}

func (cir *CartItemRequest) GetDBObject() *CartItem {
	return &CartItem{ProductID: cir.ProductID, Quantity: cir.Quantity, Discount: cir.Discount}
}
