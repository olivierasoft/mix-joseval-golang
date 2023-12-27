package router

import (
	"github.com/gin-gonic/gin"
	"github.com/olivierasoft/mix-joseval-golang.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	authGroup := router.Group("/api/authentication")
	{
		authGroup.GET("/login", handler.Login)
	}
}
