package myerror

import "github.com/gin-gonic/gin"

func Errors(ctx *gin.Context, err error, str string, status int) {
	ctx.JSON(status, gin.H{
		"message": str,
		"error":   err.Error(),
	})
	return
}
