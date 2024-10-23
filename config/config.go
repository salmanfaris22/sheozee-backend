package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main/models"
)

var db *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=poomon dbname=ecommerce port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to the database", err)
		return
	}

	db.AutoMigrate(
		&models.User{},
		&models.Address{},
		&models.Cart{},
		&models.CartItem{},
		&models.JWTToken{},
		&models.Order{},
		&models.Product{},
		&models.Review{},
		&models.Review{},
		&models.Wishlist{},
		&models.WishlistItem{},
	)

}

func GetDB() *gorm.DB {
	return db
}
