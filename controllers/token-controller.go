package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/models"
)

func ValidateRefreshToken(token string, ctx *gin.Context) (string, error) {
	userID, err := ctx.Cookie("userId")
	if err != nil {
		ctx.Abort()
		return "", errors.New("Unauthorized")
	}
	var user models.User
	err = db.Model(&models.User{}).Where("id=?", userID).First(&user).Error
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
