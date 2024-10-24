package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"main/controllers"
	myerror "main/error"
	"net/http"
	"time"
)

var secretKey = []byte("your_secret_key")

func TokenAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			myerror.Errors(ctx, errors.New("unauthorized"), "Unauthorized User", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		valid, err := validateToken(token)
		if err != nil || !valid {
			refreshToken, err := ctx.Cookie("refreshToken")
			if err != nil {
				myerror.Errors(ctx, err, "Refresh token not found", http.StatusUnauthorized)
				ctx.Abort()
				return
			}

			// Validate and refresh the token
			newToken, err := controllers.ValidateRefreshToken(refreshToken, ctx)
			if err != nil {
				myerror.Errors(ctx, err, "Invalid refresh token", http.StatusUnauthorized)
				ctx.Abort()
				return
			}

			if newToken != "" {
				ctx.SetCookie("token", newToken, int(15*time.Minute), "/", "localhost", true, true)
				ctx.Next()
				return
			}

			myerror.Errors(ctx, errors.New("invalid or expired token"), "Invalid or expired token", http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func validateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
