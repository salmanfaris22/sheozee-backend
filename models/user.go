package models

import (
	"time"

	"gorm.io/gorm"
)

// User model definition with GORM types and JSON tags
type User struct {
	ID                uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName         string         `json:"first_name" validate:"required,min=2,max=100"`
	LastName          string         `json:"last_name" validate:"required,min=2,max=100"`
	Email             string         `json:"email" validate:"required,email"`
	Password          string         `json:"password" validate:"required,min=8,max=255"`
	Phone             string         `json:"phone" validate:"required,len=15"`
	Role              string         `gorm:"type:varchar(50);default:'user'" json:"role"`
	ProfilePictureURL string         `gorm:"type:varchar(255)" json:"profile_picture_url,omitempty"`
	IsVerified        bool           `gorm:"default:false" json:"is_verified"`
	IsActive          bool           `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Cart              Cart           `gorm:"foreignKey:UserID" json:"cart"`
	Orders            []Order        `gorm:"foreignKey:UserID" json:"orders"`
	Addresses         []Address      `gorm:"foreignKey:UserID" json:"addresses"`
	Wishlist          Wishlist       `gorm:"foreignKey:UserID" json:"wishlist"`
	Products          []Product      `gorm:"foreignKey:UserID" json:"products"`
	Reviews           []Review       `gorm:"foreignKey:UserID" json:"reviews"`
	Tokens            []JWTToken     `gorm:"foreignKey:UserID" json:"token"`
}
