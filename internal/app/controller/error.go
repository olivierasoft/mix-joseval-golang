package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/domain/exception"
)

func sendError(c *gin.Context, status int, message string, show bool) {
	errorEntity := exception.HttpError{
		Code:    status,
		Message: message,
		Show:    show,
	}

	c.JSON(status, &errorEntity)
}
