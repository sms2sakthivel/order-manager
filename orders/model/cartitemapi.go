package model

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
}

func (cir *CartItemRequest) GetDBObject() *CartItem {
	return &CartItem{ProductID: cir.ProductID, Quantity: cir.Quantity, Discount: cir.Discount}
}
