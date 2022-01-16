package router

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/services"
)

func New(service *services.Service) *gin.Engine {
	router := gin.New()
	setupRoutes(router, service)
	return router
}

func setupRoutes(r *gin.Engine, s *services.Service) {
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "up",
		})
	})
}
