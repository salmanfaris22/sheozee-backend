package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	myerror "main/error"
	"main/models"
	"net/http"
)

type UserController struct {
}

func (uc UserController) Register(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		myerror.Errors(ctx, err, "invalid user informations", http.StatusBadRequest)
		return
	}
	validat := validator.New()
	err = validat.Struct(&user)
	if err != nil {
		myerror.Errors(ctx, err, "invalid user informations", http.StatusBadRequest)
		return
	}
	ctx.JSON(200, gin.H{
		"message": user,
	})
}
