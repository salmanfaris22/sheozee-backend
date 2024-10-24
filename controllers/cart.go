package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	myerror "main/error"
	"main/models"
	"net/http"
	"strconv"
)

type Cart struct {
}

func (c Cart) AddToCart(ctx *gin.Context) {
	userID, err := ctx.Cookie("userId")
	if err != nil {
		myerror.Errors(ctx, err, "Cat This add toCart", http.StatusInternalServerError)
		return
	}
	var input struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	err = ctx.BindJSON(&input)
	if err != nil {
		myerror.Errors(ctx, err, "can't This add toCart", http.StatusInternalServerError)
		return
	}
	var product models.Product
	err = db.First(&product, input.ProductID).Error
	if err != nil || input.Quantity <= 0 {
		myerror.Errors(ctx, err, "cant find Product", http.StatusInternalServerError)
		return
	}
	if product.Stock == 0 || product.IsAvailable == false {
		fmt.Println("kn")
		myerror.Errors(ctx, errors.New("out of stock"), "stoke not uvalible", http.StatusBadRequest)
		return
	}
	var cart models.Cart
	err = db.Where("user_id=?", userID).First(&cart).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			conId, err := strconv.Atoi(userID)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				myerror.Errors(ctx, err, "Error converting string to int:", http.StatusInternalServerError)
				return
			}
			cart.UserID = uint(conId)
			err = db.Create(&cart).Error
			if err != nil {
				myerror.Errors(ctx, err, "cart creat error", http.StatusInternalServerError)
				return
			}
		} else {
			myerror.Errors(ctx, err, "Unable to fetch cart", http.StatusInternalServerError)
			return
		}
	}

	if input.Quantity <= 0 {
		myerror.Errors(ctx, err, "invalid quandty", http.StatusInternalServerError)
		return
	}

	var cartItem models.CartItem
	err = db.Where("cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&cartItem).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			cartItem.Quantity = input.Quantity
			cartItem.ProductID = input.ProductID
			cartItem.CartID = cart.ID
			err = db.Create(&cartItem).Error
			if err != nil {
				myerror.Errors(ctx, err, "cart creat error", http.StatusInternalServerError)
				return
			}
			ctx.JSON(200, gin.H{
				"message": "Cart item added successfully",
			})
			return
		} else {
			myerror.Errors(ctx, err, "Unable to fetch cart model", http.StatusInternalServerError)
			return
		}
	} else {
		cartItem.Quantity += cartItem.Quantity
		err = db.Save(&cartItem).Error
		if err != nil {
			myerror.Errors(ctx, err, "cann't save quatry", http.StatusInternalServerError)
			return
		}
		ctx.JSON(200, gin.H{
			"message": "Cart item updated successfully",
		})
	}

}

func (c Cart) GetCartItems(ctx *gin.Context) {
	userID, err := ctx.Cookie("userId")
	if err != nil {
		myerror.Errors(ctx, err, "Cat This add toCart", http.StatusInternalServerError)
		return
	}
	var carts models.Cart
	err = db.Where("user_id=?", userID).First(&carts).Error
	if err != nil {
		myerror.Errors(ctx, err, "no cart Items", http.StatusBadRequest)
	}
	var cartItems []models.CartItem
	err = db.Where("cart_id=?", userID).Find(&cartItems).Error
	if err != nil {
		myerror.Errors(ctx, err, "cant find cart item", http.StatusBadRequest)
	}
	ctx.JSON(200, gin.H{
		"message": "Cart item updated successfully",
	})

}
