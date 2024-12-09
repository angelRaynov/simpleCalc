package calculator

import "github.com/gin-gonic/gin"

type Handler interface {
	EvaluateHandler(c *gin.Context)
	ValidateHandler(c *gin.Context)
	ErrorHandler(c *gin.Context)
}
