package handlers

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/models"
	"mggers-reports-api/services"
	"net/http"
)

func Create(service *services.ReportService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.ReportRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "json: " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		report, err := service.CreateReport(c.Request.Context(), request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  err.Error(),
				"status": http.StatusInternalServerError,
			})
			return
		}

		c.JSON(http.StatusCreated, report)
	}
}

