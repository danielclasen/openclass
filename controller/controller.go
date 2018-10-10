package controller

import "github.com/gin-gonic/gin"

func handleError(c *gin.Context, err error, statusCode int) bool {
	if err != nil {
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return true
	}
	return false
}
