package models

import (
	"github.com/lib/pq"
	"time"
)

type Product struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `json:"user_id"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Price       float64        `json:"price" gorm:"not null"`
	Stock       int            `json:"stock" gorm:"not null"`
	IsAvailable bool           `json:"is_available" gorm:"default:true"`
	CompanyName string         `json:"company_name"`
	Brand       string         `json:"brand"`
	Size        pq.StringArray `gorm:"type:text[]"`
	Images      pq.StringArray `json:"images" gorm:"type:text[]"`
	Category    string         `json:"category"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Filter struct {
	MinPrice    *float64 `form:"min_price"`
	MaxPrice    *float64 `form:"max_price"`
	IsAvailable *bool    `form:"is_available"`
	Category    string   `form:"category"`
	Brand       string   `form:"brand"`
}
