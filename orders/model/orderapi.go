package model

// Order API Models
type OrderCreateRequest struct {
	CartID uint `json:"cart_id" binding:"required"`
}

type OrderUpdateRequest struct {
	ID     uint `json:"order_id"`
	CartID uint `json:"cart_id"`
}

type OrderResponse struct {
	ID     uint         `json:"order_id"`
	CartID uint         `json:"cart_id"`
	Cart   CartResponse `json:"cart,omitempty"`
}

func (ocr *OrderCreateRequest) GetOrderDBObject() *Order {
	return &Order{CartID: ocr.CartID}
}

func (our *OrderUpdateRequest) GetOrderDBObject() *Order {
	return &Order{ID: our.ID, CartID: our.CartID}
}
