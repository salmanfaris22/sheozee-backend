package models

import "time"

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"` //foreignKey
	Total     float64   `json:"total" gorm:"not null"`
	Status    string    `json:"status" gorm:"default:'pending'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
