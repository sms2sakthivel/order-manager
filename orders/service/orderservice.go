package service

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"github.com/sms2sakthivel/order-manager/orders/repository"
	"gorm.io/gorm"
)

type OrderService struct {
	CartService CartService
	Repo        repository.OrderRepository
}

func (os *OrderService) InitCartService(db *gorm.DB) {
	os.CartService.Init(db)
}
func (s *OrderService) GetOrders() ([]model.OrderResponse, error) {
	var orders []model.OrderResponse = []model.OrderResponse{}
	dbOrders, err := s.Repo.GetAllOrders()
	if err != nil {
		return nil, err
	}
	for _, order := range dbOrders {
		orders = append(orders, *order.GetAPIResponseObject())
	}
	return orders, nil
}

func (s *OrderService) GetOrder(id uint) (*model.OrderResponse, error) {
	dbOrder, err := s.Repo.GetOrderByID(id)
	if err != nil {
		return nil, err
	}
	return dbOrder.GetAPIResponseObject(), nil
}

func (s *OrderService) CreateOrder(orderReq *model.OrderCreateRequest) (*model.OrderResponse, error) {
	order := orderReq.GetOrderDBObject()
	err := s.Repo.CreateOrder(order)
	return order.GetAPIResponseObject(), err
}

func (s *OrderService) UpdateOrder(orderReq *model.OrderUpdateRequest) (*model.OrderResponse, error) {
	order := orderReq.GetOrderDBObject()
	err := s.Repo.UpdateOrder(order)
	return order.GetAPIResponseObject(), err
}

func (s *OrderService) DeleteOrder(id uint) error {
	return s.Repo.DeleteOrder(id)
}
