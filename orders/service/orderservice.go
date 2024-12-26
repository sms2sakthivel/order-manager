package service

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"github.com/sms2sakthivel/order-manager/orders/repository"
)

type OrderService struct {
	Repo repository.OrderRepository
}

func (s *OrderService) GetOrders() ([]model.Order, error) {
	return s.Repo.GetAllOrders()
}

func (s *OrderService) GetOrder(id uint) (*model.Order, error) {
	return s.Repo.GetOrderByID(id)
}

func (s *OrderService) CreateOrder(order *model.Order) error {
	return s.Repo.CreateOrder(order)
}

func (s *OrderService) UpdateOrder(order *model.Order) error {
	return s.Repo.UpdateOrder(order)
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.Repo.DeleteOrder(id)
}
