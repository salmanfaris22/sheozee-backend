package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"main/config"
	myerror "main/error"
	"main/models"
	"main/services"
	"main/validation"
	"net/http"
	"time"
)

var db = config.ConnectDB()

type UserController struct {
}

func (uc UserController) Register(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		myerror.Errors(ctx, err, "invalid user informations", http.StatusBadRequest)
		return
	}

	err = validation.ValidateUser(ctx, user)
	if err != nil {
		myerror.Errors(ctx, err, "validation failed", http.StatusBadRequest)
		return
	}

	token, err := user.SetAccessToken()
	if err != nil {
		myerror.Errors(ctx, err, "cant Set Acces Token", http.StatusInternalServerError)
		return
	}

	refresh, err := user.SetRefreshToken()
	if err != nil {
		myerror.Errors(ctx, err, "cant Set refresh Token", http.StatusInternalServerError)
		return
	}

	user.Token = token
	err = db.Create(&user).Error
	if err != nil {
		myerror.Errors(ctx, err, "invalid user informations", http.StatusBadRequest)
		return
	}
	ctx.Set("userType", user.Role)
	ctx.Set("user_Id", user.ID)
	ctx.SetCookie("token", user.Token, int(24*time.Hour.Seconds()), "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", refresh, int(24*time.Hour.Seconds()), "/", "localhost", true, true)
	ctx.JSON(200, gin.H{
		"message":      "register successfully",
		"user":         user.FirstName,
		"token":        user.Token,
		"refreshToken": refresh,
	})
}

func (uc UserController) Login(ctx *gin.Context) {
	var newUser models.User
	err := ctx.BindJSON(&newUser)
	if err != nil {
		myerror.Errors(ctx, err, "invalid user informations", http.StatusBadRequest)
		return
	}

	var user models.User
	err = db.Where("email=?", newUser.Email).First(&user).Error
	if err != nil {
		myerror.Errors(ctx, err, "cant find email id", http.StatusBadRequest)
		return
	}

	pass := services.CheckPasswordHash(newUser.Password, user.Password)
	if !pass {
		myerror.Errors(ctx, errors.New("hash pass don't match"), "password not match", http.StatusBadRequest)
		return
	}

	refresh, err := user.SetRefreshToken()
	if err != nil {
		myerror.Errors(ctx, err, "cant Set refresh Token", http.StatusInternalServerError)
		return
	}

	err = db.Model(&models.User{}).Where("email=?", newUser.Email).Update("token", user.Token).Error
	if err != nil {
		myerror.Errors(ctx, err, "cant set token", http.StatusInternalServerError)
		return
	}
	ctx.Set("userType", user.Role)
	ctx.Set("user_Id", user.ID)
	ctx.SetCookie("refreshToken", refresh, int(24*time.Hour.Seconds()), "/", "localhost", true, true)

	ctx.JSON(200, gin.H{
		"message":      "login successfully",
		"user":         user.FirstName,
		"refreshToken": refresh,
		"token":        user.Token,
	})
}

func (uc UserController) LogOut(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "logout successfully",
	})
}
