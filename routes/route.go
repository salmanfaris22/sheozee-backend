package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/controllers"
	"main/middleware"
)

func UserRoutes(r *gin.Engine) {
	var rout controllers.UserController
	var products controllers.Product
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600,
		AllowCredentials: true,
	}))

	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", rout.Register)
		authRoute.POST("/login", rout.Login)

	}
	product := r.Group("/products")
	{

		product.GET("/", products.GetAllProduct)
		product.GET("/:id", products.GetProduct)
		product.GET("/search", products.SearchProduct)
		product.GET("/filter", products.FilterProduct)
	}

	user := r.Group("/user", middleware.TokenAuthMiddleware())
	{
		var cart controllers.Cart
		user.POST("/addCart", cart.AddToCart)
		user.GET("/cartItems", cart.GetCartItems)
	}

	r.POST("/logout", rout.LogOut)
}
