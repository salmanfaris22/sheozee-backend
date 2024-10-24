package routes

import (
	"github.com/gin-contrib/cors"
	"main/controllers"
	"main/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	var rout controllers.UserController
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
		product.GET("/", controllers.GetAllProduct)
		product.GET("/:id", controllers.GetProduct)
		product.GET("/search", controllers.SearchProduct)
		product.GET("/filter", controllers.FilterProduct)
	}

	user := r.Group("/user", middleware.TokenAuthMiddleware())
	{
		user.GET("/hy", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "gru",
			})
		})
	}

	r.POST("/logout", rout.LogOut)
}
