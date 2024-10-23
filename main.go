package main

import (
	"github.com/gin-gonic/gin"
	"main/config"
	"main/routes"
)

func main() {

	config.ConnectDB()
	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()

}
