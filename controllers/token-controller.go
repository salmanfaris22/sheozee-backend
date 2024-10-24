package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"main/models"
)

func ValidateRefreshToken(token string, ctx *gin.Context) (string, error) {
	userID, exists := ctx.Get("user_Id")
	fmt.Println(userID)
	if !exists {
		ctx.Abort()
		return "", errors.New("Unauthorized")
	}
	var user models.User
	err := db.Model(&models.User{}).Where("id=?", userID).First(&user).Error
	if err != nil {
		ctx.Abort()
		return "", errors.New("Unauthorized")
	}

	if user.Token != token {
		return "", errors.New("Token Not Match")
	}
	newToken, err := user.SetRefreshToken()
	if user.Token != token {
		return "", err
	}
	return newToken, err
}
