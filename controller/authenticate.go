package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olivierasoft/mix-joseval-golang.git/business"
)

func Login(ctx *gin.Context) {

	business.UserLogin("clientId")

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Hello",
	})
}
