package routes

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	var rout controllers.UserController
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", rout.Register)
		authRoute.POST("/login", rout.Login)
		authRoute.POST("/logout", rout.LogOut)
	}
}
