package service

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"github.com/sms2sakthivel/order-manager/orders/repository"
	"gorm.io/gorm"
)

type CartService struct {
	Repo repository.CartRepository
}

func (cs *CartService) Init(db *gorm.DB) {
	cs.Repo = &repository.GormCartRepository{}
	cs.Repo.Init(db)
}

func (s *CartService) GetCarts() ([]model.CartResponse, error) {
	dbCarts, err := s.Repo.GetAllCarts()
	if err != nil {
		return nil, err
	}
	var carts []model.CartResponse = []model.CartResponse{}
	for _, cart := range dbCarts {
		carts = append(carts, *cart.GetAPIResponseObject())
	}
	return carts, nil
}

func (s *CartService) GetCart(id uint) (*model.CartResponse, error) {
	dbCart, err := s.Repo.GetCartByID(id)
	if err != nil {
		return nil, err
	}
	return dbCart.GetAPIResponseObject(), err
}

func (s *CartService) CreateCart(cartReq *model.CartCreateRequest) (*model.CartResponse, error) {
	_, err := GetUserByID(cartReq.UserID)
	if err != nil {
		return nil, err
	}
	cart := cartReq.GetDBObject()
	var productIds []uint = []uint{}
	for _, cartItem := range cart.CartItems {
		productIds = append(productIds, cartItem.ProductID)
	}
	for _, productID := range productIds {
		_, err = GetProductByID(productID)
		if err != nil {
			return nil, err
		}
	}
	err = s.Repo.CreateCart(cart)
	return cart.GetAPIResponseObject(), err
}

func (s *CartService) UpdateCart(cartReq *model.CartUpdateRequest) (*model.CartResponse, error) {
	_, err := GetUserByID(cartReq.UserID)
	if err != nil {
		return nil, err
	}
	cart := cartReq.GetDBObject()
	var productIds []uint = []uint{}
	for _, cartItem := range cart.CartItems {
		productIds = append(productIds, cartItem.ProductID)
	}
	for _, productID := range productIds {
		_, err = GetProductByID(productID)
		if err != nil {
			return nil, err
		}
	}
	err = s.Repo.UpdateCart(cart)
	return cart.GetAPIResponseObject(), err
}

func (s *CartService) DeleteCart(id uint) error {
	return s.Repo.DeleteCart(id)
}
