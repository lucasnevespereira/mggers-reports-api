package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	}
}
