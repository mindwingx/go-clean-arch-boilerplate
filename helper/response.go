package helper

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, statusCode int, data map[string]interface{}) {
	c.JSON(statusCode, map[string]interface{}{
		"status_code": statusCode,
		"error":       nil,
		"data":        data,
	})
	c.Abort()
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, map[string]interface{}{
		"status_code": statusCode,
		"error":       err.Error(),
		"data":        nil,
	})
}
