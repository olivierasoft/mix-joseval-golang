package controller

import (
	"github.com/gin-gonic/gin"
)

func OAuthEntry(c *gin.Context) {
	c.JSON(200, "Hail")
}
