package model

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
	CartItems []CartItemResponse `json:"cart_items"`
	UserID    uint               `json:"user_id"`
	CartValue uint               `json:"cart_value"`
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
