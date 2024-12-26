package model

import users "github.com/sms2sakthivel/user-manager/users/model"

type OrderCreateRequest struct {
	CartID uint `json:"cart_id" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}

type OrderResponse struct {
	ID     uint        `json:"order_id"`
	CartID uint        `json:"cart_id"`
	UserID uint        `json:"user_id"`
	Cart   *Cart       `json:"cart,omitempty"`
	User   *users.User `json:"user,omitempty"`
}
