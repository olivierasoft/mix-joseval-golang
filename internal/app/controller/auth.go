package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/app/service"
)

func Login(c *gin.Context) {

	code := c.Query("code")

	if len(code) == 0 {
		sendError(c, http.StatusBadRequest, "O código enviado é inválido", true)
		return
	}

	jwt, err := service.Login(code)

	if err != nil {
		sendError(c, http.StatusBadRequest, err.Error(), false)
		return
	}

	c.JSON(200, map[string]string{
		"authorization": *jwt,
	})
}
