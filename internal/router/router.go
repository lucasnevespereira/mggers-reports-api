package router

import (
	"github.com/gin-gonic/gin"
	"mggers-reports-api/internal/handlers"
	"mggers-reports-api/internal/services"
)

func Init(service *services.Service) *gin.Engine {
	router := gin.New()
	setupRoutes(router, service)
	return router
}

func setupRoutes(r *gin.Engine, s *services.Service) {
	r.GET("/", handlers.GetHealth())
	reports := r.Group("/reports")
	reports.POST("/create", handlers.CreateReport(s))
	reports.GET("/reports", handlers.GetReports(s))
	reports.GET("/delete/:reportId", handlers.DeleteReport(s))
}
