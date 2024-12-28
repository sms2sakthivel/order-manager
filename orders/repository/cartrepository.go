package repository

import (
	"github.com/sms2sakthivel/order-manager/orders/model"
	"gorm.io/gorm"
)

type GormCartRepository struct {
	CartItemRepo GormCartItemRepository
	DB           *gorm.DB
}

func (gcr *GormCartRepository) Init(db *gorm.DB) {
	gcr.DB = db
	gcr.CartItemRepo = GormCartItemRepository{DB: db}
}

func (repo *GormCartRepository) GetAllCarts() ([]model.Cart, error) {
	var carts []model.Cart
	err := repo.DB.Preload("CartItems").Find(&carts).Error
	return carts, err
}

func (repo *GormCartRepository) GetCartByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	err := repo.DB.Preload("CartItems").First(&cart, id).Error
	return &cart, err
}

func (repo *GormCartRepository) CreateCart(cart *model.Cart) error {
	err := repo.CartItemRepo.CreateCartItems(&cart.CartItems)
	if err != nil {
		return err
	}
	return repo.DB.Create(cart).Error
}

func (repo *GormCartRepository) UpdateCart(cart *model.Cart) error {
	return repo.DB.Save(cart).Error
}

func (repo *GormCartRepository) DeleteCart(id uint) error {
	return repo.DB.Delete(&model.Cart{}, id).Error
}
