package router

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/handlers"
	"mggers-reports-api/services"
)

func Init(services *services.Services) *gin.Engine {
	r := gin.New()
	setupRoutes(r, services)
	return r
}

func setupRoutes(r *gin.Engine, services *services.Services) {
	r.GET("/", handlers.GetHealth)
	r.POST("/api/create", handlers.Create(services.ReportsService))
}