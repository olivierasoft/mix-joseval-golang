package config

import (
	"github.com/gin-gonic/gin"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/app/controller"
)

func BootstrapGinFramework() {
	g := gin.Default()

	configureRouter(g)

	g.Run()
}

func configureRouter(g *gin.Engine) {
	auth := g.Group("/api/authentication")
	{
		auth.POST("/login", controller.OAuthEntry)
	}

}
