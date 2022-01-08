package handlers

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/models"
	"mggers-reports-api/services"
	"net/http"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func Create(service services.ReportsService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.CreateReport

		c.ShouldBindJSON(&request)

		res, err := service.Create(c.Request.Context(), request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, res)


	}
}