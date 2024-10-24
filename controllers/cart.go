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
	userID := ctx.Param("id")

	storeID, err := ctx.Cookie("userId")
	if err != nil {
		myerror.Errors(ctx, err, "can't add toCart", http.StatusInternalServerError)
		return
	}
	if userID != storeID {
		myerror.Errors(ctx, errors.New("user wrong user"), "cant get Cart", http.StatusInternalServerError)
		return
	}

	id := ctx.Query("product_id")
	ProductID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		myerror.Errors(ctx, errors.New("select product"), "cant Get Product", http.StatusInternalServerError)
		return
	}
	Quantity, err := strconv.Atoi(ctx.Query("qty"))
	if err != nil {
		Quantity = 1
	}
	var product models.Product
	err = db.First(&product, ProductID).Error
	var cart models.Cart
	if err != nil {
		myerror.Errors(ctx, errors.New("quatyu unvalible"), "cant find Product", http.StatusInternalServerError)
		return
	}
	if product.Stock == 0 || product.IsAvailable == false {
		myerror.Errors(ctx, errors.New("out of stock"), "stoke not uvalible", http.StatusBadRequest)
		return
	}

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
	var cartItem models.CartItem

	err = db.Where("cart_id = ? AND product_id = ?", cart.ID, product.ID).First(&cartItem).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound && Quantity > 0 {
			cartItem.Quantity = Quantity

			cartItem.ProductID = uint(ProductID)
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

		if Quantity <= 0 {
			cartItem.Quantity += Quantity
			fmt.Println(cartItem.ProductID)
			if cartItem.Quantity <= 0 {

				err = db.Delete(&cartItem).Error
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to remove item from cart"})
					return
				}
			}
			ctx.JSON(200, gin.H{
				"message": "Cart item updated successfully",
			})
			return

		} else {
			cartItem.Quantity += Quantity
		}

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
	userID := ctx.Param("id")
	storeID, err := ctx.Cookie("userId")
	if err != nil {
		myerror.Errors(ctx, err, "Cat This add toCart", http.StatusInternalServerError)
		return
	}
	if userID != storeID {
		myerror.Errors(ctx, errors.New("user wrong user"), "cant get Cart", http.StatusInternalServerError)
		return
	}
	var carts models.Cart
	err = db.Where("user_id=?", userID).First(&carts).Error
	if err != nil {
		myerror.Errors(ctx, err, "no cart Items", http.StatusBadRequest)
		return
	}
	var cartItems []models.CartItemWithProduct

	//err = db.Where("cart_id=?", userID).Find(&cartItems).Error
	//if err != nil {
	//	myerror.Errors(ctx, err, "cant find cart item", http.StatusBadRequest)
	//}

	err = db.Table("cart_items").
		Select("cart_items.*, products.* ").
		Joins("JOIN products ON products.id=cart_items.product_id ").
		Where("cart_items.cart_id=?", carts.ID).Scan(&cartItems).Error
	if err != nil {
		myerror.Errors(ctx, err, "can't fins cart itesms", http.StatusBadRequest)
		return
	}
	sum := c.PrizeSum(cartItems)
	ctx.JSON(200, gin.H{
		"message": cartItems,
		"cart":    carts,
		"totole":  sum,
	})

}

func (c Cart) PrizeSum(item []models.CartItemWithProduct) float64 {
	sum := 0.0
	for _, v := range item {
		itemTotle := v.Price * float64(v.Quantity)
		sum += itemTotle
	}
	return sum
}
