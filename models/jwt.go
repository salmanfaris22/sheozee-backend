package models

import "time"

type JWTToken struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"` //foreignKey  user
	Token     string    `json:"token" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}
