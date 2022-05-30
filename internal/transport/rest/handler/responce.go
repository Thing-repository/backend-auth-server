package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.WithFields(logrus.Fields{
		"module":      "handler",
		"func":        "newErrorResponse",
		"status_code": statusCode,
		"message":     message,
	})
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
}
