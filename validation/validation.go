package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/models"
)

func ValidateUser(ctx *gin.Context, user models.User) error {
	err := validator.New().Struct(&user)
	if err != nil {
		return err
	}
	return nil
}
