package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `json:"user_id"` //foreignKey  user
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null"`
	Stock       int       `json:"stock" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
