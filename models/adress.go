package models

import "time"

type Address struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `json:"user_id"` //foreignKey user
	Street    string    `json:"street" gorm:"not null"`
	City      string    `json:"city" gorm:"not null"`
	State     string    `json:"state" gorm:"not null"`
	ZipCode   string    `json:"zip_code" gorm:"not null"`
	Country   string    `json:"country" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
