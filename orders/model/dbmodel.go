package model

import (
	products "github.com/sms2sakthivel/product-manager/products/model"
	users "github.com/sms2sakthivel/user-manager/users/model"
)

type Order struct {
	ID     uint `gorm:"primaryKey"`
	CartID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`

	// Foreign key constraints
	Cart Cart       `gorm:"forignKey:CartID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
	User users.User `gorm:"forignKey:UserID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

type Cart struct {
	ID        uint       `gorm:"primaryKey"`
	CartItems []CartItem `gorm:"many2many:cart_items_map;"`
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
	Discount  int  `gorm:"not null"`

	// Foreign key constraints
	Product products.Product `gorm:"forignKey:ProductID;reference:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;not null"`
}

type CartItemsMap struct {
	CartID     uint `gorm:"primaryKey"` // Foreign key to Cart
	CartItemID uint `gorm:"primaryKey"` // Foreign key to CartItem

	// Foreign key constraints
	Cart     Cart     `gorm:"foreignKey:CartID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
	CartItem CartItem `gorm:"foreignKey:CartItemID;references:ID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE"`
}
