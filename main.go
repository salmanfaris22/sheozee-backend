package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
)

func main() {

	r := gin.Default()
	routes.UserRoutes(r)
	r.Run()

}
