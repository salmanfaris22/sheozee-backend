package models

import (
	"fmt"
	"main/services"
	"time"

	"gorm.io/gorm"
)

// User model definition with GORM types and JSON tags
type User struct {
	ID                uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email             string         `gorm:"unique;type:varchar(100)" json:"email" validate:"required,email"`
	FirstName         string         `json:"first_name" validate:"required,min=2,max=100"`
	LastName          string         `json:"last_name" validate:"required,min=2,max=100"`
	Password          string         `json:"password" validate:"required,min=8,max=255"`
	Phone             string         `json:"phone" validate:"required,min=10"`
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
	Token             string         `gorm:"type:varchar(255);default:null" json:"token,omitempty"`
	//RefreshToken      string         `json:"refresh_token,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	str, err := services.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = str
	fmt.Println("sss")
	return
}

func (u *User) SetAccessToken() (string, error) {
	token, err := services.GenerateAccessToken(u.Email)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *User) SetRefreshToken() (string, error) {
	token, err := services.GenerateRefreshToken(u.Email)
	if err != nil {
		return "", err
	}

	return token, err
}
