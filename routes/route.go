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

	}
	product := r.Group("/products")
	{
		product.GET("/", controllers.GetAllProduct)
		product.GET("/:id", controllers.GetProduct)
		product.GET("/search", controllers.SearchProduct)
		product.GET("/filter", controllers.FilterProduct)
	}

	r.POST("/logout", rout.LogOut)
}
