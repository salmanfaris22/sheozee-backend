package models

import "time"

type Cart struct {
	ID        uint       `gorm:"primaryKey"`
	UserID    uint       `json:"user_id"` // foreignKey User
	Items     []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `json:"cart_id"` //foreignKey Cart
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity" gorm:"default:1"`
}
