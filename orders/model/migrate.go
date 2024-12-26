package model

import (
	products "github.com/sms2sakthivel/product-manager/products/model"
	users "github.com/sms2sakthivel/user-manager/users/model"
	"gorm.io/gorm"
)

func Automigrate(DB *gorm.DB) error {
	err := users.Automigrate(DB)
	if err != nil {
		return err
	}
	err = products.Automigrate(DB)
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&CartItem{}, &Cart{}, &Order{})
}
