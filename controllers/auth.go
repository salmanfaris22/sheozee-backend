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
	"strconv"
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

	accesToken, err := user.SetRefreshToken()
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
	SetCookieS(ctx, accesToken, user)
	ctx.JSON(200, gin.H{
		"message":      "register successfully",
		"user":         user.FirstName,
		"accesToken":   accesToken,
		"refreshToken": user.Token,
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

	accesToken, err := user.SetRefreshToken()
	if err != nil {
		myerror.Errors(ctx, err, "cant Set refresh Token", http.StatusInternalServerError)
		return
	}

	err = db.Model(&models.User{}).Where("email=?", newUser.Email).Update("token", user.Token).Error
	if err != nil {
		myerror.Errors(ctx, err, "cant set token", http.StatusInternalServerError)
		return
	}
	SetCookieS(ctx, accesToken, user)
	//ctx.SetCookie("token", accesToken, int(15*time.Minute), "/", "localhost", true, true)
	ctx.JSON(200, gin.H{
		"message":      "login successfully",
		"user":         user.FirstName,
		"refreshToken": user.Token,
		"accesToken":   accesToken,
	})
}

func SetCookieS(ctx *gin.Context, accestoken string, user models.User) {
	ctx.Header("Authorization", accestoken)
	userID := strconv.FormatUint(uint64(user.ID), 10)
	ctx.SetCookie("userId", userID, int(7*24*time.Hour), "/", "localhost", true, true)
	ctx.SetCookie("refreshToken", user.Token, int(7*24*time.Hour), "/", "localhost", true, true)
}

func ClearCookies(ctx *gin.Context) {
	ctx.Header("Authorization", "")
	ctx.SetCookie("userId", "", -1, "/", "", false, true)
	ctx.SetCookie("refreshToken", "", -1, "/", "", false, true)
}
func (uc UserController) LogOut(ctx *gin.Context) {
	ClearCookies(ctx)
	ctx.JSON(200, gin.H{
		"message": "logout successfully",
	})
}
