package router

import "github.com/gin-gonic/gin"

type Router interface {
	Start()
}

type impel struct {
	gin *gin.Engine
}

func NewRouter() Router {
	return &impel{
		gin: gin.New(),
	}
}
