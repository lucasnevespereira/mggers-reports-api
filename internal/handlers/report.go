package handlers

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/internal/models"
	"mggers-reports-api/internal/services"
	"net/http"
)


func CreateReport(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request models.ReportRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":  "bind json" + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

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

func GetReports(service *services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		res, err := service.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"results": res,
		})
	}
}
